# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetCloudProjectCapabilitiesContainerregistryResult',
    'AwaitableGetCloudProjectCapabilitiesContainerregistryResult',
    'get_cloud_project_capabilities_containerregistry',
    'get_cloud_project_capabilities_containerregistry_output',
]

@pulumi.output_type
class GetCloudProjectCapabilitiesContainerregistryResult:
    """
    A collection of values returned by getCloudProjectCapabilitiesContainerregistry.
    """
    def __init__(__self__, id=None, results=None, service_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetCloudProjectCapabilitiesContainerregistryResultResult']:
        return pulumi.get(self, "results")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")


class AwaitableGetCloudProjectCapabilitiesContainerregistryResult(GetCloudProjectCapabilitiesContainerregistryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudProjectCapabilitiesContainerregistryResult(
            id=self.id,
            results=self.results,
            service_name=self.service_name)


def get_cloud_project_capabilities_containerregistry(service_name: Optional[str] = None,
                                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudProjectCapabilitiesContainerregistryResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ovh:index/getCloudProjectCapabilitiesContainerregistry:getCloudProjectCapabilitiesContainerregistry', __args__, opts=opts, typ=GetCloudProjectCapabilitiesContainerregistryResult).value

    return AwaitableGetCloudProjectCapabilitiesContainerregistryResult(
        id=__ret__.id,
        results=__ret__.results,
        service_name=__ret__.service_name)


@_utilities.lift_output_func(get_cloud_project_capabilities_containerregistry)
def get_cloud_project_capabilities_containerregistry_output(service_name: Optional[pulumi.Input[str]] = None,
                                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCloudProjectCapabilitiesContainerregistryResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
