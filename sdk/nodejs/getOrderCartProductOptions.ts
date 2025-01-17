// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getOrderCartProductOptions(args: GetOrderCartProductOptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetOrderCartProductOptionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOrderCartProductOptions:getOrderCartProductOptions", {
        "cartId": args.cartId,
        "catalogName": args.catalogName,
        "planCode": args.planCode,
        "product": args.product,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrderCartProductOptions.
 */
export interface GetOrderCartProductOptionsArgs {
    cartId: string;
    catalogName?: string;
    planCode: string;
    product: string;
}

/**
 * A collection of values returned by getOrderCartProductOptions.
 */
export interface GetOrderCartProductOptionsResult {
    readonly cartId: string;
    readonly catalogName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly planCode: string;
    readonly product: string;
    readonly results: outputs.GetOrderCartProductOptionsResult[];
}

export function getOrderCartProductOptionsOutput(args: GetOrderCartProductOptionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrderCartProductOptionsResult> {
    return pulumi.output(args).apply(a => getOrderCartProductOptions(a, opts))
}

/**
 * A collection of arguments for invoking getOrderCartProductOptions.
 */
export interface GetOrderCartProductOptionsOutputArgs {
    cartId: pulumi.Input<string>;
    catalogName?: pulumi.Input<string>;
    planCode: pulumi.Input<string>;
    product: pulumi.Input<string>;
}
