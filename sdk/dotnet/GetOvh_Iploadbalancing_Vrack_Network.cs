// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOvh_Iploadbalancing_Vrack_Network
    {
        public static Task<GetOvh_Iploadbalancing_Vrack_NetworkResult> InvokeAsync(GetOvh_Iploadbalancing_Vrack_NetworkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOvh_Iploadbalancing_Vrack_NetworkResult>("ovh:index/getOvh_Iploadbalancing_Vrack_Network:getOvh_Iploadbalancing_Vrack_Network", args ?? new GetOvh_Iploadbalancing_Vrack_NetworkArgs(), options.WithVersion());
    }


    public sealed class GetOvh_Iploadbalancing_Vrack_NetworkArgs : Pulumi.InvokeArgs
    {
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        [Input("vrackNetworkId", required: true)]
        public int VrackNetworkId { get; set; }

        public GetOvh_Iploadbalancing_Vrack_NetworkArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOvh_Iploadbalancing_Vrack_NetworkResult
    {
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string NatIp;
        public readonly string ServiceName;
        public readonly string Subnet;
        public readonly int Vlan;
        public readonly int VrackNetworkId;

        [OutputConstructor]
        private GetOvh_Iploadbalancing_Vrack_NetworkResult(
            string displayName,

            string id,

            string natIp,

            string serviceName,

            string subnet,

            int vlan,

            int vrackNetworkId)
        {
            DisplayName = displayName;
            Id = id;
            NatIp = natIp;
            ServiceName = serviceName;
            Subnet = subnet;
            Vlan = vlan;
            VrackNetworkId = vrackNetworkId;
        }
    }
}
