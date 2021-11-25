// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    [OvhResourceType("ovh:index/ipService:IpService")]
    public partial class IpService : Pulumi.CustomResource
    {
        [Output("canBeTerminated")]
        public Output<bool> CanBeTerminated { get; private set; } = null!;

        [Output("country")]
        public Output<string> Country { get; private set; } = null!;

        /// <summary>
        /// Custom description on your ip
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// Details about an Order
        /// </summary>
        [Output("orders")]
        public Output<ImmutableArray<Outputs.IpServiceOrder>> Orders { get; private set; } = null!;

        [Output("organisationId")]
        public Output<string> OrganisationId { get; private set; } = null!;

        /// <summary>
        /// Ovh Subsidiary
        /// </summary>
        [Output("ovhSubsidiary")]
        public Output<string> OvhSubsidiary { get; private set; } = null!;

        /// <summary>
        /// Ovh payment mode
        /// </summary>
        [Output("paymentMean")]
        public Output<string> PaymentMean { get; private set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Output("plan")]
        public Output<Outputs.IpServicePlan> Plan { get; private set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Output("planOptions")]
        public Output<ImmutableArray<Outputs.IpServicePlanOption>> PlanOptions { get; private set; } = null!;

        /// <summary>
        /// Routage information
        /// </summary>
        [Output("routedTos")]
        public Output<ImmutableArray<Outputs.IpServiceRoutedTo>> RoutedTos { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Possible values for ip type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IpService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpService(string name, IpServiceArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/ipService:IpService", name, args ?? new IpServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpService(string name, Input<string> id, IpServiceState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/ipService:IpService", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IpService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpService Get(string name, Input<string> id, IpServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new IpService(name, id, state, options);
        }
    }

    public sealed class IpServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom description on your ip
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Ovh Subsidiary
        /// </summary>
        [Input("ovhSubsidiary", required: true)]
        public Input<string> OvhSubsidiary { get; set; } = null!;

        /// <summary>
        /// Ovh payment mode
        /// </summary>
        [Input("paymentMean", required: true)]
        public Input<string> PaymentMean { get; set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Input("plan", required: true)]
        public Input<Inputs.IpServicePlanArgs> Plan { get; set; } = null!;

        [Input("planOptions")]
        private InputList<Inputs.IpServicePlanOptionArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.IpServicePlanOptionArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.IpServicePlanOptionArgs>());
            set => _planOptions = value;
        }

        public IpServiceArgs()
        {
        }
    }

    public sealed class IpServiceState : Pulumi.ResourceArgs
    {
        [Input("canBeTerminated")]
        public Input<bool>? CanBeTerminated { get; set; }

        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Custom description on your ip
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("orders")]
        private InputList<Inputs.IpServiceOrderGetArgs>? _orders;

        /// <summary>
        /// Details about an Order
        /// </summary>
        public InputList<Inputs.IpServiceOrderGetArgs> Orders
        {
            get => _orders ?? (_orders = new InputList<Inputs.IpServiceOrderGetArgs>());
            set => _orders = value;
        }

        [Input("organisationId")]
        public Input<string>? OrganisationId { get; set; }

        /// <summary>
        /// Ovh Subsidiary
        /// </summary>
        [Input("ovhSubsidiary")]
        public Input<string>? OvhSubsidiary { get; set; }

        /// <summary>
        /// Ovh payment mode
        /// </summary>
        [Input("paymentMean")]
        public Input<string>? PaymentMean { get; set; }

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Input("plan")]
        public Input<Inputs.IpServicePlanGetArgs>? Plan { get; set; }

        [Input("planOptions")]
        private InputList<Inputs.IpServicePlanOptionGetArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.IpServicePlanOptionGetArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.IpServicePlanOptionGetArgs>());
            set => _planOptions = value;
        }

        [Input("routedTos")]
        private InputList<Inputs.IpServiceRoutedToGetArgs>? _routedTos;

        /// <summary>
        /// Routage information
        /// </summary>
        public InputList<Inputs.IpServiceRoutedToGetArgs> RoutedTos
        {
            get => _routedTos ?? (_routedTos = new InputList<Inputs.IpServiceRoutedToGetArgs>());
            set => _routedTos = value;
        }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Possible values for ip type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public IpServiceState()
        {
        }
    }
}