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
    public sealed class Ovh_iploadbalancing_http_routeAction
    {
        public readonly int? Status;
        public readonly string? Target;
        public readonly string Type;

        [OutputConstructor]
        private Ovh_iploadbalancing_http_routeAction(
            int? status,

            string? target,

            string type)
        {
            Status = status;
            Target = target;
            Type = type;
        }
    }
}