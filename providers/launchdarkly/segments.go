package launchdarkly

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SegmentGenerator struct {
	LaunchDarklyService
}

func (g *SegmentGenerator) InitResources() error {
	ldClient := g.createClient()
	g.createCustomRoleResources(ldClient)
	return nil
}

func (g *SegmentGenerator) createCustomRoleResources(ldClient LaunchDarklyClient) error {
	ctx := ldClient.ctx
	client := ldClient.client
	projectKey := ldClient.projectKey
	env_key := ldClient.environmentKey
	userSegments, _, err := client.UserSegmentsApi.GetUserSegments(ctx, projectKey, env_key, nil)
	if err != nil {
		return err
	}

	for i, userSegment := range userSegments.Items {

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			userSegment.Key,
			userSegment.Name,
			"launchdarkly_segment",
			g.ProviderName,
			[]string{}))

		g.Resources[i].InstanceState.Attributes["key"] = userSegment.Key
		g.Resources[i].InstanceState.Attributes["name"] = userSegment.Name
		g.Resources[i].InstanceState.Attributes["project_key"] = projectKey
		g.Resources[i].InstanceState.Attributes["env_key"] = env_key

	}

	return nil
}
