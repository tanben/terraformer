### Use with LaunchDarkly

Usage:

```
./terraformer {plan | import} launchdarkly --api-key {REST API KEY}  --proj-key {PROJECT KEY} --env-key {ENVIRONMENT KEY}  -r {project| flags | environment}

```

Example:

Importing LaunchDarkly Project

```
./terraformer import launchdarkly --api-key "REST-AA-111"  --proj-key "sample-demo-proj" -r project
```

Importing LaunchDarkly Flags

```
./terraformer import launchdarkly --api-key "REST-AA-111"  --proj-key "sample-demo-proj"  --env-key  development -r flags
```


List of supported LaunchDarkly resources:

*   `launchdarkly_project`
*   `launchdarkly_feature_flag`
*   `launchdarkly_feature_flag_environment`
*   `launchdarkly_custom_role`
*   `launchdarkly_segment`

Docs

* [LaunchDarkly Provider Doc](https://registry.terraform.io/providers/launchdarkly/launchdarkly/latest/docs)

* [LaunchDarkly REST API](https://apidocs.launchdarkly.com/#section/Overview)

* [REST API Client Go Doc](https://pkg.go.dev/github.com/launchdarkly/api-client-go@v5.3.0+incompatible#section-readme)

Repo
* [REST API Client Repo](https://github.com/launchdarkly/api-client-go)

* [LaunchDarkly Provider Repo](https://github.com/launchdarkly/terraform-provider-launchdarkly)