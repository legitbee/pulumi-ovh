// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/ovh_dedicated_ceph_acl:ovh_dedicated_ceph_acl")]
    public partial class Ovh_dedicated_ceph_acl : Pulumi.CustomResource
    {
        [Output("family")]
        public Output<string> Family { get; private set; } = null!;

        [Output("netmask")]
        public Output<string> Netmask { get; private set; } = null!;

        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a Ovh_dedicated_ceph_acl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ovh_dedicated_ceph_acl(string name, Ovh_dedicated_ceph_aclArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_dedicated_ceph_acl:ovh_dedicated_ceph_acl", name, args ?? new Ovh_dedicated_ceph_aclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ovh_dedicated_ceph_acl(string name, Input<string> id, Ovh_dedicated_ceph_aclState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_dedicated_ceph_acl:ovh_dedicated_ceph_acl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ovh_dedicated_ceph_acl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ovh_dedicated_ceph_acl Get(string name, Input<string> id, Ovh_dedicated_ceph_aclState? state = null, CustomResourceOptions? options = null)
        {
            return new Ovh_dedicated_ceph_acl(name, id, state, options);
        }
    }

    public sealed class Ovh_dedicated_ceph_aclArgs : Pulumi.ResourceArgs
    {
        [Input("netmask", required: true)]
        public Input<string> Netmask { get; set; } = null!;

        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public Ovh_dedicated_ceph_aclArgs()
        {
        }
    }

    public sealed class Ovh_dedicated_ceph_aclState : Pulumi.ResourceArgs
    {
        [Input("family")]
        public Input<string>? Family { get; set; }

        [Input("netmask")]
        public Input<string>? Netmask { get; set; }

        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public Ovh_dedicated_ceph_aclState()
        {
        }
    }
}
