// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getOvh_Me_Paymentmean_Bankaccount(args?: GetOvh_Me_Paymentmean_BankaccountArgs, opts?: pulumi.InvokeOptions): Promise<GetOvh_Me_Paymentmean_BankaccountResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOvh_Me_Paymentmean_Bankaccount:getOvh_Me_Paymentmean_Bankaccount", {
        "descriptionRegexp": args.descriptionRegexp,
        "state": args.state,
        "useDefault": args.useDefault,
        "useOldest": args.useOldest,
    }, opts);
}

/**
 * A collection of arguments for invoking getOvh_Me_Paymentmean_Bankaccount.
 */
export interface GetOvh_Me_Paymentmean_BankaccountArgs {
    descriptionRegexp?: string;
    state?: string;
    useDefault?: boolean;
    useOldest?: boolean;
}

/**
 * A collection of values returned by getOvh_Me_Paymentmean_Bankaccount.
 */
export interface GetOvh_Me_Paymentmean_BankaccountResult {
    readonly default: boolean;
    readonly description: string;
    readonly descriptionRegexp?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly state: string;
    readonly useDefault?: boolean;
    readonly useOldest?: boolean;
}

export function getOvh_Me_Paymentmean_BankaccountOutput(args?: GetOvh_Me_Paymentmean_BankaccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOvh_Me_Paymentmean_BankaccountResult> {
    return pulumi.output(args).apply(a => getOvh_Me_Paymentmean_Bankaccount(a, opts))
}

/**
 * A collection of arguments for invoking getOvh_Me_Paymentmean_Bankaccount.
 */
export interface GetOvh_Me_Paymentmean_BankaccountOutputArgs {
    descriptionRegexp?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    useDefault?: pulumi.Input<boolean>;
    useOldest?: pulumi.Input<boolean>;
}
