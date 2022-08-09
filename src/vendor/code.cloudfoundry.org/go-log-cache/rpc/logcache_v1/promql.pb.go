// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/v1/promql.proto

package logcache_v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PromQL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PromQL) Reset() {
	*x = PromQL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_promql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL) ProtoMessage() {}

func (x *PromQL) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_promql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL.ProtoReflect.Descriptor instead.
func (*PromQL) Descriptor() ([]byte, []int) {
	return file_api_v1_promql_proto_rawDescGZIP(), []int{0}
}

type PromQL_InstantQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Time  string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *PromQL_InstantQueryRequest) Reset() {
	*x = PromQL_InstantQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_promql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL_InstantQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL_InstantQueryRequest) ProtoMessage() {}

func (x *PromQL_InstantQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_promql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL_InstantQueryRequest.ProtoReflect.Descriptor instead.
func (*PromQL_InstantQueryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_promql_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PromQL_InstantQueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *PromQL_InstantQueryRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type PromQL_RangeQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Start string `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End   string `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	Step  string `protobuf:"bytes,4,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *PromQL_RangeQueryRequest) Reset() {
	*x = PromQL_RangeQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_promql_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL_RangeQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL_RangeQueryRequest) ProtoMessage() {}

func (x *PromQL_RangeQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_promql_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL_RangeQueryRequest.ProtoReflect.Descriptor instead.
func (*PromQL_RangeQueryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_promql_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PromQL_RangeQueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *PromQL_RangeQueryRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *PromQL_RangeQueryRequest) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *PromQL_RangeQueryRequest) GetStep() string {
	if x != nil {
		return x.Step
	}
	return ""
}

type PromQL_InstantQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*PromQL_InstantQueryResult_Scalar
	//	*PromQL_InstantQueryResult_Vector
	//	*PromQL_InstantQueryResult_Matrix
	Result isPromQL_InstantQueryResult_Result `protobuf_oneof:"Result"`
}

func (x *PromQL_InstantQueryResult) Reset() {
	*x = PromQL_InstantQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_promql_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL_InstantQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL_InstantQueryResult) ProtoMessage() {}

func (x *PromQL_InstantQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_promql_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL_InstantQueryResult.ProtoReflect.Descriptor instead.
func (*PromQL_InstantQueryResult) Descriptor() ([]byte, []int) {
	return file_api_v1_promql_proto_rawDescGZIP(), []int{0, 2}
}

func (m *PromQL_InstantQueryResult) GetResult() isPromQL_InstantQueryResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *PromQL_InstantQueryResult) GetScalar() *PromQL_Scalar {
	if x, ok := x.GetResult().(*PromQL_InstantQueryResult_Scalar); ok {
		return x.Scalar
	}
	return nil
}

func (x *PromQL_InstantQueryResult) GetVector() *PromQL_Vector {
	if x, ok := x.GetResult().(*PromQL_InstantQueryResult_Vector); ok {
		return x.Vector
	}
	return nil
}

func (x *PromQL_InstantQueryResult) GetMatrix() *PromQL_Matrix {
	if x, ok := x.GetResult().(*PromQL_InstantQueryResult_Matrix); ok {
		return x.Matrix
	}
	return nil
}

type isPromQL_InstantQueryResult_Result interface {
	isPromQL_InstantQueryResult_Result()
}

type PromQL_InstantQueryResult_Scalar struct {
	Scalar *PromQL_Scalar `protobuf:"bytes,1,opt,name=scalar,proto3,oneof"`
}

type PromQL_InstantQueryResult_Vector struct {
	Vector *PromQL_Vector `protobuf:"bytes,2,opt,name=vector,proto3,oneof"`
}

type PromQL_InstantQueryResult_Matrix struct {
	Matrix *PromQL_Matrix `protobuf:"bytes,3,opt,name=matrix,proto3,oneof"`
}

func (*PromQL_InstantQueryResult_Scalar) isPromQL_InstantQueryResult_Result() {}

func (*PromQL_InstantQueryResult_Vector) isPromQL_InstantQueryResult_Result() {}

func (*PromQL_InstantQueryResult_Matrix) isPromQL_InstantQueryResult_Result() {}

type PromQL_RangeQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*PromQL_RangeQueryResult_Matrix
	Result isPromQL_RangeQueryResult_Result `protobuf_oneof:"Result"`
}

func (x *PromQL_RangeQueryResult) Reset() {
	*x = PromQL_RangeQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_promql_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL_RangeQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL_RangeQueryResult) ProtoMessage() {}

func (x *PromQL_RangeQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_promql_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL_RangeQueryResult.ProtoReflect.Descriptor instead.
func (*PromQL_RangeQueryResult) Descriptor() ([]byte, []int) {
	return file_api_v1_promql_proto_rawDescGZIP(), []int{0, 3}
}

func (m *PromQL_RangeQueryResult) GetResult() isPromQL_RangeQueryResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *PromQL_RangeQueryResult) GetMatrix() *PromQL_Matrix {
	if x, ok := x.GetResult().(*PromQL_RangeQueryResult_Matrix); ok {
		return x.Matrix
	}
	return nil
}

type isPromQL_RangeQueryResult_Result interface {
	isPromQL_RangeQueryResult_Result()
}

type PromQL_RangeQueryResult_Matrix struct {
	Matrix *PromQL_Matrix `protobuf:"bytes,1,opt,name=matrix,proto3,oneof"`
}

func (*PromQL_RangeQueryResult_Matrix) isPromQL_RangeQueryResult_Result() {}

type PromQL_Scalar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  string  `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PromQL_Scalar) Reset() {
	*x = PromQL_Scalar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_promql_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL_Scalar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL_Scalar) ProtoMessage() {}

func (x *PromQL_Scalar) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_promql_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL_Scalar.ProtoReflect.Descriptor instead.
func (*PromQL_Scalar) Descriptor() ([]byte, []int) {
	return file_api_v1_promql_proto_rawDescGZIP(), []int{0, 4}
}

func (x *PromQL_Scalar) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *PromQL_Scalar) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type PromQL_Vector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Samples []*PromQL_Sample `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *PromQL_Vector) Reset() {
	*x = PromQL_Vector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_promql_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL_Vector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL_Vector) ProtoMessage() {}

func (x *PromQL_Vector) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_promql_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL_Vector.ProtoReflect.Descriptor instead.
func (*PromQL_Vector) Descriptor() ([]byte, []int) {
	return file_api_v1_promql_proto_rawDescGZIP(), []int{0, 5}
}

func (x *PromQL_Vector) GetSamples() []*PromQL_Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

type PromQL_Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  string  `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PromQL_Point) Reset() {
	*x = PromQL_Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_promql_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL_Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL_Point) ProtoMessage() {}

