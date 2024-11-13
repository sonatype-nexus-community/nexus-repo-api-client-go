/*
Sonatype Nexus Repository Manager

Testing SecurityManagementSecretsEncryptionAPIService

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

func Test_sonatyperepo_SecurityManagementSecretsEncryptionAPIService(t *testing.T) {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)

	t.Run("Test SecurityManagementSecretsEncryptionAPIService ReEncrypt", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementSecretsEncryptionAPI.ReEncrypt(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
