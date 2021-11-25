// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getOvh_Me_Ssh_Keys(opts?: pulumi.InvokeOptions): Promise<GetOvh_Me_Ssh_KeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOvh_Me_Ssh_Keys:getOvh_Me_Ssh_Keys", {
    }, opts);
}

/**
 * A collection of values returned by getOvh_Me_Ssh_Keys.
 */
export interface GetOvh_Me_Ssh_KeysResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly names: string[];
}
