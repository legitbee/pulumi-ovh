// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getOvh_Cloud_Project_Regions(args: GetOvh_Cloud_Project_RegionsArgs, opts?: pulumi.InvokeOptions): Promise<GetOvh_Cloud_Project_RegionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getOvh_Cloud_Project_Regions:getOvh_Cloud_Project_Regions", {
        "hasServicesUps": args.hasServicesUps,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getOvh_Cloud_Project_Regions.
 */
export interface GetOvh_Cloud_Project_RegionsArgs {
    hasServicesUps?: string[];
    serviceName: string;
}

/**
 * A collection of values returned by getOvh_Cloud_Project_Regions.
 */
export interface GetOvh_Cloud_Project_RegionsResult {
    readonly hasServicesUps?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly names: string[];
    readonly serviceName: string;
}

export function getOvh_Cloud_Project_RegionsOutput(args: GetOvh_Cloud_Project_RegionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOvh_Cloud_Project_RegionsResult> {
    return pulumi.output(args).apply(a => getOvh_Cloud_Project_Regions(a, opts))
}

/**
 * A collection of arguments for invoking getOvh_Cloud_Project_Regions.
 */
export interface GetOvh_Cloud_Project_RegionsOutputArgs {
    hasServicesUps?: pulumi.Input<pulumi.Input<string>[]>;
    serviceName: pulumi.Input<string>;
}