// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class Ovh_cloud_project extends pulumi.CustomResource {
    /**
     * Get an existing Ovh_cloud_project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ovh_cloud_projectState, opts?: pulumi.CustomResourceOptions): Ovh_cloud_project {
        return new Ovh_cloud_project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/ovh_cloud_project:ovh_cloud_project';

    /**
     * Returns true if the given object is an instance of Ovh_cloud_project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ovh_cloud_project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ovh_cloud_project.__pulumiType;
    }

    public /*out*/ readonly access!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string>;
    /**
     * Details about an Order
     */
    public /*out*/ readonly orders!: pulumi.Output<outputs.Ovh_cloud_projectOrder[]>;
    /**
     * Ovh Subsidiary
     */
    public readonly ovhSubsidiary!: pulumi.Output<string>;
    /**
     * Ovh payment mode
     */
    public readonly paymentMean!: pulumi.Output<string>;
    /**
     * Product Plan to order
     */
    public readonly plan!: pulumi.Output<outputs.Ovh_cloud_projectPlan>;
    /**
     * Product Plan to order
     */
    public readonly planOptions!: pulumi.Output<outputs.Ovh_cloud_projectPlanOption[] | undefined>;
    public /*out*/ readonly projectId!: pulumi.Output<string>;
    public /*out*/ readonly projectName!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Ovh_cloud_project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Ovh_cloud_projectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ovh_cloud_projectArgs | Ovh_cloud_projectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ovh_cloud_projectState | undefined;
            inputs["access"] = state ? state.access : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["orders"] = state ? state.orders : undefined;
            inputs["ovhSubsidiary"] = state ? state.ovhSubsidiary : undefined;
            inputs["paymentMean"] = state ? state.paymentMean : undefined;
            inputs["plan"] = state ? state.plan : undefined;
            inputs["planOptions"] = state ? state.planOptions : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["projectName"] = state ? state.projectName : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as Ovh_cloud_projectArgs | undefined;
            if ((!args || args.ovhSubsidiary === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ovhSubsidiary'");
            }
            if ((!args || args.paymentMean === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentMean'");
            }
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["ovhSubsidiary"] = args ? args.ovhSubsidiary : undefined;
            inputs["paymentMean"] = args ? args.paymentMean : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["planOptions"] = args ? args.planOptions : undefined;
            inputs["access"] = undefined /*out*/;
            inputs["orders"] = undefined /*out*/;
            inputs["projectId"] = undefined /*out*/;
            inputs["projectName"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Ovh_cloud_project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ovh_cloud_project resources.
 */
export interface Ovh_cloud_projectState {
    access?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * Details about an Order
     */
    orders?: pulumi.Input<pulumi.Input<inputs.Ovh_cloud_projectOrder>[]>;
    /**
     * Ovh Subsidiary
     */
    ovhSubsidiary?: pulumi.Input<string>;
    /**
     * Ovh payment mode
     */
    paymentMean?: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan?: pulumi.Input<inputs.Ovh_cloud_projectPlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.Ovh_cloud_projectPlanOption>[]>;
    projectId?: pulumi.Input<string>;
    projectName?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ovh_cloud_project resource.
 */
export interface Ovh_cloud_projectArgs {
    description?: pulumi.Input<string>;
    /**
     * Ovh Subsidiary
     */
    ovhSubsidiary: pulumi.Input<string>;
    /**
     * Ovh payment mode
     */
    paymentMean: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan: pulumi.Input<inputs.Ovh_cloud_projectPlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.Ovh_cloud_projectPlanOption>[]>;
}