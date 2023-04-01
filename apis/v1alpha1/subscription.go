// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SubscriptionSpec defines the desired state of Subscription.
//
// A wrapper type for the attributes of an Amazon SNS subscription.
type SubscriptionSpec struct {
	DeliveryPolicy *string `json:"deliveryPolicy,omitempty"`
	// The endpoint that you want to receive notifications. Endpoints vary by protocol:
	//
	//   - For the http protocol, the (public) endpoint is a URL beginning with
	//     http://.
	//
	//   - For the https protocol, the (public) endpoint is a URL beginning with
	//     https://.
	//
	//   - For the email protocol, the endpoint is an email address.
	//
	//   - For the email-json protocol, the endpoint is an email address.
	//
	//   - For the sms protocol, the endpoint is a phone number of an SMS-enabled
	//     device.
	//
	//   - For the sqs protocol, the endpoint is the ARN of an Amazon SQS queue.
	//
	//   - For the application protocol, the endpoint is the EndpointArn of a mobile
	//     app and device.
	//
	//   - For the lambda protocol, the endpoint is the ARN of an Lambda function.
	//
	//   - For the firehose protocol, the endpoint is the ARN of an Amazon Kinesis
	//     Data Firehose delivery stream.
	Endpoint     *string `json:"endpoint,omitempty"`
	FilterPolicy *string `json:"filterPolicy,omitempty"`
	// The protocol that you want to use. Supported protocols include:
	//
	//   - http – delivery of JSON-encoded message via HTTP POST
	//
	//   - https – delivery of JSON-encoded message via HTTPS POST
	//
	//   - email – delivery of message via SMTP
	//
	//   - email-json – delivery of JSON-encoded message via SMTP
	//
	//   - sms – delivery of message via SMS
	//
	//   - sqs – delivery of JSON-encoded message to an Amazon SQS queue
	//
	//   - application – delivery of JSON-encoded message to an EndpointArn for
	//     a mobile app and device
	//
	//   - lambda – delivery of JSON-encoded message to an Lambda function
	//
	//   - firehose – delivery of JSON-encoded message to an Amazon Kinesis Data
	//     Firehose delivery stream.
	//
	// +kubebuilder:validation:Required
	Protocol            *string `json:"protocol"`
	RawMessageDelivery  *string `json:"rawMessageDelivery,omitempty"`
	RedrivePolicy       *string `json:"redrivePolicy,omitempty"`
	SubscriptionRoleARN *string `json:"subscriptionRoleARN,omitempty"`
	// The ARN of the topic you want to subscribe to.
	TopicARN *string                                  `json:"topicARN,omitempty"`
	TopicRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"topicRef,omitempty"`
}

// SubscriptionStatus defines the observed state of Subscription
type SubscriptionStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// +kubebuilder:validation:Optional
	ConfirmationWasAuthenticated *string `json:"confirmationWasAuthenticated,omitempty"`
	// +kubebuilder:validation:Optional
	EffectiveDeliveryPolicy *string `json:"effectiveDeliveryPolicy,omitempty"`
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty"`
	// +kubebuilder:validation:Optional
	PendingConfirmation *string `json:"pendingConfirmation,omitempty"`
}

// Subscription is the Schema for the Subscriptions API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionSpec   `json:"spec,omitempty"`
	Status            SubscriptionStatus `json:"status,omitempty"`
}

// SubscriptionList contains a list of Subscription
// +kubebuilder:object:root=true
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}