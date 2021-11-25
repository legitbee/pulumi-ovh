// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetMePaymentmeanBankaccount
    {
        public static Task<GetMePaymentmeanBankaccountResult> InvokeAsync(GetMePaymentmeanBankaccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMePaymentmeanBankaccountResult>("ovh:index/getMePaymentmeanBankaccount:getMePaymentmeanBankaccount", args ?? new GetMePaymentmeanBankaccountArgs(), options.WithVersion());
    }


    public sealed class GetMePaymentmeanBankaccountArgs : Pulumi.InvokeArgs
    {
        [Input("descriptionRegexp")]
        public string? DescriptionRegexp { get; set; }

        [Input("state")]
        public string? State { get; set; }

        [Input("useDefault")]
        public bool? UseDefault { get; set; }

        [Input("useOldest")]
        public bool? UseOldest { get; set; }

        public GetMePaymentmeanBankaccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMePaymentmeanBankaccountResult
    {
        public readonly bool Default;
        public readonly string Description;
        public readonly string? DescriptionRegexp;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string State;
        public readonly bool? UseDefault;
        public readonly bool? UseOldest;

        [OutputConstructor]
        private GetMePaymentmeanBankaccountResult(
            bool @default,

            string description,

            string? descriptionRegexp,

            string id,

            string state,

            bool? useDefault,

            bool? useOldest)
        {
            Default = @default;
            Description = description;
            DescriptionRegexp = descriptionRegexp;
            Id = id;
            State = state;
            UseDefault = useDefault;
            UseOldest = useOldest;
        }
    }
}