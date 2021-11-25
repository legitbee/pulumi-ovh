module github.com/legitbee/pulumi-ovh/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
)

require (
	github.com/hashicorp/hcl/v2 v2.3.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/ovh/terraform-provider-ovh v0.16.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.9.0
	github.com/pulumi/pulumi/sdk/v3 v3.14.1-0.20211007222624-789e39219452
	github.com/zclconf/go-cty v1.8.4 // indirect
)
