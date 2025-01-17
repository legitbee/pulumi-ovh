// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/ipReverseResource:IpReverseResource")]
    public partial class IpReverseResource : Pulumi.CustomResource
    {
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        [Output("ipReverse")]
        public Output<string> IpReverse { get; private set; } = null!;

        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;


        /// <summary>
        /// Create a IpReverseResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpReverseResource(string name, IpReverseResourceArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/ipReverseResource:IpReverseResource", name, args ?? new IpReverseResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpReverseResource(string name, Input<string> id, IpReverseResourceState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/ipReverseResource:IpReverseResource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpReverseResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpReverseResource Get(string name, Input<string> id, IpReverseResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new IpReverseResource(name, id, state, options);
        }
    }

    public sealed class IpReverseResourceArgs : Pulumi.ResourceArgs
    {
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        [Input("ipReverse", required: true)]
        public Input<string> IpReverse { get; set; } = null!;

        [Input("reverse", required: true)]
        public Input<string> Reverse { get; set; } = null!;

        public IpReverseResourceArgs()
        {
        }
    }

    public sealed class IpReverseResourceState : Pulumi.ResourceArgs
    {
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("ipReverse")]
        public Input<string>? IpReverse { get; set; }

        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        public IpReverseResourceState()
        {
        }
    }
}
