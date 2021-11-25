# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['Ovh_ip_reverseArgs', 'Ovh_ip_reverse']

@pulumi.input_type
class Ovh_ip_reverseArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 ip_reverse: pulumi.Input[str],
                 reverse: pulumi.Input[str]):
        """
        The set of arguments for constructing a Ovh_ip_reverse resource.
        """
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "ip_reverse", ip_reverse)
        pulumi.set(__self__, "reverse", reverse)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="ipReverse")
    def ip_reverse(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ip_reverse")

    @ip_reverse.setter
    def ip_reverse(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_reverse", value)

    @property
    @pulumi.getter
    def reverse(self) -> pulumi.Input[str]:
        return pulumi.get(self, "reverse")

    @reverse.setter
    def reverse(self, value: pulumi.Input[str]):
        pulumi.set(self, "reverse", value)


@pulumi.input_type
class _Ovh_ip_reverseState:
    def __init__(__self__, *,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_reverse: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ovh_ip_reverse resources.
        """
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ip_reverse is not None:
            pulumi.set(__self__, "ip_reverse", ip_reverse)
        if reverse is not None:
            pulumi.set(__self__, "reverse", reverse)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="ipReverse")
    def ip_reverse(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_reverse")

    @ip_reverse.setter
    def ip_reverse(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_reverse", value)

    @property
    @pulumi.getter
    def reverse(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "reverse")

    @reverse.setter
    def reverse(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reverse", value)


class Ovh_ip_reverse(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_reverse: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Ovh_ip_reverse resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Ovh_ip_reverseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Ovh_ip_reverse resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param Ovh_ip_reverseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Ovh_ip_reverseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_reverse: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[str]] = None,
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
            __props__ = Ovh_ip_reverseArgs.__new__(Ovh_ip_reverseArgs)

            if ip is None and not opts.urn:
                raise TypeError("Missing required property 'ip'")
            __props__.__dict__["ip"] = ip
            if ip_reverse is None and not opts.urn:
                raise TypeError("Missing required property 'ip_reverse'")
            __props__.__dict__["ip_reverse"] = ip_reverse
            if reverse is None and not opts.urn:
                raise TypeError("Missing required property 'reverse'")
            __props__.__dict__["reverse"] = reverse
        super(Ovh_ip_reverse, __self__).__init__(
            'ovh:index/ovh_ip_reverse:ovh_ip_reverse',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ip: Optional[pulumi.Input[str]] = None,
            ip_reverse: Optional[pulumi.Input[str]] = None,
            reverse: Optional[pulumi.Input[str]] = None) -> 'Ovh_ip_reverse':
        """
        Get an existing Ovh_ip_reverse resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Ovh_ip_reverseState.__new__(_Ovh_ip_reverseState)

        __props__.__dict__["ip"] = ip
        __props__.__dict__["ip_reverse"] = ip_reverse
        __props__.__dict__["reverse"] = reverse
        return Ovh_ip_reverse(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="ipReverse")
    def ip_reverse(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ip_reverse")

    @property
    @pulumi.getter
    def reverse(self) -> pulumi.Output[str]:
        return pulumi.get(self, "reverse")
