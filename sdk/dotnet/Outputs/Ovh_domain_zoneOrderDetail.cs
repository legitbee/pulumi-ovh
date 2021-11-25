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
    public sealed class Ovh_domain_zoneOrderDetail
    {
        public readonly string? Description;
        public readonly string? Domain;
        public readonly int? OrderDetailId;
        public readonly string? Quantity;

        [OutputConstructor]
        private Ovh_domain_zoneOrderDetail(
            string? description,

            string? domain,

            int? orderDetailId,

            string? quantity)
        {
            Description = description;
            Domain = domain;
            OrderDetailId = orderDetailId;
            Quantity = quantity;
        }
    }
}