func (x *PromQL_Point) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_promql_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL_Point.ProtoReflect.Descriptor instead.
func (*PromQL_Point) Descriptor() ([]byte, []int) {
	return file_api_v1_promql_proto_rawDescGZIP(), []int{0, 6}
}

func (x *PromQL_Point) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *PromQL_Point) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type PromQL_Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric map[string]string `protobuf:"bytes,1,rep,name=metric,proto3" json:"metric,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Point  *PromQL_Point     `protobuf:"bytes,2,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *PromQL_Sample) Reset() {
	*x = PromQL_Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_promql_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL_Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL_Sample) ProtoMessage() {}

func (x *PromQL_Sample) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_promql_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL_Sample.ProtoReflect.Descriptor instead.
func (*PromQL_Sample) Descriptor() ([]byte, []int) {
	return file_api_v1_promql_proto_rawDescGZIP(), []int{0, 7}
}

func (x *PromQL_Sample) GetMetric() map[string]string {
	if x != nil {
		return x.Metric
	}
	return nil
}

func (x *PromQL_Sample) GetPoint() *PromQL_Point {
	if x != nil {
		return x.Point
	}
	return nil
}

type PromQL_Matrix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Series []*PromQL_Series `protobuf:"bytes,1,rep,name=series,proto3" json:"series,omitempty"`
}

