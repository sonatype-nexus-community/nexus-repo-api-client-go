/*
Sonatype Nexus Repository Manager

Testing SecurityManagementRealmsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sonatyperepo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func Test_sonatyperepo_SecurityManagementRealmsAPIService(t *testing.T) {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)

	t.Run("Test SecurityManagementRealmsAPIService GetActiveRealms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SecurityManagementRealmsAPI.GetActiveRealms(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementRealmsAPIService GetRealms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SecurityManagementRealmsAPI.GetRealms(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementRealmsAPIService SetActiveRealms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementRealmsAPI.SetActiveRealms(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
