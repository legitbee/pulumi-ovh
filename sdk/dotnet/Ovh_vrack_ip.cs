// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/ovh_vrack_ip:ovh_vrack_ip")]
    public partial class Ovh_vrack_ip : Pulumi.CustomResource
    {
        /// <summary>
        /// Your IP block.
        /// </summary>
        [Output("block")]
        public Output<string> Block { get; private set; } = null!;

        /// <summary>
        /// Your gateway
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// Your IP block
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The internal name of your vrack
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Where you want your block announced on the network
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Ovh_vrack_ip resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ovh_vrack_ip(string name, Ovh_vrack_ipArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_vrack_ip:ovh_vrack_ip", name, args ?? new Ovh_vrack_ipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ovh_vrack_ip(string name, Input<string> id, Ovh_vrack_ipState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_vrack_ip:ovh_vrack_ip", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ovh_vrack_ip resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ovh_vrack_ip Get(string name, Input<string> id, Ovh_vrack_ipState? state = null, CustomResourceOptions? options = null)
        {
            return new Ovh_vrack_ip(name, id, state, options);
        }
    }

    public sealed class Ovh_vrack_ipArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Your IP block.
        /// </summary>
        [Input("block", required: true)]
        public Input<string> Block { get; set; } = null!;

        /// <summary>
        /// The internal name of your vrack
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public Ovh_vrack_ipArgs()
        {
        }
    }

    public sealed class Ovh_vrack_ipState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Your IP block.
        /// </summary>
        [Input("block")]
        public Input<string>? Block { get; set; }

        /// <summary>
        /// Your gateway
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Your IP block
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The internal name of your vrack
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Where you want your block announced on the network
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public Ovh_vrack_ipState()
        {
        }
    }
}
