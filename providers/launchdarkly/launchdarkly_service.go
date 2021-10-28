package launchdarkly

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	ldapi "github.com/launchdarkly/api-client-go"
)

type LaunchDarklyClient struct {
	apiKey         string
	projectKey     string
	environmentKey string
	ctx            context.Context
	client         *ldapi.APIClient
}

type LaunchDarklyService struct { //nolint
	terraformutils.Service
}

func (s *LaunchDarklyService) createClient() LaunchDarklyClient {

	args := s.GetArgs()
	apiKey := args["apiKey"].(string)
	projectKey := args["projectKey"].(string)
	environmentKey := args["environmentKey"].(string)
	client := ldapi.NewAPIClient(ldapi.NewConfiguration())
	ctx := context.WithValue(context.Background(), ldapi.ContextAPIKey, ldapi.APIKey{
		Key: apiKey,
	})
	ldClient := LaunchDarklyClient{
		apiKey,
		projectKey,
		environmentKey,
		ctx,
		client,
	}

	return ldClient
}
