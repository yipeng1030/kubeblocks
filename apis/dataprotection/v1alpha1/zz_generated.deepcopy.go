//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupList) DeepCopyInto(out *BackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupList.
func (in *BackupList) DeepCopy() *BackupList {
	if in == nil {
		return nil
	}
	out := new(BackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLogStatus) DeepCopyInto(out *BackupLogStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.StopTime != nil {
		in, out := &in.StopTime, &out.StopTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLogStatus.
func (in *BackupLogStatus) DeepCopy() *BackupLogStatus {
	if in == nil {
		return nil
	}
	out := new(BackupLogStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicy) DeepCopyInto(out *BackupPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicy.
func (in *BackupPolicy) DeepCopy() *BackupPolicy {
	if in == nil {
		return nil
	}
	out := new(BackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyHook) DeepCopyInto(out *BackupPolicyHook) {
	*out = *in
	if in.PreCommands != nil {
		in, out := &in.PreCommands, &out.PreCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostCommands != nil {
		in, out := &in.PostCommands, &out.PostCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyHook.
func (in *BackupPolicyHook) DeepCopy() *BackupPolicyHook {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyList) DeepCopyInto(out *BackupPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyList.
func (in *BackupPolicyList) DeepCopy() *BackupPolicyList {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicySecret) DeepCopyInto(out *BackupPolicySecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicySecret.
func (in *BackupPolicySecret) DeepCopy() *BackupPolicySecret {
	if in == nil {
		return nil
	}
	out := new(BackupPolicySecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicySpec) DeepCopyInto(out *BackupPolicySpec) {
	*out = *in
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(RetentionSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Schedule.DeepCopyInto(&out.Schedule)
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SnapshotPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Datafile != nil {
		in, out := &in.Datafile, &out.Datafile
		*out = new(CommonBackupPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Logfile != nil {
		in, out := &in.Logfile, &out.Logfile
		*out = new(CommonBackupPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicySpec.
func (in *BackupPolicySpec) DeepCopy() *BackupPolicySpec {
	if in == nil {
		return nil
	}
	out := new(BackupPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyStatus) DeepCopyInto(out *BackupPolicyStatus) {
	*out = *in
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		*out = (*in).DeepCopy()
	}
	if in.LastSuccessfulTime != nil {
		in, out := &in.LastSuccessfulTime, &out.LastSuccessfulTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyStatus.
func (in *BackupPolicyStatus) DeepCopy() *BackupPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRepo) DeepCopyInto(out *BackupRepo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRepo.
func (in *BackupRepo) DeepCopy() *BackupRepo {
	if in == nil {
		return nil
	}
	out := new(BackupRepo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupRepo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRepoList) DeepCopyInto(out *BackupRepoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupRepo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRepoList.
func (in *BackupRepoList) DeepCopy() *BackupRepoList {
	if in == nil {
		return nil
	}
	out := new(BackupRepoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupRepoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRepoSpec) DeepCopyInto(out *BackupRepoSpec) {
	*out = *in
	out.VolumeCapacity = in.VolumeCapacity.DeepCopy()
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(corev1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRepoSpec.
func (in *BackupRepoSpec) DeepCopy() *BackupRepoSpec {
	if in == nil {
		return nil
	}
	out := new(BackupRepoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRepoStatus) DeepCopyInto(out *BackupRepoStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GeneratedCSIDriverSecret != nil {
		in, out := &in.GeneratedCSIDriverSecret, &out.GeneratedCSIDriverSecret
		*out = new(corev1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRepoStatus.
func (in *BackupRepoStatus) DeepCopy() *BackupRepoStatus {
	if in == nil {
		return nil
	}
	out := new(BackupRepoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSnapshotStatus) DeepCopyInto(out *BackupSnapshotStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSnapshotStatus.
func (in *BackupSnapshotStatus) DeepCopy() *BackupSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(BackupSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = (*in).DeepCopy()
	}
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.CompletionTimestamp != nil {
		in, out := &in.CompletionTimestamp, &out.CompletionTimestamp
		*out = (*in).DeepCopy()
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.AvailableReplicas != nil {
		in, out := &in.AvailableReplicas, &out.AvailableReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Manifests != nil {
		in, out := &in.Manifests, &out.Manifests
		*out = new(ManifestsStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatusUpdate) DeepCopyInto(out *BackupStatusUpdate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatusUpdate.
func (in *BackupStatusUpdate) DeepCopy() *BackupStatusUpdate {
	if in == nil {
		return nil
	}
	out := new(BackupStatusUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupTool) DeepCopyInto(out *BackupTool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupTool.
func (in *BackupTool) DeepCopy() *BackupTool {
	if in == nil {
		return nil
	}
	out := new(BackupTool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupTool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupToolList) DeepCopyInto(out *BackupToolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupTool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupToolList.
func (in *BackupToolList) DeepCopy() *BackupToolList {
	if in == nil {
		return nil
	}
	out := new(BackupToolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupToolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupToolManifestsStatus) DeepCopyInto(out *BackupToolManifestsStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupToolManifestsStatus.
func (in *BackupToolManifestsStatus) DeepCopy() *BackupToolManifestsStatus {
	if in == nil {
		return nil
	}
	out := new(BackupToolManifestsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupToolRestoreCommand) DeepCopyInto(out *BackupToolRestoreCommand) {
	*out = *in
	if in.RestoreCommands != nil {
		in, out := &in.RestoreCommands, &out.RestoreCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncrementalRestoreCommands != nil {
		in, out := &in.IncrementalRestoreCommands, &out.IncrementalRestoreCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupToolRestoreCommand.
func (in *BackupToolRestoreCommand) DeepCopy() *BackupToolRestoreCommand {
	if in == nil {
		return nil
	}
	out := new(BackupToolRestoreCommand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupToolSpec) DeepCopyInto(out *BackupToolSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]corev1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackupCommands != nil {
		in, out := &in.BackupCommands, &out.BackupCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncrementalBackupCommands != nil {
		in, out := &in.IncrementalBackupCommands, &out.IncrementalBackupCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Physical.DeepCopyInto(&out.Physical)
	if in.Logical != nil {
		in, out := &in.Logical, &out.Logical
		*out = new(LogicalConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupToolSpec.
func (in *BackupToolSpec) DeepCopy() *BackupToolSpec {
	if in == nil {
		return nil
	}
	out := new(BackupToolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupToolStatus) DeepCopyInto(out *BackupToolStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupToolStatus.
func (in *BackupToolStatus) DeepCopy() *BackupToolStatus {
	if in == nil {
		return nil
	}
	out := new(BackupToolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasePolicy) DeepCopyInto(out *BasePolicy) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
	if in.BackupStatusUpdates != nil {
		in, out := &in.BackupStatusUpdates, &out.BackupStatusUpdates
		*out = make([]BackupStatusUpdate, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasePolicy.
func (in *BasePolicy) DeepCopy() *BasePolicy {
	if in == nil {
		return nil
	}
	out := new(BasePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonBackupPolicy) DeepCopyInto(out *CommonBackupPolicy) {
	*out = *in
	in.BasePolicy.DeepCopyInto(&out.BasePolicy)
	in.PersistentVolumeClaim.DeepCopyInto(&out.PersistentVolumeClaim)
	if in.BackupRepoName != nil {
		in, out := &in.BackupRepoName, &out.BackupRepoName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonBackupPolicy.
func (in *CommonBackupPolicy) DeepCopy() *CommonBackupPolicy {
	if in == nil {
		return nil
	}
	out := new(CommonBackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalConfig) DeepCopyInto(out *LogicalConfig) {
	*out = *in
	in.BackupToolRestoreCommand.DeepCopyInto(&out.BackupToolRestoreCommand)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalConfig.
func (in *LogicalConfig) DeepCopy() *LogicalConfig {
	if in == nil {
		return nil
	}
	out := new(LogicalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestsStatus) DeepCopyInto(out *ManifestsStatus) {
	*out = *in
	if in.BackupLog != nil {
		in, out := &in.BackupLog, &out.BackupLog
		*out = new(BackupLogStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(BackupSnapshotStatus)
		**out = **in
	}
	if in.BackupTool != nil {
		in, out := &in.BackupTool, &out.BackupTool
		*out = new(BackupToolManifestsStatus)
		**out = **in
	}
	if in.UserContext != nil {
		in, out := &in.UserContext, &out.UserContext
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestsStatus.
func (in *ManifestsStatus) DeepCopy() *ManifestsStatus {
	if in == nil {
		return nil
	}
	out := new(ManifestsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaim) DeepCopyInto(out *PersistentVolumeClaim) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	out.InitCapacity = in.InitCapacity.DeepCopy()
	if in.PersistentVolumeConfigMap != nil {
		in, out := &in.PersistentVolumeConfigMap, &out.PersistentVolumeConfigMap
		*out = new(PersistentVolumeConfigMap)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaim.
func (in *PersistentVolumeClaim) DeepCopy() *PersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeConfigMap) DeepCopyInto(out *PersistentVolumeConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeConfigMap.
func (in *PersistentVolumeConfigMap) DeepCopy() *PersistentVolumeConfigMap {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreJob) DeepCopyInto(out *RestoreJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreJob.
func (in *RestoreJob) DeepCopy() *RestoreJob {
	if in == nil {
		return nil
	}
	out := new(RestoreJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RestoreJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreJobList) DeepCopyInto(out *RestoreJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RestoreJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreJobList.
func (in *RestoreJobList) DeepCopy() *RestoreJobList {
	if in == nil {
		return nil
	}
	out := new(RestoreJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RestoreJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreJobSpec) DeepCopyInto(out *RestoreJobSpec) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
	if in.TargetVolumes != nil {
		in, out := &in.TargetVolumes, &out.TargetVolumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TargetVolumeMounts != nil {
		in, out := &in.TargetVolumeMounts, &out.TargetVolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreJobSpec.
func (in *RestoreJobSpec) DeepCopy() *RestoreJobSpec {
	if in == nil {
		return nil
	}
	out := new(RestoreJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreJobStatus) DeepCopyInto(out *RestoreJobStatus) {
	*out = *in
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = (*in).DeepCopy()
	}
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.CompletionTimestamp != nil {
		in, out := &in.CompletionTimestamp, &out.CompletionTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreJobStatus.
func (in *RestoreJobStatus) DeepCopy() *RestoreJobStatus {
	if in == nil {
		return nil
	}
	out := new(RestoreJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionSpec) DeepCopyInto(out *RetentionSpec) {
	*out = *in
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionSpec.
func (in *RetentionSpec) DeepCopy() *RetentionSpec {
	if in == nil {
		return nil
	}
	out := new(RetentionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Schedule) DeepCopyInto(out *Schedule) {
	*out = *in
	if in.StartWindowMinutes != nil {
		in, out := &in.StartWindowMinutes, &out.StartWindowMinutes
		*out = new(int64)
		**out = **in
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SchedulePolicy)
		**out = **in
	}
	if in.Datafile != nil {
		in, out := &in.Datafile, &out.Datafile
		*out = new(SchedulePolicy)
		**out = **in
	}
	if in.Logfile != nil {
		in, out := &in.Logfile, &out.Logfile
		*out = new(SchedulePolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Schedule.
func (in *Schedule) DeepCopy() *Schedule {
	if in == nil {
		return nil
	}
	out := new(Schedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulePolicy) DeepCopyInto(out *SchedulePolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulePolicy.
func (in *SchedulePolicy) DeepCopy() *SchedulePolicy {
	if in == nil {
		return nil
	}
	out := new(SchedulePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotPolicy) DeepCopyInto(out *SnapshotPolicy) {
	*out = *in
	in.BasePolicy.DeepCopyInto(&out.BasePolicy)
	if in.Hooks != nil {
		in, out := &in.Hooks, &out.Hooks
		*out = new(BackupPolicyHook)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotPolicy.
func (in *SnapshotPolicy) DeepCopy() *SnapshotPolicy {
	if in == nil {
		return nil
	}
	out := new(SnapshotPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetCluster) DeepCopyInto(out *TargetCluster) {
	*out = *in
	if in.LabelsSelector != nil {
		in, out := &in.LabelsSelector, &out.LabelsSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(BackupPolicySecret)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetCluster.
func (in *TargetCluster) DeepCopy() *TargetCluster {
	if in == nil {
		return nil
	}
	out := new(TargetCluster)
	in.DeepCopyInto(out)
	return out
}
