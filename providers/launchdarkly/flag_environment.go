package launchdarkly

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/antihax/optional"
	ldapi "github.com/launchdarkly/api-client-go"
)

type FlagEnvironmentGenerator struct {
	LaunchDarklyService
}

func (g *FlagEnvironmentGenerator) InitResources() error {

	ldClient := g.createClient()
	g.createFlagEnvironmentResources(ldClient)
	return nil
}

func (g *FlagEnvironmentGenerator) createFlagEnvironmentResources(ldClient LaunchDarklyClient) error {
	ctx := ldClient.ctx
	projectKey := ldClient.projectKey
	environmentKey := ldClient.environmentKey
	client := ldClient.client
	flagsOptions := &ldapi.FeatureFlagsApiGetFeatureFlagsOpts{
		Env: optional.NewInterface(environmentKey),
	}

	flags, _, err := client.FeatureFlagsApi.GetFeatureFlags(ctx, projectKey, flagsOptions)
	if err != nil {
		return err
	}

	for i, flag := range flags.Items {

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			flag.Key,
			flag.Name,
			"launchdarkly_feature_flag_environment",
			g.ProviderName,
			[]string{}))

		g.Resources[i].InstanceState.Attributes["flag_id"] = fmt.Sprintf("%s/%s", projectKey, flag.Key)
		g.Resources[i].InstanceState.Attributes["env_key"] = environmentKey
	}

	return nil
}
