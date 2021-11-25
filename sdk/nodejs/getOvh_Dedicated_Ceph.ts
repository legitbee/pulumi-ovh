// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getOvh_Dedicated_Ceph(args: GetOvh_Dedicated_CephArgs, opts?: pulumi.InvokeOptions): Promise<GetOvh_Dedicated_CephResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOvh_Dedicated_Ceph:getOvh_Dedicated_Ceph", {
        "cephVersion": args.cephVersion,
        "serviceName": args.serviceName,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getOvh_Dedicated_Ceph.
 */
export interface GetOvh_Dedicated_CephArgs {
    cephVersion?: string;
    serviceName: string;
    status?: string;
}

/**
 * A collection of values returned by getOvh_Dedicated_Ceph.
 */
export interface GetOvh_Dedicated_CephResult {
    readonly cephMons: string[];
    readonly cephVersion: string;
    readonly crushTunables: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly label: string;
    readonly region: string;
    readonly serviceName: string;
    readonly size: number;
    readonly state: string;
    readonly status: string;
}

export function getOvh_Dedicated_CephOutput(args: GetOvh_Dedicated_CephOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOvh_Dedicated_CephResult> {
    return pulumi.output(args).apply(a => getOvh_Dedicated_Ceph(a, opts))
}

/**
 * A collection of arguments for invoking getOvh_Dedicated_Ceph.
 */
export interface GetOvh_Dedicated_CephOutputArgs {
    cephVersion?: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}