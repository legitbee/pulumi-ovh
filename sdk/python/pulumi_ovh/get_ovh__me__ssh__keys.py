# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetOvh_Me_Ssh_KeysResult',
    'AwaitableGetOvh_Me_Ssh_KeysResult',
    'get_ovh__me__ssh__keys',
]

@pulumi.output_type
class GetOvh_Me_Ssh_KeysResult:
    """
    A collection of values returned by getOvh_Me_Ssh_Keys.
    """
    def __init__(__self__, id=None, names=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")


class AwaitableGetOvh_Me_Ssh_KeysResult(GetOvh_Me_Ssh_KeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOvh_Me_Ssh_KeysResult(
            id=self.id,
            names=self.names)


def get_ovh__me__ssh__keys(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOvh_Me_Ssh_KeysResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ovh:index/getOvh_Me_Ssh_Keys:getOvh_Me_Ssh_Keys', __args__, opts=opts, typ=GetOvh_Me_Ssh_KeysResult).value

    return AwaitableGetOvh_Me_Ssh_KeysResult(
        id=__ret__.id,
        names=__ret__.names)
