package corenlp_test

import (
	"context"
	"testing"
	"time"

	"github.com/krixi/go-corenlp"
	"github.com/stretchr/testify/assert"
)

func TestCoreNLP(t *testing.T) {
	tests := []struct {
		description string
		request     string
		test        func(ctx context.Context, req string, c corenlp.Client)
	}{
		{
			description: "Annotate returns list of annotated sentences",
			request:     "The quick brown fox jumps over the lazy dog.",
			test: func(ctx context.Context, req string, c corenlp.Client) {

				root, err := c.Annotate(ctx, req)
				assert.NoError(t, err)
				assert.NotNil(t, root)

				assert.Len(t, root.Sentences, 1)

				sentence := root.Sentences[0]
				assert.Equal(t, 0, sentence.Index)
				assert.NotEmpty(t, sentence.Dependencies)
				assert.NotEmpty(t, sentence.EnhancedPlusPlusDependencies)
				assert.NotEmpty(t, sentence.OpenIETuples)
				assert.NotEmpty(t, sentence.Tokens)
				assert.NotEmpty(t, sentence.RawParse)
			},
		},
		{
			description: "Annotate returns multiple sentences",
			request:     "How much wood would a woodchuck chuck if a woodchuck could chuck wood? A woodchuck would chuck all that he could chuck if a woodchuck could chuck wood.",
			test: func(ctx context.Context, req string, c corenlp.Client) {

				root, err := c.Annotate(ctx, req)
				assert.NoError(t, err)
				assert.NotNil(t, root)

				assert.Len(t, root.Sentences, 2)
				assert.Equal(t, 0, root.Sentences[0].Index)
				assert.Equal(t, 1, root.Sentences[1].Index)
			},
		},
		{
			description: "KBP and EntityMentions are returned",
			request:     "Barack Obama was the 44th president of the United States",
			test: func(ctx context.Context, req string, c corenlp.Client) {

				root, err := c.Annotate(ctx, req)
				assert.NoError(t, err)
				assert.NotNil(t, root)

				assert.Len(t, root.Sentences, 1)
				sentence := root.Sentences[0]
				assert.NotEmpty(t, sentence.KBP)
				assert.NotEmpty(t, sentence.EntityMentions)
			},
		},
		{
			description: "Corefs are returned",
			request:     "President Xi Jinping of Chaina, on his first state visit to the United States, showed off his familiarity with American history and pop culture on Tuesday night.",
			test: func(ctx context.Context, req string, c corenlp.Client) {

				root, err := c.Annotate(ctx, req)
				assert.NoError(t, err)
				assert.NotNil(t, root)
				assert.NotEmpty(t, root.Corefs)
			},
		},
	}

	// Initialize the test client.
	httpClient := corenlp.NewHTTPClient("http://localhost:9000")
	c := corenlp.NewClient(httpClient)

	for _, scenario := range tests {
		t.Run(scenario.description, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			scenario.test(ctx, scenario.request, *c)
		})
	}
}
