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
    public sealed class GetOvh_Order_Cart_Product_PlanSelectedPricePriceResult
    {
        public readonly string CurrencyCode;
        public readonly string Text;
        public readonly double Value;

        [OutputConstructor]
        private GetOvh_Order_Cart_Product_PlanSelectedPricePriceResult(
            string currencyCode,

            string text,

            double value)
        {
            CurrencyCode = currencyCode;
            Text = text;
            Value = value;
        }
    }
}