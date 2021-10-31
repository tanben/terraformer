package launchdarkly

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type CustomRoleGenerator struct {
	LaunchDarklyService
}

func (g *CustomRoleGenerator) InitResources() error {
	ldClient := g.createClient()
	g.createCustomRoleResources(ldClient)
	return nil
}

func (g *CustomRoleGenerator) createCustomRoleResources(ldClient LaunchDarklyClient) error {
	ctx := ldClient.ctx
	client := ldClient.client
	customRoles, _, err := client.CustomRolesApi.GetCustomRoles(ctx)
	if err != nil {
		return err
	}

	for i, customRole := range customRoles.Items {

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			customRole.Key,
			customRole.Name,
			"launchdarkly_custom_role",
			g.ProviderName,
			[]string{}))

		g.Resources[i].InstanceState.Attributes["customRoleKey"] = customRole.Key

	}

	return nil
}
