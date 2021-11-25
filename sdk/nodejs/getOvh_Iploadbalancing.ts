// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getOvh_Iploadbalancing(args?: GetOvh_IploadbalancingArgs, opts?: pulumi.InvokeOptions): Promise<GetOvh_IploadbalancingResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOvh_Iploadbalancing:getOvh_Iploadbalancing", {
        "displayName": args.displayName,
        "ipLoadbalancing": args.ipLoadbalancing,
        "ipv4": args.ipv4,
        "ipv6": args.ipv6,
        "offer": args.offer,
        "serviceName": args.serviceName,
        "sslConfiguration": args.sslConfiguration,
        "state": args.state,
        "vrackEligibility": args.vrackEligibility,
        "vrackName": args.vrackName,
        "zones": args.zones,
    }, opts);
}

/**
 * A collection of arguments for invoking getOvh_Iploadbalancing.
 */
export interface GetOvh_IploadbalancingArgs {
    displayName?: string;
    ipLoadbalancing?: string;
    ipv4?: string;
    ipv6?: string;
    offer?: string;
    serviceName?: string;
    sslConfiguration?: string;
    state?: string;
    vrackEligibility?: boolean;
    vrackName?: string;
    zones?: string[];
}

/**
 * A collection of values returned by getOvh_Iploadbalancing.
 */
export interface GetOvh_IploadbalancingResult {
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipLoadbalancing: string;
    readonly ipv4: string;
    readonly ipv6: string;
    readonly metricsToken: string;
    readonly offer: string;
    readonly orderableZones: outputs.GetOvh_IploadbalancingOrderableZone[];
    readonly serviceName: string;
    readonly sslConfiguration: string;
    readonly state: string;
    readonly vrackEligibility: boolean;
    readonly vrackName: string;
    readonly zones: string[];
}

export function getOvh_IploadbalancingOutput(args?: GetOvh_IploadbalancingOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOvh_IploadbalancingResult> {
    return pulumi.output(args).apply(a => getOvh_Iploadbalancing(a, opts))
}

/**
 * A collection of arguments for invoking getOvh_Iploadbalancing.
 */
export interface GetOvh_IploadbalancingOutputArgs {
    displayName?: pulumi.Input<string>;
    ipLoadbalancing?: pulumi.Input<string>;
    ipv4?: pulumi.Input<string>;
    ipv6?: pulumi.Input<string>;
    offer?: pulumi.Input<string>;
    serviceName?: pulumi.Input<string>;
    sslConfiguration?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    vrackEligibility?: pulumi.Input<boolean>;
    vrackName?: pulumi.Input<string>;
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}
