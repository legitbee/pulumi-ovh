// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOvh_Cloud_Project_Containerregistry
    {
        public static Task<GetOvh_Cloud_Project_ContainerregistryResult> InvokeAsync(GetOvh_Cloud_Project_ContainerregistryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOvh_Cloud_Project_ContainerregistryResult>("ovh:index/getOvh_Cloud_Project_Containerregistry:getOvh_Cloud_Project_Containerregistry", args ?? new GetOvh_Cloud_Project_ContainerregistryArgs(), options.WithVersion());
    }


    public sealed class GetOvh_Cloud_Project_ContainerregistryArgs : Pulumi.InvokeArgs
    {
        [Input("registryId", required: true)]
        public string RegistryId { get; set; } = null!;

        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetOvh_Cloud_Project_ContainerregistryArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOvh_Cloud_Project_ContainerregistryResult
    {
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ProjectId;
        public readonly string Region;
        public readonly string RegistryId;
        public readonly string ServiceName;
        public readonly int Size;
        public readonly string Status;
        public readonly string UpdatedAt;
        public readonly string Url;
        public readonly string Version;

        [OutputConstructor]
        private GetOvh_Cloud_Project_ContainerregistryResult(
            string createdAt,

            string id,

            string name,

            string projectId,

            string region,

            string registryId,

            string serviceName,

            int size,

            string status,

            string updatedAt,

            string url,

            string version)
        {
            CreatedAt = createdAt;
            Id = id;
            Name = name;
            ProjectId = projectId;
            Region = region;
            RegistryId = registryId;
            ServiceName = serviceName;
            Size = size;
            Status = status;
            UpdatedAt = updatedAt;
            Url = url;
            Version = version;
        }
    }
}
