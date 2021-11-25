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
    public sealed class GetOvh_Me_Installation_TemplatePartitionSchemeResult
    {
        public readonly ImmutableArray<Outputs.GetOvh_Me_Installation_TemplatePartitionSchemeHardwareRaidResult> HardwareRaids;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetOvh_Me_Installation_TemplatePartitionSchemePartitionResult> Partitions;
        public readonly int Priority;

        [OutputConstructor]
        private GetOvh_Me_Installation_TemplatePartitionSchemeResult(
            ImmutableArray<Outputs.GetOvh_Me_Installation_TemplatePartitionSchemeHardwareRaidResult> hardwareRaids,

            string name,

            ImmutableArray<Outputs.GetOvh_Me_Installation_TemplatePartitionSchemePartitionResult> partitions,

            int priority)
        {
            HardwareRaids = hardwareRaids;
            Name = name;
            Partitions = partitions;
            Priority = priority;
        }
    }
}
