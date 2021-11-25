# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['Ovh_me_ssh_keyArgs', 'Ovh_me_ssh_key']

@pulumi.input_type
class Ovh_me_ssh_keyArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 key_name: pulumi.Input[str],
                 default: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Ovh_me_ssh_key resource.
        :param pulumi.Input[str] key: ASCII encoded public Ssh key
        :param pulumi.Input[str] key_name: Name of this public Ssh key
        :param pulumi.Input[bool] default: True when this public Ssh key is used for rescue mode and reinstallations
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "key_name", key_name)
        if default is not None:
            pulumi.set(__self__, "default", default)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        ASCII encoded public Ssh key
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Input[str]:
        """
        Name of this public Ssh key
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[bool]]:
        """
        True when this public Ssh key is used for rescue mode and reinstallations
        """
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default", value)


@pulumi.input_type
class _Ovh_me_ssh_keyState:
    def __init__(__self__, *,
                 default: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ovh_me_ssh_key resources.
        :param pulumi.Input[bool] default: True when this public Ssh key is used for rescue mode and reinstallations
        :param pulumi.Input[str] key: ASCII encoded public Ssh key
        :param pulumi.Input[str] key_name: Name of this public Ssh key
        """
        if default is not None:
            pulumi.set(__self__, "default", default)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)

    @property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[bool]]:
        """
        True when this public Ssh key is used for rescue mode and reinstallations
        """
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        ASCII encoded public Ssh key
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of this public Ssh key
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)


class Ovh_me_ssh_key(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Ovh_me_ssh_key resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] default: True when this public Ssh key is used for rescue mode and reinstallations
        :param pulumi.Input[str] key: ASCII encoded public Ssh key
        :param pulumi.Input[str] key_name: Name of this public Ssh key
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Ovh_me_ssh_keyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Ovh_me_ssh_key resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param Ovh_me_ssh_keyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Ovh_me_ssh_keyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = Ovh_me_ssh_keyArgs.__new__(Ovh_me_ssh_keyArgs)

            __props__.__dict__["default"] = default
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            if key_name is None and not opts.urn:
                raise TypeError("Missing required property 'key_name'")
            __props__.__dict__["key_name"] = key_name
        super(Ovh_me_ssh_key, __self__).__init__(
            'ovh:index/ovh_me_ssh_key:ovh_me_ssh_key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            key_name: Optional[pulumi.Input[str]] = None) -> 'Ovh_me_ssh_key':
        """
        Get an existing Ovh_me_ssh_key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] default: True when this public Ssh key is used for rescue mode and reinstallations
        :param pulumi.Input[str] key: ASCII encoded public Ssh key
        :param pulumi.Input[str] key_name: Name of this public Ssh key
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Ovh_me_ssh_keyState.__new__(_Ovh_me_ssh_keyState)

        __props__.__dict__["default"] = default
        __props__.__dict__["key"] = key
        __props__.__dict__["key_name"] = key_name
        return Ovh_me_ssh_key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def default(self) -> pulumi.Output[bool]:
        """
        True when this public Ssh key is used for rescue mode and reinstallations
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        ASCII encoded public Ssh key
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Output[str]:
        """
        Name of this public Ssh key
        """
        return pulumi.get(self, "key_name")

