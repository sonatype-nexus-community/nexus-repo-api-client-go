/*
Sonatype Nexus Repository Manager

Testing CleanupPoliciesAPIService

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

func Test_v3_CleanupPoliciesAPIService(t *testing.T) {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)

	t.Run("Test CleanupPoliciesAPIService Create1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CleanupPoliciesAPI.Create1(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CleanupPoliciesAPIService DeletePolicyByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.CleanupPoliciesAPI.DeletePolicyByName(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CleanupPoliciesAPIService GetAll", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CleanupPoliciesAPI.GetAll(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CleanupPoliciesAPIService GetCleanupPolicyByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.CleanupPoliciesAPI.GetCleanupPolicyByName(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CleanupPoliciesAPIService Update1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyName string

		httpRes, err := apiClient.CleanupPoliciesAPI.Update1(context.Background(), policyName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
