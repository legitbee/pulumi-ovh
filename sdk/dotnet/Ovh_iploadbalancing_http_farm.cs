// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/ovh_iploadbalancing_http_farm:ovh_iploadbalancing_http_farm")]
    public partial class Ovh_iploadbalancing_http_farm : Pulumi.CustomResource
    {
        [Output("balance")]
        public Output<string?> Balance { get; private set; } = null!;

        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        [Output("probe")]
        public Output<Outputs.Ovh_iploadbalancing_http_farmProbe?> Probe { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        [Output("stickiness")]
        public Output<string?> Stickiness { get; private set; } = null!;

        [Output("vrackNetworkId")]
        public Output<int?> VrackNetworkId { get; private set; } = null!;

        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Ovh_iploadbalancing_http_farm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ovh_iploadbalancing_http_farm(string name, Ovh_iploadbalancing_http_farmArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_iploadbalancing_http_farm:ovh_iploadbalancing_http_farm", name, args ?? new Ovh_iploadbalancing_http_farmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ovh_iploadbalancing_http_farm(string name, Input<string> id, Ovh_iploadbalancing_http_farmState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_iploadbalancing_http_farm:ovh_iploadbalancing_http_farm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ovh_iploadbalancing_http_farm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ovh_iploadbalancing_http_farm Get(string name, Input<string> id, Ovh_iploadbalancing_http_farmState? state = null, CustomResourceOptions? options = null)
        {
            return new Ovh_iploadbalancing_http_farm(name, id, state, options);
        }
    }

    public sealed class Ovh_iploadbalancing_http_farmArgs : Pulumi.ResourceArgs
    {
        [Input("balance")]
        public Input<string>? Balance { get; set; }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("probe")]
        public Input<Inputs.Ovh_iploadbalancing_http_farmProbeArgs>? Probe { get; set; }

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        [Input("stickiness")]
        public Input<string>? Stickiness { get; set; }

        [Input("vrackNetworkId")]
        public Input<int>? VrackNetworkId { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public Ovh_iploadbalancing_http_farmArgs()
        {
        }
    }

    public sealed class Ovh_iploadbalancing_http_farmState : Pulumi.ResourceArgs
    {
        [Input("balance")]
        public Input<string>? Balance { get; set; }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("probe")]
        public Input<Inputs.Ovh_iploadbalancing_http_farmProbeGetArgs>? Probe { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        [Input("stickiness")]
        public Input<string>? Stickiness { get; set; }

        [Input("vrackNetworkId")]
        public Input<int>? VrackNetworkId { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public Ovh_iploadbalancing_http_farmState()
        {
        }
    }
}
