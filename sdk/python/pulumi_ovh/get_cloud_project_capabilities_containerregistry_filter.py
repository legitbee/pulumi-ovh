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
    'GetCloudProjectCapabilitiesContainerregistryFilterResult',
    'AwaitableGetCloudProjectCapabilitiesContainerregistryFilterResult',
    'get_cloud_project_capabilities_containerregistry_filter',
    'get_cloud_project_capabilities_containerregistry_filter_output',
]

@pulumi.output_type
class GetCloudProjectCapabilitiesContainerregistryFilterResult:
    """
    A collection of values returned by getCloudProjectCapabilitiesContainerregistryFilter.
    """
    def __init__(__self__, code=None, created_at=None, features=None, id=None, name=None, plan_name=None, region=None, registry_limits=None, service_name=None, updated_at=None):
        if code and not isinstance(code, str):
            raise TypeError("Expected argument 'code' to be a str")
        pulumi.set(__self__, "code", code)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if features and not isinstance(features, list):
            raise TypeError("Expected argument 'features' to be a list")
        pulumi.set(__self__, "features", features)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if plan_name and not isinstance(plan_name, str):
            raise TypeError("Expected argument 'plan_name' to be a str")
        pulumi.set(__self__, "plan_name", plan_name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if registry_limits and not isinstance(registry_limits, list):
            raise TypeError("Expected argument 'registry_limits' to be a list")
        pulumi.set(__self__, "registry_limits", registry_limits)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def code(self) -> str:
        return pulumi.get(self, "code")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def features(self) -> Sequence['outputs.GetCloudProjectCapabilitiesContainerregistryFilterFeatureResult']:
        return pulumi.get(self, "features")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="planName")
    def plan_name(self) -> str:
        return pulumi.get(self, "plan_name")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="registryLimits")
    def registry_limits(self) -> Sequence['outputs.GetCloudProjectCapabilitiesContainerregistryFilterRegistryLimitResult']:
        return pulumi.get(self, "registry_limits")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")


class AwaitableGetCloudProjectCapabilitiesContainerregistryFilterResult(GetCloudProjectCapabilitiesContainerregistryFilterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudProjectCapabilitiesContainerregistryFilterResult(
            code=self.code,
            created_at=self.created_at,
            features=self.features,
            id=self.id,
            name=self.name,
            plan_name=self.plan_name,
            region=self.region,
            registry_limits=self.registry_limits,
            service_name=self.service_name,
            updated_at=self.updated_at)


def get_cloud_project_capabilities_containerregistry_filter(plan_name: Optional[str] = None,
                                                            region: Optional[str] = None,
                                                            service_name: Optional[str] = None,
                                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudProjectCapabilitiesContainerregistryFilterResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['planName'] = plan_name
    __args__['region'] = region
    __args__['serviceName'] = service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ovh:index/getCloudProjectCapabilitiesContainerregistryFilter:getCloudProjectCapabilitiesContainerregistryFilter', __args__, opts=opts, typ=GetCloudProjectCapabilitiesContainerregistryFilterResult).value

    return AwaitableGetCloudProjectCapabilitiesContainerregistryFilterResult(
        code=__ret__.code,
        created_at=__ret__.created_at,
        features=__ret__.features,
        id=__ret__.id,
        name=__ret__.name,
        plan_name=__ret__.plan_name,
        region=__ret__.region,
        registry_limits=__ret__.registry_limits,
        service_name=__ret__.service_name,
        updated_at=__ret__.updated_at)


@_utilities.lift_output_func(get_cloud_project_capabilities_containerregistry_filter)
def get_cloud_project_capabilities_containerregistry_filter_output(plan_name: Optional[pulumi.Input[str]] = None,
                                                                   region: Optional[pulumi.Input[str]] = None,
                                                                   service_name: Optional[pulumi.Input[str]] = None,
                                                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCloudProjectCapabilitiesContainerregistryFilterResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
