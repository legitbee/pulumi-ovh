// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOvh_Cloud_Project_Containerregistries
    {
        public static Task<GetOvh_Cloud_Project_ContainerregistriesResult> InvokeAsync(GetOvh_Cloud_Project_ContainerregistriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOvh_Cloud_Project_ContainerregistriesResult>("ovh:index/getOvh_Cloud_Project_Containerregistries:getOvh_Cloud_Project_Containerregistries", args ?? new GetOvh_Cloud_Project_ContainerregistriesArgs(), options.WithVersion());
    }


    public sealed class GetOvh_Cloud_Project_ContainerregistriesArgs : Pulumi.InvokeArgs
    {
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetOvh_Cloud_Project_ContainerregistriesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOvh_Cloud_Project_ContainerregistriesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetOvh_Cloud_Project_ContainerregistriesResultResult> Results;
        public readonly string ServiceName;

        [OutputConstructor]
        private GetOvh_Cloud_Project_ContainerregistriesResult(
            string id,

            ImmutableArray<Outputs.GetOvh_Cloud_Project_ContainerregistriesResultResult> results,

            string serviceName)
        {
            Id = id;
            Results = results;
            ServiceName = serviceName;
        }
    }
}
