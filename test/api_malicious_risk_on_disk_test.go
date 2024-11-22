/*
Sonatype Nexus Repository Manager

Testing MaliciousRiskOnDiskAPIService

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

func Test_sonatyperepo_MaliciousRiskOnDiskAPIService(t *testing.T) {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)

	t.Run("Test MaliciousRiskOnDiskAPIService GetEnabledRegistries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.MaliciousRiskOnDiskAPI.GetEnabledRegistries(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MaliciousRiskOnDiskAPIService GetMaliciousRiskOnDiskCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.MaliciousRiskOnDiskAPI.GetMaliciousRiskOnDiskCount(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}