// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getOrderCartProductPlan(args: GetOrderCartProductPlanArgs, opts?: pulumi.InvokeOptions): Promise<GetOrderCartProductPlanResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOrderCartProductPlan:getOrderCartProductPlan", {
        "cartId": args.cartId,
        "catalogName": args.catalogName,
        "planCode": args.planCode,
        "priceCapacity": args.priceCapacity,
        "product": args.product,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrderCartProductPlan.
 */
export interface GetOrderCartProductPlanArgs {
    cartId: string;
    catalogName?: string;
    planCode: string;
    priceCapacity: string;
    product: string;
}

/**
 * A collection of values returned by getOrderCartProductPlan.
 */
export interface GetOrderCartProductPlanResult {
    readonly cartId: string;
    readonly catalogName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly planCode: string;
    readonly priceCapacity: string;
    readonly prices: outputs.GetOrderCartProductPlanPrice[];
    readonly product: string;
    readonly productName: string;
    readonly productType: string;
    readonly selectedPrices: outputs.GetOrderCartProductPlanSelectedPrice[];
}

export function getOrderCartProductPlanOutput(args: GetOrderCartProductPlanOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrderCartProductPlanResult> {
    return pulumi.output(args).apply(a => getOrderCartProductPlan(a, opts))
}

/**
 * A collection of arguments for invoking getOrderCartProductPlan.
 */
export interface GetOrderCartProductPlanOutputArgs {
    cartId: pulumi.Input<string>;
    catalogName?: pulumi.Input<string>;
    planCode: pulumi.Input<string>;
    priceCapacity: pulumi.Input<string>;
    product: pulumi.Input<string>;
}