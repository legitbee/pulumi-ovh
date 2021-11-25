// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getCloudProjectCapabilitiesContainerregistryFilter(args: GetCloudProjectCapabilitiesContainerregistryFilterArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudProjectCapabilitiesContainerregistryFilterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getCloudProjectCapabilitiesContainerregistryFilter:getCloudProjectCapabilitiesContainerregistryFilter", {
        "planName": args.planName,
        "region": args.region,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudProjectCapabilitiesContainerregistryFilter.
 */
export interface GetCloudProjectCapabilitiesContainerregistryFilterArgs {
    planName: string;
    region: string;
    serviceName: string;
}

/**
 * A collection of values returned by getCloudProjectCapabilitiesContainerregistryFilter.
 */
export interface GetCloudProjectCapabilitiesContainerregistryFilterResult {
    readonly code: string;
    readonly createdAt: string;
    readonly features: outputs.GetCloudProjectCapabilitiesContainerregistryFilterFeature[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly planName: string;
    readonly region: string;
    readonly registryLimits: outputs.GetCloudProjectCapabilitiesContainerregistryFilterRegistryLimit[];
    readonly serviceName: string;
    readonly updatedAt: string;
}

export function getCloudProjectCapabilitiesContainerregistryFilterOutput(args: GetCloudProjectCapabilitiesContainerregistryFilterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudProjectCapabilitiesContainerregistryFilterResult> {
    return pulumi.output(args).apply(a => getCloudProjectCapabilitiesContainerregistryFilter(a, opts))
}

/**
 * A collection of arguments for invoking getCloudProjectCapabilitiesContainerregistryFilter.
 */
export interface GetCloudProjectCapabilitiesContainerregistryFilterOutputArgs {
    planName: pulumi.Input<string>;
    region: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
}