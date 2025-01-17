// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getMeInstallationTemplate(args: GetMeInstallationTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetMeInstallationTemplateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("ovh:index/getMeInstallationTemplate:getMeInstallationTemplate", {
        "templateName": args.templateName,
    }, opts);
}

/**
 * A collection of arguments for invoking getMeInstallationTemplate.
 */
export interface GetMeInstallationTemplateArgs {
    templateName: string;
}

/**
 * A collection of values returned by getMeInstallationTemplate.
 */
export interface GetMeInstallationTemplateResult {
    readonly availableLanguages: string[];
    readonly beta: boolean;
    readonly bitFormat: number;
    readonly category: string;
    readonly customizations: outputs.GetMeInstallationTemplateCustomization[];
    readonly defaultLanguage: string;
    readonly deprecated: boolean;
    readonly description: string;
    readonly distribution: string;
    readonly family: string;
    readonly filesystems: string[];
    readonly hardRaidConfiguration: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lastModification: string;
    readonly lvmReady: boolean;
    readonly partitionSchemes: outputs.GetMeInstallationTemplatePartitionScheme[];
    readonly supportsDistributionKernel: boolean;
    readonly supportsGptLabel: boolean;
    readonly supportsRtm: boolean;
    readonly supportsSqlServer: boolean;
    readonly supportsUefi: string;
    readonly templateName: string;
}

export function getMeInstallationTemplateOutput(args: GetMeInstallationTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMeInstallationTemplateResult> {
    return pulumi.output(args).apply(a => getMeInstallationTemplate(a, opts))
}

/**
 * A collection of arguments for invoking getMeInstallationTemplate.
 */
export interface GetMeInstallationTemplateOutputArgs {
    templateName: pulumi.Input<string>;
}
