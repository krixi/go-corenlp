package corenlp_test

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/krixi/go-corenlp"
)

func TestHTTPClient(t *testing.T) {

	tests := []struct {
		description string
		request     string
		test        func(ctx context.Context, req string, c corenlp.HTTPClient)
	}{
		{
			description: "Basic parse is returned",
			request:     "The quick brown fox jumps over the lazy dog.",
			test: func(ctx context.Context, req string, c corenlp.HTTPClient) {

				resp, err := c.Run(ctx, req)
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				defer resp.Close()

				rawjson, err := ioutil.ReadAll(resp)
				assert.NoError(t, err)
				assert.True(t, len(rawjson) > 100)
			},
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.description, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			c := corenlp.NewHTTPClient("http://localhost:9000")

			scenario.test(ctx, scenario.request, *c)

		})
	}
}
