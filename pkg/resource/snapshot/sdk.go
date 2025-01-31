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

package snapshot

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/elasticache"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/elasticache-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ElastiCache{}
	_ = &svcapitypes.Snapshot{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeSnapshotsOutput
	resp, err = rm.sdkapi.DescribeSnapshotsWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeSnapshots", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "CacheClusterNotFound" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.Snapshots {
		if elem.ARN != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.ARN)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.AutoMinorVersionUpgrade != nil {
			ko.Status.AutoMinorVersionUpgrade = elem.AutoMinorVersionUpgrade
		} else {
			ko.Status.AutoMinorVersionUpgrade = nil
		}
		if elem.AutomaticFailover != nil {
			ko.Status.AutomaticFailover = elem.AutomaticFailover
		} else {
			ko.Status.AutomaticFailover = nil
		}
		if elem.CacheClusterCreateTime != nil {
			ko.Status.CacheClusterCreateTime = &metav1.Time{*elem.CacheClusterCreateTime}
		} else {
			ko.Status.CacheClusterCreateTime = nil
		}
		if elem.CacheClusterId != nil {
			ko.Spec.CacheClusterID = elem.CacheClusterId
		} else {
			ko.Spec.CacheClusterID = nil
		}
		if elem.CacheNodeType != nil {
			ko.Status.CacheNodeType = elem.CacheNodeType
		} else {
			ko.Status.CacheNodeType = nil
		}
		if elem.CacheParameterGroupName != nil {
			ko.Status.CacheParameterGroupName = elem.CacheParameterGroupName
		} else {
			ko.Status.CacheParameterGroupName = nil
		}
		if elem.CacheSubnetGroupName != nil {
			ko.Status.CacheSubnetGroupName = elem.CacheSubnetGroupName
		} else {
			ko.Status.CacheSubnetGroupName = nil
		}
		if elem.Engine != nil {
			ko.Status.Engine = elem.Engine
		} else {
			ko.Status.Engine = nil
		}
		if elem.EngineVersion != nil {
			ko.Status.EngineVersion = elem.EngineVersion
		} else {
			ko.Status.EngineVersion = nil
		}
		if elem.KmsKeyId != nil {
			ko.Spec.KMSKeyID = elem.KmsKeyId
		} else {
			ko.Spec.KMSKeyID = nil
		}
		if elem.NodeSnapshots != nil {
			f11 := []*svcapitypes.NodeSnapshot{}
			for _, f11iter := range elem.NodeSnapshots {
				f11elem := &svcapitypes.NodeSnapshot{}
				if f11iter.CacheClusterId != nil {
					f11elem.CacheClusterID = f11iter.CacheClusterId
				}
				if f11iter.CacheNodeCreateTime != nil {
					f11elem.CacheNodeCreateTime = &metav1.Time{*f11iter.CacheNodeCreateTime}
				}
				if f11iter.CacheNodeId != nil {
					f11elem.CacheNodeID = f11iter.CacheNodeId
				}
				if f11iter.CacheSize != nil {
					f11elem.CacheSize = f11iter.CacheSize
				}
				if f11iter.NodeGroupConfiguration != nil {
					f11elemf4 := &svcapitypes.NodeGroupConfiguration{}
					if f11iter.NodeGroupConfiguration.NodeGroupId != nil {
						f11elemf4.NodeGroupID = f11iter.NodeGroupConfiguration.NodeGroupId
					}
					if f11iter.NodeGroupConfiguration.PrimaryAvailabilityZone != nil {
						f11elemf4.PrimaryAvailabilityZone = f11iter.NodeGroupConfiguration.PrimaryAvailabilityZone
					}
					if f11iter.NodeGroupConfiguration.PrimaryOutpostArn != nil {
						f11elemf4.PrimaryOutpostARN = f11iter.NodeGroupConfiguration.PrimaryOutpostArn
					}
					if f11iter.NodeGroupConfiguration.ReplicaAvailabilityZones != nil {
						f11elemf4f3 := []*string{}
						for _, f11elemf4f3iter := range f11iter.NodeGroupConfiguration.ReplicaAvailabilityZones {
							var f11elemf4f3elem string
							f11elemf4f3elem = *f11elemf4f3iter
							f11elemf4f3 = append(f11elemf4f3, &f11elemf4f3elem)
						}
						f11elemf4.ReplicaAvailabilityZones = f11elemf4f3
					}
					if f11iter.NodeGroupConfiguration.ReplicaCount != nil {
						f11elemf4.ReplicaCount = f11iter.NodeGroupConfiguration.ReplicaCount
					}
					if f11iter.NodeGroupConfiguration.ReplicaOutpostArns != nil {
						f11elemf4f5 := []*string{}
						for _, f11elemf4f5iter := range f11iter.NodeGroupConfiguration.ReplicaOutpostArns {
							var f11elemf4f5elem string
							f11elemf4f5elem = *f11elemf4f5iter
							f11elemf4f5 = append(f11elemf4f5, &f11elemf4f5elem)
						}
						f11elemf4.ReplicaOutpostARNs = f11elemf4f5
					}
					if f11iter.NodeGroupConfiguration.Slots != nil {
						f11elemf4.Slots = f11iter.NodeGroupConfiguration.Slots
					}
					f11elem.NodeGroupConfiguration = f11elemf4
				}
				if f11iter.NodeGroupId != nil {
					f11elem.NodeGroupID = f11iter.NodeGroupId
				}
				if f11iter.SnapshotCreateTime != nil {
					f11elem.SnapshotCreateTime = &metav1.Time{*f11iter.SnapshotCreateTime}
				}
				f11 = append(f11, f11elem)
			}
			ko.Status.NodeSnapshots = f11
		} else {
			ko.Status.NodeSnapshots = nil
		}
		if elem.NumCacheNodes != nil {
			ko.Status.NumCacheNodes = elem.NumCacheNodes
		} else {
			ko.Status.NumCacheNodes = nil
		}
		if elem.NumNodeGroups != nil {
			ko.Status.NumNodeGroups = elem.NumNodeGroups
		} else {
			ko.Status.NumNodeGroups = nil
		}
		if elem.Port != nil {
			ko.Status.Port = elem.Port
		} else {
			ko.Status.Port = nil
		}
		if elem.PreferredAvailabilityZone != nil {
			ko.Status.PreferredAvailabilityZone = elem.PreferredAvailabilityZone
		} else {
			ko.Status.PreferredAvailabilityZone = nil
		}
		if elem.PreferredMaintenanceWindow != nil {
			ko.Status.PreferredMaintenanceWindow = elem.PreferredMaintenanceWindow
		} else {
			ko.Status.PreferredMaintenanceWindow = nil
		}
		if elem.PreferredOutpostArn != nil {
			ko.Status.PreferredOutpostARN = elem.PreferredOutpostArn
		} else {
			ko.Status.PreferredOutpostARN = nil
		}
		if elem.ReplicationGroupDescription != nil {
			ko.Status.ReplicationGroupDescription = elem.ReplicationGroupDescription
		} else {
			ko.Status.ReplicationGroupDescription = nil
		}
		if elem.ReplicationGroupId != nil {
			ko.Spec.ReplicationGroupID = elem.ReplicationGroupId
		} else {
			ko.Spec.ReplicationGroupID = nil
		}
		if elem.SnapshotName != nil {
			ko.Spec.SnapshotName = elem.SnapshotName
		} else {
			ko.Spec.SnapshotName = nil
		}
		if elem.SnapshotRetentionLimit != nil {
			ko.Status.SnapshotRetentionLimit = elem.SnapshotRetentionLimit
		} else {
			ko.Status.SnapshotRetentionLimit = nil
		}
		if elem.SnapshotSource != nil {
			ko.Status.SnapshotSource = elem.SnapshotSource
		} else {
			ko.Status.SnapshotSource = nil
		}
		if elem.SnapshotStatus != nil {
			ko.Status.SnapshotStatus = elem.SnapshotStatus
		} else {
			ko.Status.SnapshotStatus = nil
		}
		if elem.SnapshotWindow != nil {
			ko.Status.SnapshotWindow = elem.SnapshotWindow
		} else {
			ko.Status.SnapshotWindow = nil
		}
		if elem.TopicArn != nil {
			ko.Status.TopicARN = elem.TopicArn
		} else {
			ko.Status.TopicARN = nil
		}
		if elem.VpcId != nil {
			ko.Status.VPCID = elem.VpcId
		} else {
			ko.Status.VPCID = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	// custom set output from response
	ko, err = rm.CustomDescribeSnapshotSetOutput(ctx, r, resp, ko)
	if err != nil {
		return nil, err
	}
	return &resource{ko}, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeSnapshotsInput, error) {
	res := &svcsdk.DescribeSnapshotsInput{}

	if r.ko.Spec.SnapshotName != nil {
		res.SetSnapshotName(*r.ko.Spec.SnapshotName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	created, err = rm.CustomCreateSnapshot(ctx, desired)
	if created != nil || err != nil {
		return created, err
	}
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateSnapshotOutput
	_ = resp
	resp, err = rm.sdkapi.CreateSnapshotWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateSnapshot", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Snapshot.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Snapshot.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Snapshot.AutoMinorVersionUpgrade != nil {
		ko.Status.AutoMinorVersionUpgrade = resp.Snapshot.AutoMinorVersionUpgrade
	} else {
		ko.Status.AutoMinorVersionUpgrade = nil
	}
	if resp.Snapshot.AutomaticFailover != nil {
		ko.Status.AutomaticFailover = resp.Snapshot.AutomaticFailover
	} else {
		ko.Status.AutomaticFailover = nil
	}
	if resp.Snapshot.CacheClusterCreateTime != nil {
		ko.Status.CacheClusterCreateTime = &metav1.Time{*resp.Snapshot.CacheClusterCreateTime}
	} else {
		ko.Status.CacheClusterCreateTime = nil
	}
	if resp.Snapshot.CacheNodeType != nil {
		ko.Status.CacheNodeType = resp.Snapshot.CacheNodeType
	} else {
		ko.Status.CacheNodeType = nil
	}
	if resp.Snapshot.CacheParameterGroupName != nil {
		ko.Status.CacheParameterGroupName = resp.Snapshot.CacheParameterGroupName
	} else {
		ko.Status.CacheParameterGroupName = nil
	}
	if resp.Snapshot.CacheSubnetGroupName != nil {
		ko.Status.CacheSubnetGroupName = resp.Snapshot.CacheSubnetGroupName
	} else {
		ko.Status.CacheSubnetGroupName = nil
	}
	if resp.Snapshot.Engine != nil {
		ko.Status.Engine = resp.Snapshot.Engine
	} else {
		ko.Status.Engine = nil
	}
	if resp.Snapshot.EngineVersion != nil {
		ko.Status.EngineVersion = resp.Snapshot.EngineVersion
	} else {
		ko.Status.EngineVersion = nil
	}
	if resp.Snapshot.NodeSnapshots != nil {
		f11 := []*svcapitypes.NodeSnapshot{}
		for _, f11iter := range resp.Snapshot.NodeSnapshots {
			f11elem := &svcapitypes.NodeSnapshot{}
			if f11iter.CacheClusterId != nil {
				f11elem.CacheClusterID = f11iter.CacheClusterId
			}
			if f11iter.CacheNodeCreateTime != nil {
				f11elem.CacheNodeCreateTime = &metav1.Time{*f11iter.CacheNodeCreateTime}
			}
			if f11iter.CacheNodeId != nil {
				f11elem.CacheNodeID = f11iter.CacheNodeId
			}
			if f11iter.CacheSize != nil {
				f11elem.CacheSize = f11iter.CacheSize
			}
			if f11iter.NodeGroupConfiguration != nil {
				f11elemf4 := &svcapitypes.NodeGroupConfiguration{}
				if f11iter.NodeGroupConfiguration.NodeGroupId != nil {
					f11elemf4.NodeGroupID = f11iter.NodeGroupConfiguration.NodeGroupId
				}
				if f11iter.NodeGroupConfiguration.PrimaryAvailabilityZone != nil {
					f11elemf4.PrimaryAvailabilityZone = f11iter.NodeGroupConfiguration.PrimaryAvailabilityZone
				}
				if f11iter.NodeGroupConfiguration.PrimaryOutpostArn != nil {
					f11elemf4.PrimaryOutpostARN = f11iter.NodeGroupConfiguration.PrimaryOutpostArn
				}
				if f11iter.NodeGroupConfiguration.ReplicaAvailabilityZones != nil {
					f11elemf4f3 := []*string{}
					for _, f11elemf4f3iter := range f11iter.NodeGroupConfiguration.ReplicaAvailabilityZones {
						var f11elemf4f3elem string
						f11elemf4f3elem = *f11elemf4f3iter
						f11elemf4f3 = append(f11elemf4f3, &f11elemf4f3elem)
					}
					f11elemf4.ReplicaAvailabilityZones = f11elemf4f3
				}
				if f11iter.NodeGroupConfiguration.ReplicaCount != nil {
					f11elemf4.ReplicaCount = f11iter.NodeGroupConfiguration.ReplicaCount
				}
				if f11iter.NodeGroupConfiguration.ReplicaOutpostArns != nil {
					f11elemf4f5 := []*string{}
					for _, f11elemf4f5iter := range f11iter.NodeGroupConfiguration.ReplicaOutpostArns {
						var f11elemf4f5elem string
						f11elemf4f5elem = *f11elemf4f5iter
						f11elemf4f5 = append(f11elemf4f5, &f11elemf4f5elem)
					}
					f11elemf4.ReplicaOutpostARNs = f11elemf4f5
				}
				if f11iter.NodeGroupConfiguration.Slots != nil {
					f11elemf4.Slots = f11iter.NodeGroupConfiguration.Slots
				}
				f11elem.NodeGroupConfiguration = f11elemf4
			}
			if f11iter.NodeGroupId != nil {
				f11elem.NodeGroupID = f11iter.NodeGroupId
			}
			if f11iter.SnapshotCreateTime != nil {
				f11elem.SnapshotCreateTime = &metav1.Time{*f11iter.SnapshotCreateTime}
			}
			f11 = append(f11, f11elem)
		}
		ko.Status.NodeSnapshots = f11
	} else {
		ko.Status.NodeSnapshots = nil
	}
	if resp.Snapshot.NumCacheNodes != nil {
		ko.Status.NumCacheNodes = resp.Snapshot.NumCacheNodes
	} else {
		ko.Status.NumCacheNodes = nil
	}
	if resp.Snapshot.NumNodeGroups != nil {
		ko.Status.NumNodeGroups = resp.Snapshot.NumNodeGroups
	} else {
		ko.Status.NumNodeGroups = nil
	}
	if resp.Snapshot.Port != nil {
		ko.Status.Port = resp.Snapshot.Port
	} else {
		ko.Status.Port = nil
	}
	if resp.Snapshot.PreferredAvailabilityZone != nil {
		ko.Status.PreferredAvailabilityZone = resp.Snapshot.PreferredAvailabilityZone
	} else {
		ko.Status.PreferredAvailabilityZone = nil
	}
	if resp.Snapshot.PreferredMaintenanceWindow != nil {
		ko.Status.PreferredMaintenanceWindow = resp.Snapshot.PreferredMaintenanceWindow
	} else {
		ko.Status.PreferredMaintenanceWindow = nil
	}
	if resp.Snapshot.PreferredOutpostArn != nil {
		ko.Status.PreferredOutpostARN = resp.Snapshot.PreferredOutpostArn
	} else {
		ko.Status.PreferredOutpostARN = nil
	}
	if resp.Snapshot.ReplicationGroupDescription != nil {
		ko.Status.ReplicationGroupDescription = resp.Snapshot.ReplicationGroupDescription
	} else {
		ko.Status.ReplicationGroupDescription = nil
	}
	if resp.Snapshot.SnapshotRetentionLimit != nil {
		ko.Status.SnapshotRetentionLimit = resp.Snapshot.SnapshotRetentionLimit
	} else {
		ko.Status.SnapshotRetentionLimit = nil
	}
	if resp.Snapshot.SnapshotSource != nil {
		ko.Status.SnapshotSource = resp.Snapshot.SnapshotSource
	} else {
		ko.Status.SnapshotSource = nil
	}
	if resp.Snapshot.SnapshotStatus != nil {
		ko.Status.SnapshotStatus = resp.Snapshot.SnapshotStatus
	} else {
		ko.Status.SnapshotStatus = nil
	}
	if resp.Snapshot.SnapshotWindow != nil {
		ko.Status.SnapshotWindow = resp.Snapshot.SnapshotWindow
	} else {
		ko.Status.SnapshotWindow = nil
	}
	if resp.Snapshot.TopicArn != nil {
		ko.Status.TopicARN = resp.Snapshot.TopicArn
	} else {
		ko.Status.TopicARN = nil
	}
	if resp.Snapshot.VpcId != nil {
		ko.Status.VPCID = resp.Snapshot.VpcId
	} else {
		ko.Status.VPCID = nil
	}

	rm.setStatusDefaults(ko)
	// custom set output from response
	ko, err = rm.CustomCreateSnapshotSetOutput(ctx, desired, resp, ko)
	if err != nil {
		return nil, err
	}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateSnapshotInput, error) {
	res := &svcsdk.CreateSnapshotInput{}

	if r.ko.Spec.CacheClusterID != nil {
		res.SetCacheClusterId(*r.ko.Spec.CacheClusterID)
	}
	if r.ko.Spec.KMSKeyID != nil {
		res.SetKmsKeyId(*r.ko.Spec.KMSKeyID)
	}
	if r.ko.Spec.ReplicationGroupID != nil {
		res.SetReplicationGroupId(*r.ko.Spec.ReplicationGroupID)
	}
	if r.ko.Spec.SnapshotName != nil {
		res.SetSnapshotName(*r.ko.Spec.SnapshotName)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	return rm.customUpdateSnapshot(ctx, desired, latest, delta)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteSnapshotOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteSnapshotWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteSnapshot", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteSnapshotInput, error) {
	res := &svcsdk.DeleteSnapshotInput{}

	if r.ko.Spec.SnapshotName != nil {
		res.SetSnapshotName(*r.ko.Spec.SnapshotName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Snapshot,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Message()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Message()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	// custom update conditions
	customUpdate := rm.CustomUpdateConditions(ko, r, err)
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil || customUpdate {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "InvalidParameter",
		"InvalidParameterValue",
		"InvalidParameterCombination",
		"SnapshotAlreadyExistsFault",
		"CacheClusterNotFound",
		"ReplicationGroupNotFoundFault",
		"SnapshotQuotaExceededFault",
		"SnapshotFeatureNotSupportedFault":
		return true
	default:
		return false
	}
}
