// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetDbaasLogsOutputGraylogStream
    {
        public static Task<GetDbaasLogsOutputGraylogStreamResult> InvokeAsync(GetDbaasLogsOutputGraylogStreamArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDbaasLogsOutputGraylogStreamResult>("ovh:index/getDbaasLogsOutputGraylogStream:getDbaasLogsOutputGraylogStream", args ?? new GetDbaasLogsOutputGraylogStreamArgs(), options.WithVersion());
    }


    public sealed class GetDbaasLogsOutputGraylogStreamArgs : Pulumi.InvokeArgs
    {
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        [Input("title", required: true)]
        public string Title { get; set; } = null!;

        public GetDbaasLogsOutputGraylogStreamArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDbaasLogsOutputGraylogStreamResult
    {
        public readonly bool CanAlert;
        public readonly string ColdStorageCompression;
        public readonly string ColdStorageContent;
        public readonly bool ColdStorageEnabled;
        public readonly bool ColdStorageNotifyEnabled;
        public readonly int ColdStorageRetention;
        public readonly string ColdStorageTarget;
        public readonly string CreatedAt;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IndexingEnabled;
        public readonly int IndexingMaxSize;
        public readonly bool IndexingNotifyEnabled;
        public readonly bool IsEditable;
        public readonly bool IsShareable;
        public readonly int NbAlertCondition;
        public readonly int NbArchive;
        public readonly string ParentStreamId;
        public readonly bool PauseIndexingOnMaxSize;
        public readonly string RetentionId;
        public readonly string ServiceName;
        public readonly string StreamId;
        public readonly string Title;
        public readonly string UpdatedAt;
        public readonly string WebSocketEnabled;

        [OutputConstructor]
        private GetDbaasLogsOutputGraylogStreamResult(
            bool canAlert,

            string coldStorageCompression,

            string coldStorageContent,

            bool coldStorageEnabled,

            bool coldStorageNotifyEnabled,

            int coldStorageRetention,

            string coldStorageTarget,

            string createdAt,

            string description,

            string id,

            bool indexingEnabled,

            int indexingMaxSize,

            bool indexingNotifyEnabled,

            bool isEditable,

            bool isShareable,

            int nbAlertCondition,

            int nbArchive,

            string parentStreamId,

            bool pauseIndexingOnMaxSize,

            string retentionId,

            string serviceName,

            string streamId,

            string title,

            string updatedAt,

            string webSocketEnabled)
        {
            CanAlert = canAlert;
            ColdStorageCompression = coldStorageCompression;
            ColdStorageContent = coldStorageContent;
            ColdStorageEnabled = coldStorageEnabled;
            ColdStorageNotifyEnabled = coldStorageNotifyEnabled;
            ColdStorageRetention = coldStorageRetention;
            ColdStorageTarget = coldStorageTarget;
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            IndexingEnabled = indexingEnabled;
            IndexingMaxSize = indexingMaxSize;
            IndexingNotifyEnabled = indexingNotifyEnabled;
            IsEditable = isEditable;
            IsShareable = isShareable;
            NbAlertCondition = nbAlertCondition;
            NbArchive = nbArchive;
            ParentStreamId = parentStreamId;
            PauseIndexingOnMaxSize = pauseIndexingOnMaxSize;
            RetentionId = retentionId;
            ServiceName = serviceName;
            StreamId = streamId;
            Title = title;
            UpdatedAt = updatedAt;
            WebSocketEnabled = webSocketEnabled;
        }
    }
}