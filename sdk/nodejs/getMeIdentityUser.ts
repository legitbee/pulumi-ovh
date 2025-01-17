// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getMeIdentityUser(args: GetMeIdentityUserArgs, opts?: pulumi.InvokeOptions): Promise<GetMeIdentityUserResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getMeIdentityUser:getMeIdentityUser", {
        "user": args.user,
    }, opts);
}

/**
 * A collection of arguments for invoking getMeIdentityUser.
 */
export interface GetMeIdentityUserArgs {
    user: string;
}

/**
 * A collection of values returned by getMeIdentityUser.
 */
export interface GetMeIdentityUserResult {
    readonly creation: string;
    readonly description: string;
    readonly email: string;
    readonly group: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lastUpdate: string;
    readonly login: string;
    readonly passwordLastUpdate: string;
    readonly status: string;
    readonly user: string;
}

export function getMeIdentityUserOutput(args: GetMeIdentityUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMeIdentityUserResult> {
    return pulumi.output(args).apply(a => getMeIdentityUser(a, opts))
}

/**
 * A collection of arguments for invoking getMeIdentityUser.
 */
export interface GetMeIdentityUserOutputArgs {
    user: pulumi.Input<string>;
}
