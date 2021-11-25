// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOvh_Iploadbalancing_Vrack_Networks
    {
        public static Task<GetOvh_Iploadbalancing_Vrack_NetworksResult> InvokeAsync(GetOvh_Iploadbalancing_Vrack_NetworksArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOvh_Iploadbalancing_Vrack_NetworksResult>("ovh:index/getOvh_Iploadbalancing_Vrack_Networks:getOvh_Iploadbalancing_Vrack_Networks", args ?? new GetOvh_Iploadbalancing_Vrack_NetworksArgs(), options.WithVersion());
    }


    public sealed class GetOvh_Iploadbalancing_Vrack_NetworksArgs : Pulumi.InvokeArgs
    {
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        [Input("subnet")]
        public string? Subnet { get; set; }

        [Input("vlanId")]
        public int? VlanId { get; set; }

        public GetOvh_Iploadbalancing_Vrack_NetworksArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOvh_Iploadbalancing_Vrack_NetworksResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<int> Results;
        public readonly string ServiceName;
        public readonly string? Subnet;
        public readonly int? VlanId;

        [OutputConstructor]
        private GetOvh_Iploadbalancing_Vrack_NetworksResult(
            string id,

            ImmutableArray<int> results,

            string serviceName,

            string? subnet,

            int? vlanId)
        {
            Id = id;
            Results = results;
            ServiceName = serviceName;
            Subnet = subnet;
            VlanId = vlanId;
        }
    }
}