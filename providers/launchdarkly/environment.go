package launchdarkly

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type EnvironmentGenerator struct {
	LaunchDarklyService
}

func (g *EnvironmentGenerator) InitResources() error {

	ldClient := g.createClient()
	g.createEnvironmentResources(ldClient)
	return nil
}

func (g *EnvironmentGenerator) createEnvironmentResources(ldClient LaunchDarklyClient) error {
	ctx := ldClient.ctx
	projectKey := ldClient.projectKey
	environmentKey := ldClient.environmentKey
	client := ldClient.client

	environment, _, err := client.EnvironmentsApi.GetEnvironment(ctx, projectKey, environmentKey)
	if err != nil {
		return err
	}

	resource := terraformutils.NewSimpleResource(
		environment.Key,
		environment.Name,
		"launchdarkly_environment",
		g.ProviderName,
		[]string{})

	resource.InstanceState.Attributes["key"] = environment.Key
	resource.InstanceState.Attributes["project_key"] = projectKey
	g.Resources = append(g.Resources, resource)
	return nil
}
