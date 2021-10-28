package launchdarkly

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ProjectGenerator struct {
	LaunchDarklyService
}

func (g *ProjectGenerator) InitResources() error {
	ldClient := g.createClient()
	g.createProjectResources(ldClient)
	return nil
}
func (g *ProjectGenerator) createProjectResources(ldClient LaunchDarklyClient) error {
	ctx := ldClient.ctx
	projectKey := ldClient.projectKey
	client := ldClient.client
	project, _, err := client.ProjectsApi.GetProject(ctx, projectKey)
	if err != nil {
		return err
	}

	resource := terraformutils.NewSimpleResource(
		project.Id,
		project.Name,
		"launchdarkly_project",
		g.ProviderName,
		[]string{"tags."})

	resource.ResourceName = project.Key
	resource.InstanceState.Attributes["key"] = project.Key

	g.Resources = append(g.Resources, resource)

	return nil
}
