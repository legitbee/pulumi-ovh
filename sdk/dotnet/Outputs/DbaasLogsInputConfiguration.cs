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
    public sealed class DbaasLogsInputConfiguration
    {
        public readonly Outputs.DbaasLogsInputConfigurationFlowgger? Flowgger;
        public readonly Outputs.DbaasLogsInputConfigurationLogstash? Logstash;

        [OutputConstructor]
        private DbaasLogsInputConfiguration(
            Outputs.DbaasLogsInputConfigurationFlowgger? flowgger,

            Outputs.DbaasLogsInputConfigurationLogstash? logstash)
        {
            Flowgger = flowgger;
            Logstash = logstash;
        }
    }
}
