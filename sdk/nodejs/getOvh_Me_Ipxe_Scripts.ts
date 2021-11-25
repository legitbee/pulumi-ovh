// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getOvh_Me_Ipxe_Scripts(opts?: pulumi.InvokeOptions): Promise<GetOvh_Me_Ipxe_ScriptsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOvh_Me_Ipxe_Scripts:getOvh_Me_Ipxe_Scripts", {
    }, opts);
}

/**
 * A collection of values returned by getOvh_Me_Ipxe_Scripts.
 */
export interface GetOvh_Me_Ipxe_ScriptsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly results: string[];
}