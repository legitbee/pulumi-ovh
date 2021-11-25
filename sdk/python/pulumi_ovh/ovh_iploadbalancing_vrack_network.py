# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['Ovh_iploadbalancing_vrack_networkArgs', 'Ovh_iploadbalancing_vrack_network']

@pulumi.input_type
class Ovh_iploadbalancing_vrack_networkArgs:
    def __init__(__self__, *,
                 nat_ip: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 subnet: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 farm_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 vlan: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Ovh_iploadbalancing_vrack_network resource.
        :param pulumi.Input[str] nat_ip: An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must
               be in the private network and reserved for the Load Balancer
        :param pulumi.Input[str] service_name: The internal name of your IPloadbalancer
        :param pulumi.Input[str] subnet: IP block of the private network in the vRack
        :param pulumi.Input[str] display_name: Human readable name for your vrack network
        :param pulumi.Input[Sequence[pulumi.Input[int]]] farm_ids: This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp
               farms `vrack_network_id` attribute
        :param pulumi.Input[int] vlan: VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
        """
        pulumi.set(__self__, "nat_ip", nat_ip)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "subnet", subnet)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if farm_ids is not None:
            pulumi.set(__self__, "farm_ids", farm_ids)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)

    @property
    @pulumi.getter(name="natIp")
    def nat_ip(self) -> pulumi.Input[str]:
        """
        An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must
        be in the private network and reserved for the Load Balancer
        """
        return pulumi.get(self, "nat_ip")

    @nat_ip.setter
    def nat_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "nat_ip", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The internal name of your IPloadbalancer
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Input[str]:
        """
        IP block of the private network in the vRack
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable name for your vrack network
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="farmIds")
    def farm_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp
        farms `vrack_network_id` attribute
        """
        return pulumi.get(self, "farm_ids")

    @farm_ids.setter
    def farm_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "farm_ids", value)

    @property
    @pulumi.getter
    def vlan(self) -> Optional[pulumi.Input[int]]:
        """
        VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vlan", value)


