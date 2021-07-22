// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/edgify/edgifypb/prediction.proto

package edgifypb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PredictionItem struct {
	Data                 []string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PredictionItem) Reset()         { *m = PredictionItem{} }
func (m *PredictionItem) String() string { return proto.CompactTextString(m) }
func (*PredictionItem) ProtoMessage()    {}
func (*PredictionItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b081fc5a5c788afb, []int{0}
}

func (m *PredictionItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionItem.Unmarshal(m, b)
}
func (m *PredictionItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionItem.Marshal(b, m, deterministic)
}
func (m *PredictionItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionItem.Merge(m, src)
}
func (m *PredictionItem) XXX_Size() int {
	return xxx_messageInfo_PredictionItem.Size(m)
}
func (m *PredictionItem) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionItem.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionItem proto.InternalMessageInfo

func (m *PredictionItem) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

type Prediction struct {
	Uuid                 string            `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Duration             uint32            `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	PredictedAt          string            `protobuf:"bytes,3,opt,name=predicted_at,json=predictedAt,proto3" json:"predicted_at,omitempty"`
	ModelId              uint32            `protobuf:"varint,4,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	Predictions          []*PredictionItem `protobuf:"bytes,5,rep,name=predictions,proto3" json:"predictions,omitempty"`
	Image                *Image            `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Source               string            `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	Certain              bool              `protobuf:"varint,8,opt,name=certain,proto3" json:"certain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Prediction) Reset()         { *m = Prediction{} }
func (m *Prediction) String() string { return proto.CompactTextString(m) }
func (*Prediction) ProtoMessage()    {}
func (*Prediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_b081fc5a5c788afb, []int{1}
}

func (m *Prediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prediction.Unmarshal(m, b)
}
func (m *Prediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prediction.Marshal(b, m, deterministic)
}
func (m *Prediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prediction.Merge(m, src)
}
func (m *Prediction) XXX_Size() int {
	return xxx_messageInfo_Prediction.Size(m)
}
func (m *Prediction) XXX_DiscardUnknown() {
	xxx_messageInfo_Prediction.DiscardUnknown(m)
}

var xxx_messageInfo_Prediction proto.InternalMessageInfo

func (m *Prediction) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Prediction) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Prediction) GetPredictedAt() string {
	if m != nil {
		return m.PredictedAt
	}
	return ""
}

func (m *Prediction) GetModelId() uint32 {
	if m != nil {
		return m.ModelId
	}
	return 0
}

func (m *Prediction) GetPredictions() []*PredictionItem {
	if m != nil {
		return m.Predictions
	}
	return nil
}

func (m *Prediction) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Prediction) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Prediction) GetCertain() bool {
	if m != nil {
		return m.Certain
	}
	return false
}

type Image struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Image                []byte   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	GroupId              string   `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_b081fc5a5c788afb, []int{2}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Image) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Image) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

type ModelDeployment struct {
	Id                   uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ModelId              uint32               `protobuf:"varint,2,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	ModelPath            string               `protobuf:"bytes,3,opt,name=model_path,json=modelPath,proto3" json:"model_path,omitempty"`
	ConfigPath           string               `protobuf:"bytes,4,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
	Verified             bool                 `protobuf:"varint,5,opt,name=verified,proto3" json:"verified,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ModelDeployment) Reset()         { *m = ModelDeployment{} }
func (m *ModelDeployment) String() string { return proto.CompactTextString(m) }
func (*ModelDeployment) ProtoMessage()    {}
func (*ModelDeployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b081fc5a5c788afb, []int{3}
}

func (m *ModelDeployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelDeployment.Unmarshal(m, b)
}
func (m *ModelDeployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelDeployment.Marshal(b, m, deterministic)
}
func (m *ModelDeployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelDeployment.Merge(m, src)
}
func (m *ModelDeployment) XXX_Size() int {
	return xxx_messageInfo_ModelDeployment.Size(m)
}
func (m *ModelDeployment) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelDeployment.DiscardUnknown(m)
}

var xxx_messageInfo_ModelDeployment proto.InternalMessageInfo

func (m *ModelDeployment) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ModelDeployment) GetModelId() uint32 {
	if m != nil {
		return m.ModelId
	}
	return 0
}

func (m *ModelDeployment) GetModelPath() string {
	if m != nil {
		return m.ModelPath
	}
	return ""
}

