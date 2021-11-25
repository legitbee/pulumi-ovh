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
    public sealed class Ovh_domain_zoneOrder
    {
        public readonly string? Date;
        public readonly ImmutableArray<Outputs.Ovh_domain_zoneOrderDetail> Details;
        public readonly string? ExpirationDate;
        public readonly int? OrderId;

        [OutputConstructor]
        private Ovh_domain_zoneOrder(
            string? date,

            ImmutableArray<Outputs.Ovh_domain_zoneOrderDetail> details,

            string? expirationDate,

            int? orderId)
        {
            Date = date;
            Details = details;
            ExpirationDate = expirationDate;
            OrderId = orderId;
        }
    }
}