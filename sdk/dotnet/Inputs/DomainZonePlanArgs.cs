// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Inputs
{

    public sealed class DomainZonePlanArgs : Pulumi.ResourceArgs
    {
        [Input("catalogName")]
        public Input<string>? CatalogName { get; set; }

        [Input("configurations")]
        private InputList<Inputs.DomainZonePlanConfigurationArgs>? _configurations;
        public InputList<Inputs.DomainZonePlanConfigurationArgs> Configurations
        {
            get => _configurations ?? (_configurations = new InputList<Inputs.DomainZonePlanConfigurationArgs>());
            set => _configurations = value;
        }

        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        [Input("planCode", required: true)]
        public Input<string> PlanCode { get; set; } = null!;

        [Input("pricingMode", required: true)]
        public Input<string> PricingMode { get; set; } = null!;

        public DomainZonePlanArgs()
        {
        }
    }
}