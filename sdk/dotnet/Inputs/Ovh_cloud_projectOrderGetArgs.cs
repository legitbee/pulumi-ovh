// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Inputs
{

    public sealed class Ovh_cloud_projectOrderGetArgs : Pulumi.ResourceArgs
    {
        [Input("date")]
        public Input<string>? Date { get; set; }

        [Input("details")]
        private InputList<Inputs.Ovh_cloud_projectOrderDetailGetArgs>? _details;
        public InputList<Inputs.Ovh_cloud_projectOrderDetailGetArgs> Details
        {
            get => _details ?? (_details = new InputList<Inputs.Ovh_cloud_projectOrderDetailGetArgs>());
            set => _details = value;
        }

        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        [Input("orderId")]
        public Input<int>? OrderId { get; set; }

        public Ovh_cloud_projectOrderGetArgs()
        {
        }
    }
}
