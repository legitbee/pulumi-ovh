// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Inputs
{

    public sealed class Ovh_dbaas_logs_inputConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("flowgger")]
        public Input<Inputs.Ovh_dbaas_logs_inputConfigurationFlowggerArgs>? Flowgger { get; set; }

        [Input("logstash")]
        public Input<Inputs.Ovh_dbaas_logs_inputConfigurationLogstashArgs>? Logstash { get; set; }

        public Ovh_dbaas_logs_inputConfigurationArgs()
        {
        }
    }
}
