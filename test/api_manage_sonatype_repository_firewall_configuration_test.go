/*
Sonatype Nexus Repository Manager

Testing ManageSonatypeRepositoryFirewallConfigurationAPIService

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

func Test_sonatyperepo_ManageSonatypeRepositoryFirewallConfigurationAPIService(t *testing.T) {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService DisableIq", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.DisableIq(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService EnableIq", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.EnableIq(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService GetConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.GetConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService UpdateConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.UpdateConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService VerifyConnection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.VerifyConnection(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
