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
    public sealed class Ovh_vrackPlanOptionConfiguration
    {
        public readonly string Label;
        public readonly string Value;

        [OutputConstructor]
        private Ovh_vrackPlanOptionConfiguration(
            string label,

            string value)
        {
            Label = label;
            Value = value;
        }
    }
}
