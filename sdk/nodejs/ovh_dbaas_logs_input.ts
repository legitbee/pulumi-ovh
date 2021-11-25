// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class Ovh_dbaas_logs_input extends pulumi.CustomResource {
    /**
     * Get an existing Ovh_dbaas_logs_input resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ovh_dbaas_logs_inputState, opts?: pulumi.CustomResourceOptions): Ovh_dbaas_logs_input {
        return new Ovh_dbaas_logs_input(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/ovh_dbaas_logs_input:ovh_dbaas_logs_input';

    /**
     * Returns true if the given object is an instance of Ovh_dbaas_logs_input.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ovh_dbaas_logs_input {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ovh_dbaas_logs_input.__pulumiType;
    }

    /**
     * IP blocks
     */
    public readonly allowedNetworks!: pulumi.Output<string[]>;
    /**
     * Input configuration
     */
    public readonly configuration!: pulumi.Output<outputs.Ovh_dbaas_logs_inputConfiguration>;
    /**
     * Input creation
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Input description
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Input engine ID
     */
    public readonly engineId!: pulumi.Output<string>;
    /**
     * Port
     */
    public readonly exposedPort!: pulumi.Output<string>;
    /**
     * Hostname
     */
    public /*out*/ readonly hostname!: pulumi.Output<string>;
    /**
     * Input ID
     */
    public /*out*/ readonly inputId!: pulumi.Output<string>;
    /**
     * Indicate if input need to be restarted
     */
    public /*out*/ readonly isRestartRequired!: pulumi.Output<boolean>;
    /**
     * Number of instance running
     */
    public readonly nbInstance!: pulumi.Output<number>;
    /**
     * Input IP address
     */
    public /*out*/ readonly publicAddress!: pulumi.Output<string>;
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Input SSL certificate
     */
    public /*out*/ readonly sslCertificate!: pulumi.Output<string>;
    /**
     * init: configuration required, pending: ready to start, running: available
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Associated Graylog stream
     */
    public readonly streamId!: pulumi.Output<string>;
    /**
     * Input title
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Input last update
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Ovh_dbaas_logs_input resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Ovh_dbaas_logs_inputArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ovh_dbaas_logs_inputArgs | Ovh_dbaas_logs_inputState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ovh_dbaas_logs_inputState | undefined;
            inputs["allowedNetworks"] = state ? state.allowedNetworks : undefined;
            inputs["configuration"] = state ? state.configuration : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["engineId"] = state ? state.engineId : undefined;
            inputs["exposedPort"] = state ? state.exposedPort : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["inputId"] = state ? state.inputId : undefined;
            inputs["isRestartRequired"] = state ? state.isRestartRequired : undefined;
            inputs["nbInstance"] = state ? state.nbInstance : undefined;
            inputs["publicAddress"] = state ? state.publicAddress : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["sslCertificate"] = state ? state.sslCertificate : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["streamId"] = state ? state.streamId : undefined;
            inputs["title"] = state ? state.title : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as Ovh_dbaas_logs_inputArgs | undefined;
            if ((!args || args.configuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuration'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.engineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineId'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.streamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'streamId'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            inputs["allowedNetworks"] = args ? args.allowedNetworks : undefined;
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["engineId"] = args ? args.engineId : undefined;
            inputs["exposedPort"] = args ? args.exposedPort : undefined;
            inputs["nbInstance"] = args ? args.nbInstance : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["streamId"] = args ? args.streamId : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["hostname"] = undefined /*out*/;
            inputs["inputId"] = undefined /*out*/;
            inputs["isRestartRequired"] = undefined /*out*/;
            inputs["publicAddress"] = undefined /*out*/;
            inputs["sslCertificate"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Ovh_dbaas_logs_input.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ovh_dbaas_logs_input resources.
 */
export interface Ovh_dbaas_logs_inputState {
    /**
     * IP blocks
     */
    allowedNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Input configuration
     */
    configuration?: pulumi.Input<inputs.Ovh_dbaas_logs_inputConfiguration>;
    /**
     * Input creation
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Input description
     */
    description?: pulumi.Input<string>;
    /**
     * Input engine ID
     */
    engineId?: pulumi.Input<string>;
    /**
     * Port
     */
    exposedPort?: pulumi.Input<string>;
    /**
     * Hostname
     */
    hostname?: pulumi.Input<string>;
    /**
     * Input ID
     */
    inputId?: pulumi.Input<string>;
    /**
     * Indicate if input need to be restarted
     */
    isRestartRequired?: pulumi.Input<boolean>;
    /**
     * Number of instance running
     */
    nbInstance?: pulumi.Input<number>;
    /**
     * Input IP address
     */
    publicAddress?: pulumi.Input<string>;
    serviceName?: pulumi.Input<string>;
    /**
     * Input SSL certificate
     */
    sslCertificate?: pulumi.Input<string>;
    /**
     * init: configuration required, pending: ready to start, running: available
     */
    status?: pulumi.Input<string>;
    /**
     * Associated Graylog stream
     */
    streamId?: pulumi.Input<string>;
    /**
     * Input title
     */
    title?: pulumi.Input<string>;
    /**
     * Input last update
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ovh_dbaas_logs_input resource.
 */
export interface Ovh_dbaas_logs_inputArgs {
    /**
     * IP blocks
     */
    allowedNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Input configuration
     */
    configuration: pulumi.Input<inputs.Ovh_dbaas_logs_inputConfiguration>;
    /**
     * Input description
     */
    description: pulumi.Input<string>;
    /**
     * Input engine ID
     */
    engineId: pulumi.Input<string>;
    /**
     * Port
     */
    exposedPort?: pulumi.Input<string>;
    /**
     * Number of instance running
     */
    nbInstance?: pulumi.Input<number>;
    serviceName: pulumi.Input<string>;
    /**
     * Associated Graylog stream
     */
    streamId: pulumi.Input<string>;
    /**
     * Input title
     */
    title: pulumi.Input<string>;
}