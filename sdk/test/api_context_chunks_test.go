/*
FastAPI- Private GPT

Testing ContextChunksAPIService

*/

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/RomanGod6/privategpt-go/sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ContextChunksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContextChunksAPIService ChunksRetrievalV1ChunksPost", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ContextChunksAPI.ChunksRetrievalV1ChunksPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
