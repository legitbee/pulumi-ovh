// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Inputs
{

    public sealed class Ovh_iploadbalancing_http_routeActionArgs : Pulumi.ResourceArgs
    {
        [Input("status")]
        public Input<int>? Status { get; set; }

        [Input("target")]
        public Input<string>? Target { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public Ovh_iploadbalancing_http_routeActionArgs()
        {
        }
    }
}
