// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class VrackCloudproject extends pulumi.CustomResource {
    /**
     * Get an existing VrackCloudproject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VrackCloudprojectState, opts?: pulumi.CustomResourceOptions): VrackCloudproject {
        return new VrackCloudproject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/vrackCloudproject:VrackCloudproject';

    /**
     * Returns true if the given object is an instance of VrackCloudproject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VrackCloudproject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VrackCloudproject.__pulumiType;
    }

    public readonly projectId!: pulumi.Output<string>;
    /**
     * Service name of the resource representing the id of the cloud project.
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a VrackCloudproject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VrackCloudprojectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VrackCloudprojectArgs | VrackCloudprojectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VrackCloudprojectState | undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as VrackCloudprojectArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VrackCloudproject.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VrackCloudproject resources.
 */
export interface VrackCloudprojectState {
    projectId?: pulumi.Input<string>;
    /**
     * Service name of the resource representing the id of the cloud project.
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VrackCloudproject resource.
 */
export interface VrackCloudprojectArgs {
    projectId: pulumi.Input<string>;
    /**
     * Service name of the resource representing the id of the cloud project.
     */
    serviceName: pulumi.Input<string>;
}