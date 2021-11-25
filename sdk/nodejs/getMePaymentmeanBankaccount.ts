// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getMePaymentmeanBankaccount(args?: GetMePaymentmeanBankaccountArgs, opts?: pulumi.InvokeOptions): Promise<GetMePaymentmeanBankaccountResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getMePaymentmeanBankaccount:getMePaymentmeanBankaccount", {
        "descriptionRegexp": args.descriptionRegexp,
        "state": args.state,
        "useDefault": args.useDefault,
        "useOldest": args.useOldest,
    }, opts);
}

/**
 * A collection of arguments for invoking getMePaymentmeanBankaccount.
 */
export interface GetMePaymentmeanBankaccountArgs {
    descriptionRegexp?: string;
    state?: string;
    useDefault?: boolean;
    useOldest?: boolean;
}

/**
 * A collection of values returned by getMePaymentmeanBankaccount.
 */
export interface GetMePaymentmeanBankaccountResult {
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

export function getMePaymentmeanBankaccountOutput(args?: GetMePaymentmeanBankaccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMePaymentmeanBankaccountResult> {
    return pulumi.output(args).apply(a => getMePaymentmeanBankaccount(a, opts))
}

/**
 * A collection of arguments for invoking getMePaymentmeanBankaccount.
 */
export interface GetMePaymentmeanBankaccountOutputArgs {
    descriptionRegexp?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    useDefault?: pulumi.Input<boolean>;
    useOldest?: pulumi.Input<boolean>;
}