// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetMeIdentityUser
    {
        public static Task<GetMeIdentityUserResult> InvokeAsync(GetMeIdentityUserArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMeIdentityUserResult>("ovh:index/getMeIdentityUser:getMeIdentityUser", args ?? new GetMeIdentityUserArgs(), options.WithVersion());
    }


    public sealed class GetMeIdentityUserArgs : Pulumi.InvokeArgs
    {
        [Input("user", required: true)]
        public string User { get; set; } = null!;

        public GetMeIdentityUserArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMeIdentityUserResult
    {
        public readonly string Creation;
        public readonly string Description;
        public readonly string Email;
        public readonly string Group;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LastUpdate;
        public readonly string Login;
        public readonly string PasswordLastUpdate;
        public readonly string Status;
        public readonly string User;

        [OutputConstructor]
        private GetMeIdentityUserResult(
            string creation,

            string description,

            string email,

            string group,

            string id,

            string lastUpdate,

            string login,

            string passwordLastUpdate,

            string status,

            string user)
        {
            Creation = creation;
            Description = description;
            Email = email;
            Group = group;
            Id = id;
            LastUpdate = lastUpdate;
            Login = login;
            PasswordLastUpdate = passwordLastUpdate;
            Status = status;
            User = user;
        }
    }
}