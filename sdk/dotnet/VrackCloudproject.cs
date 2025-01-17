// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/vrackCloudproject:VrackCloudproject")]
    public partial class VrackCloudproject : Pulumi.CustomResource
    {
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Service name of the resource representing the id of the cloud project.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a VrackCloudproject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VrackCloudproject(string name, VrackCloudprojectArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/vrackCloudproject:VrackCloudproject", name, args ?? new VrackCloudprojectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VrackCloudproject(string name, Input<string> id, VrackCloudprojectState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/vrackCloudproject:VrackCloudproject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VrackCloudproject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VrackCloudproject Get(string name, Input<string> id, VrackCloudprojectState? state = null, CustomResourceOptions? options = null)
        {
            return new VrackCloudproject(name, id, state, options);
        }
    }

    public sealed class VrackCloudprojectArgs : Pulumi.ResourceArgs
    {
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Service name of the resource representing the id of the cloud project.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public VrackCloudprojectArgs()
        {
        }
    }

    public sealed class VrackCloudprojectState : Pulumi.ResourceArgs
    {
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Service name of the resource representing the id of the cloud project.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public VrackCloudprojectState()
        {
        }
    }
}
