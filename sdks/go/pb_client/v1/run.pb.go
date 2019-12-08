// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/run.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Run kind enum
type RunKind int32

const (
	RunKind_job      RunKind = 0
	RunKind_service  RunKind = 1
	RunKind_dag      RunKind = 2
	RunKind_parallel RunKind = 3
)

var RunKind_name = map[int32]string{
	0: "job",
	1: "service",
	2: "dag",
	3: "parallel",
}

var RunKind_value = map[string]int32{
	"job":      0,
	"service":  1,
	"dag":      2,
	"parallel": 3,
}

func (x RunKind) String() string {
	return proto.EnumName(RunKind_name, int32(x))
}

func (RunKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{0}
}

// Run meta info specification
type RunMetaInfo struct {
	// Optional flag to tell if the run has a service
	Service bool `protobuf:"varint,1,opt,name=service,proto3" json:"service,omitempty"`
	// Optional an indicator if the run has a concurrency
	Concurrency int32 `protobuf:"varint,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	// Optional the parallel kind
	ParallelKind string `protobuf:"bytes,3,opt,name=parallel_kind,json=parallelKind,proto3" json:"parallel_kind,omitempty"`
	// Optional the run kind
	RunKind              string   `protobuf:"bytes,4,opt,name=run_kind,json=runKind,proto3" json:"run_kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunMetaInfo) Reset()         { *m = RunMetaInfo{} }
func (m *RunMetaInfo) String() string { return proto.CompactTextString(m) }
func (*RunMetaInfo) ProtoMessage()    {}
func (*RunMetaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{0}
}

func (m *RunMetaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunMetaInfo.Unmarshal(m, b)
}
func (m *RunMetaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunMetaInfo.Marshal(b, m, deterministic)
}
func (m *RunMetaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunMetaInfo.Merge(m, src)
}
func (m *RunMetaInfo) XXX_Size() int {
	return xxx_messageInfo_RunMetaInfo.Size(m)
}
func (m *RunMetaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RunMetaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RunMetaInfo proto.InternalMessageInfo

func (m *RunMetaInfo) GetService() bool {
	if m != nil {
		return m.Service
	}
	return false
}

func (m *RunMetaInfo) GetConcurrency() int32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func (m *RunMetaInfo) GetParallelKind() string {
	if m != nil {
		return m.ParallelKind
	}
	return ""
}

func (m *RunMetaInfo) GetRunKind() string {
	if m != nil {
		return m.RunKind
	}
	return ""
}

// Run specification
type Run struct {
	// UUID
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Optional name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Optional description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Optional Tags of this entity
	Tags []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	// Optional if the entity has been deleted
	Deleted bool `protobuf:"varint,5,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Required name of user started this entity
	User string `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	// Required name of owner of this entity
	Owner string `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	// Required project name
	Project string `protobuf:"bytes,8,opt,name=project,proto3" json:"project,omitempty"`
	// Optional time when the entityt was created
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Optional last time the entity was updated
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Optional last time the entity was started
	StartedAt *timestamp.Timestamp `protobuf:"bytes,11,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// Optional last time the entity was started
	FinishedAt *timestamp.Timestamp `protobuf:"bytes,12,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	// Optional flag to tell if this entity is managed by the platform
	IsManaged string `protobuf:"bytes,13,opt,name=is_managed,json=isManaged,proto3" json:"is_managed,omitempty"`
	// Optional content of the entity's spec
	Content string `protobuf:"bytes,14,opt,name=content,proto3" json:"content,omitempty"`
	// Optional latest status of this entity
	Status string `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	// Optional a readme text describing this entity
	Readme string `protobuf:"bytes,16,opt,name=readme,proto3" json:"readme,omitempty"`
	// Optional if this entity was bookmarked
	Bookmarked bool `protobuf:"varint,17,opt,name=bookmarked,proto3" json:"bookmarked,omitempty"`
	// Optional run meta info
	MetaInfo *RunMetaInfo `protobuf:"bytes,18,opt,name=meta_info,json=metaInfo,proto3" json:"meta_info,omitempty"`
	// Optional kind to tell the nature of this run
	Kind RunKind `protobuf:"varint,19,opt,name=kind,proto3,enum=v1.RunKind" json:"kind,omitempty"`
	// Optional hub reference versioned component
	Hub string `protobuf:"bytes,20,opt,name=hub,proto3" json:"hub,omitempty"`
	// Optional inputs of this entity
	Inputs *_struct.Struct `protobuf:"bytes,21,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// Optional outputs of this entity
	Outputs *_struct.Struct `protobuf:"bytes,22,opt,name=outputs,proto3" json:"outputs,omitempty"`
	// Optional run environment tracked
	RunEnv *_struct.Struct `protobuf:"bytes,23,opt,name=run_env,json=runEnv,proto3" json:"run_env,omitempty"`
	// Is resume
	IsResume bool `protobuf:"varint,24,opt,name=is_resume,json=isResume,proto3" json:"is_resume,omitempty"`
	// Is clone
	IsClone bool `protobuf:"varint,25,opt,name=is_clone,json=isClone,proto3" json:"is_clone,omitempty"`
	// Optional if this run was restarted/copied/resumed/cached
	CloningStrategy string `protobuf:"bytes,26,opt,name=cloning_strategy,json=cloningStrategy,proto3" json:"cloning_strategy,omitempty"`
	// Optional uuid of the pipeline
	Pipeline string `protobuf:"bytes,27,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	// Optional uuid of the original run
	Original string `protobuf:"bytes,28,opt,name=original,proto3" json:"original,omitempty"`
	// Optional name of the pipeline
	PipelineName string `protobuf:"bytes,29,opt,name=pipeline_name,json=pipelineName,proto3" json:"pipeline_name,omitempty"`
	// Optional name of the original run
	OriginalName         string   `protobuf:"bytes,30,opt,name=original_name,json=originalName,proto3" json:"original_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Run) Reset()         { *m = Run{} }
func (m *Run) String() string { return proto.CompactTextString(m) }
func (*Run) ProtoMessage()    {}
func (*Run) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{1}
}

func (m *Run) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Run.Unmarshal(m, b)
}
func (m *Run) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Run.Marshal(b, m, deterministic)
}
func (m *Run) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Run.Merge(m, src)
}
func (m *Run) XXX_Size() int {
	return xxx_messageInfo_Run.Size(m)
}
func (m *Run) XXX_DiscardUnknown() {
	xxx_messageInfo_Run.DiscardUnknown(m)
}

var xxx_messageInfo_Run proto.InternalMessageInfo

func (m *Run) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Run) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Run) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Run) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Run) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *Run) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Run) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Run) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Run) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Run) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Run) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *Run) GetFinishedAt() *timestamp.Timestamp {
	if m != nil {
		return m.FinishedAt
	}
	return nil
}

func (m *Run) GetIsManaged() string {
	if m != nil {
		return m.IsManaged
	}
	return ""
}

func (m *Run) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Run) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Run) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

func (m *Run) GetBookmarked() bool {
	if m != nil {
		return m.Bookmarked
	}
	return false
}

func (m *Run) GetMetaInfo() *RunMetaInfo {
	if m != nil {
		return m.MetaInfo
	}
	return nil
}

func (m *Run) GetKind() RunKind {
	if m != nil {
		return m.Kind
	}
	return RunKind_job
}

func (m *Run) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

func (m *Run) GetInputs() *_struct.Struct {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Run) GetOutputs() *_struct.Struct {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Run) GetRunEnv() *_struct.Struct {
	if m != nil {
		return m.RunEnv
	}
	return nil
}

func (m *Run) GetIsResume() bool {
	if m != nil {
		return m.IsResume
	}
	return false
}

func (m *Run) GetIsClone() bool {
	if m != nil {
		return m.IsClone
	}
	return false
}

func (m *Run) GetCloningStrategy() string {
	if m != nil {
		return m.CloningStrategy
	}
	return ""
}

func (m *Run) GetPipeline() string {
	if m != nil {
		return m.Pipeline
	}
	return ""
}

func (m *Run) GetOriginal() string {
	if m != nil {
		return m.Original
	}
	return ""
}

func (m *Run) GetPipelineName() string {
	if m != nil {
		return m.PipelineName
	}
	return ""
}

func (m *Run) GetOriginalName() string {
	if m != nil {
		return m.OriginalName
	}
	return ""
}

// Request data to create run
type RunBodyRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Project where the experiement will be assigned
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Run object
	Run                  *Run     `protobuf:"bytes,3,opt,name=run,proto3" json:"run,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunBodyRequest) Reset()         { *m = RunBodyRequest{} }
func (m *RunBodyRequest) String() string { return proto.CompactTextString(m) }
func (*RunBodyRequest) ProtoMessage()    {}
func (*RunBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{2}
}

func (m *RunBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunBodyRequest.Unmarshal(m, b)
}
func (m *RunBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunBodyRequest.Marshal(b, m, deterministic)
}
func (m *RunBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunBodyRequest.Merge(m, src)
}
func (m *RunBodyRequest) XXX_Size() int {
	return xxx_messageInfo_RunBodyRequest.Size(m)
}
func (m *RunBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunBodyRequest proto.InternalMessageInfo

func (m *RunBodyRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RunBodyRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *RunBodyRequest) GetRun() *Run {
	if m != nil {
		return m.Run
	}
	return nil
}

// Request data to update run
type EntityRunBodyRequest struct {
	// Enitity Run
	Entity *ProjectEntityResourceRequest `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	// Run object
	Run                  *Run     `protobuf:"bytes,2,opt,name=run,proto3" json:"run,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityRunBodyRequest) Reset()         { *m = EntityRunBodyRequest{} }
func (m *EntityRunBodyRequest) String() string { return proto.CompactTextString(m) }
func (*EntityRunBodyRequest) ProtoMessage()    {}
func (*EntityRunBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{3}
}

func (m *EntityRunBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityRunBodyRequest.Unmarshal(m, b)
}
func (m *EntityRunBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityRunBodyRequest.Marshal(b, m, deterministic)
}
func (m *EntityRunBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityRunBodyRequest.Merge(m, src)
}
func (m *EntityRunBodyRequest) XXX_Size() int {
	return xxx_messageInfo_EntityRunBodyRequest.Size(m)
}
func (m *EntityRunBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityRunBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EntityRunBodyRequest proto.InternalMessageInfo

func (m *EntityRunBodyRequest) GetEntity() *ProjectEntityResourceRequest {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *EntityRunBodyRequest) GetRun() *Run {
	if m != nil {
		return m.Run
	}
	return nil
}

// Contains list runs
type ListRunsResponse struct {
	// Count of the entities
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// List of all entities
	Results []*Run `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Previous page
	Previous string `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
	// Next page
	Next                 string   `protobuf:"bytes,4,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRunsResponse) Reset()         { *m = ListRunsResponse{} }
func (m *ListRunsResponse) String() string { return proto.CompactTextString(m) }
func (*ListRunsResponse) ProtoMessage()    {}
func (*ListRunsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{4}
}

func (m *ListRunsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRunsResponse.Unmarshal(m, b)
}
func (m *ListRunsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRunsResponse.Marshal(b, m, deterministic)
}
func (m *ListRunsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRunsResponse.Merge(m, src)
}
func (m *ListRunsResponse) XXX_Size() int {
	return xxx_messageInfo_ListRunsResponse.Size(m)
}
func (m *ListRunsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRunsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRunsResponse proto.InternalMessageInfo

func (m *ListRunsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListRunsResponse) GetResults() []*Run {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListRunsResponse) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *ListRunsResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

// Run Settings specification
type RunSettingsCatalog struct {
	// Uuid
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Name
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunSettingsCatalog) Reset()         { *m = RunSettingsCatalog{} }
func (m *RunSettingsCatalog) String() string { return proto.CompactTextString(m) }
func (*RunSettingsCatalog) ProtoMessage()    {}
func (*RunSettingsCatalog) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{5}
}

func (m *RunSettingsCatalog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunSettingsCatalog.Unmarshal(m, b)
}
func (m *RunSettingsCatalog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunSettingsCatalog.Marshal(b, m, deterministic)
}
func (m *RunSettingsCatalog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunSettingsCatalog.Merge(m, src)
}
func (m *RunSettingsCatalog) XXX_Size() int {
	return xxx_messageInfo_RunSettingsCatalog.Size(m)
}
func (m *RunSettingsCatalog) XXX_DiscardUnknown() {
	xxx_messageInfo_RunSettingsCatalog.DiscardUnknown(m)
}

var xxx_messageInfo_RunSettingsCatalog proto.InternalMessageInfo

func (m *RunSettingsCatalog) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *RunSettingsCatalog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RunSettings struct {
	// Namespace
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Agent
	Agent *RunSettingsCatalog `protobuf:"bytes,2,opt,name=agent,proto3" json:"agent,omitempty"`
	// Queue
	Queue *RunSettingsCatalog `protobuf:"bytes,3,opt,name=queue,proto3" json:"queue,omitempty"`
	// Logs Store
	LogsStore *RunSettingsCatalog `protobuf:"bytes,4,opt,name=logs_store,json=logsStore,proto3" json:"logs_store,omitempty"`
	// Outputs Store
	OutputsStore *RunSettingsCatalog `protobuf:"bytes,5,opt,name=outputs_store,json=outputsStore,proto3" json:"outputs_store,omitempty"`
	// Init artifact Stores
	InitArtifactsStores []*RunSettingsCatalog `protobuf:"bytes,6,rep,name=init_artifacts_stores,json=initArtifactsStores,proto3" json:"init_artifacts_stores,omitempty"`
	// Artifacts Store
	ArtifactsStores []*RunSettingsCatalog `protobuf:"bytes,7,rep,name=artifacts_stores,json=artifactsStores,proto3" json:"artifacts_stores,omitempty"`
	// git Accesses
	GitAccesses []*RunSettingsCatalog `protobuf:"bytes,8,rep,name=git_accesses,json=gitAccesses,proto3" json:"git_accesses,omitempty"`
	// Registry Access
	RegistryAccess *RunSettingsCatalog `protobuf:"bytes,9,opt,name=registry_access,json=registryAccess,proto3" json:"registry_access,omitempty"`
	// K8S secrets
	K8SSecrets []*RunSettingsCatalog `protobuf:"bytes,10,rep,name=k8s_secrets,json=k8sSecrets,proto3" json:"k8s_secrets,omitempty"`
	// K8S config maps
	K8SConfigMaps        []*RunSettingsCatalog `protobuf:"bytes,11,rep,name=k8s_config_maps,json=k8sConfigMaps,proto3" json:"k8s_config_maps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RunSettings) Reset()         { *m = RunSettings{} }
func (m *RunSettings) String() string { return proto.CompactTextString(m) }
func (*RunSettings) ProtoMessage()    {}
func (*RunSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{6}
}

func (m *RunSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunSettings.Unmarshal(m, b)
}
func (m *RunSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunSettings.Marshal(b, m, deterministic)
}
func (m *RunSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunSettings.Merge(m, src)
}
func (m *RunSettings) XXX_Size() int {
	return xxx_messageInfo_RunSettings.Size(m)
}
func (m *RunSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RunSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RunSettings proto.InternalMessageInfo

func (m *RunSettings) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *RunSettings) GetAgent() *RunSettingsCatalog {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *RunSettings) GetQueue() *RunSettingsCatalog {
	if m != nil {
		return m.Queue
	}
	return nil
}

func (m *RunSettings) GetLogsStore() *RunSettingsCatalog {
	if m != nil {
		return m.LogsStore
	}
	return nil
}

func (m *RunSettings) GetOutputsStore() *RunSettingsCatalog {
	if m != nil {
		return m.OutputsStore
	}
	return nil
}

func (m *RunSettings) GetInitArtifactsStores() []*RunSettingsCatalog {
	if m != nil {
		return m.InitArtifactsStores
	}
	return nil
}

func (m *RunSettings) GetArtifactsStores() []*RunSettingsCatalog {
	if m != nil {
		return m.ArtifactsStores
	}
	return nil
}

func (m *RunSettings) GetGitAccesses() []*RunSettingsCatalog {
	if m != nil {
		return m.GitAccesses
	}
	return nil
}

func (m *RunSettings) GetRegistryAccess() *RunSettingsCatalog {
	if m != nil {
		return m.RegistryAccess
	}
	return nil
}

func (m *RunSettings) GetK8SSecrets() []*RunSettingsCatalog {
	if m != nil {
		return m.K8SSecrets
	}
	return nil
}

func (m *RunSettings) GetK8SConfigMaps() []*RunSettingsCatalog {
	if m != nil {
		return m.K8SConfigMaps
	}
	return nil
}

func init() {
	proto.RegisterEnum("v1.RunKind", RunKind_name, RunKind_value)
	proto.RegisterType((*RunMetaInfo)(nil), "v1.RunMetaInfo")
	proto.RegisterType((*Run)(nil), "v1.Run")
	proto.RegisterType((*RunBodyRequest)(nil), "v1.RunBodyRequest")
	proto.RegisterType((*EntityRunBodyRequest)(nil), "v1.EntityRunBodyRequest")
	proto.RegisterType((*ListRunsResponse)(nil), "v1.ListRunsResponse")
	proto.RegisterType((*RunSettingsCatalog)(nil), "v1.RunSettingsCatalog")
	proto.RegisterType((*RunSettings)(nil), "v1.RunSettings")
}

func init() { proto.RegisterFile("v1/run.proto", fileDescriptor_46ec9b83e76db539) }

var fileDescriptor_46ec9b83e76db539 = []byte{
	// 1075 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x4b, 0x73, 0x1b, 0x45,
	0x10, 0x46, 0x96, 0x6d, 0x49, 0xbd, 0x92, 0x25, 0x26, 0xaf, 0x89, 0xf3, 0x12, 0xca, 0xc5, 0x50,
	0x89, 0x84, 0x4d, 0x01, 0x4e, 0x85, 0x82, 0x12, 0xa9, 0x1c, 0x78, 0x98, 0xa2, 0xd6, 0xdc, 0x38,
	0x6c, 0x8d, 0x76, 0x5b, 0x9b, 0x89, 0xa4, 0x99, 0xcd, 0x3c, 0x14, 0x5c, 0xfc, 0x02, 0xee, 0xfc,
	0x5f, 0xa8, 0x79, 0xac, 0x63, 0x9c, 0x58, 0xe6, 0xa4, 0xe9, 0xaf, 0xbf, 0xaf, 0x7b, 0x66, 0x7a,
	0xba, 0x57, 0xd0, 0x5d, 0x1f, 0x4e, 0x94, 0x15, 0xe3, 0x4a, 0x49, 0x23, 0xc9, 0xd6, 0xfa, 0x70,
	0xff, 0x7e, 0x29, 0x65, 0xb9, 0xc4, 0x09, 0xab, 0xf8, 0x84, 0x09, 0x21, 0x0d, 0x33, 0x5c, 0x0a,
	0x1d, 0x18, 0xe7, 0x5e, 0x6f, 0xcd, 0xec, 0x7c, 0xa2, 0x8d, 0xb2, 0xb9, 0x89, 0xde, 0x47, 0x97,
	0xbd, 0x86, 0xaf, 0x50, 0x1b, 0xb6, 0xaa, 0x22, 0xe1, 0x89, 0xff, 0xc9, 0x9f, 0x96, 0x28, 0x9e,
	0xea, 0xb7, 0xac, 0x2c, 0x51, 0x4d, 0x64, 0xe5, 0x13, 0x7c, 0x20, 0x59, 0x6f, 0x7d, 0x38, 0x99,
	0x31, 0x8d, 0xc1, 0x1c, 0xfd, 0xd5, 0x80, 0x24, 0xb5, 0xe2, 0x04, 0x0d, 0xfb, 0x41, 0xcc, 0x25,
	0xa1, 0xd0, 0xd2, 0xa8, 0xd6, 0x3c, 0x47, 0xda, 0x18, 0x36, 0x0e, 0xda, 0x69, 0x6d, 0x92, 0x21,
	0x24, 0xb9, 0x14, 0xb9, 0x55, 0x0a, 0x45, 0x7e, 0x46, 0xb7, 0x86, 0x8d, 0x83, 0x9d, 0xf4, 0x22,
	0x44, 0x1e, 0x43, 0xaf, 0x62, 0x8a, 0x2d, 0x97, 0xb8, 0xcc, 0x16, 0x5c, 0x14, 0xb4, 0x39, 0x6c,
	0x1c, 0x74, 0xd2, 0x6e, 0x0d, 0xfe, 0xc4, 0x45, 0x41, 0xee, 0x42, 0x5b, 0x59, 0x11, 0xfc, 0xdb,
	0xde, 0xdf, 0x52, 0x56, 0x38, 0xd7, 0xe8, 0xef, 0x36, 0x34, 0x53, 0x2b, 0x08, 0x81, 0x6d, 0x6b,
	0x79, 0xe1, 0x37, 0xd0, 0x49, 0xfd, 0xda, 0x61, 0x82, 0xad, 0xd0, 0xa7, 0xed, 0xa4, 0x7e, 0xed,
	0x76, 0x54, 0xa0, 0xce, 0x15, 0xf7, 0x87, 0x8d, 0xd9, 0x2e, 0x42, 0x4e, 0x65, 0x58, 0xa9, 0xe9,
	0xf6, 0xb0, 0xe9, 0x54, 0x6e, 0xed, 0x4e, 0x58, 0xe0, 0x12, 0x0d, 0x16, 0x74, 0x27, 0x9c, 0x30,
	0x9a, 0x3e, 0xaf, 0x46, 0x45, 0x77, 0x63, 0x5e, 0x8d, 0x8a, 0xdc, 0x84, 0x1d, 0xf9, 0x56, 0xa0,
	0xa2, 0x2d, 0x0f, 0x06, 0xc3, 0xc5, 0xa8, 0x94, 0x7c, 0x8d, 0xb9, 0xa1, 0xed, 0x70, 0x86, 0x68,
	0x92, 0x67, 0x00, 0xb9, 0x42, 0x66, 0xb0, 0xc8, 0x98, 0xa1, 0x9d, 0x61, 0xe3, 0x20, 0x39, 0xda,
	0x1f, 0x87, 0x12, 0x8e, 0xeb, 0x12, 0x8e, 0x7f, 0xab, 0x4b, 0x98, 0x76, 0x22, 0x7b, 0xea, 0xa5,
	0xb6, 0x2a, 0x6a, 0x29, 0x5c, 0x2f, 0x8d, 0xec, 0x20, 0xd5, 0x86, 0xa9, 0x28, 0x4d, 0xae, 0x97,
	0x46, 0xf6, 0xd4, 0x90, 0xe7, 0x90, 0xcc, 0xb9, 0xe0, 0xfa, 0x55, 0xd0, 0x76, 0xaf, 0xd5, 0x42,
	0x4d, 0x9f, 0x1a, 0xf2, 0x00, 0x80, 0xeb, 0x6c, 0xc5, 0x04, 0x2b, 0xb1, 0xa0, 0x3d, 0x7f, 0x15,
	0x1d, 0xae, 0x4f, 0x02, 0xe0, 0xae, 0x29, 0x97, 0xc2, 0xa0, 0x30, 0x74, 0x2f, 0x5c, 0x53, 0x34,
	0xc9, 0x6d, 0xd8, 0xd5, 0x86, 0x19, 0xab, 0x69, 0xdf, 0x3b, 0xa2, 0xe5, 0x70, 0x85, 0xac, 0x58,
	0x21, 0x1d, 0x04, 0x3c, 0x58, 0xe4, 0x21, 0xc0, 0x4c, 0xca, 0xc5, 0x8a, 0xa9, 0x05, 0x16, 0xf4,
	0x63, 0x5f, 0xb7, 0x0b, 0x08, 0x79, 0x02, 0x9d, 0x15, 0x1a, 0x96, 0x71, 0x31, 0x97, 0x94, 0xf8,
	0x33, 0xf4, 0xc7, 0xeb, 0xc3, 0xf1, 0x85, 0xa7, 0x9d, 0xb6, 0x57, 0xf5, 0x23, 0x7f, 0x04, 0xdb,
	0xfe, 0xfd, 0xdd, 0x18, 0x36, 0x0e, 0xf6, 0x8e, 0x92, 0x48, 0x74, 0x6f, 0x30, 0xf5, 0x0e, 0x32,
	0x80, 0xe6, 0x2b, 0x3b, 0xa3, 0x37, 0xfd, 0x1e, 0xdc, 0x92, 0x4c, 0x60, 0x97, 0x8b, 0xca, 0x1a,
	0x4d, 0x6f, 0xf9, 0xe8, 0x77, 0xde, 0xbb, 0xa1, 0x53, 0xdf, 0xb4, 0x69, 0xa4, 0x91, 0x43, 0x68,
	0x49, 0x6b, 0xbc, 0xe2, 0xf6, 0x66, 0x45, 0xcd, 0x23, 0x9f, 0x83, 0x6b, 0x85, 0x0c, 0xc5, 0x9a,
	0xde, 0xb9, 0x26, 0x89, 0xb2, 0xe2, 0xa5, 0x58, 0x93, 0x7b, 0xd0, 0xe1, 0x3a, 0x53, 0xa8, 0xed,
	0x0a, 0x29, 0xf5, 0xb7, 0xd2, 0xe6, 0x3a, 0xf5, 0xb6, 0xeb, 0x34, 0xae, 0xb3, 0x7c, 0x29, 0x05,
	0xd2, 0xbb, 0xe1, 0xa5, 0x73, 0xfd, 0xc2, 0x99, 0xe4, 0x53, 0x18, 0x38, 0x9c, 0x8b, 0x32, 0xd3,
	0x46, 0x31, 0x83, 0xe5, 0x19, 0xdd, 0xf7, 0x87, 0xed, 0x47, 0xfc, 0x34, 0xc2, 0x64, 0x1f, 0xda,
	0x15, 0xaf, 0x70, 0xc9, 0x05, 0xd2, 0x7b, 0x9e, 0x72, 0x6e, 0x3b, 0x9f, 0x54, 0xbc, 0xe4, 0x82,
	0x2d, 0xe9, 0xfd, 0xe0, 0xab, 0x6d, 0x3f, 0x0c, 0x22, 0x2f, 0xf3, 0x9d, 0xfb, 0x20, 0x0e, 0x83,
	0x08, 0xfe, 0xe2, 0x3a, 0xf8, 0x31, 0xf4, 0x6a, 0x41, 0x20, 0x3d, 0x0c, 0xa4, 0x1a, 0x74, 0xa4,
	0xd1, 0xef, 0xb0, 0x97, 0x5a, 0xf1, 0xbd, 0x2c, 0xce, 0x52, 0x7c, 0x63, 0x51, 0x9b, 0x77, 0x4d,
	0xd9, 0xb8, 0xa2, 0x29, 0xb7, 0xfe, 0xdb, 0x94, 0x77, 0xa1, 0xa9, 0x6c, 0x18, 0x10, 0xc9, 0x51,
	0x2b, 0x96, 0x3b, 0x75, 0xd8, 0x68, 0x01, 0x37, 0x5f, 0x0a, 0xc3, 0xcd, 0xd9, 0xa5, 0x14, 0xc7,
	0xb0, 0x8b, 0x1e, 0xf7, 0x39, 0x92, 0xa3, 0xa1, 0x53, 0xfd, 0x1a, 0xe2, 0x45, 0x01, 0x6a, 0x69,
	0x55, 0x8e, 0x51, 0x91, 0x46, 0x7e, 0x9d, 0x6c, 0xeb, 0x03, 0xc9, 0xfe, 0x84, 0xc1, 0xcf, 0x5c,
	0x9b, 0xd4, 0x0a, 0x57, 0xa3, 0x4a, 0x0a, 0x8d, 0xee, 0x2c, 0xb9, 0xb4, 0xc2, 0xf8, 0x3c, 0x3b,
	0x69, 0x30, 0xc8, 0x27, 0xd0, 0x72, 0x55, 0x5d, 0x1a, 0x4d, 0xb7, 0x86, 0xcd, 0x8b, 0x81, 0x6a,
	0xdc, 0x17, 0x46, 0xe1, 0x9a, 0x4b, 0xab, 0xe3, 0xe8, 0x3b, 0xb7, 0xfd, 0xb4, 0xc4, 0x3f, 0x4c,
	0x1c, 0xb0, 0x7e, 0x3d, 0xfa, 0x06, 0x48, 0x6a, 0xc5, 0x29, 0x1a, 0xc3, 0x45, 0xa9, 0x5f, 0x30,
	0xc3, 0x96, 0xb2, 0xfc, 0xbf, 0xb3, 0x76, 0xf4, 0xcf, 0xb6, 0xff, 0x4e, 0xd4, 0x72, 0x72, 0x1f,
	0x3a, 0x0e, 0xd7, 0x15, 0x8b, 0x5f, 0x8a, 0x4e, 0xfa, 0x0e, 0x20, 0x4f, 0x60, 0x87, 0x95, 0xae,
	0xed, 0xc3, 0x2d, 0xdc, 0x8e, 0x9b, 0xbf, 0x94, 0x3c, 0x0d, 0x24, 0xc7, 0x7e, 0x63, 0xd1, 0x62,
	0x2c, 0xd0, 0x95, 0x6c, 0x4f, 0x22, 0x5f, 0x02, 0x2c, 0x65, 0xa9, 0x33, 0x6d, 0xa4, 0x42, 0x7f,
	0xc2, 0xab, 0x25, 0x1d, 0xc7, 0x3c, 0x75, 0x44, 0xf2, 0x1c, 0x7a, 0xb1, 0xcf, 0xa2, 0x72, 0x67,
	0xa3, 0xb2, 0x1b, 0xc9, 0x41, 0xfc, 0x23, 0xdc, 0xe2, 0x82, 0x9b, 0x8c, 0x29, 0xc3, 0xe7, 0x2c,
	0xaf, 0x63, 0x68, 0xba, 0xeb, 0x8b, 0x73, 0x55, 0x90, 0x1b, 0x4e, 0x34, 0xad, 0x35, 0x3e, 0x94,
	0x26, 0x53, 0x18, 0xbc, 0x17, 0xa6, 0xb5, 0x31, 0x4c, 0x9f, 0x5d, 0x0a, 0xf1, 0x0c, 0xba, 0xa5,
	0xdb, 0x4d, 0x9e, 0xa3, 0xd6, 0xa8, 0x69, 0x7b, 0xa3, 0x3c, 0x29, 0xb9, 0x99, 0x46, 0x2a, 0xf9,
	0x0e, 0xfa, 0x0a, 0x4b, 0xae, 0x8d, 0x3a, 0x8b, 0xfa, 0xf8, 0x91, 0xba, 0x4a, 0xbd, 0x57, 0xd3,
	0x43, 0x08, 0xf2, 0x35, 0x24, 0x8b, 0x63, 0x9d, 0x69, 0xcc, 0x15, 0x1a, 0x4d, 0x61, 0x63, 0x6a,
	0x58, 0x1c, 0xeb, 0xd3, 0xc0, 0x24, 0xdf, 0x42, 0xdf, 0x09, 0x73, 0x29, 0xe6, 0xbc, 0xcc, 0x56,
	0xac, 0xd2, 0x34, 0xd9, 0x28, 0xee, 0x2d, 0x8e, 0xf5, 0x0b, 0xcf, 0x3e, 0x61, 0x95, 0xfe, 0xec,
	0x2b, 0x68, 0xc5, 0x21, 0x4d, 0x5a, 0xd0, 0x7c, 0x2d, 0x67, 0x83, 0x8f, 0x48, 0x72, 0xfe, 0x6f,
	0x65, 0xd0, 0x70, 0x68, 0xc1, 0xca, 0xc1, 0x16, 0xe9, 0x42, 0xbb, 0xfe, 0xcb, 0x31, 0x68, 0xce,
	0x76, 0xfd, 0xf0, 0xfc, 0xe2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x0b, 0xf9, 0xa0, 0x96,
	0x09, 0x00, 0x00,
}