func (m *ModelDeployment) GetConfigPath() string {
	if m != nil {
		return m.ConfigPath
	}
	return ""
}

func (m *ModelDeployment) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *ModelDeployment) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type GroundTruth struct {
	Prediction           *Prediction `protobuf:"bytes,1,opt,name=prediction,proto3" json:"prediction,omitempty"`
	Label                string      `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Source               string      `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GroundTruth) Reset()         { *m = GroundTruth{} }
func (m *GroundTruth) String() string { return proto.CompactTextString(m) }
func (*GroundTruth) ProtoMessage()    {}
func (*GroundTruth) Descriptor() ([]byte, []int) {
	return fileDescriptor_b081fc5a5c788afb, []int{4}
}

func (m *GroundTruth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroundTruth.Unmarshal(m, b)
}
func (m *GroundTruth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroundTruth.Marshal(b, m, deterministic)
}
func (m *GroundTruth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroundTruth.Merge(m, src)
}
func (m *GroundTruth) XXX_Size() int {
	return xxx_messageInfo_GroundTruth.Size(m)
}
func (m *GroundTruth) XXX_DiscardUnknown() {
	xxx_messageInfo_GroundTruth.DiscardUnknown(m)
}

var xxx_messageInfo_GroundTruth proto.InternalMessageInfo

func (m *GroundTruth) GetPrediction() *Prediction {
	if m != nil {
		return m.Prediction
	}
	return nil
}

func (m *GroundTruth) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *GroundTruth) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type PredictionRequest struct {
	Image                *Image   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	WithoutPrediction    bool     `protobuf:"varint,2,opt,name=withoutPrediction,proto3" json:"withoutPrediction,omitempty"`
	Source               string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PredictionRequest) Reset()         { *m = PredictionRequest{} }
func (m *PredictionRequest) String() string { return proto.CompactTextString(m) }
func (*PredictionRequest) ProtoMessage()    {}
func (*PredictionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b081fc5a5c788afb, []int{5}
}

func (m *PredictionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionRequest.Unmarshal(m, b)
}
func (m *PredictionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionRequest.Marshal(b, m, deterministic)
}
func (m *PredictionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionRequest.Merge(m, src)
}
func (m *PredictionRequest) XXX_Size() int {
	return xxx_messageInfo_PredictionRequest.Size(m)
}
func (m *PredictionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionRequest proto.InternalMessageInfo

func (m *PredictionRequest) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *PredictionRequest) GetWithoutPrediction() bool {
	if m != nil {
		return m.WithoutPrediction
	}
	return false
}