func (x *PromQL_Matrix) Reset() {
	*x = PromQL_Matrix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_promql_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL_Matrix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL_Matrix) ProtoMessage() {}

func (x *PromQL_Matrix) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_promql_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL_Matrix.ProtoReflect.Descriptor instead.
func (*PromQL_Matrix) Descriptor() ([]byte, []int) {
	return file_api_v1_promql_proto_rawDescGZIP(), []int{0, 8}
}

func (x *PromQL_Matrix) GetSeries() []*PromQL_Series {
	if x != nil {
		return x.Series
	}
	return nil
}

type PromQL_Series struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric map[string]string `protobuf:"bytes,1,rep,name=metric,proto3" json:"metric,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Points []*PromQL_Point   `protobuf:"bytes,2,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *PromQL_Series) Reset() {
	*x = PromQL_Series{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_promql_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL_Series) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL_Series) ProtoMessage() {}

func (x *PromQL_Series) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_promql_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL_Series.ProtoReflect.Descriptor instead.
func (*PromQL_Series) Descriptor() ([]byte, []int) {
	return file_api_v1_promql_proto_rawDescGZIP(), []int{0, 9}
}

func (x *PromQL_Series) GetMetric() map[string]string {
	if x != nil {
		return x.Metric
	}
	return nil
}

func (x *PromQL_Series) GetPoints() []*PromQL_Point {
	if x != nil {
		return x.Points
	}
	return nil
}

var File_api_v1_promql_proto protoreflect.FileDescriptor

var file_api_v1_promql_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x71, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9c, 0x08, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x1a, 0x3f, 0x0a, 0x13, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x65, 0x0a, 0x11,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x1a, 0xc0, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x63,
	0x61, 0x6c, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x6f, 0x67,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e,
	0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x48, 0x00, 0x52, 0x06, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72,
	0x12, 0x34, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x06,
	0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e, 0x4d, 0x61, 0x74, 0x72,
	0x69, 0x78, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x42, 0x08, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x52, 0x0a, 0x10, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x6f, 0x67,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e,
	0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x42, 0x08, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x32, 0x0a, 0x06, 0x53, 0x63,
	0x61, 0x6c, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x3e,
	0x0a, 0x06, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x1a, 0x31,
	0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0xb4, 0x01, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c,
	0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x51,
	0x4c, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x2f, 0x0a, 0x05,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x6f,
	0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c,
	0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x39, 0x0a,
	0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x06, 0x4d, 0x61, 0x74, 0x72,
	0x69, 0x78, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x1a, 0xb6, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x3e, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x12, 0x31, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0xff, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65,
	0x72, 0x12, 0x76, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x27, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6c, 0x6f, 0x67,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x76, 0x0a, 0x0a, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x25, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x6d, 0x51, 0x4c, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_promql_proto_rawDescOnce sync.Once
	file_api_v1_promql_proto_rawDescData = file_api_v1_promql_proto_rawDesc
)

func file_api_v1_promql_proto_rawDescGZIP() []byte {
	file_api_v1_promql_proto_rawDescOnce.Do(func() {
		file_api_v1_promql_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_promql_proto_rawDescData)
	})
	return file_api_v1_promql_proto_rawDescData
}

