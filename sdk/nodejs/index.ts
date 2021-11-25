// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./getOvhCloud_Project_Capabilities_Containerregistry";
export * from "./getOvh_Cloud_Project_Capabilities_Containerregistry_Filter";
export * from "./getOvh_Cloud_Project_Containerregistries";
export * from "./getOvh_Cloud_Project_Containerregistry";
export * from "./getOvh_Cloud_Project_Containerregistry_Users";
export * from "./getOvh_Cloud_Project_Kube";
export * from "./getOvh_Cloud_Project_Region";
export * from "./getOvh_Cloud_Project_Regions";
export * from "./getOvh_Dbaas_Logs_Input_Engine";
export * from "./getOvh_Dbaas_Logs_Output_Graylog_Stream";
export * from "./getOvh_Dedicated_Ceph";
export * from "./getOvh_Dedicated_Installation_Templates";
export * from "./getOvh_Dedicated_Server";
export * from "./getOvh_Dedicated_Server_Boots";
export * from "./getOvh_Dedicated_Servers";
export * from "./getOvh_Domain_Zone";
export * from "./getOvh_Ip_Service";
export * from "./getOvh_Iploadbalancing";
export * from "./getOvh_Iploadbalancing_Vrack_Network";
export * from "./getOvh_Iploadbalancing_Vrack_Networks";
export * from "./getOvh_Me_Identity_User";
export * from "./getOvh_Me_Identity_Users";
export * from "./getOvh_Me_Installation_Template";
export * from "./getOvh_Me_Installation_Templates";
export * from "./getOvh_Me_Ipxe_Script";
export * from "./getOvh_Me_Ipxe_Scripts";
export * from "./getOvh_Me_Paymentmean_Bankaccount";
export * from "./getOvh_Me_Paymentmean_Creditcard";
export * from "./getOvh_Me_Ssh_Key";
export * from "./getOvh_Me_Ssh_Keys";
export * from "./getOvh_Order_Cart";
export * from "./getOvh_Order_Cart_Product";
export * from "./getOvh_Order_Cart_Product_Options";
export * from "./getOvh_Order_Cart_Product_Options_Plan";
export * from "./getOvh_Order_Cart_Product_Plan";
export * from "./getOvh_Vps";
export * from "./getOvh_Vracks";
export * from "./ovh_cloud_project";
export * from "./ovh_cloud_project_containerregistry";
export * from "./ovh_cloud_project_containerregistry_user";
export * from "./ovh_cloud_project_kube";
export * from "./ovh_cloud_project_kube_nodepool";
export * from "./ovh_cloud_project_network_private";
export * from "./ovh_cloud_project_network_private_subnet";
export * from "./ovh_cloud_project_user";
export * from "./ovh_dbaas_logs_input";
export * from "./ovh_dbaas_logs_output_graylog_stream";
export * from "./ovh_dedicated_ceph_acl";
export * from "./ovh_dedicated_server_install_task";
export * from "./ovh_dedicated_server_reboot_task";
export * from "./ovh_dedicated_server_update";
export * from "./ovh_domain_zone";
export * from "./ovh_domain_zone_record";
export * from "./ovh_domain_zone_redirection";
export * from "./ovh_ip_reverse";
export * from "./ovh_ip_service";
export * from "./ovh_iploadbalancing";
export * from "./ovh_iploadbalancing_http_farm";
export * from "./ovh_iploadbalancing_http_farm_server";
export * from "./ovh_iploadbalancing_http_frontend";
export * from "./ovh_iploadbalancing_http_route";
export * from "./ovh_iploadbalancing_http_route_rule";
export * from "./ovh_iploadbalancing_refresh";
export * from "./ovh_iploadbalancing_tcp_farm";
export * from "./ovh_iploadbalancing_tcp_farm_server";
export * from "./ovh_iploadbalancing_tcp_frontend";
export * from "./ovh_iploadbalancing_tcp_route";
export * from "./ovh_iploadbalancing_tcp_route_rule";
export * from "./ovh_iploadbalancing_vrack_network";
export * from "./ovh_me_identity_user";
export * from "./ovh_me_installation_template";
export * from "./ovh_me_installation_template_partition_scheme";
export * from "./ovh_me_installation_template_partition_scheme_hardware_raid";
export * from "./ovh_me_installation_template_partition_scheme_partition";
export * from "./ovh_me_ipxe_script";
export * from "./ovh_me_ssh_key";
export * from "./ovh_vrack";
export * from "./ovh_vrack_cloudproject";
export * from "./ovh_vrack_dedicated_server";
export * from "./ovh_vrack_dedicated_server_interface";
export * from "./ovh_vrack_ip";
export * from "./ovh_vrack_iploadbalancing";
export * from "./provider";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { Ovh_cloud_project } from "./ovh_cloud_project";
import { Ovh_cloud_project_containerregistry } from "./ovh_cloud_project_containerregistry";
import { Ovh_cloud_project_containerregistry_user } from "./ovh_cloud_project_containerregistry_user";
import { Ovh_cloud_project_kube } from "./ovh_cloud_project_kube";
import { Ovh_cloud_project_kube_nodepool } from "./ovh_cloud_project_kube_nodepool";
import { Ovh_cloud_project_network_private } from "./ovh_cloud_project_network_private";
import { Ovh_cloud_project_network_private_subnet } from "./ovh_cloud_project_network_private_subnet";
import { Ovh_cloud_project_user } from "./ovh_cloud_project_user";
import { Ovh_dbaas_logs_input } from "./ovh_dbaas_logs_input";
import { Ovh_dbaas_logs_output_graylog_stream } from "./ovh_dbaas_logs_output_graylog_stream";
import { Ovh_dedicated_ceph_acl } from "./ovh_dedicated_ceph_acl";
import { Ovh_dedicated_server_install_task } from "./ovh_dedicated_server_install_task";
import { Ovh_dedicated_server_reboot_task } from "./ovh_dedicated_server_reboot_task";
import { Ovh_dedicated_server_update } from "./ovh_dedicated_server_update";
import { Ovh_domain_zone } from "./ovh_domain_zone";
import { Ovh_domain_zone_record } from "./ovh_domain_zone_record";
import { Ovh_domain_zone_redirection } from "./ovh_domain_zone_redirection";
import { Ovh_ip_reverse } from "./ovh_ip_reverse";
import { Ovh_ip_service } from "./ovh_ip_service";
import { Ovh_iploadbalancing } from "./ovh_iploadbalancing";
import { Ovh_iploadbalancing_http_farm } from "./ovh_iploadbalancing_http_farm";
import { Ovh_iploadbalancing_http_farm_server } from "./ovh_iploadbalancing_http_farm_server";
import { Ovh_iploadbalancing_http_frontend } from "./ovh_iploadbalancing_http_frontend";
import { Ovh_iploadbalancing_http_route } from "./ovh_iploadbalancing_http_route";
import { Ovh_iploadbalancing_http_route_rule } from "./ovh_iploadbalancing_http_route_rule";
import { Ovh_iploadbalancing_refresh } from "./ovh_iploadbalancing_refresh";
import { Ovh_iploadbalancing_tcp_farm } from "./ovh_iploadbalancing_tcp_farm";
import { Ovh_iploadbalancing_tcp_farm_server } from "./ovh_iploadbalancing_tcp_farm_server";
import { Ovh_iploadbalancing_tcp_frontend } from "./ovh_iploadbalancing_tcp_frontend";
import { Ovh_iploadbalancing_tcp_route } from "./ovh_iploadbalancing_tcp_route";
import { Ovh_iploadbalancing_tcp_route_rule } from "./ovh_iploadbalancing_tcp_route_rule";
import { Ovh_iploadbalancing_vrack_network } from "./ovh_iploadbalancing_vrack_network";
import { Ovh_me_identity_user } from "./ovh_me_identity_user";
import { Ovh_me_installation_template } from "./ovh_me_installation_template";
import { Ovh_me_installation_template_partition_scheme } from "./ovh_me_installation_template_partition_scheme";
import { Ovh_me_installation_template_partition_scheme_hardware_raid } from "./ovh_me_installation_template_partition_scheme_hardware_raid";
import { Ovh_me_installation_template_partition_scheme_partition } from "./ovh_me_installation_template_partition_scheme_partition";
import { Ovh_me_ipxe_script } from "./ovh_me_ipxe_script";
import { Ovh_me_ssh_key } from "./ovh_me_ssh_key";
import { Ovh_vrack } from "./ovh_vrack";
import { Ovh_vrack_cloudproject } from "./ovh_vrack_cloudproject";
import { Ovh_vrack_dedicated_server } from "./ovh_vrack_dedicated_server";
import { Ovh_vrack_dedicated_server_interface } from "./ovh_vrack_dedicated_server_interface";
import { Ovh_vrack_ip } from "./ovh_vrack_ip";
import { Ovh_vrack_iploadbalancing } from "./ovh_vrack_iploadbalancing";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ovh:index/ovh_cloud_project:ovh_cloud_project":
                return new Ovh_cloud_project(name, <any>undefined, { urn })
            case "ovh:index/ovh_cloud_project_containerregistry:ovh_cloud_project_containerregistry":
                return new Ovh_cloud_project_containerregistry(name, <any>undefined, { urn })
            case "ovh:index/ovh_cloud_project_containerregistry_user:ovh_cloud_project_containerregistry_user":
                return new Ovh_cloud_project_containerregistry_user(name, <any>undefined, { urn })
            case "ovh:index/ovh_cloud_project_kube:ovh_cloud_project_kube":
                return new Ovh_cloud_project_kube(name, <any>undefined, { urn })
            case "ovh:index/ovh_cloud_project_kube_nodepool:ovh_cloud_project_kube_nodepool":
                return new Ovh_cloud_project_kube_nodepool(name, <any>undefined, { urn })
            case "ovh:index/ovh_cloud_project_network_private:ovh_cloud_project_network_private":
                return new Ovh_cloud_project_network_private(name, <any>undefined, { urn })
            case "ovh:index/ovh_cloud_project_network_private_subnet:ovh_cloud_project_network_private_subnet":
                return new Ovh_cloud_project_network_private_subnet(name, <any>undefined, { urn })
            case "ovh:index/ovh_cloud_project_user:ovh_cloud_project_user":
                return new Ovh_cloud_project_user(name, <any>undefined, { urn })
            case "ovh:index/ovh_dbaas_logs_input:ovh_dbaas_logs_input":
                return new Ovh_dbaas_logs_input(name, <any>undefined, { urn })
            case "ovh:index/ovh_dbaas_logs_output_graylog_stream:ovh_dbaas_logs_output_graylog_stream":
                return new Ovh_dbaas_logs_output_graylog_stream(name, <any>undefined, { urn })
            case "ovh:index/ovh_dedicated_ceph_acl:ovh_dedicated_ceph_acl":
                return new Ovh_dedicated_ceph_acl(name, <any>undefined, { urn })
            case "ovh:index/ovh_dedicated_server_install_task:ovh_dedicated_server_install_task":
                return new Ovh_dedicated_server_install_task(name, <any>undefined, { urn })
            case "ovh:index/ovh_dedicated_server_reboot_task:ovh_dedicated_server_reboot_task":
                return new Ovh_dedicated_server_reboot_task(name, <any>undefined, { urn })
            case "ovh:index/ovh_dedicated_server_update:ovh_dedicated_server_update":
                return new Ovh_dedicated_server_update(name, <any>undefined, { urn })
            case "ovh:index/ovh_domain_zone:ovh_domain_zone":
                return new Ovh_domain_zone(name, <any>undefined, { urn })
            case "ovh:index/ovh_domain_zone_record:ovh_domain_zone_record":
                return new Ovh_domain_zone_record(name, <any>undefined, { urn })
            case "ovh:index/ovh_domain_zone_redirection:ovh_domain_zone_redirection":
                return new Ovh_domain_zone_redirection(name, <any>undefined, { urn })
            case "ovh:index/ovh_ip_reverse:ovh_ip_reverse":
                return new Ovh_ip_reverse(name, <any>undefined, { urn })
            case "ovh:index/ovh_ip_service:ovh_ip_service":
                return new Ovh_ip_service(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing:ovh_iploadbalancing":
                return new Ovh_iploadbalancing(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_http_farm:ovh_iploadbalancing_http_farm":
                return new Ovh_iploadbalancing_http_farm(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_http_farm_server:ovh_iploadbalancing_http_farm_server":
                return new Ovh_iploadbalancing_http_farm_server(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_http_frontend:ovh_iploadbalancing_http_frontend":
                return new Ovh_iploadbalancing_http_frontend(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_http_route:ovh_iploadbalancing_http_route":
                return new Ovh_iploadbalancing_http_route(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_http_route_rule:ovh_iploadbalancing_http_route_rule":
                return new Ovh_iploadbalancing_http_route_rule(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_refresh:ovh_iploadbalancing_refresh":
                return new Ovh_iploadbalancing_refresh(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_tcp_farm:ovh_iploadbalancing_tcp_farm":
                return new Ovh_iploadbalancing_tcp_farm(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_tcp_farm_server:ovh_iploadbalancing_tcp_farm_server":
                return new Ovh_iploadbalancing_tcp_farm_server(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_tcp_frontend:ovh_iploadbalancing_tcp_frontend":
                return new Ovh_iploadbalancing_tcp_frontend(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_tcp_route:ovh_iploadbalancing_tcp_route":
                return new Ovh_iploadbalancing_tcp_route(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_tcp_route_rule:ovh_iploadbalancing_tcp_route_rule":
                return new Ovh_iploadbalancing_tcp_route_rule(name, <any>undefined, { urn })
            case "ovh:index/ovh_iploadbalancing_vrack_network:ovh_iploadbalancing_vrack_network":
                return new Ovh_iploadbalancing_vrack_network(name, <any>undefined, { urn })
            case "ovh:index/ovh_me_identity_user:ovh_me_identity_user":
                return new Ovh_me_identity_user(name, <any>undefined, { urn })
            case "ovh:index/ovh_me_installation_template:ovh_me_installation_template":
                return new Ovh_me_installation_template(name, <any>undefined, { urn })
            case "ovh:index/ovh_me_installation_template_partition_scheme:ovh_me_installation_template_partition_scheme":
                return new Ovh_me_installation_template_partition_scheme(name, <any>undefined, { urn })
            case "ovh:index/ovh_me_installation_template_partition_scheme_hardware_raid:ovh_me_installation_template_partition_scheme_hardware_raid":
                return new Ovh_me_installation_template_partition_scheme_hardware_raid(name, <any>undefined, { urn })
            case "ovh:index/ovh_me_installation_template_partition_scheme_partition:ovh_me_installation_template_partition_scheme_partition":
                return new Ovh_me_installation_template_partition_scheme_partition(name, <any>undefined, { urn })
            case "ovh:index/ovh_me_ipxe_script:ovh_me_ipxe_script":
                return new Ovh_me_ipxe_script(name, <any>undefined, { urn })
            case "ovh:index/ovh_me_ssh_key:ovh_me_ssh_key":
                return new Ovh_me_ssh_key(name, <any>undefined, { urn })
            case "ovh:index/ovh_vrack:ovh_vrack":
                return new Ovh_vrack(name, <any>undefined, { urn })
            case "ovh:index/ovh_vrack_cloudproject:ovh_vrack_cloudproject":
                return new Ovh_vrack_cloudproject(name, <any>undefined, { urn })
            case "ovh:index/ovh_vrack_dedicated_server:ovh_vrack_dedicated_server":
                return new Ovh_vrack_dedicated_server(name, <any>undefined, { urn })
            case "ovh:index/ovh_vrack_dedicated_server_interface:ovh_vrack_dedicated_server_interface":
                return new Ovh_vrack_dedicated_server_interface(name, <any>undefined, { urn })
            case "ovh:index/ovh_vrack_ip:ovh_vrack_ip":
                return new Ovh_vrack_ip(name, <any>undefined, { urn })
            case "ovh:index/ovh_vrack_iploadbalancing:ovh_vrack_iploadbalancing":
                return new Ovh_vrack_iploadbalancing(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ovh", "index/ovh_cloud_project", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_cloud_project_containerregistry", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_cloud_project_containerregistry_user", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_cloud_project_kube", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_cloud_project_kube_nodepool", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_cloud_project_network_private", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_cloud_project_network_private_subnet", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_cloud_project_user", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_dbaas_logs_input", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_dbaas_logs_output_graylog_stream", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_dedicated_ceph_acl", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_dedicated_server_install_task", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_dedicated_server_reboot_task", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_dedicated_server_update", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_domain_zone", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_domain_zone_record", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_domain_zone_redirection", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_ip_reverse", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_ip_service", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_http_farm", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_http_farm_server", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_http_frontend", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_http_route", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_http_route_rule", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_refresh", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_tcp_farm", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_tcp_farm_server", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_tcp_frontend", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_tcp_route", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_tcp_route_rule", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_iploadbalancing_vrack_network", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_me_identity_user", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_me_installation_template", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_me_installation_template_partition_scheme", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_me_installation_template_partition_scheme_hardware_raid", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_me_installation_template_partition_scheme_partition", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_me_ipxe_script", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_me_ssh_key", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_vrack", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_vrack_cloudproject", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_vrack_dedicated_server", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_vrack_dedicated_server_interface", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_vrack_ip", _module)
pulumi.runtime.registerResourceModule("ovh", "index/ovh_vrack_iploadbalancing", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("ovh", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:ovh") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
