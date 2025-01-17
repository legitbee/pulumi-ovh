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
    public sealed class GetDedicatedServerVniResult
    {
        public readonly bool Enabled;
        public readonly string Mode;
        public readonly string Name;
        public readonly ImmutableArray<string> Nics;
        public readonly string ServerName;
        public readonly string Uuid;
        public readonly string Vrack;

        [OutputConstructor]
        private GetDedicatedServerVniResult(
            bool enabled,

            string mode,

            string name,

            ImmutableArray<string> nics,

            string serverName,

            string uuid,

            string vrack)
        {
            Enabled = enabled;
            Mode = mode;
            Name = name;
            Nics = nics;
            ServerName = serverName;
            Uuid = uuid;
            Vrack = vrack;
        }
    }
}