var file_api_v1_promql_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_v1_promql_proto_goTypes = []interface{}{
	(*PromQL)(nil),                     // 0: logcache.v1.PromQL
	(*PromQL_InstantQueryRequest)(nil), // 1: logcache.v1.PromQL.InstantQueryRequest
	(*PromQL_RangeQueryRequest)(nil),   // 2: logcache.v1.PromQL.RangeQueryRequest
	(*PromQL_InstantQueryResult)(nil),  // 3: logcache.v1.PromQL.InstantQueryResult
	(*PromQL_RangeQueryResult)(nil),    // 4: logcache.v1.PromQL.RangeQueryResult
	(*PromQL_Scalar)(nil),              // 5: logcache.v1.PromQL.Scalar
	(*PromQL_Vector)(nil),              // 6: logcache.v1.PromQL.Vector
	(*PromQL_Point)(nil),               // 7: logcache.v1.PromQL.Point
	(*PromQL_Sample)(nil),              // 8: logcache.v1.PromQL.Sample
	(*PromQL_Matrix)(nil),              // 9: logcache.v1.PromQL.Matrix
	(*PromQL_Series)(nil),              // 10: logcache.v1.PromQL.Series
	nil,                                // 11: logcache.v1.PromQL.Sample.MetricEntry
	nil,                                // 12: logcache.v1.PromQL.Series.MetricEntry
}
var file_api_v1_promql_proto_depIdxs = []int32{
	5,  // 0: logcache.v1.PromQL.InstantQueryResult.scalar:type_name -> logcache.v1.PromQL.Scalar
	6,  // 1: logcache.v1.PromQL.InstantQueryResult.vector:type_name -> logcache.v1.PromQL.Vector
	9,  // 2: logcache.v1.PromQL.InstantQueryResult.matrix:type_name -> logcache.v1.PromQL.Matrix
	9,  // 3: logcache.v1.PromQL.RangeQueryResult.matrix:type_name -> logcache.v1.PromQL.Matrix
	8,  // 4: logcache.v1.PromQL.Vector.samples:type_name -> logcache.v1.PromQL.Sample
	11, // 5: logcache.v1.PromQL.Sample.metric:type_name -> logcache.v1.PromQL.Sample.MetricEntry
	7,  // 6: logcache.v1.PromQL.Sample.point:type_name -> logcache.v1.PromQL.Point
	10, // 7: logcache.v1.PromQL.Matrix.series:type_name -> logcache.v1.PromQL.Series
	12, // 8: logcache.v1.PromQL.Series.metric:type_name -> logcache.v1.PromQL.Series.MetricEntry
	7,  // 9: logcache.v1.PromQL.Series.points:type_name -> logcache.v1.PromQL.Point
	1,  // 10: logcache.v1.PromQLQuerier.InstantQuery:input_type -> logcache.v1.PromQL.InstantQueryRequest
	2,  // 11: logcache.v1.PromQLQuerier.RangeQuery:input_type -> logcache.v1.PromQL.RangeQueryRequest
	3,  // 12: logcache.v1.PromQLQuerier.InstantQuery:output_type -> logcache.v1.PromQL.InstantQueryResult
	4,  // 13: logcache.v1.PromQLQuerier.RangeQuery:output_type -> logcache.v1.PromQL.RangeQueryResult
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_v1_promql_proto_init() }
func file_api_v1_promql_proto_init() {
	if File_api_v1_promql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_promql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_promql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL_InstantQueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_promql_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL_RangeQueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_promql_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL_InstantQueryResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_promql_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL_RangeQueryResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_promql_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL_Scalar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_promql_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL_Vector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_promql_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL_Point); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_promql_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL_Sample); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_promql_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL_Matrix); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_promql_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL_Series); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_v1_promql_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*PromQL_InstantQueryResult_Scalar)(nil),
		(*PromQL_InstantQueryResult_Vector)(nil),
		(*PromQL_InstantQueryResult_Matrix)(nil),
	}
	file_api_v1_promql_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*PromQL_RangeQueryResult_Matrix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_promql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_promql_proto_goTypes,
		DependencyIndexes: file_api_v1_promql_proto_depIdxs,
		MessageInfos:      file_api_v1_promql_proto_msgTypes,
	}.Build()
	File_api_v1_promql_proto = out.File
	file_api_v1_promql_proto_rawDesc = nil
	file_api_v1_promql_proto_goTypes = nil
	file_api_v1_promql_proto_depIdxs = nil
}
