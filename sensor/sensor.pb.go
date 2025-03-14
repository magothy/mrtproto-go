// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: sensor.proto

package sensor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Range_Type int32

const (
	Range_UNKNOWN    Range_Type = 0
	Range_ULTRASOUND Range_Type = 1
	Range_INFRARED   Range_Type = 2
	Range_LASER      Range_Type = 3
)

// Enum value maps for Range_Type.
var (
	Range_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "ULTRASOUND",
		2: "INFRARED",
		3: "LASER",
	}
	Range_Type_value = map[string]int32{
		"UNKNOWN":    0,
		"ULTRASOUND": 1,
		"INFRARED":   2,
		"LASER":      3,
	}
)

func (x Range_Type) Enum() *Range_Type {
	p := new(Range_Type)
	*p = x
	return p
}

func (x Range_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Range_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_sensor_proto_enumTypes[0].Descriptor()
}

func (Range_Type) Type() protoreflect.EnumType {
	return &file_sensor_proto_enumTypes[0]
}

func (x Range_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Range_Type.Descriptor instead.
func (Range_Type) EnumDescriptor() ([]byte, []int) {
	return file_sensor_proto_rawDescGZIP(), []int{1, 0}
}

type Pose struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	LatDeg     float64                `protobuf:"fixed64,1,opt,name=lat_deg,json=latDeg,proto3" json:"lat_deg,omitempty"`
	LonDeg     float64                `protobuf:"fixed64,2,opt,name=lon_deg,json=lonDeg,proto3" json:"lon_deg,omitempty"`
	HeadingDeg float32                `protobuf:"fixed32,3,opt,name=heading_deg,json=headingDeg,proto3" json:"heading_deg,omitempty"`
	// Position covariance [m^2] defined relative to a tangential plane
	//
	//	through the reported position. The components are East, North, and
	//	Up (ENU), in row-major order.
	PositionCovariance []float32 `protobuf:"fixed32,4,rep,packed,name=position_covariance,json=positionCovariance,proto3" json:"position_covariance,omitempty"`
	HeadingErrorDeg    *float32  `protobuf:"fixed32,5,opt,name=heading_error_deg,json=headingErrorDeg,proto3,oneof" json:"heading_error_deg,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Pose) Reset() {
	*x = Pose{}
	mi := &file_sensor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pose) ProtoMessage() {}

func (x *Pose) ProtoReflect() protoreflect.Message {
	mi := &file_sensor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pose.ProtoReflect.Descriptor instead.
func (*Pose) Descriptor() ([]byte, []int) {
	return file_sensor_proto_rawDescGZIP(), []int{0}
}

func (x *Pose) GetLatDeg() float64 {
	if x != nil {
		return x.LatDeg
	}
	return 0
}

func (x *Pose) GetLonDeg() float64 {
	if x != nil {
		return x.LonDeg
	}
	return 0
}

func (x *Pose) GetHeadingDeg() float32 {
	if x != nil {
		return x.HeadingDeg
	}
	return 0
}

func (x *Pose) GetPositionCovariance() []float32 {
	if x != nil {
		return x.PositionCovariance
	}
	return nil
}

func (x *Pose) GetHeadingErrorDeg() float32 {
	if x != nil && x.HeadingErrorDeg != nil {
		return *x.HeadingErrorDeg
	}
	return 0
}

type Range struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	TtagSystem        *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ttag_system,json=ttagSystem,proto3" json:"ttag_system,omitempty"`
	TtagSteadyNs      uint64                 `protobuf:"varint,2,opt,name=ttag_steady_ns,json=ttagSteadyNs,proto3" json:"ttag_steady_ns,omitempty"`
	Type              Range_Type             `protobuf:"varint,3,opt,name=type,proto3,enum=magothy.protobuf.sensor.Range_Type" json:"type,omitempty"`
	RangeM            float64                `protobuf:"fixed64,4,opt,name=range_m,json=rangeM,proto3" json:"range_m,omitempty"`
	RangeUncertaintyM float64                `protobuf:"fixed64,5,opt,name=range_uncertainty_m,json=rangeUncertaintyM,proto3" json:"range_uncertainty_m,omitempty"`
	MaxRangeM         float64                `protobuf:"fixed64,8,opt,name=max_range_m,json=maxRangeM,proto3" json:"max_range_m,omitempty"`
	FieldOfViewDeg    float64                `protobuf:"fixed64,6,opt,name=field_of_view_deg,json=fieldOfViewDeg,proto3" json:"field_of_view_deg,omitempty"`
	Pose              *Pose                  `protobuf:"bytes,7,opt,name=pose,proto3,oneof" json:"pose,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Range) Reset() {
	*x = Range{}
	mi := &file_sensor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Range) ProtoMessage() {}

func (x *Range) ProtoReflect() protoreflect.Message {
	mi := &file_sensor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Range.ProtoReflect.Descriptor instead.
func (*Range) Descriptor() ([]byte, []int) {
	return file_sensor_proto_rawDescGZIP(), []int{1}
}

func (x *Range) GetTtagSystem() *timestamppb.Timestamp {
	if x != nil {
		return x.TtagSystem
	}
	return nil
}

func (x *Range) GetTtagSteadyNs() uint64 {
	if x != nil {
		return x.TtagSteadyNs
	}
	return 0
}

func (x *Range) GetType() Range_Type {
	if x != nil {
		return x.Type
	}
	return Range_UNKNOWN
}

func (x *Range) GetRangeM() float64 {
	if x != nil {
		return x.RangeM
	}
	return 0
}

func (x *Range) GetRangeUncertaintyM() float64 {
	if x != nil {
		return x.RangeUncertaintyM
	}
	return 0
}

func (x *Range) GetMaxRangeM() float64 {
	if x != nil {
		return x.MaxRangeM
	}
	return 0
}

func (x *Range) GetFieldOfViewDeg() float64 {
	if x != nil {
		return x.FieldOfViewDeg
	}
	return 0
}

func (x *Range) GetPose() *Pose {
	if x != nil {
		return x.Pose
	}
	return nil
}

var File_sensor_proto protoreflect.FileDescriptor

var file_sensor_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x6d, 0x61, 0x67, 0x6f, 0x74, 0x68, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x5f, 0x64, 0x65, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x44, 0x65, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f,
	0x6e, 0x5f, 0x64, 0x65, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6c, 0x6f, 0x6e,
	0x44, 0x65, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64,
	0x65, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x67, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x02, 0x52, 0x12, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x11, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x48, 0x00, 0x52, 0x0f, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x44, 0x65, 0x67, 0x88, 0x01, 0x01, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x67, 0x22, 0xb6, 0x03, 0x0a,
	0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x74, 0x61, 0x67, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x74, 0x74, 0x61, 0x67, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x74, 0x61, 0x67, 0x5f, 0x73, 0x74, 0x65, 0x61,
	0x64, 0x79, 0x5f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x74, 0x61,
	0x67, 0x53, 0x74, 0x65, 0x61, 0x64, 0x79, 0x4e, 0x73, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6d, 0x61, 0x67, 0x6f, 0x74, 0x68,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x12, 0x2e, 0x0a, 0x13, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x75, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x79,
	0x5f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x55,
	0x6e, 0x63, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x79, 0x4d, 0x12, 0x1e, 0x0a, 0x0b, 0x6d,
	0x61, 0x78, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x12, 0x29, 0x0a, 0x11, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x66, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x64, 0x65, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x66, 0x56,
	0x69, 0x65, 0x77, 0x44, 0x65, 0x67, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x67, 0x6f, 0x74, 0x68, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x50,
	0x6f, 0x73, 0x65, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x88, 0x01, 0x01, 0x22, 0x3c,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4c, 0x54, 0x52, 0x41, 0x53, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x46, 0x52, 0x41, 0x52, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x41, 0x53, 0x45, 0x52, 0x10, 0x03, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x70, 0x6f, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_sensor_proto_rawDescOnce sync.Once
	file_sensor_proto_rawDescData []byte
)

func file_sensor_proto_rawDescGZIP() []byte {
	file_sensor_proto_rawDescOnce.Do(func() {
		file_sensor_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sensor_proto_rawDesc), len(file_sensor_proto_rawDesc)))
	})
	return file_sensor_proto_rawDescData
}

var file_sensor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sensor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sensor_proto_goTypes = []any{
	(Range_Type)(0),               // 0: magothy.protobuf.sensor.Range.Type
	(*Pose)(nil),                  // 1: magothy.protobuf.sensor.Pose
	(*Range)(nil),                 // 2: magothy.protobuf.sensor.Range
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_sensor_proto_depIdxs = []int32{
	3, // 0: magothy.protobuf.sensor.Range.ttag_system:type_name -> google.protobuf.Timestamp
	0, // 1: magothy.protobuf.sensor.Range.type:type_name -> magothy.protobuf.sensor.Range.Type
	1, // 2: magothy.protobuf.sensor.Range.pose:type_name -> magothy.protobuf.sensor.Pose
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sensor_proto_init() }
func file_sensor_proto_init() {
	if File_sensor_proto != nil {
		return
	}
	file_sensor_proto_msgTypes[0].OneofWrappers = []any{}
	file_sensor_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sensor_proto_rawDesc), len(file_sensor_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sensor_proto_goTypes,
		DependencyIndexes: file_sensor_proto_depIdxs,
		EnumInfos:         file_sensor_proto_enumTypes,
		MessageInfos:      file_sensor_proto_msgTypes,
	}.Build()
	File_sensor_proto = out.File
	file_sensor_proto_goTypes = nil
	file_sensor_proto_depIdxs = nil
}
