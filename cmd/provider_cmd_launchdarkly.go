package cmd

import (
	launchdarkly_terraformer "github.com/GoogleCloudPlatform/terraformer/providers/launchdarkly"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdLaunchDarklyImporter(options ImportOptions) *cobra.Command {
	var apiKey, projectKey, environmentKey string

	cmd := &cobra.Command{
		Use:   "launchdarkly",
		Short: "Import current state to Terraform configuration from LaunchDarkly",
		Long:  "Import current state to Terraform configuration from LaunchDarkly",
		RunE: func(cmd *cobra.Command, args []string) error {

			provider := newLaunchDarklyProvider()
			err := Import(provider, options, []string{apiKey, projectKey, environmentKey})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newLaunchDarklyProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "", "")

	cmd.PersistentFlags().StringVarP(&apiKey, "api-key", "", "", "Your LaunchDarkly REST API key")
	cmd.PersistentFlags().StringVarP(&projectKey, "proj-key", "", "", "The LaunchDarkly Project key")
	cmd.PersistentFlags().StringVarP(&environmentKey, "env-key", "", "", "The LaunchDarkly Environment key")

	return cmd
}

func newLaunchDarklyProvider() terraformutils.ProviderGenerator {
	return &launchdarkly_terraformer.LaunchDarklyProvider{}
}