func (m *PredictionRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type PredictionResponse struct {
	Prediction           *Prediction `protobuf:"bytes,1,opt,name=prediction,proto3" json:"prediction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PredictionResponse) Reset()         { *m = PredictionResponse{} }
func (m *PredictionResponse) String() string { return proto.CompactTextString(m) }
func (*PredictionResponse) ProtoMessage()    {}
func (*PredictionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b081fc5a5c788afb, []int{6}
}

func (m *PredictionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionResponse.Unmarshal(m, b)
}
func (m *PredictionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionResponse.Marshal(b, m, deterministic)
}
func (m *PredictionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionResponse.Merge(m, src)
}
func (m *PredictionResponse) XXX_Size() int {
	return xxx_messageInfo_PredictionResponse.Size(m)
}
func (m *PredictionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionResponse proto.InternalMessageInfo

func (m *PredictionResponse) GetPrediction() *Prediction {
	if m != nil {
		return m.Prediction
	}
	return nil
}

type GroundTruthRequest struct {
	GroundTruth          *GroundTruth `protobuf:"bytes,1,opt,name=ground_truth,json=groundTruth,proto3" json:"ground_truth,omitempty"`
	Source               string       `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GroundTruthRequest) Reset()         { *m = GroundTruthRequest{} }
func (m *GroundTruthRequest) String() string { return proto.CompactTextString(m) }
func (*GroundTruthRequest) ProtoMessage()    {}
func (*GroundTruthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b081fc5a5c788afb, []int{7}
}

func (m *GroundTruthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroundTruthRequest.Unmarshal(m, b)
}
func (m *GroundTruthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroundTruthRequest.Marshal(b, m, deterministic)
}
func (m *GroundTruthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroundTruthRequest.Merge(m, src)
}
func (m *GroundTruthRequest) XXX_Size() int {
	return xxx_messageInfo_GroundTruthRequest.Size(m)
}
func (m *GroundTruthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GroundTruthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GroundTruthRequest proto.InternalMessageInfo

func (m *GroundTruthRequest) GetGroundTruth() *GroundTruth {
	if m != nil {
		return m.GroundTruth
	}
	return nil
}

func (m *GroundTruthRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type GroundTruthResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroundTruthResponse) Reset()         { *m = GroundTruthResponse{} }
func (m *GroundTruthResponse) String() string { return proto.CompactTextString(m) }
func (*GroundTruthResponse) ProtoMessage()    {}
func (*GroundTruthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b081fc5a5c788afb, []int{8}
}

func (m *GroundTruthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroundTruthResponse.Unmarshal(m, b)
}
func (m *GroundTruthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroundTruthResponse.Marshal(b, m, deterministic)
}
func (m *GroundTruthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroundTruthResponse.Merge(m, src)
}
func (m *GroundTruthResponse) XXX_Size() int {
	return xxx_messageInfo_GroundTruthResponse.Size(m)
}
func (m *GroundTruthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GroundTruthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GroundTruthResponse proto.InternalMessageInfo

type GetCurrentModelDeploymentRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentModelDeploymentRequest) Reset()         { *m = GetCurrentModelDeploymentRequest{} }
func (m *GetCurrentModelDeploymentRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentModelDeploymentRequest) ProtoMessage()    {}
func (*GetCurrentModelDeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b081fc5a5c788afb, []int{9}
}

func (m *GetCurrentModelDeploymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentModelDeploymentRequest.Unmarshal(m, b)
}
func (m *GetCurrentModelDeploymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentModelDeploymentRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrentModelDeploymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentModelDeploymentRequest.Merge(m, src)
}
func (m *GetCurrentModelDeploymentRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentModelDeploymentRequest.Size(m)
}
func (m *GetCurrentModelDeploymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentModelDeploymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentModelDeploymentRequest proto.InternalMessageInfo

type GetCurrentModelDeploymentResponse struct {
	ModelDeployment      *ModelDeployment `protobuf:"bytes,1,opt,name=model_deployment,json=modelDeployment,proto3" json:"model_deployment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetCurrentModelDeploymentResponse) Reset()         { *m = GetCurrentModelDeploymentResponse{} }
func (m *GetCurrentModelDeploymentResponse) String() string { return proto.CompactTextString(m) }
func (*GetCurrentModelDeploymentResponse) ProtoMessage()    {}
func (*GetCurrentModelDeploymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b081fc5a5c788afb, []int{10}
}

func (m *GetCurrentModelDeploymentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentModelDeploymentResponse.Unmarshal(m, b)
}
func (m *GetCurrentModelDeploymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentModelDeploymentResponse.Marshal(b, m, deterministic)
}
func (m *GetCurrentModelDeploymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentModelDeploymentResponse.Merge(m, src)
}
func (m *GetCurrentModelDeploymentResponse) XXX_Size() int {
	return xxx_messageInfo_GetCurrentModelDeploymentResponse.Size(m)
}
func (m *GetCurrentModelDeploymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentModelDeploymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentModelDeploymentResponse proto.InternalMessageInfo

func (m *GetCurrentModelDeploymentResponse) GetModelDeployment() *ModelDeployment {
	if m != nil {
		return m.ModelDeployment
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictionItem)(nil), "edgify.PredictionItem")
	proto.RegisterType((*Prediction)(nil), "edgify.Prediction")
	proto.RegisterType((*Image)(nil), "edgify.Image")
	proto.RegisterType((*ModelDeployment)(nil), "edgify.ModelDeployment")
	proto.RegisterType((*GroundTruth)(nil), "edgify.GroundTruth")
	proto.RegisterType((*PredictionRequest)(nil), "edgify.PredictionRequest")
	proto.RegisterType((*PredictionResponse)(nil), "edgify.PredictionResponse")
	proto.RegisterType((*GroundTruthRequest)(nil), "edgify.GroundTruthRequest")
	proto.RegisterType((*GroundTruthResponse)(nil), "edgify.GroundTruthResponse")
	proto.RegisterType((*GetCurrentModelDeploymentRequest)(nil), "edgify.GetCurrentModelDeploymentRequest")
	proto.RegisterType((*GetCurrentModelDeploymentResponse)(nil), "edgify.GetCurrentModelDeploymentResponse")
}

func init() {
	proto.RegisterFile("src/edgify/edgifypb/prediction.proto", fileDescriptor_b081fc5a5c788afb)
}

var fileDescriptor_b081fc5a5c788afb = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x3e, 0x76, 0x9a, 0xc4, 0x19, 0x37, 0xed, 0xe9, 0xf6, 0x9c, 0x1e, 0xc7, 0x47, 0xa8, 0xae,
	0xe9, 0x45, 0x90, 0x50, 0x22, 0x05, 0x09, 0xc1, 0x65, 0x5b, 0x50, 0x1b, 0x09, 0x50, 0x65, 0x7a,
	0xc5, 0x4d, 0xe5, 0x78, 0x37, 0xce, 0x4a, 0xb1, 0xd7, 0xac, 0xd7, 0x45, 0xbd, 0xe1, 0x05, 0x78,
	0x06, 0x9e, 0x8a, 0x17, 0x42, 0xde, 0x5d, 0x3b, 0x4e, 0x9b, 0x16, 0xc4, 0x55, 0x76, 0x66, 0xbe,
	0xf9, 0xf9, 0xbe, 0x99, 0xc8, 0x70, 0x9c, 0xf3, 0x68, 0x4c, 0x70, 0x4c, 0xe7, 0xb7, 0xfa, 0x27,
	0x9b, 0x8d, 0x33, 0x4e, 0x30, 0x8d, 0x04, 0x65, 0xe9, 0x28, 0xe3, 0x4c, 0x30, 0xd4, 0x51, 0x21,
	0xf7, 0x30, 0x66, 0x2c, 0x5e, 0x92, 0xb1, 0xf4, 0xce, 0x8a, 0xf9, 0x58, 0xd0, 0x84, 0xe4, 0x22,
	0x4c, 0x32, 0x05, 0xf4, 0x8f, 0x61, 0xe7, 0xb2, 0x4e, 0x9e, 0x0a, 0x92, 0x20, 0x04, 0x5b, 0x38,
	0x14, 0xa1, 0x63, 0x78, 0xad, 0x61, 0x2f, 0x90, 0x6f, 0xff, 0x9b, 0x09, 0xb0, 0x82, 0x95, 0x90,
	0xa2, 0xa0, 0xd8, 0x31, 0x3c, 0xa3, 0x84, 0x94, 0x6f, 0xe4, 0x82, 0x85, 0x0b, 0x1e, 0x96, 0x71,
	0xc7, 0xf4, 0x8c, 0x61, 0x3f, 0xa8, 0x6d, 0x74, 0x04, 0xdb, 0x7a, 0x42, 0x82, 0xaf, 0x43, 0xe1,
	0xb4, 0x64, 0x9e, 0x5d, 0xfb, 0x4e, 0x04, 0x1a, 0x80, 0x95, 0x30, 0x4c, 0x96, 0xd7, 0x14, 0x3b,
	0x5b, 0x32, 0xbd, 0x2b, 0xed, 0x29, 0x46, 0xaf, 0xc0, 0x5e, 0xf1, 0xcb, 0x9d, 0xb6, 0xd7, 0x1a,
	0xda, 0x93, 0x83, 0x91, 0x62, 0x38, 0x5a, 0x9f, 0x3e, 0x68, 0x42, 0xd1, 0x53, 0x68, 0xd3, 0x24,
	0x8c, 0x89, 0xd3, 0xf1, 0x8c, 0xa1, 0x3d, 0xe9, 0x57, 0x39, 0xd3, 0xd2, 0x19, 0xa8, 0x18, 0x3a,
	0x80, 0x4e, 0xce, 0x0a, 0x1e, 0x11, 0xa7, 0x2b, 0xc7, 0xd2, 0x16, 0x72, 0xa0, 0x1b, 0x11, 0x2e,
	0x42, 0x9a, 0x3a, 0x96, 0x67, 0x0c, 0xad, 0xa0, 0x32, 0xfd, 0x77, 0xd0, 0x96, 0x15, 0x36, 0xea,
	0xf0, 0x4f, 0xd5, 0xb3, 0x14, 0x61, 0xbb, 0x6a, 0x32, 0x00, 0x2b, 0xe6, 0xac, 0xc8, 0x4a, 0x7a,
	0x8a, 0x7d, 0x57, 0xda, 0x53, 0xec, 0xff, 0x30, 0x60, 0xf7, 0x7d, 0x49, 0xf5, 0x0d, 0xc9, 0x96,
	0xec, 0x36, 0x21, 0xa9, 0x40, 0x3b, 0x60, 0xea, 0xb2, 0xfd, 0xc0, 0xa4, 0x78, 0x4d, 0x1d, 0x73,
	0x5d, 0x9d, 0x27, 0x00, 0x2a, 0x94, 0x85, 0x62, 0xa1, 0x6b, 0xf7, 0xa4, 0xe7, 0x32, 0x14, 0x0b,
	0x74, 0x08, 0x76, 0xc4, 0xd2, 0x39, 0x8d, 0x55, 0x7c, 0x4b, 0xc6, 0x41, 0xb9, 0x24, 0xc0, 0x05,
	0xeb, 0x86, 0x70, 0x3a, 0xa7, 0x04, 0x3b, 0x6d, 0xc9, 0xb3, 0xb6, 0xd1, 0x6b, 0x80, 0x88, 0x93,
	0x50, 0x6f, 0x4d, 0x89, 0xe8, 0x8e, 0xd4, 0x49, 0x8d, 0xaa, 0x93, 0x1a, 0x5d, 0x55, 0x27, 0x15,
	0xf4, 0x34, 0xfa, 0x44, 0xf8, 0x0c, 0xec, 0x73, 0xce, 0x8a, 0x14, 0x5f, 0xf1, 0x42, 0x2c, 0xd0,
	0x04, 0x60, 0xb5, 0x18, 0x49, 0xcc, 0x9e, 0xa0, 0xfb, 0x2b, 0x0c, 0x1a, 0xa8, 0x52, 0xc9, 0x65,
	0x38, 0x23, 0x4b, 0xc9, 0xb8, 0x17, 0x28, 0xa3, 0xb1, 0xae, 0x56, 0x73, 0x5d, 0xfe, 0x57, 0xd8,
	0x6b, 0xd4, 0x21, 0x9f, 0x0b, 0x92, 0x8b, 0xd5, 0x01, 0x18, 0x8f, 0x1c, 0xc0, 0x73, 0xd8, 0xfb,
	0x42, 0xc5, 0x82, 0x15, 0x62, 0x55, 0x40, 0xf6, 0xb4, 0x82, 0xfb, 0x81, 0x07, 0xfb, 0x5f, 0x00,
	0x6a, 0xf6, 0xcf, 0x33, 0x96, 0xe6, 0xe4, 0x4f, 0x78, 0xfb, 0x18, 0x50, 0x43, 0xba, 0x8a, 0xca,
	0x4b, 0xd8, 0x8e, 0xa5, 0xf7, 0x5a, 0x94, 0x6e, 0x5d, 0x6b, 0xbf, 0xaa, 0xd5, 0xcc, 0xb0, 0xe3,
	0x86, 0xf2, 0xab, 0x79, 0xcd, 0xb5, 0x79, 0xff, 0x85, 0xfd, 0xb5, 0x2e, 0x6a, 0x60, 0xdf, 0x07,
	0xef, 0x9c, 0x88, 0xb3, 0x82, 0x73, 0x92, 0x8a, 0x3b, 0x67, 0xa9, 0x47, 0xf1, 0x63, 0x38, 0x7a,
	0x04, 0xa3, 0x99, 0x9f, 0xc2, 0xdf, 0xea, 0x2e, 0x71, 0x1d, 0xd3, 0x33, 0xff, 0x57, 0xcd, 0x7c,
	0x37, 0x75, 0x37, 0x59, 0x77, 0x4c, 0xbe, 0x9b, 0xd0, 0x7f, 0x2b, 0xb1, 0x1f, 0x09, 0xbf, 0xa1,
	0x11, 0x41, 0x17, 0xd0, 0x3f, 0x27, 0xcd, 0x75, 0x0c, 0x36, 0x88, 0xa9, 0xc6, 0x74, 0xdd, 0x4d,
	0x21, 0x4d, 0xf3, 0x2f, 0xf4, 0x01, 0xf6, 0xce, 0xe4, 0xb5, 0x36, 0xcf, 0xd4, 0xdd, 0x24, 0xa7,
	0x2e, 0xf7, 0xff, 0xc6, 0x58, 0x5d, 0x8f, 0xc3, 0xe0, 0x41, 0x51, 0xd0, 0xb0, 0xce, 0xfd, 0x85,
	0xb6, 0xee, 0xb3, 0xdf, 0x40, 0x56, 0x3d, 0x4f, 0xe1, 0x93, 0x55, 0x7d, 0x02, 0x66, 0x1d, 0xf9,
	0x7f, 0x7c, 0xf1, 0x33, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xc9, 0x10, 0x93, 0x20, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EdgifyServiceClient is the client API for EdgifyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EdgifyServiceClient interface {
	GetPrediction(ctx context.Context, in *PredictionRequest, opts ...grpc.CallOption) (*PredictionResponse, error)
	CreateGroundTruth(ctx context.Context, in *GroundTruthRequest, opts ...grpc.CallOption) (*GroundTruthResponse, error)
	GetCurrentModelDeployment(ctx context.Context, in *GetCurrentModelDeploymentRequest, opts ...grpc.CallOption) (*GetCurrentModelDeploymentResponse, error)
}

type edgifyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEdgifyServiceClient(cc grpc.ClientConnInterface) EdgifyServiceClient {
	return &edgifyServiceClient{cc}
}

func (c *edgifyServiceClient) GetPrediction(ctx context.Context, in *PredictionRequest, opts ...grpc.CallOption) (*PredictionResponse, error) {
	out := new(PredictionResponse)
	err := c.cc.Invoke(ctx, "/edgify.EdgifyService/GetPrediction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edgifyServiceClient) CreateGroundTruth(ctx context.Context, in *GroundTruthRequest, opts ...grpc.CallOption) (*GroundTruthResponse, error) {
	out := new(GroundTruthResponse)
	err := c.cc.Invoke(ctx, "/edgify.EdgifyService/CreateGroundTruth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edgifyServiceClient) GetCurrentModelDeployment(ctx context.Context, in *GetCurrentModelDeploymentRequest, opts ...grpc.CallOption) (*GetCurrentModelDeploymentResponse, error) {
	out := new(GetCurrentModelDeploymentResponse)
	err := c.cc.Invoke(ctx, "/edgify.EdgifyService/GetCurrentModelDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EdgifyServiceServer is the server API for EdgifyService service.
type EdgifyServiceServer interface {
	GetPrediction(context.Context, *PredictionRequest) (*PredictionResponse, error)
	CreateGroundTruth(context.Context, *GroundTruthRequest) (*GroundTruthResponse, error)
	GetCurrentModelDeployment(context.Context, *GetCurrentModelDeploymentRequest) (*GetCurrentModelDeploymentResponse, error)
}

// UnimplementedEdgifyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEdgifyServiceServer struct {
}

func (*UnimplementedEdgifyServiceServer) GetPrediction(ctx context.Context, req *PredictionRequest) (*PredictionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrediction not implemented")
}
func (*UnimplementedEdgifyServiceServer) CreateGroundTruth(ctx context.Context, req *GroundTruthRequest) (*GroundTruthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroundTruth not implemented")
}
func (*UnimplementedEdgifyServiceServer) GetCurrentModelDeployment(ctx context.Context, req *GetCurrentModelDeploymentRequest) (*GetCurrentModelDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentModelDeployment not implemented")
}

func RegisterEdgifyServiceServer(s *grpc.Server, srv EdgifyServiceServer) {
	s.RegisterService(&_EdgifyService_serviceDesc, srv)
}

func _EdgifyService_GetPrediction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgifyServiceServer).GetPrediction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgify.EdgifyService/GetPrediction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgifyServiceServer).GetPrediction(ctx, req.(*PredictionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EdgifyService_CreateGroundTruth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroundTruthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgifyServiceServer).CreateGroundTruth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgify.EdgifyService/CreateGroundTruth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgifyServiceServer).CreateGroundTruth(ctx, req.(*GroundTruthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EdgifyService_GetCurrentModelDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentModelDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgifyServiceServer).GetCurrentModelDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgify.EdgifyService/GetCurrentModelDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgifyServiceServer).GetCurrentModelDeployment(ctx, req.(*GetCurrentModelDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EdgifyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "edgify.EdgifyService",
	HandlerType: (*EdgifyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrediction",
			Handler:    _EdgifyService_GetPrediction_Handler,
		},
		{
			MethodName: "CreateGroundTruth",
			Handler:    _EdgifyService_CreateGroundTruth_Handler,
		},
		{
			MethodName: "GetCurrentModelDeployment",
			Handler:    _EdgifyService_GetCurrentModelDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/edgify/edgifypb/prediction.proto",
}