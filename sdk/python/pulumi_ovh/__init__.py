# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .cloud_project import *
from .cloud_project_containerregistry import *
from .cloud_project_containerregistry_user import *
from .cloud_project_kube import *
from .cloud_project_kube_nodepool import *
from .cloud_project_network_private import *
from .cloud_project_network_private_subnet import *
from .cloud_project_user import *
from .dbaas_logs_input import *
from .dbaas_logs_output_graylog_stream import *
from .dedicated_ceph_acl import *
from .dedicated_server_install_task import *
from .dedicated_server_reboot_task import *
from .dedicated_server_update import *
from .domain_zone import *
from .domain_zone_record import *
from .domain_zone_redirection import *
from .get_cloud_project_capabilities_containerregistry import *
from .get_cloud_project_capabilities_containerregistry_filter import *
from .get_cloud_project_containerregistries import *
from .get_cloud_project_containerregistry import *
from .get_cloud_project_containerregistry_users import *
from .get_cloud_project_kube import *
from .get_cloud_project_region import *
from .get_cloud_project_regions import *
from .get_dbaas_logs_input_engine import *
from .get_dbaas_logs_output_graylog_stream import *
from .get_dedicated_ceph import *
from .get_dedicated_installation_templates import *
from .get_dedicated_server import *
from .get_dedicated_server_boots import *
from .get_dedicated_servers import *
from .get_domain_zone import *
from .get_ip_service import *
from .get_iploadbalancing import *
from .get_iploadbalancing_vrack_network import *
from .get_iploadbalancing_vrack_networks import *
from .get_me_identity_user import *
from .get_me_identity_users import *
from .get_me_installation_template import *
from .get_me_installation_templates import *
from .get_me_ipxe_script import *
from .get_me_ipxe_scripts import *
from .get_me_paymentmean_bankaccount import *
from .get_me_paymentmean_creditcard import *
from .get_me_ssh_key import *
from .get_me_ssh_keys import *
from .get_order_cart import *
from .get_order_cart_product import *
from .get_order_cart_product_options import *
from .get_order_cart_product_options_plan import *
from .get_order_cart_product_plan import *
from .get_vps import *
from .get_vracks import *
from .ip_reverse import *
from .ip_service import *
from .iploadbalancing import *
from .iploadbalancing_http_farm import *
from .iploadbalancing_http_farm_server import *
from .iploadbalancing_http_frontend import *
from .iploadbalancing_http_route import *
from .iploadbalancing_http_route_rule import *
from .iploadbalancing_refresh import *
from .iploadbalancing_tcp_farm import *
from .iploadbalancing_tcp_farm_server import *
from .iploadbalancing_tcp_frontend import *
from .iploadbalancing_tcp_route import *
from .iploadbalancing_tcp_route_rule import *
from .iploadbalancing_vrack_network import *
from .me_identity_user import *
from .me_installation_template import *
from .me_installation_template_partition_scheme import *
from .me_installation_template_partition_scheme_hardware_raid import *
from .me_installation_template_partition_scheme_partition import *
from .me_ipxe_script import *
from .me_ssh_key import *
from .provider import *
from .vrack import *
from .vrack_cloudproject import *
from .vrack_dedicated_server import *
from .vrack_dedicated_server_interface import *
from .vrack_ip import *
from .vrack_iploadbalancing import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_ovh.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_ovh.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "ovh",
  "mod": "index/cloudProject",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/cloudProject:CloudProject": "CloudProject"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/cloudProjectContainerregistry",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/cloudProjectContainerregistry:CloudProjectContainerregistry": "CloudProjectContainerregistry"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/cloudProjectContainerregistryUser",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/cloudProjectContainerregistryUser:CloudProjectContainerregistryUser": "CloudProjectContainerregistryUser"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/cloudProjectKube",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/cloudProjectKube:CloudProjectKube": "CloudProjectKube"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/cloudProjectKubeNodepool",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/cloudProjectKubeNodepool:CloudProjectKubeNodepool": "CloudProjectKubeNodepool"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/cloudProjectNetworkPrivate",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/cloudProjectNetworkPrivate:CloudProjectNetworkPrivate": "CloudProjectNetworkPrivate"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/cloudProjectNetworkPrivateSubnet",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/cloudProjectNetworkPrivateSubnet:CloudProjectNetworkPrivateSubnet": "CloudProjectNetworkPrivateSubnet"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/cloudProjectUser",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/cloudProjectUser:CloudProjectUser": "CloudProjectUser"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/dbaasLogsInput",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/dbaasLogsInput:DbaasLogsInput": "DbaasLogsInput"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/dbaasLogsOutputGraylogStream",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/dbaasLogsOutputGraylogStream:DbaasLogsOutputGraylogStream": "DbaasLogsOutputGraylogStream"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/dedicatedCephAcl",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/dedicatedCephAcl:DedicatedCephAcl": "DedicatedCephAcl"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/dedicatedServerInstallTask",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/dedicatedServerInstallTask:DedicatedServerInstallTask": "DedicatedServerInstallTask"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/dedicatedServerRebootTask",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/dedicatedServerRebootTask:DedicatedServerRebootTask": "DedicatedServerRebootTask"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/dedicatedServerUpdate",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/dedicatedServerUpdate:DedicatedServerUpdate": "DedicatedServerUpdate"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/domainZone",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/domainZone:DomainZone": "DomainZone"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/domainZoneRecord",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/domainZoneRecord:DomainZoneRecord": "DomainZoneRecord"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/domainZoneRedirection",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/domainZoneRedirection:DomainZoneRedirection": "DomainZoneRedirection"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/ipReverse",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/ipReverse:IpReverse": "IpReverse"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/ipService",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/ipService:IpService": "IpService"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancing",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancing:Iploadbalancing": "Iploadbalancing"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingHttpFarm",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingHttpFarm:IploadbalancingHttpFarm": "IploadbalancingHttpFarm"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingHttpFarmServer",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingHttpFarmServer:IploadbalancingHttpFarmServer": "IploadbalancingHttpFarmServer"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingHttpFrontend",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingHttpFrontend:IploadbalancingHttpFrontend": "IploadbalancingHttpFrontend"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingHttpRoute",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingHttpRoute:IploadbalancingHttpRoute": "IploadbalancingHttpRoute"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingHttpRouteRule",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingHttpRouteRule:IploadbalancingHttpRouteRule": "IploadbalancingHttpRouteRule"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingRefresh",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingRefresh:IploadbalancingRefresh": "IploadbalancingRefresh"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingTcpFarm",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingTcpFarm:IploadbalancingTcpFarm": "IploadbalancingTcpFarm"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingTcpFarmServer",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingTcpFarmServer:IploadbalancingTcpFarmServer": "IploadbalancingTcpFarmServer"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingTcpFrontend",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingTcpFrontend:IploadbalancingTcpFrontend": "IploadbalancingTcpFrontend"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingTcpRoute",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingTcpRoute:IploadbalancingTcpRoute": "IploadbalancingTcpRoute"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingTcpRouteRule",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingTcpRouteRule:IploadbalancingTcpRouteRule": "IploadbalancingTcpRouteRule"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/iploadbalancingVrackNetwork",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/iploadbalancingVrackNetwork:IploadbalancingVrackNetwork": "IploadbalancingVrackNetwork"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/meIdentityUser",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/meIdentityUser:MeIdentityUser": "MeIdentityUser"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/meInstallationTemplate",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/meInstallationTemplate:MeInstallationTemplate": "MeInstallationTemplate"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/meInstallationTemplatePartitionScheme",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/meInstallationTemplatePartitionScheme:MeInstallationTemplatePartitionScheme": "MeInstallationTemplatePartitionScheme"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/meInstallationTemplatePartitionSchemeHardwareRaid",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/meInstallationTemplatePartitionSchemeHardwareRaid:MeInstallationTemplatePartitionSchemeHardwareRaid": "MeInstallationTemplatePartitionSchemeHardwareRaid"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/meInstallationTemplatePartitionSchemePartition",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/meInstallationTemplatePartitionSchemePartition:MeInstallationTemplatePartitionSchemePartition": "MeInstallationTemplatePartitionSchemePartition"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/meIpxeScript",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/meIpxeScript:MeIpxeScript": "MeIpxeScript"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/meSshKey",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/meSshKey:MeSshKey": "MeSshKey"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/vrack",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/vrack:Vrack": "Vrack"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/vrackCloudproject",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/vrackCloudproject:VrackCloudproject": "VrackCloudproject"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/vrackDedicatedServer",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/vrackDedicatedServer:VrackDedicatedServer": "VrackDedicatedServer"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/vrackDedicatedServerInterface",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/vrackDedicatedServerInterface:VrackDedicatedServerInterface": "VrackDedicatedServerInterface"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/vrackIp",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/vrackIp:VrackIp": "VrackIp"
  }
 },
 {
  "pkg": "ovh",
  "mod": "index/vrackIploadbalancing",
  "fqn": "pulumi_ovh",
  "classes": {
   "ovh:index/vrackIploadbalancing:VrackIploadbalancing": "VrackIploadbalancing"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "ovh",
  "token": "pulumi:providers:ovh",
  "fqn": "pulumi_ovh",
  "class": "Provider"
 }
]
"""
)
