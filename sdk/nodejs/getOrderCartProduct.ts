// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getOrderCartProduct(args: GetOrderCartProductArgs, opts?: pulumi.InvokeOptions): Promise<GetOrderCartProductResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOrderCartProduct:getOrderCartProduct", {
        "cartId": args.cartId,
        "product": args.product,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrderCartProduct.
 */
export interface GetOrderCartProductArgs {
    cartId: string;
    product: string;
}

/**
 * A collection of values returned by getOrderCartProduct.
 */
export interface GetOrderCartProductResult {
    readonly cartId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly product: string;
    readonly results: outputs.GetOrderCartProductResult[];
}

export function getOrderCartProductOutput(args: GetOrderCartProductOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrderCartProductResult> {
    return pulumi.output(args).apply(a => getOrderCartProduct(a, opts))
}

/**
 * A collection of arguments for invoking getOrderCartProduct.
 */
export interface GetOrderCartProductOutputArgs {
    cartId: pulumi.Input<string>;
    product: pulumi.Input<string>;
}