package launchdarkly

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type LaunchDarklyProvider struct { //nolint
	terraformutils.Provider
	apiKey         string
	projectKey     string
	environmentKey string
}

func (p *LaunchDarklyProvider) Init(args []string) error {
	p.apiKey = args[0]
	p.projectKey = args[1]
	p.environmentKey = args[2]

	return nil
}

func (p *LaunchDarklyProvider) GetName() string {
	return "launchdarkly"
}

func (p *LaunchDarklyProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (LaunchDarklyProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *LaunchDarklyProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {

	return map[string]terraformutils.ServiceGenerator{
		"project":     &ProjectGenerator{},
		"flags":       &FlagsGenerator{},
		"environment": &FlagEnvironmentGenerator{},
	}
}

func (p *LaunchDarklyProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("launchdarkly: " + serviceName + " not supported service")
	}

	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())

	p.Service.SetArgs(map[string]interface{}{
		"apiKey":         p.apiKey,
		"projectKey":     p.projectKey,
		"environmentKey": p.environmentKey,
	})

	return nil
}

func (p *LaunchDarklyProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"access_token": cty.StringVal(p.apiKey),
	})
}
func (p *LaunchDarklyProvider) GetBasicConfig() cty.Value {
	return p.GetConfig()
}
