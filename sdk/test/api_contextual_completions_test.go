/*
FastAPI- Private GPT

Testing ContextualCompletionsAPIService

*/

package sdk

import (
	"context"
	"testing"

	"github.com/RomanGod6/privategpt-go/sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ContextualCompletionsAPIService(t *testing.T) {

	configuration := sdk.NewConfiguration()
	apiClient := sdk.NewAPIClient(configuration)

	t.Run("Test ContextualCompletionsAPIService ChatCompletionV1ChatCompletionsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ContextualCompletionsAPI.ChatCompletionV1ChatCompletionsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContextualCompletionsAPIService PromptCompletionV1CompletionsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ContextualCompletionsAPI.PromptCompletionV1CompletionsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
