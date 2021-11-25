# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetOvh_Cloud_Project_KubeResult',
    'AwaitableGetOvh_Cloud_Project_KubeResult',
    'get_ovh__cloud__project__kube',
    'get_ovh__cloud__project__kube_output',
]

@pulumi.output_type
class GetOvh_Cloud_Project_KubeResult:
    """
    A collection of values returned by getOvh_Cloud_Project_Kube.
    """
    def __init__(__self__, control_plane_is_up_to_date=None, id=None, is_up_to_date=None, kube_id=None, name=None, next_upgrade_versions=None, nodes_url=None, private_network_id=None, region=None, service_name=None, status=None, update_policy=None, url=None, version=None):
        if control_plane_is_up_to_date and not isinstance(control_plane_is_up_to_date, bool):
            raise TypeError("Expected argument 'control_plane_is_up_to_date' to be a bool")
        pulumi.set(__self__, "control_plane_is_up_to_date", control_plane_is_up_to_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_up_to_date and not isinstance(is_up_to_date, bool):
            raise TypeError("Expected argument 'is_up_to_date' to be a bool")
        pulumi.set(__self__, "is_up_to_date", is_up_to_date)
        if kube_id and not isinstance(kube_id, str):
            raise TypeError("Expected argument 'kube_id' to be a str")
        pulumi.set(__self__, "kube_id", kube_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if next_upgrade_versions and not isinstance(next_upgrade_versions, list):
            raise TypeError("Expected argument 'next_upgrade_versions' to be a list")
        pulumi.set(__self__, "next_upgrade_versions", next_upgrade_versions)
        if nodes_url and not isinstance(nodes_url, str):
            raise TypeError("Expected argument 'nodes_url' to be a str")
        pulumi.set(__self__, "nodes_url", nodes_url)
        if private_network_id and not isinstance(private_network_id, str):
            raise TypeError("Expected argument 'private_network_id' to be a str")
        pulumi.set(__self__, "private_network_id", private_network_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if update_policy and not isinstance(update_policy, str):
            raise TypeError("Expected argument 'update_policy' to be a str")
        pulumi.set(__self__, "update_policy", update_policy)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="controlPlaneIsUpToDate")
    def control_plane_is_up_to_date(self) -> bool:
        return pulumi.get(self, "control_plane_is_up_to_date")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isUpToDate")
    def is_up_to_date(self) -> bool:
        return pulumi.get(self, "is_up_to_date")

    @property
    @pulumi.getter(name="kubeId")
    def kube_id(self) -> str:
        return pulumi.get(self, "kube_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nextUpgradeVersions")
    def next_upgrade_versions(self) -> Sequence[str]:
        return pulumi.get(self, "next_upgrade_versions")

    @property
    @pulumi.getter(name="nodesUrl")
    def nodes_url(self) -> str:
        return pulumi.get(self, "nodes_url")

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> str:
        return pulumi.get(self, "private_network_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatePolicy")
    def update_policy(self) -> Optional[str]:
        return pulumi.get(self, "update_policy")

    @property
    @pulumi.getter
    def url(self) -> str:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")


class AwaitableGetOvh_Cloud_Project_KubeResult(GetOvh_Cloud_Project_KubeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOvh_Cloud_Project_KubeResult(
            control_plane_is_up_to_date=self.control_plane_is_up_to_date,
            id=self.id,
            is_up_to_date=self.is_up_to_date,
            kube_id=self.kube_id,
            name=self.name,
            next_upgrade_versions=self.next_upgrade_versions,
            nodes_url=self.nodes_url,
            private_network_id=self.private_network_id,
            region=self.region,
            service_name=self.service_name,
            status=self.status,
            update_policy=self.update_policy,
            url=self.url,
            version=self.version)


def get_ovh__cloud__project__kube(kube_id: Optional[str] = None,
                                  name: Optional[str] = None,
                                  region: Optional[str] = None,
                                  service_name: Optional[str] = None,
                                  update_policy: Optional[str] = None,
                                  version: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOvh_Cloud_Project_KubeResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['kubeId'] = kube_id
    __args__['name'] = name
    __args__['region'] = region
    __args__['serviceName'] = service_name
    __args__['updatePolicy'] = update_policy
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ovh:index/getOvh_Cloud_Project_Kube:getOvh_Cloud_Project_Kube', __args__, opts=opts, typ=GetOvh_Cloud_Project_KubeResult).value

    return AwaitableGetOvh_Cloud_Project_KubeResult(
        control_plane_is_up_to_date=__ret__.control_plane_is_up_to_date,
        id=__ret__.id,
        is_up_to_date=__ret__.is_up_to_date,
        kube_id=__ret__.kube_id,
        name=__ret__.name,
        next_upgrade_versions=__ret__.next_upgrade_versions,
        nodes_url=__ret__.nodes_url,
        private_network_id=__ret__.private_network_id,
        region=__ret__.region,
        service_name=__ret__.service_name,
        status=__ret__.status,
        update_policy=__ret__.update_policy,
        url=__ret__.url,
        version=__ret__.version)


@_utilities.lift_output_func(get_ovh__cloud__project__kube)
def get_ovh__cloud__project__kube_output(kube_id: Optional[pulumi.Input[str]] = None,
                                         name: Optional[pulumi.Input[Optional[str]]] = None,
                                         region: Optional[pulumi.Input[Optional[str]]] = None,
                                         service_name: Optional[pulumi.Input[str]] = None,
                                         update_policy: Optional[pulumi.Input[Optional[str]]] = None,
                                         version: Optional[pulumi.Input[Optional[str]]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOvh_Cloud_Project_KubeResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
