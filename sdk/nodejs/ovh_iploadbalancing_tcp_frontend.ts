// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Ovh_iploadbalancing_tcp_frontend extends pulumi.CustomResource {
    /**
     * Get an existing Ovh_iploadbalancing_tcp_frontend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ovh_iploadbalancing_tcp_frontendState, opts?: pulumi.CustomResourceOptions): Ovh_iploadbalancing_tcp_frontend {
        return new Ovh_iploadbalancing_tcp_frontend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/ovh_iploadbalancing_tcp_frontend:ovh_iploadbalancing_tcp_frontend';

    /**
     * Returns true if the given object is an instance of Ovh_iploadbalancing_tcp_frontend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ovh_iploadbalancing_tcp_frontend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ovh_iploadbalancing_tcp_frontend.__pulumiType;
    }

    public readonly allowedSources!: pulumi.Output<string[] | undefined>;
    public readonly dedicatedIpfos!: pulumi.Output<string[] | undefined>;
    public readonly defaultFarmId!: pulumi.Output<number>;
    public readonly defaultSslId!: pulumi.Output<number>;
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    public readonly displayName!: pulumi.Output<string | undefined>;
    public readonly port!: pulumi.Output<string>;
    public readonly serviceName!: pulumi.Output<string>;
    public readonly ssl!: pulumi.Output<boolean | undefined>;
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Ovh_iploadbalancing_tcp_frontend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Ovh_iploadbalancing_tcp_frontendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ovh_iploadbalancing_tcp_frontendArgs | Ovh_iploadbalancing_tcp_frontendState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ovh_iploadbalancing_tcp_frontendState | undefined;
            inputs["allowedSources"] = state ? state.allowedSources : undefined;
            inputs["dedicatedIpfos"] = state ? state.dedicatedIpfos : undefined;
            inputs["defaultFarmId"] = state ? state.defaultFarmId : undefined;
            inputs["defaultSslId"] = state ? state.defaultSslId : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["ssl"] = state ? state.ssl : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as Ovh_iploadbalancing_tcp_frontendArgs | undefined;
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["allowedSources"] = args ? args.allowedSources : undefined;
            inputs["dedicatedIpfos"] = args ? args.dedicatedIpfos : undefined;
            inputs["defaultFarmId"] = args ? args.defaultFarmId : undefined;
            inputs["defaultSslId"] = args ? args.defaultSslId : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["ssl"] = args ? args.ssl : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Ovh_iploadbalancing_tcp_frontend.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ovh_iploadbalancing_tcp_frontend resources.
 */
export interface Ovh_iploadbalancing_tcp_frontendState {
    allowedSources?: pulumi.Input<pulumi.Input<string>[]>;
    dedicatedIpfos?: pulumi.Input<pulumi.Input<string>[]>;
    defaultFarmId?: pulumi.Input<number>;
    defaultSslId?: pulumi.Input<number>;
    disabled?: pulumi.Input<boolean>;
    displayName?: pulumi.Input<string>;
    port?: pulumi.Input<string>;
    serviceName?: pulumi.Input<string>;
    ssl?: pulumi.Input<boolean>;
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ovh_iploadbalancing_tcp_frontend resource.
 */
export interface Ovh_iploadbalancing_tcp_frontendArgs {
    allowedSources?: pulumi.Input<pulumi.Input<string>[]>;
    dedicatedIpfos?: pulumi.Input<pulumi.Input<string>[]>;
    defaultFarmId?: pulumi.Input<number>;
    defaultSslId?: pulumi.Input<number>;
    disabled?: pulumi.Input<boolean>;
    displayName?: pulumi.Input<string>;
    port: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
    ssl?: pulumi.Input<boolean>;
    zone: pulumi.Input<string>;
}