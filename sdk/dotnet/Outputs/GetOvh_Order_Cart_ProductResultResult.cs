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
    public sealed class GetOvh_Order_Cart_ProductResultResult
    {
        public readonly string PlanCode;
        public readonly ImmutableArray<Outputs.GetOvh_Order_Cart_ProductResultPriceResult> Prices;
        public readonly string ProductName;
        public readonly string ProductType;

        [OutputConstructor]
        private GetOvh_Order_Cart_ProductResultResult(
            string planCode,

            ImmutableArray<Outputs.GetOvh_Order_Cart_ProductResultPriceResult> prices,

            string productName,

            string productType)
        {
            PlanCode = planCode;
            Prices = prices;
            ProductName = productName;
            ProductType = productType;
        }
    }
}
