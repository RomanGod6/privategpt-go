/*
FastAPI- Private GPT

Testing EmbeddingsAPIService

*/

package sdk

import (
	"context"
	"testing"

	sdk "github.com/RomanGod6/privategpt-go/sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_EmbeddingsAPIService(t *testing.T) {

	configuration := sdk.NewConfiguration()
	apiClient := sdk.NewAPIClient(configuration)

	t.Run("Test EmbeddingsAPIService EmbeddingsGenerationV1EmbeddingsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EmbeddingsAPI.EmbeddingsGenerationV1EmbeddingsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
