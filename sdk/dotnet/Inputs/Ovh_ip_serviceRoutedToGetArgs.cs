// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Inputs
{

    public sealed class Ovh_ip_serviceRoutedToGetArgs : Pulumi.ResourceArgs
    {
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public Ovh_ip_serviceRoutedToGetArgs()
        {
        }
    }
}