// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Ovh_dedicated_server_reboot_task extends pulumi.CustomResource {
    /**
     * Get an existing Ovh_dedicated_server_reboot_task resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ovh_dedicated_server_reboot_taskState, opts?: pulumi.CustomResourceOptions): Ovh_dedicated_server_reboot_task {
        return new Ovh_dedicated_server_reboot_task(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/ovh_dedicated_server_reboot_task:ovh_dedicated_server_reboot_task';

    /**
     * Returns true if the given object is an instance of Ovh_dedicated_server_reboot_task.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ovh_dedicated_server_reboot_task {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ovh_dedicated_server_reboot_task.__pulumiType;
    }

    /**
     * Details of this task
     */
    public /*out*/ readonly comment!: pulumi.Output<string>;
    /**
     * Completion date
     */
    public /*out*/ readonly doneDate!: pulumi.Output<string>;
    /**
     * Function name
     */
    public /*out*/ readonly function!: pulumi.Output<string>;
    /**
     * Change this value to recreate a reboot task.
     */
    public readonly keepers!: pulumi.Output<string[]>;
    /**
     * Last update
     */
    public /*out*/ readonly lastUpdate!: pulumi.Output<string>;
    /**
     * The internal name of your dedicated server.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Task Creation date
     */
    public /*out*/ readonly startDate!: pulumi.Output<string>;
    /**
     * Task status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Ovh_dedicated_server_reboot_task resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Ovh_dedicated_server_reboot_taskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ovh_dedicated_server_reboot_taskArgs | Ovh_dedicated_server_reboot_taskState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ovh_dedicated_server_reboot_taskState | undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["doneDate"] = state ? state.doneDate : undefined;
            inputs["function"] = state ? state.function : undefined;
            inputs["keepers"] = state ? state.keepers : undefined;
            inputs["lastUpdate"] = state ? state.lastUpdate : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["startDate"] = state ? state.startDate : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as Ovh_dedicated_server_reboot_taskArgs | undefined;
            if ((!args || args.keepers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keepers'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["keepers"] = args ? args.keepers : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["comment"] = undefined /*out*/;
            inputs["doneDate"] = undefined /*out*/;
            inputs["function"] = undefined /*out*/;
            inputs["lastUpdate"] = undefined /*out*/;
            inputs["startDate"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Ovh_dedicated_server_reboot_task.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ovh_dedicated_server_reboot_task resources.
 */
export interface Ovh_dedicated_server_reboot_taskState {
    /**
     * Details of this task
     */
    comment?: pulumi.Input<string>;
    /**
     * Completion date
     */
    doneDate?: pulumi.Input<string>;
    /**
     * Function name
     */
    function?: pulumi.Input<string>;
    /**
     * Change this value to recreate a reboot task.
     */
    keepers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Last update
     */
    lastUpdate?: pulumi.Input<string>;
    /**
     * The internal name of your dedicated server.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Task Creation date
     */
    startDate?: pulumi.Input<string>;
    /**
     * Task status
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ovh_dedicated_server_reboot_task resource.
 */
export interface Ovh_dedicated_server_reboot_taskArgs {
    /**
     * Change this value to recreate a reboot task.
     */
    keepers: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The internal name of your dedicated server.
     */
    serviceName: pulumi.Input<string>;
}