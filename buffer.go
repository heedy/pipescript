package pipescript

import (
	"container/list"
	"errors"
	"fmt"
)

var ErrBeforeStart = errors.New("There is no value for this query yet")

const bufferPageSize = 3
const backPages = 1
const uintMax = ^uint64(0)

type bufferPage struct {
	DP []*Datapoint // Array of datapoints that make up page
	I  uint64       // Page index
}

type Buffer struct {
	Input Iterator

	Pages *list.List
	// The last page, after which there are no more datapoints
	// Initially set to max int64
	EndPage uint64
	// An error that was encountered
	Error error

	Iterators []*BufferIterator
}

func NewBuffer(n Iterator) *Buffer {
	return &Buffer{
		Input:     n,
		Pages:     list.New(),
		EndPage:   uintMax,
		Iterators: make([]*BufferIterator, 0, 1),
	}
}

func (b *Buffer) Iterator() *BufferIterator {
	it := &BufferIterator{
		Buf:  b,
		Elem: b.Pages.Front(),
		J:    0,
	}
	b.Iterators = append(b.Iterators, it)
	return it
}

func (b *Buffer) nextPage() {
	// Gets a new page of data. Assumes that it was called by an iterator

	// Find the next page index
	le := b.Pages.Back()
	nextIndex := uint64(0)
	if le != nil {
		nextIndex = le.Value.(*bufferPage).I + 1
	}

	// First, check the lowest page index of any iterator
	lowestIndex := uintMax
	for _, it := range b.Iterators {
		if it.Elem == nil {
			lowestIndex = 0
			break
		}
		curpage := it.Elem.Value.(*bufferPage).I
		if curpage < lowestIndex {
			lowestIndex = curpage
		}
	}

	// Now, determine if we need to add a page, or if we can use an existing page
	// by taking it from the beginning of the list
	fp := b.Pages.Front()
	if fp == nil || fp.Value.(*bufferPage).I+backPages >= lowestIndex {
		// Need to add a new page
		bp := &bufferPage{
			DP: make([]*Datapoint, bufferPageSize),
		}

		b.Pages.PushBack(bp)
	} else {
		b.Pages.MoveToBack(fp)
	}

	// We now have a valid page at the back
	bp := b.Pages.Back().Value.(*bufferPage)
	bp.I = nextIndex

	// ... and now populate it with values
	for i := range bp.DP {
		if bp.DP[i] == nil {
			bp.DP[i] = &Datapoint{}
		}
		bp.DP[i], b.Error = b.Input.Next(bp.DP[i])
		if b.Error != nil {
			return
		}
		if bp.DP[i] == nil {
			// This is the end of the stream - clear the rest of the array,
			// and set the end index
			for j := i + 1; j < len(bp.DP); j++ {
				bp.DP[j] = nil
			}
			b.EndPage = nextIndex
			break
		}
	}

}

type BufferIterator struct {
	Buf  *Buffer
	Elem *list.Element
	J    int
}

func (i *BufferIterator) init() {
	if i.Elem == nil {
		i.Elem = i.Buf.Pages.Front()
		if i.Elem == nil && i.Buf.Error == nil {
			i.Buf.nextPage()
		}
		i.Elem = i.Buf.Pages.Front()
	}
}

func (i *BufferIterator) Next() (*Datapoint, error) {
	i.init()
	if i.Buf.Error != nil {
		return nil, i.Buf.Error
	}
	bp := i.Elem.Value.(*bufferPage)

	if i.J >= bufferPageSize {
		// We are at the end of a page. Move to the next one

		// ...but if it is the last page, just return nil
		if bp.I >= i.Buf.EndPage {
			return nil, nil
		}

		ne := i.Elem.Next()
		if ne == nil {
			// Ask the buffer for another page
			i.Buf.nextPage()
			if i.Buf.Error != nil {
				return nil, i.Buf.Error
			}
			ne = i.Elem.Next()
		}
		i.Elem = ne
		bp = i.Elem.Value.(*bufferPage)
		i.J = 0
	}
	dp := bp.DP[i.J]
	i.J++
	return dp, nil
}

func (i *BufferIterator) Peek(idx int) (*Datapoint, error) {
	i.init()
	if i.Buf.Error != nil {
		return nil, i.Buf.Error
	}

	// The actual index relative to the start of the page
	idx = idx + i.J
	if idx >= 0 {
		pages := uint64(idx / bufferPageSize)
		pageIndex := idx % bufferPageSize

		curp := i.Elem
		pval := curp.Value.(*bufferPage)
		// If peeking past end of stream, return nils
		if pval.I+pages > i.Buf.EndPage {
			return nil, nil
		}
		// Move to the desired page
		for j := uint64(0); j < pages; j++ {
			nextp := curp.Next()
			if nextp == nil {
				i.Buf.nextPage()
				if i.Buf.Error != nil {
					return nil, i.Buf.Error
				}
				nextp = curp.Next()
				// Check once more if we're peeking past end of stream,
				// since we might have just reached end
				if pval.I+pages-j > i.Buf.EndPage {
					return nil, nil
				}
			}
			curp = nextp
			pval = curp.Value.(*bufferPage)
		}

		// Now return the datapoint at the desired index
		return pval.DP[pageIndex], nil
	}

	// The index is negative, so we look *back*
	pageIndex := bufferPageSize + idx%bufferPageSize
	pages := uint64(1 - idx/bufferPageSize)
	if pageIndex == 0 {
		pages--
	}
	// See if we have the capability to look back this far
	curp := i.Elem
	curv := curp.Value.(*bufferPage)
	if curv.I < pages || i.Buf.Pages.Front().Value.(*bufferPage).I > curv.I-pages {
		if curv.I < pages {
			return nil, ErrBeforeStart
		}
		return nil, fmt.Errorf("Peeking to relative index %d failed", idx-i.J)
	}
	// We can look back that far, so go directly to the page
	for j := uint64(0); j < pages; j++ {
		curp = curp.Prev()
	}
	curv = curp.Value.(*bufferPage)
	return curv.DP[pageIndex], nil
}
