// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOvh_Dedicated_Ceph
    {
        public static Task<GetOvh_Dedicated_CephResult> InvokeAsync(GetOvh_Dedicated_CephArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOvh_Dedicated_CephResult>("ovh:index/getOvh_Dedicated_Ceph:getOvh_Dedicated_Ceph", args ?? new GetOvh_Dedicated_CephArgs(), options.WithVersion());
    }


    public sealed class GetOvh_Dedicated_CephArgs : Pulumi.InvokeArgs
    {
        [Input("cephVersion")]
        public string? CephVersion { get; set; }

        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        [Input("status")]
        public string? Status { get; set; }

        public GetOvh_Dedicated_CephArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOvh_Dedicated_CephResult
    {
        public readonly ImmutableArray<string> CephMons;
        public readonly string CephVersion;
        public readonly string CrushTunables;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Label;
        public readonly string Region;
        public readonly string ServiceName;
        public readonly double Size;
        public readonly string State;
        public readonly string Status;

        [OutputConstructor]
        private GetOvh_Dedicated_CephResult(
            ImmutableArray<string> cephMons,

            string cephVersion,

            string crushTunables,

            string id,

            string label,

            string region,

            string serviceName,

            double size,

            string state,

            string status)
        {
            CephMons = cephMons;
            CephVersion = cephVersion;
            CrushTunables = crushTunables;
            Id = id;
            Label = label;
            Region = region;
            ServiceName = serviceName;
            Size = size;
            State = state;
            Status = status;
        }
    }
}