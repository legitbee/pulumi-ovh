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
    public sealed class DbaasLogsInputConfigurationFlowgger
    {
        public readonly string LogFormat;
        public readonly string LogFraming;

        [OutputConstructor]
        private DbaasLogsInputConfigurationFlowgger(
            string logFormat,

            string logFraming)
        {
            LogFormat = logFormat;
            LogFraming = logFraming;
        }
    }
}