@pulumi.input_type
class _Ovh_iploadbalancing_vrack_networkState:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 farm_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 nat_ip: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[int]] = None,
                 vrack_network_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Ovh_iploadbalancing_vrack_network resources.
        :param pulumi.Input[str] display_name: Human readable name for your vrack network
        :param pulumi.Input[Sequence[pulumi.Input[int]]] farm_ids: This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp
               farms `vrack_network_id` attribute
        :param pulumi.Input[str] nat_ip: An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must
               be in the private network and reserved for the Load Balancer
        :param pulumi.Input[str] service_name: The internal name of your IPloadbalancer
        :param pulumi.Input[str] subnet: IP block of the private network in the vRack
        :param pulumi.Input[int] vlan: VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
        :param pulumi.Input[int] vrack_network_id: Internal Load Balancer identifier of the vRack private network
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if farm_ids is not None:
            pulumi.set(__self__, "farm_ids", farm_ids)
        if nat_ip is not None:
            pulumi.set(__self__, "nat_ip", nat_ip)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)
        if vrack_network_id is not None:
            pulumi.set(__self__, "vrack_network_id", vrack_network_id)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable name for your vrack network
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="farmIds")
    def farm_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp
        farms `vrack_network_id` attribute
        """
        return pulumi.get(self, "farm_ids")

    @farm_ids.setter
    def farm_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "farm_ids", value)

    @property
    @pulumi.getter(name="natIp")
    def nat_ip(self) -> Optional[pulumi.Input[str]]:
        """
        An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must
        be in the private network and reserved for the Load Balancer
        """
        return pulumi.get(self, "nat_ip")

    @nat_ip.setter
    def nat_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_ip", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The internal name of your IPloadbalancer
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input[str]]:
        """
        IP block of the private network in the vRack
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter
    def vlan(self) -> Optional[pulumi.Input[int]]:
        """
        VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vlan", value)

    @property
    @pulumi.getter(name="vrackNetworkId")
    def vrack_network_id(self) -> Optional[pulumi.Input[int]]:
        """
        Internal Load Balancer identifier of the vRack private network
        """
        return pulumi.get(self, "vrack_network_id")

    @vrack_network_id.setter
    def vrack_network_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vrack_network_id", value)


class Ovh_iploadbalancing_vrack_network(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 farm_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 nat_ip: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a Ovh_iploadbalancing_vrack_network resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Human readable name for your vrack network
        :param pulumi.Input[Sequence[pulumi.Input[int]]] farm_ids: This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp
               farms `vrack_network_id` attribute
        :param pulumi.Input[str] nat_ip: An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must
               be in the private network and reserved for the Load Balancer
        :param pulumi.Input[str] service_name: The internal name of your IPloadbalancer
        :param pulumi.Input[str] subnet: IP block of the private network in the vRack
        :param pulumi.Input[int] vlan: VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Ovh_iploadbalancing_vrack_networkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Ovh_iploadbalancing_vrack_network resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param Ovh_iploadbalancing_vrack_networkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Ovh_iploadbalancing_vrack_networkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 farm_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 nat_ip: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Ovh_iploadbalancing_vrack_networkArgs.__new__(Ovh_iploadbalancing_vrack_networkArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["farm_ids"] = farm_ids
            if nat_ip is None and not opts.urn:
                raise TypeError("Missing required property 'nat_ip'")
            __props__.__dict__["nat_ip"] = nat_ip
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if subnet is None and not opts.urn:
                raise TypeError("Missing required property 'subnet'")
            __props__.__dict__["subnet"] = subnet
            __props__.__dict__["vlan"] = vlan
            __props__.__dict__["vrack_network_id"] = None
        super(Ovh_iploadbalancing_vrack_network, __self__).__init__(
            'ovh:index/ovh_iploadbalancing_vrack_network:ovh_iploadbalancing_vrack_network',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            farm_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            nat_ip: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            subnet: Optional[pulumi.Input[str]] = None,
            vlan: Optional[pulumi.Input[int]] = None,
            vrack_network_id: Optional[pulumi.Input[int]] = None) -> 'Ovh_iploadbalancing_vrack_network':
        """
        Get an existing Ovh_iploadbalancing_vrack_network resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Human readable name for your vrack network
        :param pulumi.Input[Sequence[pulumi.Input[int]]] farm_ids: This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp
               farms `vrack_network_id` attribute
        :param pulumi.Input[str] nat_ip: An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must
               be in the private network and reserved for the Load Balancer
        :param pulumi.Input[str] service_name: The internal name of your IPloadbalancer
        :param pulumi.Input[str] subnet: IP block of the private network in the vRack
        :param pulumi.Input[int] vlan: VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
        :param pulumi.Input[int] vrack_network_id: Internal Load Balancer identifier of the vRack private network
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Ovh_iploadbalancing_vrack_networkState.__new__(_Ovh_iploadbalancing_vrack_networkState)

        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["farm_ids"] = farm_ids
        __props__.__dict__["nat_ip"] = nat_ip
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["subnet"] = subnet
        __props__.__dict__["vlan"] = vlan
        __props__.__dict__["vrack_network_id"] = vrack_network_id
        return Ovh_iploadbalancing_vrack_network(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Human readable name for your vrack network
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="farmIds")
    def farm_ids(self) -> pulumi.Output[Optional[Sequence[int]]]:
        """
        This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp
        farms `vrack_network_id` attribute
        """
        return pulumi.get(self, "farm_ids")

    @property
    @pulumi.getter(name="natIp")
    def nat_ip(self) -> pulumi.Output[str]:
        """
        An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must
        be in the private network and reserved for the Load Balancer
        """
        return pulumi.get(self, "nat_ip")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your IPloadbalancer
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Output[str]:
        """
        IP block of the private network in the vRack
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def vlan(self) -> pulumi.Output[int]:
        """
        VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
        """
        return pulumi.get(self, "vlan")

    @property
    @pulumi.getter(name="vrackNetworkId")
    def vrack_network_id(self) -> pulumi.Output[int]:
        """
        Internal Load Balancer identifier of the vRack private network
        """
        return pulumi.get(self, "vrack_network_id")

