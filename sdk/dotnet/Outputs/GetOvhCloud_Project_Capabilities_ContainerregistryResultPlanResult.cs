// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Outputs
{

    [OutputType]
    public sealed class GetOvhCloud_Project_Capabilities_ContainerregistryResultPlanResult
    {
        public readonly string Code;
        public readonly string CreatedAt;
        public readonly ImmutableArray<Outputs.GetOvhCloud_Project_Capabilities_ContainerregistryResultPlanFeatureResult> Features;
        public readonly string Id;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetOvhCloud_Project_Capabilities_ContainerregistryResultPlanRegistryLimitResult> RegistryLimits;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetOvhCloud_Project_Capabilities_ContainerregistryResultPlanResult(
            string code,

            string createdAt,

            ImmutableArray<Outputs.GetOvhCloud_Project_Capabilities_ContainerregistryResultPlanFeatureResult> features,

            string id,

            string name,

            ImmutableArray<Outputs.GetOvhCloud_Project_Capabilities_ContainerregistryResultPlanRegistryLimitResult> registryLimits,

            string updatedAt)
        {
            Code = code;
            CreatedAt = createdAt;
            Features = features;
            Id = id;
            Name = name;
            RegistryLimits = registryLimits;
            UpdatedAt = updatedAt;
        }
    }
}