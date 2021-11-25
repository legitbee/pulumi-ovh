// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getOvh_Vps(args: GetOvh_VpsArgs, opts?: pulumi.InvokeOptions): Promise<GetOvh_VpsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOvh_Vps:getOvh_Vps", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getOvh_Vps.
 */
export interface GetOvh_VpsArgs {
    serviceName: string;
}

/**
 * A collection of values returned by getOvh_Vps.
 */
export interface GetOvh_VpsResult {
    readonly cluster: string;
    readonly datacenter: {[key: string]: string};
    readonly displayname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ips: string[];
    readonly keymap: string;
    readonly memory: number;
    readonly model: {[key: string]: string};
    readonly name: string;
    readonly netbootmode: string;
    readonly offertype: string;
    readonly serviceName: string;
    readonly slamonitoring: boolean;
    readonly state: string;
    readonly type: string;
    readonly vcore: number;
    readonly zone: string;
}

export function getOvh_VpsOutput(args: GetOvh_VpsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOvh_VpsResult> {
    return pulumi.output(args).apply(a => getOvh_Vps(a, opts))
}

/**
 * A collection of arguments for invoking getOvh_Vps.
 */
export interface GetOvh_VpsOutputArgs {
    serviceName: pulumi.Input<string>;
}
