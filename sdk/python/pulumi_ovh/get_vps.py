# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetVpsResult',
    'AwaitableGetVpsResult',
    'get_vps',
    'get_vps_output',
]

@pulumi.output_type
class GetVpsResult:
    """
    A collection of values returned by getVps.
    """
    def __init__(__self__, cluster=None, datacenter=None, displayname=None, id=None, ips=None, keymap=None, memory=None, model=None, name=None, netbootmode=None, offertype=None, service_name=None, slamonitoring=None, state=None, type=None, vcore=None, zone=None):
        if cluster and not isinstance(cluster, str):
            raise TypeError("Expected argument 'cluster' to be a str")
        pulumi.set(__self__, "cluster", cluster)
        if datacenter and not isinstance(datacenter, dict):
            raise TypeError("Expected argument 'datacenter' to be a dict")
        pulumi.set(__self__, "datacenter", datacenter)
        if displayname and not isinstance(displayname, str):
            raise TypeError("Expected argument 'displayname' to be a str")
        pulumi.set(__self__, "displayname", displayname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)
        if keymap and not isinstance(keymap, str):
            raise TypeError("Expected argument 'keymap' to be a str")
        pulumi.set(__self__, "keymap", keymap)
        if memory and not isinstance(memory, int):
            raise TypeError("Expected argument 'memory' to be a int")
        pulumi.set(__self__, "memory", memory)
        if model and not isinstance(model, dict):
            raise TypeError("Expected argument 'model' to be a dict")
        pulumi.set(__self__, "model", model)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if netbootmode and not isinstance(netbootmode, str):
            raise TypeError("Expected argument 'netbootmode' to be a str")
        pulumi.set(__self__, "netbootmode", netbootmode)
        if offertype and not isinstance(offertype, str):
            raise TypeError("Expected argument 'offertype' to be a str")
        pulumi.set(__self__, "offertype", offertype)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if slamonitoring and not isinstance(slamonitoring, bool):
            raise TypeError("Expected argument 'slamonitoring' to be a bool")
        pulumi.set(__self__, "slamonitoring", slamonitoring)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vcore and not isinstance(vcore, int):
            raise TypeError("Expected argument 'vcore' to be a int")
        pulumi.set(__self__, "vcore", vcore)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def cluster(self) -> str:
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter
    def datacenter(self) -> Mapping[str, str]:
        return pulumi.get(self, "datacenter")

    @property
    @pulumi.getter
    def displayname(self) -> str:
        return pulumi.get(self, "displayname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ips(self) -> Sequence[str]:
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter
    def keymap(self) -> str:
        return pulumi.get(self, "keymap")

    @property
    @pulumi.getter
    def memory(self) -> int:
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def model(self) -> Mapping[str, str]:
        return pulumi.get(self, "model")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def netbootmode(self) -> str:
        return pulumi.get(self, "netbootmode")

    @property
    @pulumi.getter
    def offertype(self) -> str:
        return pulumi.get(self, "offertype")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def slamonitoring(self) -> bool:
        return pulumi.get(self, "slamonitoring")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vcore(self) -> int:
        return pulumi.get(self, "vcore")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetVpsResult(GetVpsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpsResult(
            cluster=self.cluster,
            datacenter=self.datacenter,
            displayname=self.displayname,
            id=self.id,
            ips=self.ips,
            keymap=self.keymap,
            memory=self.memory,
            model=self.model,
            name=self.name,
            netbootmode=self.netbootmode,
            offertype=self.offertype,
            service_name=self.service_name,
            slamonitoring=self.slamonitoring,
            state=self.state,
            type=self.type,
            vcore=self.vcore,
            zone=self.zone)


def get_vps(service_name: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ovh:index/getVps:getVps', __args__, opts=opts, typ=GetVpsResult).value

    return AwaitableGetVpsResult(
        cluster=__ret__.cluster,
        datacenter=__ret__.datacenter,
        displayname=__ret__.displayname,
        id=__ret__.id,
        ips=__ret__.ips,
        keymap=__ret__.keymap,
        memory=__ret__.memory,
        model=__ret__.model,
        name=__ret__.name,
        netbootmode=__ret__.netbootmode,
        offertype=__ret__.offertype,
        service_name=__ret__.service_name,
        slamonitoring=__ret__.slamonitoring,
        state=__ret__.state,
        type=__ret__.type,
        vcore=__ret__.vcore,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_vps)
def get_vps_output(service_name: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
