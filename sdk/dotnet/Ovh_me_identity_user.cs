// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/ovh_me_identity_user:ovh_me_identity_user")]
    public partial class Ovh_me_identity_user : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation date of this user
        /// </summary>
        [Output("creation")]
        public Output<string> Creation { get; private set; } = null!;

        /// <summary>
        /// User description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// User's email
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// User's group
        /// </summary>
        [Output("group")]
        public Output<string?> Group { get; private set; } = null!;

        /// <summary>
        /// Last update of this user
        /// </summary>
        [Output("lastUpdate")]
        public Output<string> LastUpdate { get; private set; } = null!;

        /// <summary>
        /// User's login suffix
        /// </summary>
        [Output("login")]
        public Output<string> Login { get; private set; } = null!;

        /// <summary>
        /// User's password
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// When the user changed his password for the last time
        /// </summary>
        [Output("passwordLastUpdate")]
        public Output<string> PasswordLastUpdate { get; private set; } = null!;

        /// <summary>
        /// Current user's status
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Ovh_me_identity_user resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ovh_me_identity_user(string name, Ovh_me_identity_userArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_me_identity_user:ovh_me_identity_user", name, args ?? new Ovh_me_identity_userArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ovh_me_identity_user(string name, Input<string> id, Ovh_me_identity_userState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/ovh_me_identity_user:ovh_me_identity_user", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ovh_me_identity_user resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ovh_me_identity_user Get(string name, Input<string> id, Ovh_me_identity_userState? state = null, CustomResourceOptions? options = null)
        {
            return new Ovh_me_identity_user(name, id, state, options);
        }
    }

    public sealed class Ovh_me_identity_userArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User's email
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// User's group
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// User's login suffix
        /// </summary>
        [Input("login", required: true)]
        public Input<string> Login { get; set; } = null!;

        /// <summary>
        /// User's password
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        public Ovh_me_identity_userArgs()
        {
        }
    }

    public sealed class Ovh_me_identity_userState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation date of this user
        /// </summary>
        [Input("creation")]
        public Input<string>? Creation { get; set; }

        /// <summary>
        /// User description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User's email
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// User's group
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Last update of this user
        /// </summary>
        [Input("lastUpdate")]
        public Input<string>? LastUpdate { get; set; }

        /// <summary>
        /// User's login suffix
        /// </summary>
        [Input("login")]
        public Input<string>? Login { get; set; }

        /// <summary>
        /// User's password
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// When the user changed his password for the last time
        /// </summary>
        [Input("passwordLastUpdate")]
        public Input<string>? PasswordLastUpdate { get; set; }

        /// <summary>
        /// Current user's status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public Ovh_me_identity_userState()
        {
        }
    }
}
