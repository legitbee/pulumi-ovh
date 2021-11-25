// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/domainZoneRecord:DomainZoneRecord")]
    public partial class DomainZoneRecord : Pulumi.CustomResource
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
        /// Create a DomainZoneRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainZoneRecord(string name, DomainZoneRecordArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/domainZoneRecord:DomainZoneRecord", name, args ?? new DomainZoneRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainZoneRecord(string name, Input<string> id, DomainZoneRecordState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/domainZoneRecord:DomainZoneRecord", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainZoneRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainZoneRecord Get(string name, Input<string> id, DomainZoneRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainZoneRecord(name, id, state, options);
        }
    }

    public sealed class DomainZoneRecordArgs : Pulumi.ResourceArgs
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

        public DomainZoneRecordArgs()
        {
        }
    }

    public sealed class DomainZoneRecordState : Pulumi.ResourceArgs
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

        public DomainZoneRecordState()
        {
        }
    }
}