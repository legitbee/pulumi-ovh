// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Inputs
{

    public sealed class Ovh_iploadbalancing_http_routeRuleArgs : Pulumi.ResourceArgs
    {
        [Input("field")]
        public Input<string>? Field { get; set; }

        [Input("match")]
        public Input<string>? Match { get; set; }

        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        [Input("subField")]
        public Input<string>? SubField { get; set; }

        public Ovh_iploadbalancing_http_routeRuleArgs()
        {
        }
    }
}
