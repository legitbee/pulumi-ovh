// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/ovh_dedicated_server_install_task:ovh_dedicated_server_install_task")]
    public partial class Ovh_dedicated_server_install_task : Pulumi.CustomResource
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase
        /// </summary>
        [Output("bootidOnDestroy")]
        public Output<int?> BootidOnDestroy { get; private set; } = null!;

        /// <summary>
        /// Details of this task
        /// </summary>
        [Output("comment")]
        public Output<string> Comment { get; private set; } = null!;

        [Output("details")]
        public Output<Outputs.Ovh_dedicated_server_install_taskDetails?> Details { get; private set; } = null!;

        /// <summary>
        /// Completion date
        /// </summary>
        [Output("doneDate")]
        public Output<string> DoneDate { get; private set; } = null!;

        /// <summary>
        /// Function name
        /// </summary>
        [Output("function")]
        public Output<string> Function { get; private set; } = null!;

        /// <summary>
        /// Last update
        /// </summary>
        [Output("lastUpdate")]
        public Output<string> LastUpdate { get; private set; } = null!;

        /// <summary>
        /// Partition scheme name.
        /// </summary>
        [Output("partitionSchemeName")]
        public Output<string?> PartitionSchemeName { get; private set; } = null!;

        /// <summary>
        /// The internal name of your dedicated server.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Task Creation date
        /// </summary>
        [Output("startDate")]
        public Output<string> StartDate { get; private set; } = null!;

        /// <summary>
        /// Task status
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Template name
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;


        /// <summary>
        /// Create a Ovh_dedicated_server_install_task resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ovh_dedicated_server_install_task(string name, Ovh_dedicated_server_install_taskArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_dedicated_server_install_task:ovh_dedicated_server_install_task", name, args ?? new Ovh_dedicated_server_install_taskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ovh_dedicated_server_install_task(string name, Input<string> id, Ovh_dedicated_server_install_taskState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_dedicated_server_install_task:ovh_dedicated_server_install_task", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Ovh_dedicated_server_install_task resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ovh_dedicated_server_install_task Get(string name, Input<string> id, Ovh_dedicated_server_install_taskState? state = null, CustomResourceOptions? options = null)
        {
            return new Ovh_dedicated_server_install_task(name, id, state, options);
        }
    }

    public sealed class Ovh_dedicated_server_install_taskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase
        /// </summary>
        [Input("bootidOnDestroy")]
        public Input<int>? BootidOnDestroy { get; set; }

        [Input("details")]
        public Input<Inputs.Ovh_dedicated_server_install_taskDetailsArgs>? Details { get; set; }

        /// <summary>
        /// Partition scheme name.
        /// </summary>
        [Input("partitionSchemeName")]
        public Input<string>? PartitionSchemeName { get; set; }

        /// <summary>
        /// The internal name of your dedicated server.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Template name
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        public Ovh_dedicated_server_install_taskArgs()
        {
        }
    }

    public sealed class Ovh_dedicated_server_install_taskState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase
        /// </summary>
        [Input("bootidOnDestroy")]
        public Input<int>? BootidOnDestroy { get; set; }

        /// <summary>
        /// Details of this task
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("details")]
        public Input<Inputs.Ovh_dedicated_server_install_taskDetailsGetArgs>? Details { get; set; }

        /// <summary>
        /// Completion date
        /// </summary>
        [Input("doneDate")]
        public Input<string>? DoneDate { get; set; }

        /// <summary>
        /// Function name
        /// </summary>
        [Input("function")]
        public Input<string>? Function { get; set; }

        /// <summary>
        /// Last update
        /// </summary>
        [Input("lastUpdate")]
        public Input<string>? LastUpdate { get; set; }

        /// <summary>
        /// Partition scheme name.
        /// </summary>
        [Input("partitionSchemeName")]
        public Input<string>? PartitionSchemeName { get; set; }

        /// <summary>
        /// The internal name of your dedicated server.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Task Creation date
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// Task status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Template name
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        public Ovh_dedicated_server_install_taskState()
        {
        }
    }
}
