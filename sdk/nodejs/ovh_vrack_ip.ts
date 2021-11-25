// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Ovh_vrack_ip extends pulumi.CustomResource {
    /**
     * Get an existing Ovh_vrack_ip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ovh_vrack_ipState, opts?: pulumi.CustomResourceOptions): Ovh_vrack_ip {
        return new Ovh_vrack_ip(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/ovh_vrack_ip:ovh_vrack_ip';

    /**
     * Returns true if the given object is an instance of Ovh_vrack_ip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ovh_vrack_ip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ovh_vrack_ip.__pulumiType;
    }

    /**
     * Your IP block.
     */
    public readonly block!: pulumi.Output<string>;
    /**
     * Your gateway
     */
    public /*out*/ readonly gateway!: pulumi.Output<string>;
    /**
     * Your IP block
     */
    public /*out*/ readonly ip!: pulumi.Output<string>;
    /**
     * The internal name of your vrack
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Where you want your block announced on the network
     */
    public /*out*/ readonly zone!: pulumi.Output<string>;

    /**
     * Create a Ovh_vrack_ip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Ovh_vrack_ipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ovh_vrack_ipArgs | Ovh_vrack_ipState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ovh_vrack_ipState | undefined;
            inputs["block"] = state ? state.block : undefined;
            inputs["gateway"] = state ? state.gateway : undefined;
            inputs["ip"] = state ? state.ip : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as Ovh_vrack_ipArgs | undefined;
            if ((!args || args.block === undefined) && !opts.urn) {
                throw new Error("Missing required property 'block'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["block"] = args ? args.block : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["gateway"] = undefined /*out*/;
            inputs["ip"] = undefined /*out*/;
            inputs["zone"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Ovh_vrack_ip.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ovh_vrack_ip resources.
 */
export interface Ovh_vrack_ipState {
    /**
     * Your IP block.
     */
    block?: pulumi.Input<string>;
    /**
     * Your gateway
     */
    gateway?: pulumi.Input<string>;
    /**
     * Your IP block
     */
    ip?: pulumi.Input<string>;
    /**
     * The internal name of your vrack
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Where you want your block announced on the network
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ovh_vrack_ip resource.
 */
export interface Ovh_vrack_ipArgs {
    /**
     * Your IP block.
     */
    block: pulumi.Input<string>;
    /**
     * The internal name of your vrack
     */
    serviceName: pulumi.Input<string>;
}