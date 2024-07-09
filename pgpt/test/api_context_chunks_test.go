/*
FastAPI- Private GPT

Testing ContextChunksAPIService

*/

package pgpt_test

import (
	"context"
	"testing"

	"github.com/RomanGod6/privategpt-go/pgpt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ContextChunksAPIService(t *testing.T) {

	configuration := pgpt.NewConfiguration()
	apiClient := pgpt.NewAPIClient(configuration)

	t.Run("Test ContextChunksAPIService ChunksRetrievalV1ChunksPost", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ContextChunksAPI.ChunksRetrievalV1ChunksPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
