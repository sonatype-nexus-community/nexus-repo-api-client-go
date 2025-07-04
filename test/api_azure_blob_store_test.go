/*
Sonatype Nexus Repository Manager

Testing AzureBlobStoreAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v3"
)

func Test_v3_AzureBlobStoreAPIService(t *testing.T) {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)

	t.Run("Test AzureBlobStoreAPIService VerifyConnection1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AzureBlobStoreAPI.VerifyConnection1(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
