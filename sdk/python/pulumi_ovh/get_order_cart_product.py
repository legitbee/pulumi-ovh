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
    'GetOrderCartProductResult',
    'AwaitableGetOrderCartProductResult',
    'get_order_cart_product',
    'get_order_cart_product_output',
]

@pulumi.output_type
class GetOrderCartProductResult:
    """
    A collection of values returned by getOrderCartProduct.
    """
    def __init__(__self__, cart_id=None, id=None, product=None, results=None):
        if cart_id and not isinstance(cart_id, str):
            raise TypeError("Expected argument 'cart_id' to be a str")
        pulumi.set(__self__, "cart_id", cart_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if product and not isinstance(product, str):
            raise TypeError("Expected argument 'product' to be a str")
        pulumi.set(__self__, "product", product)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

    @property
    @pulumi.getter(name="cartId")
    def cart_id(self) -> str:
        return pulumi.get(self, "cart_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def product(self) -> str:
        return pulumi.get(self, "product")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetOrderCartProductResultResult']:
        return pulumi.get(self, "results")


class AwaitableGetOrderCartProductResult(GetOrderCartProductResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrderCartProductResult(
            cart_id=self.cart_id,
            id=self.id,
            product=self.product,
            results=self.results)


def get_order_cart_product(cart_id: Optional[str] = None,
                           product: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrderCartProductResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['cartId'] = cart_id
    __args__['product'] = product
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ovh:index/getOrderCartProduct:getOrderCartProduct', __args__, opts=opts, typ=GetOrderCartProductResult).value

    return AwaitableGetOrderCartProductResult(
        cart_id=__ret__.cart_id,
        id=__ret__.id,
        product=__ret__.product,
        results=__ret__.results)


@_utilities.lift_output_func(get_order_cart_product)
def get_order_cart_product_output(cart_id: Optional[pulumi.Input[str]] = None,
                                  product: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrderCartProductResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
