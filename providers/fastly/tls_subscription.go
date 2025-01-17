// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fastly

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/fastly/go-fastly/v5/fastly"
)

type TLSSubscriptionGenerator struct {
	FastlyService
}

func (g *TLSSubscriptionGenerator) loadSubscriptions(client *fastly.Client) ([]*fastly.TLSSubscription, error) {
	subscriptions, err := client.ListTLSSubscriptions(&fastly.ListTLSSubscriptionsInput{})
	if err != nil {
		return nil, err
	}
	for _, subscription := range subscriptions {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			subscription.ID,
			subscription.ID,
			"fastly_tls_subscription",
			"fastly",
			[]string{}))
	}
	return subscriptions, nil
}

func (g *TLSSubscriptionGenerator) InitResources() error {
	client, err := fastly.NewClient(g.Args["api_key"].(string))
	if err != nil {
		return err
	}
	if _, err := g.loadSubscriptions(client); err != nil {
		return err
	}

	return nil
}
