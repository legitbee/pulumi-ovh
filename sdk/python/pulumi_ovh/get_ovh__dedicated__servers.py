# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetOvh_Dedicated_ServersResult',
    'AwaitableGetOvh_Dedicated_ServersResult',
    'get_ovh__dedicated__servers',
]

@pulumi.output_type
class GetOvh_Dedicated_ServersResult:
    """
    A collection of values returned by getOvh_Dedicated_Servers.
    """
    def __init__(__self__, id=None, results=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def results(self) -> Sequence[str]:
        return pulumi.get(self, "results")


class AwaitableGetOvh_Dedicated_ServersResult(GetOvh_Dedicated_ServersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOvh_Dedicated_ServersResult(
            id=self.id,
            results=self.results)


def get_ovh__dedicated__servers(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOvh_Dedicated_ServersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ovh:index/getOvh_Dedicated_Servers:getOvh_Dedicated_Servers', __args__, opts=opts, typ=GetOvh_Dedicated_ServersResult).value

    return AwaitableGetOvh_Dedicated_ServersResult(
        id=__ret__.id,
        results=__ret__.results)
