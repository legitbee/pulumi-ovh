// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Ovh_dbaas_logs_output_graylog_stream extends pulumi.CustomResource {
    /**
     * Get an existing Ovh_dbaas_logs_output_graylog_stream resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ovh_dbaas_logs_output_graylog_streamState, opts?: pulumi.CustomResourceOptions): Ovh_dbaas_logs_output_graylog_stream {
        return new Ovh_dbaas_logs_output_graylog_stream(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/ovh_dbaas_logs_output_graylog_stream:ovh_dbaas_logs_output_graylog_stream';

    /**
     * Returns true if the given object is an instance of Ovh_dbaas_logs_output_graylog_stream.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ovh_dbaas_logs_output_graylog_stream {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ovh_dbaas_logs_output_graylog_stream.__pulumiType;
    }

    /**
     * Indicates if the current user can create alert on the stream
     */
    public /*out*/ readonly canAlert!: pulumi.Output<boolean>;
    /**
     * Cold storage compression method
     */
    public readonly coldStorageCompression!: pulumi.Output<string>;
    /**
     * ColdStorage content
     */
    public readonly coldStorageContent!: pulumi.Output<string>;
    /**
     * Is Cold storage enabled?
     */
    public readonly coldStorageEnabled!: pulumi.Output<boolean>;
    /**
     * Notify on new Cold storage archive
     */
    public readonly coldStorageNotifyEnabled!: pulumi.Output<boolean>;
    /**
     * Cold storage retention in year
     */
    public readonly coldStorageRetention!: pulumi.Output<number>;
    /**
     * ColdStorage destination
     */
    public readonly coldStorageTarget!: pulumi.Output<string>;
    /**
     * Stream creation
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Stream description
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Enable ES indexing
     */
    public readonly indexingEnabled!: pulumi.Output<boolean>;
    /**
     * Maximum indexing size (in GB)
     */
    public readonly indexingMaxSize!: pulumi.Output<number>;
    /**
     * If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
     */
    public readonly indexingNotifyEnabled!: pulumi.Output<boolean>;
    /**
     * Indicates if you are allowed to edit entry
     */
    public /*out*/ readonly isEditable!: pulumi.Output<boolean>;
    /**
     * Indicates if you are allowed to share entry
     */
    public /*out*/ readonly isShareable!: pulumi.Output<boolean>;
    /**
     * Number of alert condition
     */
    public /*out*/ readonly nbAlertCondition!: pulumi.Output<number>;
    /**
     * Number of coldstored archives
     */
    public /*out*/ readonly nbArchive!: pulumi.Output<number>;
    /**
     * Parent stream ID
     */
    public readonly parentStreamId!: pulumi.Output<string | undefined>;
    /**
     * If set, pause indexing when maximum size is reach
     */
    public readonly pauseIndexingOnMaxSize!: pulumi.Output<boolean>;
    /**
     * Retention ID
     */
    public readonly retentionId!: pulumi.Output<string>;
    /**
     * The service name
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Stream ID
     */
    public /*out*/ readonly streamId!: pulumi.Output<string>;
    /**
     * Stream description
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Stream last update
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Enable Websocket
     */
    public readonly webSocketEnabled!: pulumi.Output<boolean>;

    /**
     * Create a Ovh_dbaas_logs_output_graylog_stream resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Ovh_dbaas_logs_output_graylog_streamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ovh_dbaas_logs_output_graylog_streamArgs | Ovh_dbaas_logs_output_graylog_streamState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ovh_dbaas_logs_output_graylog_streamState | undefined;
            inputs["canAlert"] = state ? state.canAlert : undefined;
            inputs["coldStorageCompression"] = state ? state.coldStorageCompression : undefined;
            inputs["coldStorageContent"] = state ? state.coldStorageContent : undefined;
            inputs["coldStorageEnabled"] = state ? state.coldStorageEnabled : undefined;
            inputs["coldStorageNotifyEnabled"] = state ? state.coldStorageNotifyEnabled : undefined;
            inputs["coldStorageRetention"] = state ? state.coldStorageRetention : undefined;
            inputs["coldStorageTarget"] = state ? state.coldStorageTarget : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["indexingEnabled"] = state ? state.indexingEnabled : undefined;
            inputs["indexingMaxSize"] = state ? state.indexingMaxSize : undefined;
            inputs["indexingNotifyEnabled"] = state ? state.indexingNotifyEnabled : undefined;
            inputs["isEditable"] = state ? state.isEditable : undefined;
            inputs["isShareable"] = state ? state.isShareable : undefined;
            inputs["nbAlertCondition"] = state ? state.nbAlertCondition : undefined;
            inputs["nbArchive"] = state ? state.nbArchive : undefined;
            inputs["parentStreamId"] = state ? state.parentStreamId : undefined;
            inputs["pauseIndexingOnMaxSize"] = state ? state.pauseIndexingOnMaxSize : undefined;
            inputs["retentionId"] = state ? state.retentionId : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["streamId"] = state ? state.streamId : undefined;
            inputs["title"] = state ? state.title : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
            inputs["webSocketEnabled"] = state ? state.webSocketEnabled : undefined;
        } else {
            const args = argsOrState as Ovh_dbaas_logs_output_graylog_streamArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            inputs["coldStorageCompression"] = args ? args.coldStorageCompression : undefined;
            inputs["coldStorageContent"] = args ? args.coldStorageContent : undefined;
            inputs["coldStorageEnabled"] = args ? args.coldStorageEnabled : undefined;
            inputs["coldStorageNotifyEnabled"] = args ? args.coldStorageNotifyEnabled : undefined;
            inputs["coldStorageRetention"] = args ? args.coldStorageRetention : undefined;
            inputs["coldStorageTarget"] = args ? args.coldStorageTarget : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["indexingEnabled"] = args ? args.indexingEnabled : undefined;
            inputs["indexingMaxSize"] = args ? args.indexingMaxSize : undefined;
            inputs["indexingNotifyEnabled"] = args ? args.indexingNotifyEnabled : undefined;
            inputs["parentStreamId"] = args ? args.parentStreamId : undefined;
            inputs["pauseIndexingOnMaxSize"] = args ? args.pauseIndexingOnMaxSize : undefined;
            inputs["retentionId"] = args ? args.retentionId : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["webSocketEnabled"] = args ? args.webSocketEnabled : undefined;
            inputs["canAlert"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["isEditable"] = undefined /*out*/;
            inputs["isShareable"] = undefined /*out*/;
            inputs["nbAlertCondition"] = undefined /*out*/;
            inputs["nbArchive"] = undefined /*out*/;
            inputs["streamId"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Ovh_dbaas_logs_output_graylog_stream.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ovh_dbaas_logs_output_graylog_stream resources.
 */
export interface Ovh_dbaas_logs_output_graylog_streamState {
    /**
     * Indicates if the current user can create alert on the stream
     */
    canAlert?: pulumi.Input<boolean>;
    /**
     * Cold storage compression method
     */
    coldStorageCompression?: pulumi.Input<string>;
    /**
     * ColdStorage content
     */
    coldStorageContent?: pulumi.Input<string>;
    /**
     * Is Cold storage enabled?
     */
    coldStorageEnabled?: pulumi.Input<boolean>;
    /**
     * Notify on new Cold storage archive
     */
    coldStorageNotifyEnabled?: pulumi.Input<boolean>;
    /**
     * Cold storage retention in year
     */
    coldStorageRetention?: pulumi.Input<number>;
    /**
     * ColdStorage destination
     */
    coldStorageTarget?: pulumi.Input<string>;
    /**
     * Stream creation
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Stream description
     */
    description?: pulumi.Input<string>;
    /**
     * Enable ES indexing
     */
    indexingEnabled?: pulumi.Input<boolean>;
    /**
     * Maximum indexing size (in GB)
     */
    indexingMaxSize?: pulumi.Input<number>;
    /**
     * If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
     */
    indexingNotifyEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates if you are allowed to edit entry
     */
    isEditable?: pulumi.Input<boolean>;
    /**
     * Indicates if you are allowed to share entry
     */
    isShareable?: pulumi.Input<boolean>;
    /**
     * Number of alert condition
     */
    nbAlertCondition?: pulumi.Input<number>;
    /**
     * Number of coldstored archives
     */
    nbArchive?: pulumi.Input<number>;
    /**
     * Parent stream ID
     */
    parentStreamId?: pulumi.Input<string>;
    /**
     * If set, pause indexing when maximum size is reach
     */
    pauseIndexingOnMaxSize?: pulumi.Input<boolean>;
    /**
     * Retention ID
     */
    retentionId?: pulumi.Input<string>;
    /**
     * The service name
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Stream ID
     */
    streamId?: pulumi.Input<string>;
    /**
     * Stream description
     */
    title?: pulumi.Input<string>;
    /**
     * Stream last update
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Enable Websocket
     */
    webSocketEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Ovh_dbaas_logs_output_graylog_stream resource.
 */
export interface Ovh_dbaas_logs_output_graylog_streamArgs {
    /**
     * Cold storage compression method
     */
    coldStorageCompression?: pulumi.Input<string>;
    /**
     * ColdStorage content
     */
    coldStorageContent?: pulumi.Input<string>;
    /**
     * Is Cold storage enabled?
     */
    coldStorageEnabled?: pulumi.Input<boolean>;
    /**
     * Notify on new Cold storage archive
     */
    coldStorageNotifyEnabled?: pulumi.Input<boolean>;
    /**
     * Cold storage retention in year
     */
    coldStorageRetention?: pulumi.Input<number>;
    /**
     * ColdStorage destination
     */
    coldStorageTarget?: pulumi.Input<string>;
    /**
     * Stream description
     */
    description: pulumi.Input<string>;
    /**
     * Enable ES indexing
     */
    indexingEnabled?: pulumi.Input<boolean>;
    /**
     * Maximum indexing size (in GB)
     */
    indexingMaxSize?: pulumi.Input<number>;
    /**
     * If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
     */
    indexingNotifyEnabled?: pulumi.Input<boolean>;
    /**
     * Parent stream ID
     */
    parentStreamId?: pulumi.Input<string>;
    /**
     * If set, pause indexing when maximum size is reach
     */
    pauseIndexingOnMaxSize?: pulumi.Input<boolean>;
    /**
     * Retention ID
     */
    retentionId?: pulumi.Input<string>;
    /**
     * The service name
     */
    serviceName: pulumi.Input<string>;
    /**
     * Stream description
     */
    title: pulumi.Input<string>;
    /**
     * Enable Websocket
     */
    webSocketEnabled?: pulumi.Input<boolean>;
}