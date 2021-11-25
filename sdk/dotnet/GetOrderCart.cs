// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOrderCart
    {
        public static Task<GetOrderCartResult> InvokeAsync(GetOrderCartArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrderCartResult>("ovh:index/getOrderCart:getOrderCart", args ?? new GetOrderCartArgs(), options.WithVersion());
    }


    public sealed class GetOrderCartArgs : Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

        [Input("expire")]
        public string? Expire { get; set; }

        [Input("ovhSubsidiary", required: true)]
        public string OvhSubsidiary { get; set; } = null!;

        public GetOrderCartArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOrderCartResult
    {
        public readonly string CartId;
        public readonly string? Description;
        public readonly string Expire;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<int> Items;
        public readonly string OvhSubsidiary;
        public readonly bool ReadOnly;

        [OutputConstructor]
        private GetOrderCartResult(
            string cartId,

            string? description,

            string expire,

            string id,

            ImmutableArray<int> items,

            string ovhSubsidiary,

            bool readOnly)
        {
            CartId = cartId;
            Description = description;
            Expire = expire;
            Id = id;
            Items = items;
            OvhSubsidiary = ovhSubsidiary;
            ReadOnly = readOnly;
        }
    }
}