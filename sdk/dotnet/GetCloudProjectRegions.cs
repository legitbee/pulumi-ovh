// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetCloudProjectRegions
    {
        public static Task<GetCloudProjectRegionsResult> InvokeAsync(GetCloudProjectRegionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectRegionsResult>("ovh:index/getCloudProjectRegions:getCloudProjectRegions", args ?? new GetCloudProjectRegionsArgs(), options.WithVersion());
    }


    public sealed class GetCloudProjectRegionsArgs : Pulumi.InvokeArgs
    {
        [Input("hasServicesUps")]
        private List<string>? _hasServicesUps;
        public List<string> HasServicesUps
        {
            get => _hasServicesUps ?? (_hasServicesUps = new List<string>());
            set => _hasServicesUps = value;
        }

        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectRegionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudProjectRegionsResult
    {
        public readonly ImmutableArray<string> HasServicesUps;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Names;
        public readonly string ServiceName;

        [OutputConstructor]
        private GetCloudProjectRegionsResult(
            ImmutableArray<string> hasServicesUps,

            string id,

            ImmutableArray<string> names,

            string serviceName)
        {
            HasServicesUps = hasServicesUps;
            Id = id;
            Names = names;
            ServiceName = serviceName;
        }
    }
}
