package schema

import (
	"context"
	"testing"

	"github.com/go-ggz/ggz/model"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

func TestQueryShortenURL(t *testing.T) {
	assert.NoError(t, model.PrepareTestDatabase())
	user, _ := model.GetUserByID(1)
	ctx := newContextWithUser(context.TODO(), user)
	t.Run("shorten url exist", func(t *testing.T) {
		test := T{
			Query: `
{
  QueryShortenURL(slug: "abcdef") {
    url
  }
}
	  `,
			Schema: Schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"QueryShortenURL": map[string]interface{}{
						"url": "http://example.com",
					},
				},
			},
		}
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
			Context:       ctx,
		}
		testGraphql(test, params, t)
	})

	t.Run("shorten url not exist", func(t *testing.T) {
		test := T{
			Query: `
{
  QueryShortenURL(slug: "1234567890") {
    url
  }
}
`,
			Schema: Schema,
		}
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
			Context:       ctx,
		}
		testGraphqlErr(test, params, t)
	})

}
