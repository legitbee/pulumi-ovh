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
    public sealed class Ovh_iploadbalancing_http_routeRule
    {
        public readonly string? Field;
        public readonly string? Match;
        public readonly bool? Negate;
        public readonly string? Pattern;
        public readonly int? RuleId;
        public readonly string? SubField;

        [OutputConstructor]
        private Ovh_iploadbalancing_http_routeRule(
            string? field,

            string? match,

            bool? negate,

            string? pattern,

            int? ruleId,

            string? subField)
        {
            Field = field;
            Match = match;
            Negate = negate;
            Pattern = pattern;
            RuleId = ruleId;
            SubField = subField;
        }
    }
}
