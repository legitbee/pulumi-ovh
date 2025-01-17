// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetIploadbalancing
    {
        public static Task<GetIploadbalancingResult> InvokeAsync(GetIploadbalancingArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIploadbalancingResult>("ovh:index/getIploadbalancing:getIploadbalancing", args ?? new GetIploadbalancingArgs(), options.WithVersion());
    }


    public sealed class GetIploadbalancingArgs : Pulumi.InvokeArgs
    {
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("ipLoadbalancing")]
        public string? IpLoadbalancing { get; set; }

        [Input("ipv4")]
        public string? Ipv4 { get; set; }

        [Input("ipv6")]
        public string? Ipv6 { get; set; }

        [Input("offer")]
        public string? Offer { get; set; }

        [Input("serviceName")]
        public string? ServiceName { get; set; }

        [Input("sslConfiguration")]
        public string? SslConfiguration { get; set; }

        [Input("state")]
        public string? State { get; set; }

        [Input("vrackEligibility")]
        public bool? VrackEligibility { get; set; }

        [Input("vrackName")]
        public string? VrackName { get; set; }

        [Input("zones")]
        private List<string>? _zones;
        public List<string> Zones
        {
            get => _zones ?? (_zones = new List<string>());
            set => _zones = value;
        }

        public GetIploadbalancingArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIploadbalancingResult
    {
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IpLoadbalancing;
        public readonly string Ipv4;
        public readonly string Ipv6;
        public readonly string MetricsToken;
        public readonly string Offer;
        public readonly ImmutableArray<Outputs.GetIploadbalancingOrderableZoneResult> OrderableZones;
        public readonly string ServiceName;
        public readonly string SslConfiguration;
        public readonly string State;
        public readonly bool VrackEligibility;
        public readonly string VrackName;
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetIploadbalancingResult(
            string displayName,

            string id,

            string ipLoadbalancing,

            string ipv4,

            string ipv6,

            string metricsToken,

            string offer,

            ImmutableArray<Outputs.GetIploadbalancingOrderableZoneResult> orderableZones,

            string serviceName,

            string sslConfiguration,

            string state,

            bool vrackEligibility,

            string vrackName,

            ImmutableArray<string> zones)
        {
            DisplayName = displayName;
            Id = id;
            IpLoadbalancing = ipLoadbalancing;
            Ipv4 = ipv4;
            Ipv6 = ipv6;
            MetricsToken = metricsToken;
            Offer = offer;
            OrderableZones = orderableZones;
            ServiceName = serviceName;
            SslConfiguration = sslConfiguration;
            State = state;
            VrackEligibility = vrackEligibility;
            VrackName = vrackName;
            Zones = zones;
        }
    }
}
