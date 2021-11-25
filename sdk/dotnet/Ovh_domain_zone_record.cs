// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/ovh_domain_zone_record:ovh_domain_zone_record")]
    public partial class Ovh_domain_zone_record : Pulumi.CustomResource
    {
        [Output("fieldtype")]
        public Output<string> Fieldtype { get; private set; } = null!;

        [Output("subdomain")]
        public Output<string?> Subdomain { get; private set; } = null!;

        [Output("target")]
        public Output<string> Target { get; private set; } = null!;

        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;

        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Ovh_domain_zone_record resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ovh_domain_zone_record(string name, Ovh_domain_zone_recordArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_domain_zone_record:ovh_domain_zone_record", name, args ?? new Ovh_domain_zone_recordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ovh_domain_zone_record(string name, Input<string> id, Ovh_domain_zone_recordState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_domain_zone_record:ovh_domain_zone_record", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ovh_domain_zone_record resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ovh_domain_zone_record Get(string name, Input<string> id, Ovh_domain_zone_recordState? state = null, CustomResourceOptions? options = null)
        {
            return new Ovh_domain_zone_record(name, id, state, options);
        }
    }

    public sealed class Ovh_domain_zone_recordArgs : Pulumi.ResourceArgs
    {
        [Input("fieldtype", required: true)]
        public Input<string> Fieldtype { get; set; } = null!;

        [Input("subdomain")]
        public Input<string>? Subdomain { get; set; }

        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public Ovh_domain_zone_recordArgs()
        {
        }
    }

    public sealed class Ovh_domain_zone_recordState : Pulumi.ResourceArgs
    {
        [Input("fieldtype")]
        public Input<string>? Fieldtype { get; set; }

        [Input("subdomain")]
        public Input<string>? Subdomain { get; set; }

        [Input("target")]
        public Input<string>? Target { get; set; }

        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public Ovh_domain_zone_recordState()
        {
        }
    }
}