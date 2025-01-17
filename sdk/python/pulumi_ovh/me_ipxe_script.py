# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MeIpxeScriptArgs', 'MeIpxeScript']

@pulumi.input_type
class MeIpxeScriptArgs:
    def __init__(__self__, *,
                 script: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MeIpxeScript resource.
        :param pulumi.Input[str] script: Content of your IPXE script
        :param pulumi.Input[str] description: For documentation purpose only. This attribute is not passed to the OVH API as it cannot be retrieved back. Instead a
               fake description ('$name auto description') is passed at creation time.
        :param pulumi.Input[str] name: Name of your script
        """
        pulumi.set(__self__, "script", script)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def script(self) -> pulumi.Input[str]:
        """
        Content of your IPXE script
        """
        return pulumi.get(self, "script")

    @script.setter
    def script(self, value: pulumi.Input[str]):
        pulumi.set(self, "script", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        For documentation purpose only. This attribute is not passed to the OVH API as it cannot be retrieved back. Instead a
        fake description ('$name auto description') is passed at creation time.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of your script
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _MeIpxeScriptState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 script: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MeIpxeScript resources.
        :param pulumi.Input[str] description: For documentation purpose only. This attribute is not passed to the OVH API as it cannot be retrieved back. Instead a
               fake description ('$name auto description') is passed at creation time.
        :param pulumi.Input[str] name: Name of your script
        :param pulumi.Input[str] script: Content of your IPXE script
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if script is not None:
            pulumi.set(__self__, "script", script)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        For documentation purpose only. This attribute is not passed to the OVH API as it cannot be retrieved back. Instead a
        fake description ('$name auto description') is passed at creation time.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of your script
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def script(self) -> Optional[pulumi.Input[str]]:
        """
        Content of your IPXE script
        """
        return pulumi.get(self, "script")

    @script.setter
    def script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "script", value)


class MeIpxeScript(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 script: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a MeIpxeScript resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: For documentation purpose only. This attribute is not passed to the OVH API as it cannot be retrieved back. Instead a
               fake description ('$name auto description') is passed at creation time.
        :param pulumi.Input[str] name: Name of your script
        :param pulumi.Input[str] script: Content of your IPXE script
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MeIpxeScriptArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a MeIpxeScript resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param MeIpxeScriptArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MeIpxeScriptArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 script: Optional[pulumi.Input[str]] = None,
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
            __props__ = MeIpxeScriptArgs.__new__(MeIpxeScriptArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if script is None and not opts.urn:
                raise TypeError("Missing required property 'script'")
            __props__.__dict__["script"] = script
        super(MeIpxeScript, __self__).__init__(
            'ovh:index/meIpxeScript:MeIpxeScript',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            script: Optional[pulumi.Input[str]] = None) -> 'MeIpxeScript':
        """
        Get an existing MeIpxeScript resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: For documentation purpose only. This attribute is not passed to the OVH API as it cannot be retrieved back. Instead a
               fake description ('$name auto description') is passed at creation time.
        :param pulumi.Input[str] name: Name of your script
        :param pulumi.Input[str] script: Content of your IPXE script
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MeIpxeScriptState.__new__(_MeIpxeScriptState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["script"] = script
        return MeIpxeScript(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        For documentation purpose only. This attribute is not passed to the OVH API as it cannot be retrieved back. Instead a
        fake description ('$name auto description') is passed at creation time.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of your script
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def script(self) -> pulumi.Output[str]:
        """
        Content of your IPXE script
        """
        return pulumi.get(self, "script")

