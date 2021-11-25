// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getCloudProjectContainerregistries(args: GetCloudProjectContainerregistriesArgs, opts?: pulumi.InvokeOptions): Promise<GetCloudProjectContainerregistriesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getCloudProjectContainerregistries:getCloudProjectContainerregistries", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getCloudProjectContainerregistries.
 */
export interface GetCloudProjectContainerregistriesArgs {
    serviceName: string;
}

/**
 * A collection of values returned by getCloudProjectContainerregistries.
 */
export interface GetCloudProjectContainerregistriesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly results: outputs.GetCloudProjectContainerregistriesResult[];
    readonly serviceName: string;
}

export function getCloudProjectContainerregistriesOutput(args: GetCloudProjectContainerregistriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCloudProjectContainerregistriesResult> {
    return pulumi.output(args).apply(a => getCloudProjectContainerregistries(a, opts))
}

/**
 * A collection of arguments for invoking getCloudProjectContainerregistries.
 */
export interface GetCloudProjectContainerregistriesOutputArgs {
    serviceName: pulumi.Input<string>;
}
