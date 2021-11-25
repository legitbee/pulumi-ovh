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
    public sealed class CloudProjectNetworkPrivateSubnetIpPool
    {
        public readonly bool? Dhcp;
        public readonly string? End;
        public readonly string? Network;
        public readonly string? Region;
        public readonly string? Start;

        [OutputConstructor]
        private CloudProjectNetworkPrivateSubnetIpPool(
            bool? dhcp,

            string? end,

            string? network,

            string? region,

            string? start)
        {
            Dhcp = dhcp;
            End = end;
            Network = network;
            Region = region;
            Start = start;
        }
    }
}