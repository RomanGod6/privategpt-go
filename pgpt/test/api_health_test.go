/*
FastAPI- Private GPT

Testing HealthAPIService

*/

package pgpt

import (
	"context"
	"testing"

	"github.com/RomanGod6/privategpt-go/pgpt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_HealthAPIService(t *testing.T) {

	configuration := pgpt.NewConfiguration()
	apiClient := pgpt.NewAPIClient(configuration)

	t.Run("Test HealthAPIService HealthHealthGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HealthAPI.HealthHealthGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
