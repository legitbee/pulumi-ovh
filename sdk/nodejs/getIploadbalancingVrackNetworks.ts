// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getIploadbalancingVrackNetworks(args: GetIploadbalancingVrackNetworksArgs, opts?: pulumi.InvokeOptions): Promise<GetIploadbalancingVrackNetworksResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getIploadbalancingVrackNetworks:getIploadbalancingVrackNetworks", {
        "serviceName": args.serviceName,
        "subnet": args.subnet,
        "vlanId": args.vlanId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIploadbalancingVrackNetworks.
 */
export interface GetIploadbalancingVrackNetworksArgs {
    serviceName: string;
    subnet?: string;
    vlanId?: number;
}

/**
 * A collection of values returned by getIploadbalancingVrackNetworks.
 */
export interface GetIploadbalancingVrackNetworksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly results: number[];
    readonly serviceName: string;
    readonly subnet?: string;
    readonly vlanId?: number;
}

export function getIploadbalancingVrackNetworksOutput(args: GetIploadbalancingVrackNetworksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIploadbalancingVrackNetworksResult> {
    return pulumi.output(args).apply(a => getIploadbalancingVrackNetworks(a, opts))
}

/**
 * A collection of arguments for invoking getIploadbalancingVrackNetworks.
 */
export interface GetIploadbalancingVrackNetworksOutputArgs {
    serviceName: pulumi.Input<string>;
    subnet?: pulumi.Input<string>;
    vlanId?: pulumi.Input<number>;
}
