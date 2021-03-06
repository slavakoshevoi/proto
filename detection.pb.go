// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.8.0
// source: detection.proto

package detection

import (
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

type ResultResponse_Type int32

const (
	ResultResponse_EXPRESSION ResultResponse_Type = 0
	ResultResponse_OBJECT     ResultResponse_Type = 1
)

// Enum value maps for ResultResponse_Type.
var (
	ResultResponse_Type_name = map[int32]string{
		0: "EXPRESSION",
		1: "OBJECT",
	}
	ResultResponse_Type_value = map[string]int32{
		"EXPRESSION": 0,
		"OBJECT":     1,
	}
)

func (x ResultResponse_Type) Enum() *ResultResponse_Type {
	p := new(ResultResponse_Type)
	*p = x
	return p
}

func (x ResultResponse_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultResponse_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_detection_proto_enumTypes[0].Descriptor()
}

func (ResultResponse_Type) Type() protoreflect.EnumType {
	return &file_detection_proto_enumTypes[0]
}

func (x ResultResponse_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultResponse_Type.Descriptor instead.
func (ResultResponse_Type) EnumDescriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{1, 0}
}

type ImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *ImageRequest_ImageInfo `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Image    []byte                  `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *ImageRequest) Reset() {
	*x = ImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageRequest) ProtoMessage() {}

func (x *ImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageRequest.ProtoReflect.Descriptor instead.
func (*ImageRequest) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{0}
}

func (x *ImageRequest) GetMetadata() *ImageRequest_ImageInfo {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ImageRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  ResultResponse_Type    `protobuf:"varint,1,opt,name=type,proto3,enum=grpc.ResultResponse_Type" json:"type,omitempty"`
	Value string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Score float32                `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	Bbox  []*ResultResponse_BBox `protobuf:"bytes,4,rep,name=bbox,proto3" json:"bbox,omitempty"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{1}
}

func (x *ResultResponse) GetType() ResultResponse_Type {
	if x != nil {
		return x.Type
	}
	return ResultResponse_EXPRESSION
}

func (x *ResultResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ResultResponse) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *ResultResponse) GetBbox() []*ResultResponse_BBox {
	if x != nil {
		return x.Bbox
	}
	return nil
}

type ImageRequest_ImageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageType string `protobuf:"bytes,2,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
}

func (x *ImageRequest_ImageInfo) Reset() {
	*x = ImageRequest_ImageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageRequest_ImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageRequest_ImageInfo) ProtoMessage() {}

func (x *ImageRequest_ImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageRequest_ImageInfo.ProtoReflect.Descriptor instead.
func (*ImageRequest_ImageInfo) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ImageRequest_ImageInfo) GetImageType() string {
	if x != nil {
		return x.ImageType
	}
	return ""
}

type ResultResponse_BBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*ResultResponse_BBox_X
	//	*ResultResponse_BBox_Y
	//	*ResultResponse_BBox_Width
	//	*ResultResponse_BBox_Height
	Value isResultResponse_BBox_Value `protobuf_oneof:"value"`
}

func (x *ResultResponse_BBox) Reset() {
	*x = ResultResponse_BBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResponse_BBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse_BBox) ProtoMessage() {}

func (x *ResultResponse_BBox) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse_BBox.ProtoReflect.Descriptor instead.
func (*ResultResponse_BBox) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{1, 0}
}

func (m *ResultResponse_BBox) GetValue() isResultResponse_BBox_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *ResultResponse_BBox) GetX() float32 {
	if x, ok := x.GetValue().(*ResultResponse_BBox_X); ok {
		return x.X
	}
	return 0
}

func (x *ResultResponse_BBox) GetY() float32 {
	if x, ok := x.GetValue().(*ResultResponse_BBox_Y); ok {
		return x.Y
	}
	return 0
}

func (x *ResultResponse_BBox) GetWidth() float32 {
	if x, ok := x.GetValue().(*ResultResponse_BBox_Width); ok {
		return x.Width
	}
	return 0
}

func (x *ResultResponse_BBox) GetHeight() float32 {
	if x, ok := x.GetValue().(*ResultResponse_BBox_Height); ok {
		return x.Height
	}
	return 0
}

type isResultResponse_BBox_Value interface {
	isResultResponse_BBox_Value()
}

type ResultResponse_BBox_X struct {
	X float32 `protobuf:"fixed32,1,opt,name=x,proto3,oneof"`
}

type ResultResponse_BBox_Y struct {
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3,oneof"`
}

