// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getOvh_Cloud_Project_Capabilities_Containerregistry_Filter(args: GetOvh_Cloud_Project_Capabilities_Containerregistry_FilterArgs, opts?: pulumi.InvokeOptions): Promise<GetOvh_Cloud_Project_Capabilities_Containerregistry_FilterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOvh_Cloud_Project_Capabilities_Containerregistry_Filter:getOvh_Cloud_Project_Capabilities_Containerregistry_Filter", {
        "planName": args.planName,
        "region": args.region,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getOvh_Cloud_Project_Capabilities_Containerregistry_Filter.
 */
export interface GetOvh_Cloud_Project_Capabilities_Containerregistry_FilterArgs {
    planName: string;
    region: string;
    serviceName: string;
}

/**
 * A collection of values returned by getOvh_Cloud_Project_Capabilities_Containerregistry_Filter.
 */
export interface GetOvh_Cloud_Project_Capabilities_Containerregistry_FilterResult {
    readonly code: string;
    readonly createdAt: string;
    readonly features: outputs.GetOvh_Cloud_Project_Capabilities_Containerregistry_FilterFeature[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly planName: string;
    readonly region: string;
    readonly registryLimits: outputs.GetOvh_Cloud_Project_Capabilities_Containerregistry_FilterRegistryLimit[];
    readonly serviceName: string;
    readonly updatedAt: string;
}

export function getOvh_Cloud_Project_Capabilities_Containerregistry_FilterOutput(args: GetOvh_Cloud_Project_Capabilities_Containerregistry_FilterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOvh_Cloud_Project_Capabilities_Containerregistry_FilterResult> {
    return pulumi.output(args).apply(a => getOvh_Cloud_Project_Capabilities_Containerregistry_Filter(a, opts))
}

/**
 * A collection of arguments for invoking getOvh_Cloud_Project_Capabilities_Containerregistry_Filter.
 */
export interface GetOvh_Cloud_Project_Capabilities_Containerregistry_FilterOutputArgs {
    planName: pulumi.Input<string>;
    region: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
}
