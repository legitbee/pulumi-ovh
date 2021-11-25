// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getOvh_Dedicated_Servers(opts?: pulumi.InvokeOptions): Promise<GetOvh_Dedicated_ServersResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOvh_Dedicated_Servers:getOvh_Dedicated_Servers", {
    }, opts);
}

/**
 * A collection of values returned by getOvh_Dedicated_Servers.
 */
export interface GetOvh_Dedicated_ServersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly results: string[];
}
