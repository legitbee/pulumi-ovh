// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/iploadbalancingRefresh:IploadbalancingRefresh")]
    public partial class IploadbalancingRefresh : Pulumi.CustomResource
    {
        [Output("keepers")]
        public Output<ImmutableArray<string>> Keepers { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a IploadbalancingRefresh resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IploadbalancingRefresh(string name, IploadbalancingRefreshArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/iploadbalancingRefresh:IploadbalancingRefresh", name, args ?? new IploadbalancingRefreshArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IploadbalancingRefresh(string name, Input<string> id, IploadbalancingRefreshState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/iploadbalancingRefresh:IploadbalancingRefresh", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IploadbalancingRefresh resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IploadbalancingRefresh Get(string name, Input<string> id, IploadbalancingRefreshState? state = null, CustomResourceOptions? options = null)
        {
            return new IploadbalancingRefresh(name, id, state, options);
        }
    }

    public sealed class IploadbalancingRefreshArgs : Pulumi.ResourceArgs
    {
        [Input("keepers", required: true)]
        private InputList<string>? _keepers;
        public InputList<string> Keepers
        {
            get => _keepers ?? (_keepers = new InputList<string>());
            set => _keepers = value;
        }

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public IploadbalancingRefreshArgs()
        {
        }
    }

    public sealed class IploadbalancingRefreshState : Pulumi.ResourceArgs
    {
        [Input("keepers")]
        private InputList<string>? _keepers;
        public InputList<string> Keepers
        {
            get => _keepers ?? (_keepers = new InputList<string>());
            set => _keepers = value;
        }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public IploadbalancingRefreshState()
        {
        }
    }
}