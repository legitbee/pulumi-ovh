// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetCloudProjectCapabilitiesContainerregistryFilter
    {
        public static Task<GetCloudProjectCapabilitiesContainerregistryFilterResult> InvokeAsync(GetCloudProjectCapabilitiesContainerregistryFilterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectCapabilitiesContainerregistryFilterResult>("ovh:index/getCloudProjectCapabilitiesContainerregistryFilter:getCloudProjectCapabilitiesContainerregistryFilter", args ?? new GetCloudProjectCapabilitiesContainerregistryFilterArgs(), options.WithVersion());
    }


    public sealed class GetCloudProjectCapabilitiesContainerregistryFilterArgs : Pulumi.InvokeArgs
    {
        [Input("planName", required: true)]
        public string PlanName { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectCapabilitiesContainerregistryFilterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudProjectCapabilitiesContainerregistryFilterResult
    {
        public readonly string Code;
        public readonly string CreatedAt;
        public readonly ImmutableArray<Outputs.GetCloudProjectCapabilitiesContainerregistryFilterFeatureResult> Features;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string PlanName;
        public readonly string Region;
        public readonly ImmutableArray<Outputs.GetCloudProjectCapabilitiesContainerregistryFilterRegistryLimitResult> RegistryLimits;
        public readonly string ServiceName;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetCloudProjectCapabilitiesContainerregistryFilterResult(
            string code,

            string createdAt,

            ImmutableArray<Outputs.GetCloudProjectCapabilitiesContainerregistryFilterFeatureResult> features,

            string id,

            string name,

            string planName,

            string region,

            ImmutableArray<Outputs.GetCloudProjectCapabilitiesContainerregistryFilterRegistryLimitResult> registryLimits,

            string serviceName,

            string updatedAt)
        {
            Code = code;
            CreatedAt = createdAt;
            Features = features;
            Id = id;
            Name = name;
            PlanName = planName;
            Region = region;
            RegistryLimits = registryLimits;
            ServiceName = serviceName;
            UpdatedAt = updatedAt;
        }
    }
}