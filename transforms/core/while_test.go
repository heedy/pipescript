package core

import (
	"testing"

	"github.com/heedy/pipescript"
	"github.com/stretchr/testify/require"
)

func TestWhile(t *testing.T) {
	require.NoError(t, While.Register())
	// This needs much more testing
	pipescript.TestCase{
		Pipescript: "while d!=1 d",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 5},
			{Timestamp: 2, Data: 2},
			{Timestamp: 3, Data: 3},
			{Timestamp: 4, Data: 1},
			{Timestamp: 5, Data: 2},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 3, Data: 3},
			{Timestamp: 5, Data: 2},
		},
	}.Run(t)
	require.NoError(t, I.Register())
	pipescript.TestCase{
		Pipescript: "while(i%3!=0, d)",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 5},
			{Timestamp: 2, Data: 2},
			{Timestamp: 3, Data: 3},
			{Timestamp: 4, Data: 1},
			{Timestamp: 5, Data: 2},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 3, Data: 3},
			{Timestamp: 5, Data: 2},
		},
	}.Run(t)

}
