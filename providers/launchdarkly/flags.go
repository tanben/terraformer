package launchdarkly

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/antihax/optional"
	ldapi "github.com/launchdarkly/api-client-go"
)

type FlagsGenerator struct {
	LaunchDarklyService
}

func (g *FlagsGenerator) InitResources() error {
	ldClient := g.createClient()
	g.createFlagsResources(ldClient)
	return nil
}

func (g *FlagsGenerator) createFlagsResources(ldClient LaunchDarklyClient) error {
	ctx := ldClient.ctx
	projectKey := ldClient.projectKey
	environmentKey := ldClient.environmentKey
	client := ldClient.client
	options := &ldapi.FeatureFlagsApiGetFeatureFlagsOpts{
		Env: optional.NewInterface(environmentKey),
	}
	flags, _, err := client.FeatureFlagsApi.GetFeatureFlags(ctx, projectKey, options)
	if err != nil {
		return err
	}

	for i, flag := range flags.Items {

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			flag.Key,
			flag.Name,
			"launchdarkly_feature_flag",
			g.ProviderName,
			[]string{}))

		g.Resources[i].InstanceState.Attributes["key"] = flag.Key
		g.Resources[i].InstanceState.Attributes["project_key"] = projectKey
	}

	return nil
}
