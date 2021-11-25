// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Ovh_cloud_project_kube extends pulumi.CustomResource {
    /**
     * Get an existing Ovh_cloud_project_kube resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ovh_cloud_project_kubeState, opts?: pulumi.CustomResourceOptions): Ovh_cloud_project_kube {
        return new Ovh_cloud_project_kube(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/ovh_cloud_project_kube:ovh_cloud_project_kube';

    /**
     * Returns true if the given object is an instance of Ovh_cloud_project_kube.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ovh_cloud_project_kube {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ovh_cloud_project_kube.__pulumiType;
    }

    public /*out*/ readonly controlPlaneIsUpToDate!: pulumi.Output<boolean>;
    public /*out*/ readonly isUpToDate!: pulumi.Output<boolean>;
    public /*out*/ readonly kubeconfig!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly nextUpgradeVersions!: pulumi.Output<string[]>;
    public /*out*/ readonly nodesUrl!: pulumi.Output<string>;
    public readonly privateNetworkId!: pulumi.Output<string | undefined>;
    public readonly region!: pulumi.Output<string>;
    public readonly serviceName!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public /*out*/ readonly updatePolicy!: pulumi.Output<string>;
    public /*out*/ readonly url!: pulumi.Output<string>;
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Ovh_cloud_project_kube resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Ovh_cloud_project_kubeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ovh_cloud_project_kubeArgs | Ovh_cloud_project_kubeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ovh_cloud_project_kubeState | undefined;
            inputs["controlPlaneIsUpToDate"] = state ? state.controlPlaneIsUpToDate : undefined;
            inputs["isUpToDate"] = state ? state.isUpToDate : undefined;
            inputs["kubeconfig"] = state ? state.kubeconfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nextUpgradeVersions"] = state ? state.nextUpgradeVersions : undefined;
            inputs["nodesUrl"] = state ? state.nodesUrl : undefined;
            inputs["privateNetworkId"] = state ? state.privateNetworkId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["updatePolicy"] = state ? state.updatePolicy : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as Ovh_cloud_project_kubeArgs | undefined;
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["privateNetworkId"] = args ? args.privateNetworkId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["controlPlaneIsUpToDate"] = undefined /*out*/;
            inputs["isUpToDate"] = undefined /*out*/;
            inputs["kubeconfig"] = undefined /*out*/;
            inputs["nextUpgradeVersions"] = undefined /*out*/;
            inputs["nodesUrl"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["updatePolicy"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Ovh_cloud_project_kube.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ovh_cloud_project_kube resources.
 */
export interface Ovh_cloud_project_kubeState {
    controlPlaneIsUpToDate?: pulumi.Input<boolean>;
    isUpToDate?: pulumi.Input<boolean>;
    kubeconfig?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    nextUpgradeVersions?: pulumi.Input<pulumi.Input<string>[]>;
    nodesUrl?: pulumi.Input<string>;
    privateNetworkId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    serviceName?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    updatePolicy?: pulumi.Input<string>;
    url?: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ovh_cloud_project_kube resource.
 */
export interface Ovh_cloud_project_kubeArgs {
    name?: pulumi.Input<string>;
    privateNetworkId?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}
