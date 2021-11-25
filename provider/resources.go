// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ovh

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/legitbee/pulumi-ovh/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/ovh/terraform-provider-ovh/ovh"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "ovh"
	// modules:
	mainMod = "index" // the ovh module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// stringValue gets a string value from a property map if present, else ""
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	return ""
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(ovh.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "ovh",
		Description: "A Pulumi package for creating and managing ovh cloud resources.",
		Keywords:    []string{"pulumi", "ovh"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/legitbee/pulumi-ovh",
		Config:      map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: makeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },

		},
		PreConfigureCallback: preConfigureCallback,
		Resources:            map[string]*tfbridge.ResourceInfo{
			    "ovh_cloud_project": {Tok: makeResource(mainMod, "ovh_cloud_project"), Docs: &tfbridge.DocInfo{Source:"cloud_project.html.markdown"}},
			    "ovh_cloud_project_containerregistry": {Tok: makeResource(mainMod, "ovh_cloud_project_containerregistry")},
			    "ovh_cloud_project_containerregistry_user": {Tok: makeResource(mainMod, "ovh_cloud_project_containerregistry_user")},
			    "ovh_cloud_project_kube": {Tok: makeResource(mainMod, "ovh_cloud_project_kube")},
			    "ovh_cloud_project_kube_nodepool": {Tok: makeResource(mainMod, "ovh_cloud_project_kube_nodepool")},
			    "ovh_cloud_project_network_private": {Tok: makeResource(mainMod, "ovh_cloud_project_network_private")},
			    "ovh_cloud_project_network_private_subnet": {Tok: makeResource(mainMod, "ovh_cloud_project_network_private_subnet")},
			    "ovh_cloud_project_user": {Tok: makeResource(mainMod, "ovh_cloud_project_user")},
			    "ovh_dbaas_logs_input": {Tok: makeResource(mainMod, "ovh_dbaas_logs_input")},
			    "ovh_dbaas_logs_output_graylog_stream": {Tok: makeResource(mainMod, "ovh_dbaas_logs_output_graylog_stream")},
			    "ovh_dedicated_ceph_acl": {Tok: makeResource(mainMod, "ovh_dedicated_ceph_acl")},
			    "ovh_dedicated_server_install_task": {Tok: makeResource(mainMod, "ovh_dedicated_server_install_task")},
			    "ovh_dedicated_server_reboot_task": {Tok: makeResource(mainMod, "ovh_dedicated_server_reboot_task")},
			    "ovh_dedicated_server_update": {Tok: makeResource(mainMod, "ovh_dedicated_server_update")},
			    "ovh_domain_zone": {Tok: makeResource(mainMod, "ovh_domain_zone")},
			    "ovh_domain_zone_record": {Tok: makeResource(mainMod, "ovh_domain_zone_record")},
			    "ovh_domain_zone_redirection": {Tok: makeResource(mainMod, "ovh_domain_zone_redirection")},
			    "ovh_ip_reverse": {Tok: makeResource(mainMod, "ovh_ip_reverse")},
			    "ovh_ip_service": {Tok: makeResource(mainMod, "ovh_ip_service")},
			    "ovh_iploadbalancing": {Tok: makeResource(mainMod, "ovh_iploadbalancing")},
			    "ovh_iploadbalancing_http_farm": {Tok: makeResource(mainMod, "ovh_iploadbalancing_http_farm")},
			    "ovh_iploadbalancing_http_farm_server": {Tok: makeResource(mainMod, "ovh_iploadbalancing_http_farm_server")},
			    "ovh_iploadbalancing_http_frontend": {Tok: makeResource(mainMod, "ovh_iploadbalancing_http_frontend")},
			    "ovh_iploadbalancing_http_route": {Tok: makeResource(mainMod, "ovh_iploadbalancing_http_route")},
			    "ovh_iploadbalancing_http_route_rule": {Tok: makeResource(mainMod, "ovh_iploadbalancing_http_route_rule")},
			    "ovh_iploadbalancing_refresh": {Tok: makeResource(mainMod, "ovh_iploadbalancing_refresh")},
			    "ovh_iploadbalancing_tcp_farm": {Tok: makeResource(mainMod, "ovh_iploadbalancing_tcp_farm")},
			    "ovh_iploadbalancing_tcp_farm_server": {Tok: makeResource(mainMod, "ovh_iploadbalancing_tcp_farm_server")},
			    "ovh_iploadbalancing_tcp_frontend": {Tok: makeResource(mainMod, "ovh_iploadbalancing_tcp_frontend")},
			    "ovh_iploadbalancing_tcp_route": {Tok: makeResource(mainMod, "ovh_iploadbalancing_tcp_route")},
			    "ovh_iploadbalancing_tcp_route_rule": {Tok: makeResource(mainMod, "ovh_iploadbalancing_tcp_route_rule")},
			    "ovh_iploadbalancing_vrack_network": {Tok: makeResource(mainMod, "ovh_iploadbalancing_vrack_network")},
			    "ovh_me_identity_user": {Tok: makeResource(mainMod, "ovh_me_identity_user")},
			    "ovh_me_installation_template": {Tok: makeResource(mainMod, "ovh_me_installation_template")},
			    "ovh_me_installation_template_partition_scheme": {Tok: makeResource(mainMod, "ovh_me_installation_template_partition_scheme")},
			    "ovh_me_installation_template_partition_scheme_hardware_raid": {Tok: makeResource(mainMod, "ovh_me_installation_template_partition_scheme_hardware_raid")},
			    "ovh_me_installation_template_partition_scheme_partition": {Tok: makeResource(mainMod, "ovh_me_installation_template_partition_scheme_partition")},
			    "ovh_me_ipxe_script": {Tok: makeResource(mainMod, "ovh_me_ipxe_script")},
			    "ovh_me_ssh_key": {Tok: makeResource(mainMod, "ovh_me_ssh_key")},
			    "ovh_vrack": {Tok: makeResource(mainMod, "ovh_vrack")},
			    "ovh_vrack_cloudproject": {Tok: makeResource(mainMod, "ovh_vrack_cloudproject")},
			    "ovh_vrack_dedicated_server": {Tok: makeResource(mainMod, "ovh_vrack_dedicated_server")},
			    "ovh_vrack_dedicated_server_interface": {Tok: makeResource(mainMod, "ovh_vrack_dedicated_server_interface")},
			    "ovh_vrack_ip": {Tok: makeResource(mainMod, "ovh_vrack_ip")},
			    "ovh_vrack_iploadbalancing": {Tok: makeResource(mainMod, "ovh_vrack_iploadbalancing")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"ovh_cloud_project_capabilities_containerregistry": {Tok: makeDataSource(mainMod, "getOvhCloud_Project_Capabilities_Containerregistry")},
			"ovh_cloud_project_capabilities_containerregistry_filter": {Tok: makeDataSource(mainMod, "getOvh_Cloud_Project_Capabilities_Containerregistry_Filter")},
			"ovh_cloud_project_containerregistries": {Tok: makeDataSource(mainMod, "getOvh_Cloud_Project_Containerregistries")},
			"ovh_cloud_project_containerregistry": {Tok: makeDataSource(mainMod, "getOvh_Cloud_Project_Containerregistry")},
			"ovh_cloud_project_containerregistry_users": {Tok: makeDataSource(mainMod, "getOvh_Cloud_Project_Containerregistry_Users")},
			"ovh_cloud_project_kube": {Tok: makeDataSource(mainMod, "getOvh_Cloud_Project_Kube")},
			"ovh_cloud_project_region": {Tok: makeDataSource(mainMod, "getOvh_Cloud_Project_Region")},
			"ovh_cloud_project_regions": {Tok: makeDataSource(mainMod, "getOvh_Cloud_Project_Regions")},
			"ovh_dbaas_logs_input_engine": {Tok: makeDataSource(mainMod, "getOvh_Dbaas_Logs_Input_Engine")},
			"ovh_dbaas_logs_output_graylog_stream": {Tok: makeDataSource(mainMod, "getOvh_Dbaas_Logs_Output_Graylog_Stream")},
			"ovh_dedicated_ceph": {Tok: makeDataSource(mainMod, "getOvh_Dedicated_Ceph")},
			"ovh_dedicated_installation_templates": {Tok: makeDataSource(mainMod, "getOvh_Dedicated_Installation_Templates")},
			"ovh_dedicated_server": {Tok: makeDataSource(mainMod, "getOvh_Dedicated_Server")},
			"ovh_dedicated_server_boots": {Tok: makeDataSource(mainMod, "getOvh_Dedicated_Server_Boots")},
			"ovh_dedicated_servers": {Tok: makeDataSource(mainMod, "getOvh_Dedicated_Servers")},
			"ovh_domain_zone": {Tok: makeDataSource(mainMod, "getOvh_Domain_Zone")},
			"ovh_ip_service": {Tok: makeDataSource(mainMod, "getOvh_Ip_Service")},
			"ovh_iploadbalancing": {Tok: makeDataSource(mainMod, "getOvh_Iploadbalancing")},
			"ovh_iploadbalancing_vrack_network": {Tok: makeDataSource(mainMod, "getOvh_Iploadbalancing_Vrack_Network")},
			"ovh_iploadbalancing_vrack_networks": {Tok: makeDataSource(mainMod, "getOvh_Iploadbalancing_Vrack_Networks")},
			"ovh_me_identity_user": {Tok: makeDataSource(mainMod, "getOvh_Me_Identity_User")},
			"ovh_me_identity_users": {Tok: makeDataSource(mainMod, "getOvh_Me_Identity_Users")},
			"ovh_me_installation_template": {Tok: makeDataSource(mainMod, "getOvh_Me_Installation_Template")},
			"ovh_me_installation_templates": {Tok: makeDataSource(mainMod, "getOvh_Me_Installation_Templates")},
			"ovh_me_ipxe_script": {Tok: makeDataSource(mainMod, "getOvh_Me_Ipxe_Script")},
			"ovh_me_ipxe_scripts": {Tok: makeDataSource(mainMod, "getOvh_Me_Ipxe_Scripts")},
			"ovh_me_paymentmean_bankaccount": {Tok: makeDataSource(mainMod, "getOvh_Me_Paymentmean_Bankaccount")},
			"ovh_me_paymentmean_creditcard": {Tok: makeDataSource(mainMod, "getOvh_Me_Paymentmean_Creditcard")},
			"ovh_me_ssh_key": {Tok: makeDataSource(mainMod, "getOvh_Me_Ssh_Key")},
			"ovh_me_ssh_keys": {Tok: makeDataSource(mainMod, "getOvh_Me_Ssh_Keys")},
			"ovh_order_cart": {Tok: makeDataSource(mainMod, "getOvh_Order_Cart")},
			"ovh_order_cart_product": {Tok: makeDataSource(mainMod, "getOvh_Order_Cart_Product")},
			"ovh_order_cart_product_options": {Tok: makeDataSource(mainMod, "getOvh_Order_Cart_Product_Options")},
			"ovh_order_cart_product_options_plan": {Tok: makeDataSource(mainMod, "getOvh_Order_Cart_Product_Options_Plan")},
			"ovh_order_cart_product_plan": {Tok: makeDataSource(mainMod, "getOvh_Order_Cart_Product_Plan")},
			"ovh_vps": {Tok: makeDataSource(mainMod, "getOvh_Vps")},
			"ovh_vracks": {Tok: makeDataSource(mainMod, "getOvh_Vracks")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/legitbee/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