type ResultResponse_BBox_Width struct {
	Width float32 `protobuf:"fixed32,3,opt,name=width,proto3,oneof"`
}

type ResultResponse_BBox_Height struct {
	Height float32 `protobuf:"fixed32,4,opt,name=height,proto3,oneof"`
}

func (*ResultResponse_BBox_X) isResultResponse_BBox_Value() {}

func (*ResultResponse_BBox_Y) isResultResponse_BBox_Value() {}

func (*ResultResponse_BBox_Width) isResultResponse_BBox_Value() {}

func (*ResultResponse_BBox_Height) isResultResponse_BBox_Value() {}

var File_detection_proto protoreflect.FileDescriptor

var file_detection_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x2a, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x22, 0xa1, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x62, 0x6f, 0x78, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x04, 0x62, 0x62, 0x6f,
	0x78, 0x1a, 0x61, 0x0a, 0x04, 0x42, 0x42, 0x6f, 0x78, 0x12, 0x0e, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x01, 0x78, 0x12, 0x0e, 0x0a, 0x01, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x01, 0x79, 0x12, 0x16, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x12, 0x18, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a,
	0x45, 0x58, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x01, 0x32, 0x47, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_detection_proto_rawDescOnce sync.Once
	file_detection_proto_rawDescData = file_detection_proto_rawDesc
)

func file_detection_proto_rawDescGZIP() []byte {
	file_detection_proto_rawDescOnce.Do(func() {
		file_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_detection_proto_rawDescData)
	})
	return file_detection_proto_rawDescData
}

var file_detection_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_detection_proto_goTypes = []interface{}{
	(ResultResponse_Type)(0),       // 0: grpc.ResultResponse.Type
	(*ImageRequest)(nil),           // 1: grpc.ImageRequest
	(*ResultResponse)(nil),         // 2: grpc.ResultResponse
	(*ImageRequest_ImageInfo)(nil), // 3: grpc.ImageRequest.ImageInfo
	(*ResultResponse_BBox)(nil),    // 4: grpc.ResultResponse.BBox
}
var file_detection_proto_depIdxs = []int32{
	3, // 0: grpc.ImageRequest.metadata:type_name -> grpc.ImageRequest.ImageInfo
	0, // 1: grpc.ResultResponse.type:type_name -> grpc.ResultResponse.Type
	4, // 2: grpc.ResultResponse.bbox:type_name -> grpc.ResultResponse.BBox
	1, // 3: grpc.ObjectsDetector.Detect:input_type -> grpc.ImageRequest
	2, // 4: grpc.ObjectsDetector.Detect:output_type -> grpc.ResultResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_detection_proto_init() }
func file_detection_proto_init() {
	if File_detection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageRequest); i {
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
		file_detection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultResponse); i {
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
		file_detection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageRequest_ImageInfo); i {
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
		file_detection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultResponse_BBox); i {
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
	file_detection_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ResultResponse_BBox_X)(nil),
		(*ResultResponse_BBox_Y)(nil),
		(*ResultResponse_BBox_Width)(nil),
		(*ResultResponse_BBox_Height)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_detection_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_detection_proto_goTypes,
		DependencyIndexes: file_detection_proto_depIdxs,
		EnumInfos:         file_detection_proto_enumTypes,
		MessageInfos:      file_detection_proto_msgTypes,
	}.Build()
	File_detection_proto = out.File
	file_detection_proto_rawDesc = nil
	file_detection_proto_goTypes = nil
	file_detection_proto_depIdxs = nil
}
