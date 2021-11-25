// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "ovh:index/ovh_cloud_project:ovh_cloud_project":
		r = &Ovh_cloud_project{}
	case "ovh:index/ovh_cloud_project_containerregistry:ovh_cloud_project_containerregistry":
		r = &Ovh_cloud_project_containerregistry{}
	case "ovh:index/ovh_cloud_project_containerregistry_user:ovh_cloud_project_containerregistry_user":
		r = &Ovh_cloud_project_containerregistry_user{}
	case "ovh:index/ovh_cloud_project_kube:ovh_cloud_project_kube":
		r = &Ovh_cloud_project_kube{}
	case "ovh:index/ovh_cloud_project_kube_nodepool:ovh_cloud_project_kube_nodepool":
		r = &Ovh_cloud_project_kube_nodepool{}
	case "ovh:index/ovh_cloud_project_network_private:ovh_cloud_project_network_private":
		r = &Ovh_cloud_project_network_private{}
	case "ovh:index/ovh_cloud_project_network_private_subnet:ovh_cloud_project_network_private_subnet":
		r = &Ovh_cloud_project_network_private_subnet{}
	case "ovh:index/ovh_cloud_project_user:ovh_cloud_project_user":
		r = &Ovh_cloud_project_user{}
	case "ovh:index/ovh_dbaas_logs_input:ovh_dbaas_logs_input":
		r = &Ovh_dbaas_logs_input{}
	case "ovh:index/ovh_dbaas_logs_output_graylog_stream:ovh_dbaas_logs_output_graylog_stream":
		r = &Ovh_dbaas_logs_output_graylog_stream{}
	case "ovh:index/ovh_dedicated_ceph_acl:ovh_dedicated_ceph_acl":
		r = &Ovh_dedicated_ceph_acl{}
	case "ovh:index/ovh_dedicated_server_install_task:ovh_dedicated_server_install_task":
		r = &Ovh_dedicated_server_install_task{}
	case "ovh:index/ovh_dedicated_server_reboot_task:ovh_dedicated_server_reboot_task":
		r = &Ovh_dedicated_server_reboot_task{}
	case "ovh:index/ovh_dedicated_server_update:ovh_dedicated_server_update":
		r = &Ovh_dedicated_server_update{}
	case "ovh:index/ovh_domain_zone:ovh_domain_zone":
		r = &Ovh_domain_zone{}
	case "ovh:index/ovh_domain_zone_record:ovh_domain_zone_record":
		r = &Ovh_domain_zone_record{}
	case "ovh:index/ovh_domain_zone_redirection:ovh_domain_zone_redirection":
		r = &Ovh_domain_zone_redirection{}
	case "ovh:index/ovh_ip_reverse:ovh_ip_reverse":
		r = &Ovh_ip_reverse{}
	case "ovh:index/ovh_ip_service:ovh_ip_service":
		r = &Ovh_ip_service{}
	case "ovh:index/ovh_iploadbalancing:ovh_iploadbalancing":
		r = &Ovh_iploadbalancing{}
	case "ovh:index/ovh_iploadbalancing_http_farm:ovh_iploadbalancing_http_farm":
		r = &Ovh_iploadbalancing_http_farm{}
	case "ovh:index/ovh_iploadbalancing_http_farm_server:ovh_iploadbalancing_http_farm_server":
		r = &Ovh_iploadbalancing_http_farm_server{}
	case "ovh:index/ovh_iploadbalancing_http_frontend:ovh_iploadbalancing_http_frontend":
		r = &Ovh_iploadbalancing_http_frontend{}
	case "ovh:index/ovh_iploadbalancing_http_route:ovh_iploadbalancing_http_route":
		r = &Ovh_iploadbalancing_http_route{}
	case "ovh:index/ovh_iploadbalancing_http_route_rule:ovh_iploadbalancing_http_route_rule":
		r = &Ovh_iploadbalancing_http_route_rule{}
	case "ovh:index/ovh_iploadbalancing_refresh:ovh_iploadbalancing_refresh":
		r = &Ovh_iploadbalancing_refresh{}
	case "ovh:index/ovh_iploadbalancing_tcp_farm:ovh_iploadbalancing_tcp_farm":
		r = &Ovh_iploadbalancing_tcp_farm{}
	case "ovh:index/ovh_iploadbalancing_tcp_farm_server:ovh_iploadbalancing_tcp_farm_server":
		r = &Ovh_iploadbalancing_tcp_farm_server{}
	case "ovh:index/ovh_iploadbalancing_tcp_frontend:ovh_iploadbalancing_tcp_frontend":
		r = &Ovh_iploadbalancing_tcp_frontend{}
	case "ovh:index/ovh_iploadbalancing_tcp_route:ovh_iploadbalancing_tcp_route":
		r = &Ovh_iploadbalancing_tcp_route{}
	case "ovh:index/ovh_iploadbalancing_tcp_route_rule:ovh_iploadbalancing_tcp_route_rule":
		r = &Ovh_iploadbalancing_tcp_route_rule{}
	case "ovh:index/ovh_iploadbalancing_vrack_network:ovh_iploadbalancing_vrack_network":
		r = &Ovh_iploadbalancing_vrack_network{}
	case "ovh:index/ovh_me_identity_user:ovh_me_identity_user":
		r = &Ovh_me_identity_user{}
	case "ovh:index/ovh_me_installation_template:ovh_me_installation_template":
		r = &Ovh_me_installation_template{}
	case "ovh:index/ovh_me_installation_template_partition_scheme:ovh_me_installation_template_partition_scheme":
		r = &Ovh_me_installation_template_partition_scheme{}
	case "ovh:index/ovh_me_installation_template_partition_scheme_hardware_raid:ovh_me_installation_template_partition_scheme_hardware_raid":
		r = &Ovh_me_installation_template_partition_scheme_hardware_raid{}
	case "ovh:index/ovh_me_installation_template_partition_scheme_partition:ovh_me_installation_template_partition_scheme_partition":
		r = &Ovh_me_installation_template_partition_scheme_partition{}
	case "ovh:index/ovh_me_ipxe_script:ovh_me_ipxe_script":
		r = &Ovh_me_ipxe_script{}
	case "ovh:index/ovh_me_ssh_key:ovh_me_ssh_key":
		r = &Ovh_me_ssh_key{}
	case "ovh:index/ovh_vrack:ovh_vrack":
		r = &Ovh_vrack{}
	case "ovh:index/ovh_vrack_cloudproject:ovh_vrack_cloudproject":
		r = &Ovh_vrack_cloudproject{}
	case "ovh:index/ovh_vrack_dedicated_server:ovh_vrack_dedicated_server":
		r = &Ovh_vrack_dedicated_server{}
	case "ovh:index/ovh_vrack_dedicated_server_interface:ovh_vrack_dedicated_server_interface":
		r = &Ovh_vrack_dedicated_server_interface{}
	case "ovh:index/ovh_vrack_ip:ovh_vrack_ip":
		r = &Ovh_vrack_ip{}
	case "ovh:index/ovh_vrack_iploadbalancing:ovh_vrack_iploadbalancing":
		r = &Ovh_vrack_iploadbalancing{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:ovh" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_cloud_project",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_cloud_project_containerregistry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_cloud_project_containerregistry_user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_cloud_project_kube",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_cloud_project_kube_nodepool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_cloud_project_network_private",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_cloud_project_network_private_subnet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_cloud_project_user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_dbaas_logs_input",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_dbaas_logs_output_graylog_stream",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_dedicated_ceph_acl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_dedicated_server_install_task",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_dedicated_server_reboot_task",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_dedicated_server_update",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_domain_zone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_domain_zone_record",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_domain_zone_redirection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_ip_reverse",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_ip_service",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_http_farm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_http_farm_server",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_http_frontend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_http_route",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_http_route_rule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_refresh",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_tcp_farm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_tcp_farm_server",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_tcp_frontend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_tcp_route",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_tcp_route_rule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_iploadbalancing_vrack_network",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_me_identity_user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_me_installation_template",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_me_installation_template_partition_scheme",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_me_installation_template_partition_scheme_hardware_raid",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_me_installation_template_partition_scheme_partition",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_me_ipxe_script",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_me_ssh_key",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_vrack",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_vrack_cloudproject",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_vrack_dedicated_server",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_vrack_dedicated_server_interface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_vrack_ip",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"index/ovh_vrack_iploadbalancing",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"ovh",
		&pkg{version},
	)
}