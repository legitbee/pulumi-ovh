// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetCloudProjectContainerregistries
    {
        public static Task<GetCloudProjectContainerregistriesResult> InvokeAsync(GetCloudProjectContainerregistriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectContainerregistriesResult>("ovh:index/getCloudProjectContainerregistries:getCloudProjectContainerregistries", args ?? new GetCloudProjectContainerregistriesArgs(), options.WithVersion());
    }


    public sealed class GetCloudProjectContainerregistriesArgs : Pulumi.InvokeArgs
    {
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectContainerregistriesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudProjectContainerregistriesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetCloudProjectContainerregistriesResultResult> Results;
        public readonly string ServiceName;

        [OutputConstructor]
        private GetCloudProjectContainerregistriesResult(
            string id,

            ImmutableArray<Outputs.GetCloudProjectContainerregistriesResultResult> results,

            string serviceName)
        {
            Id = id;
            Results = results;
            ServiceName = serviceName;
        }
    }
}
