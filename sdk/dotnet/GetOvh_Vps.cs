// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOvh_Vps
    {
        public static Task<GetOvh_VpsResult> InvokeAsync(GetOvh_VpsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOvh_VpsResult>("ovh:index/getOvh_Vps:getOvh_Vps", args ?? new GetOvh_VpsArgs(), options.WithVersion());
    }


    public sealed class GetOvh_VpsArgs : Pulumi.InvokeArgs
    {
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetOvh_VpsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOvh_VpsResult
    {
        public readonly string Cluster;
        public readonly ImmutableDictionary<string, string> Datacenter;
        public readonly string Displayname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ips;
        public readonly string Keymap;
        public readonly int Memory;
        public readonly ImmutableDictionary<string, string> Model;
        public readonly string Name;
        public readonly string Netbootmode;
        public readonly string Offertype;
        public readonly string ServiceName;
        public readonly bool Slamonitoring;
        public readonly string State;
        public readonly string Type;
        public readonly int Vcore;
        public readonly string Zone;

        [OutputConstructor]
        private GetOvh_VpsResult(
            string cluster,

            ImmutableDictionary<string, string> datacenter,

            string displayname,

            string id,

            ImmutableArray<string> ips,

            string keymap,

            int memory,

            ImmutableDictionary<string, string> model,

            string name,

            string netbootmode,

            string offertype,

            string serviceName,

            bool slamonitoring,

            string state,

            string type,

            int vcore,

            string zone)
        {
            Cluster = cluster;
            Datacenter = datacenter;
            Displayname = displayname;
            Id = id;
            Ips = ips;
            Keymap = keymap;
            Memory = memory;
            Model = model;
            Name = name;
            Netbootmode = netbootmode;
            Offertype = offertype;
            ServiceName = serviceName;
            Slamonitoring = slamonitoring;
            State = state;
            Type = type;
            Vcore = vcore;
            Zone = zone;
        }
    }
}