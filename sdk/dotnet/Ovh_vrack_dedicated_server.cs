// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/ovh_vrack_dedicated_server:ovh_vrack_dedicated_server")]
    public partial class Ovh_vrack_dedicated_server : Pulumi.CustomResource
    {
        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;

        /// <summary>
        /// Service name of the resource representing the id of the cloud project.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a Ovh_vrack_dedicated_server resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ovh_vrack_dedicated_server(string name, Ovh_vrack_dedicated_serverArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_vrack_dedicated_server:ovh_vrack_dedicated_server", name, args ?? new Ovh_vrack_dedicated_serverArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ovh_vrack_dedicated_server(string name, Input<string> id, Ovh_vrack_dedicated_serverState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_vrack_dedicated_server:ovh_vrack_dedicated_server", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ovh_vrack_dedicated_server resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ovh_vrack_dedicated_server Get(string name, Input<string> id, Ovh_vrack_dedicated_serverState? state = null, CustomResourceOptions? options = null)
        {
            return new Ovh_vrack_dedicated_server(name, id, state, options);
        }
    }

    public sealed class Ovh_vrack_dedicated_serverArgs : Pulumi.ResourceArgs
    {
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        /// <summary>
        /// Service name of the resource representing the id of the cloud project.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public Ovh_vrack_dedicated_serverArgs()
        {
        }
    }

    public sealed class Ovh_vrack_dedicated_serverState : Pulumi.ResourceArgs
    {
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        /// <summary>
        /// Service name of the resource representing the id of the cloud project.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public Ovh_vrack_dedicated_serverState()
        {
        }
    }
}