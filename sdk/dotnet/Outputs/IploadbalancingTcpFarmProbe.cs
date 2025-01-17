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
    public sealed class IploadbalancingTcpFarmProbe
    {
        public readonly bool? ForceSsl;
        public readonly int? Interval;
        public readonly string? Match;
        public readonly string? Method;
        public readonly bool? Negate;
        public readonly string? Pattern;
        public readonly int? Port;
        public readonly string Type;
        public readonly string? Url;

        [OutputConstructor]
        private IploadbalancingTcpFarmProbe(
            bool? forceSsl,

            int? interval,

            string? match,

            string? method,

            bool? negate,

            string? pattern,

            int? port,

            string type,

            string? url)
        {
            ForceSsl = forceSsl;
            Interval = interval;
            Match = match;
            Method = method;
            Negate = negate;
            Pattern = pattern;
            Port = port;
            Type = type;
            Url = url;
        }
    }
}
