// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetMeIpxeScripts
    {
        public static Task<GetMeIpxeScriptsResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMeIpxeScriptsResult>("ovh:index/getMeIpxeScripts:getMeIpxeScripts", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetMeIpxeScriptsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Results;

        [OutputConstructor]
        private GetMeIpxeScriptsResult(
            string id,

            ImmutableArray<string> results)
        {
            Id = id;
            Results = results;
        }
    }
}
