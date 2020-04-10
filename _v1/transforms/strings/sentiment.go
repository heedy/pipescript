package strings

import (
	"encoding/json"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

var (
	sentimentWordMap map[string]int
)

// We initialize the sentiment word list from the AFINN dataset
func init() {
	err := json.Unmarshal(resources.MustAsset("data/AFINN.json"), &sentimentWordMap)
	if err != nil {
		panic(err)
	}
}

type sentimentTransform struct {
}

func (t *sentimentTransform) Copy() (pipescript.TransformInstance, error) {
	return &sentimentTransform{}, nil
}

func (t *sentimentTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	s, err := te.Datapoint.DataString()
	if err != nil {
		return nil, err
	}
	toks := Tokenize(s)

	sentiment := 0
	for i := range toks {
		if wordsentiment, ok := sentimentWordMap[toks[i]]; ok {
			sentiment += wordsentiment
		}
	}

	if len(toks) == 0 {
		return te.Set(float32(0))
	}

	// The AFINN dataset has up to +-5
	return te.Set(float32(sentiment) / float32(5*len(toks)))
}

// Sentiment is a basic sentiment analysis transform
var Sentiment = pipescript.Transform{
	Name:          "sentiment",
	Description:   "Returns a value in [-1,1] as a simple measure of a text's sentiment",
	Documentation: string(resources.MustAsset("docs/transforms/sentiment.md")),
	OutputSchema:  `{"type": "number"}`,
	InputSchema:   `{"type": "string"}`,
	OneToOne:      true,
	Stateless:     true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &sentimentTransform{}}, nil
	},
}
