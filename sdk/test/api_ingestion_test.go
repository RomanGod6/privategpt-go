/*
FastAPI- Private GPT

Testing IngestionAPIService

*/

package sdk

import (
	"context"
	"testing"

	sdk "github.com/RomanGod6/privategpt-go/sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_IngestionAPIService(t *testing.T) {

	configuration := sdk.NewConfiguration()
	apiClient := sdk.NewAPIClient(configuration)

	t.Run("Test IngestionAPIService DeleteIngestedV1IngestDocIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var docId string

		resp, httpRes, err := apiClient.IngestionAPI.DeleteIngestedV1IngestDocIdDelete(context.Background(), docId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IngestionAPIService IngestFileV1IngestFilePost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IngestionAPI.IngestFileV1IngestFilePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IngestionAPIService IngestTextV1IngestTextPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IngestionAPI.IngestTextV1IngestTextPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IngestionAPIService IngestV1IngestPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IngestionAPI.IngestV1IngestPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IngestionAPIService ListIngestedV1IngestListGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IngestionAPI.ListIngestedV1IngestListGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
