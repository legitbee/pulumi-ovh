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
    public sealed class GetOvh_Me_Installation_TemplatePartitionSchemePartitionResult
    {
        public readonly string Filesystem;
        public readonly string Mountpoint;
        public readonly int Order;
        public readonly string Raid;
        public readonly int Size;
        public readonly string Type;
        public readonly string VolumeName;

        [OutputConstructor]
        private GetOvh_Me_Installation_TemplatePartitionSchemePartitionResult(
            string filesystem,

            string mountpoint,

            int order,

            string raid,

            int size,

            string type,

            string volumeName)
        {
            Filesystem = filesystem;
            Mountpoint = mountpoint;
            Order = order;
            Raid = raid;
            Size = size;
            Type = type;
            VolumeName = volumeName;
        }
    }
}
