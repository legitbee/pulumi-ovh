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
    public sealed class Ovh_cloud_projectPlan
    {
        public readonly string? CatalogName;
        public readonly ImmutableArray<Outputs.Ovh_cloud_projectPlanConfiguration> Configurations;
        public readonly string Duration;
        public readonly string PlanCode;
        public readonly string PricingMode;

        [OutputConstructor]
        private Ovh_cloud_projectPlan(
            string? catalogName,

            ImmutableArray<Outputs.Ovh_cloud_projectPlanConfiguration> configurations,

            string duration,

            string planCode,

            string pricingMode)
        {
            CatalogName = catalogName;
            Configurations = configurations;
            Duration = duration;
            PlanCode = planCode;
            PricingMode = pricingMode;
        }
    }
}
