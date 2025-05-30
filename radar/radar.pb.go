// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: radar.proto

package radar

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

type State int32

const (
	State_STATE_OFF           State = 0
	State_STATE_STANDBY       State = 1
	State_STATE_WARMING_UP    State = 2
	State_STATE_TIMED_IDLE    State = 3
	State_STATE_STOPPING      State = 4
	State_STATE_SPINNING_DOWN State = 5
	State_STATE_STARTING      State = 6
	State_STATE_SPINNING_UP   State = 7
	State_STATE_TRANSMIT      State = 8
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "STATE_OFF",
		1: "STATE_STANDBY",
		2: "STATE_WARMING_UP",
		3: "STATE_TIMED_IDLE",
		4: "STATE_STOPPING",
		5: "STATE_SPINNING_DOWN",
		6: "STATE_STARTING",
		7: "STATE_SPINNING_UP",
		8: "STATE_TRANSMIT",
	}
	State_value = map[string]int32{
		"STATE_OFF":           0,
		"STATE_STANDBY":       1,
		"STATE_WARMING_UP":    2,
		"STATE_TIMED_IDLE":    3,
		"STATE_STOPPING":      4,
		"STATE_SPINNING_DOWN": 5,
		"STATE_STARTING":      6,
		"STATE_SPINNING_UP":   7,
		"STATE_TRANSMIT":      8,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_radar_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_radar_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_radar_proto_rawDescGZIP(), []int{0}
}

type CommandType int32

const (
	CommandType_COMMAND_NONE      CommandType = 0 // value ignored
	CommandType_COMMAND_TURN_ON   CommandType = 1 // value ignored
	CommandType_COMMAND_TURN_OFF  CommandType = 2 // value ignored
	CommandType_COMMAND_SET_RANGE CommandType = 3 // value is range in meters
	CommandType_COMMAND_SET_GAIN  CommandType = 4
	CommandType_COMMAND_SET_RAIN  CommandType = 5
	CommandType_COMMAND_SET_SEA   CommandType = 6
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0: "COMMAND_NONE",
		1: "COMMAND_TURN_ON",
		2: "COMMAND_TURN_OFF",
		3: "COMMAND_SET_RANGE",
		4: "COMMAND_SET_GAIN",
		5: "COMMAND_SET_RAIN",
		6: "COMMAND_SET_SEA",
	}
	CommandType_value = map[string]int32{
		"COMMAND_NONE":      0,
		"COMMAND_TURN_ON":   1,
		"COMMAND_TURN_OFF":  2,
		"COMMAND_SET_RANGE": 3,
		"COMMAND_SET_GAIN":  4,
		"COMMAND_SET_RAIN":  5,
		"COMMAND_SET_SEA":   6,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_radar_proto_enumTypes[1].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_radar_proto_enumTypes[1]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_radar_proto_rawDescGZIP(), []int{1}
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
	mi := &file_radar_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pose) ProtoMessage() {}

func (x *Pose) ProtoReflect() protoreflect.Message {
	mi := &file_radar_proto_msgTypes[0]
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
	return file_radar_proto_rawDescGZIP(), []int{0}
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

type Spokes struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	NumSpoke        uint32                 `protobuf:"varint,1,opt,name=num_spoke,json=numSpoke,proto3" json:"num_spoke,omitempty"` // total number of spokes in one complete scan
	FirstSpokeIndex uint32                 `protobuf:"varint,2,opt,name=first_spoke_index,json=firstSpokeIndex,proto3" json:"first_spoke_index,omitempty"`
	RangeM          float32                `protobuf:"fixed32,3,opt,name=range_m,json=rangeM,proto3" json:"range_m,omitempty"`
	Spokes          [][]byte               `protobuf:"bytes,4,rep,name=spokes,proto3" json:"spokes,omitempty"` // contains multiple spokes. each spoke is a byte
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Spokes) Reset() {
	*x = Spokes{}
	mi := &file_radar_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Spokes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spokes) ProtoMessage() {}

func (x *Spokes) ProtoReflect() protoreflect.Message {
	mi := &file_radar_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spokes.ProtoReflect.Descriptor instead.
func (*Spokes) Descriptor() ([]byte, []int) {
	return file_radar_proto_rawDescGZIP(), []int{1}
}

func (x *Spokes) GetNumSpoke() uint32 {
	if x != nil {
		return x.NumSpoke
	}
	return 0
}

func (x *Spokes) GetFirstSpokeIndex() uint32 {
	if x != nil {
		return x.FirstSpokeIndex
	}
	return 0
}

func (x *Spokes) GetRangeM() float32 {
	if x != nil {
		return x.RangeM
	}
	return 0
}

func (x *Spokes) GetSpokes() [][]byte {
	if x != nil {
		return x.Spokes
	}
	return nil
}

type Info struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TtagSteadyNs  uint64                 `protobuf:"varint,9,opt,name=ttag_steady_ns,json=ttagSteadyNs,proto3" json:"ttag_steady_ns,omitempty"`
	State         State                  `protobuf:"varint,2,opt,name=state,proto3,enum=magothy.protobuf.radar.State" json:"state,omitempty"`
	Gain          int32                  `protobuf:"varint,3,opt,name=gain,proto3" json:"gain,omitempty"`
	Rain          int32                  `protobuf:"varint,4,opt,name=rain,proto3" json:"rain,omitempty"`
	Sea           int32                  `protobuf:"varint,5,opt,name=sea,proto3" json:"sea,omitempty"`
	Range         int32                  `protobuf:"varint,6,opt,name=range,proto3" json:"range,omitempty"`
	ScanSpeed     int32                  `protobuf:"varint,7,opt,name=scan_speed,json=scanSpeed,proto3" json:"scan_speed,omitempty"`
	Spokes        *Spokes                `protobuf:"bytes,8,opt,name=spokes,proto3,oneof" json:"spokes,omitempty"`
	Pose          *Pose                  `protobuf:"bytes,10,opt,name=pose,proto3,oneof" json:"pose,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Info) Reset() {
	*x = Info{}
	mi := &file_radar_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_radar_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_radar_proto_rawDescGZIP(), []int{2}
}

func (x *Info) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Info) GetTtagSteadyNs() uint64 {
	if x != nil {
		return x.TtagSteadyNs
	}
	return 0
}

func (x *Info) GetState() State {
	if x != nil {
		return x.State
	}
	return State_STATE_OFF
}

func (x *Info) GetGain() int32 {
	if x != nil {
		return x.Gain
	}
	return 0
}

func (x *Info) GetRain() int32 {
	if x != nil {
		return x.Rain
	}
	return 0
}

func (x *Info) GetSea() int32 {
	if x != nil {
		return x.Sea
	}
	return 0
}

func (x *Info) GetRange() int32 {
	if x != nil {
		return x.Range
	}
	return 0
}

func (x *Info) GetScanSpeed() int32 {
	if x != nil {
		return x.ScanSpeed
	}
	return 0
}

func (x *Info) GetSpokes() *Spokes {
	if x != nil {
		return x.Spokes
	}
	return nil
}

func (x *Info) GetPose() *Pose {
	if x != nil {
		return x.Pose
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Command       CommandType            `protobuf:"varint,1,opt,name=command,proto3,enum=magothy.protobuf.radar.CommandType" json:"command,omitempty"`
	Value         int64                  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Command) Reset() {
	*x = Command{}
	mi := &file_radar_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_radar_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_radar_proto_rawDescGZIP(), []int{3}
}

func (x *Command) GetCommand() CommandType {
	if x != nil {
		return x.Command
	}
	return CommandType_COMMAND_NONE
}

func (x *Command) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_radar_proto protoreflect.FileDescriptor

const file_radar_proto_rawDesc = "" +
	"\n" +
	"\vradar.proto\x12\x16magothy.protobuf.radar\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd1\x01\n" +
	"\x04Pose\x12\x17\n" +
	"\alat_deg\x18\x01 \x01(\x01R\x06latDeg\x12\x17\n" +
	"\alon_deg\x18\x02 \x01(\x01R\x06lonDeg\x12\x1f\n" +
	"\vheading_deg\x18\x03 \x01(\x02R\n" +
	"headingDeg\x12/\n" +
	"\x13position_covariance\x18\x04 \x03(\x02R\x12positionCovariance\x12/\n" +
	"\x11heading_error_deg\x18\x05 \x01(\x02H\x00R\x0fheadingErrorDeg\x88\x01\x01B\x14\n" +
	"\x12_heading_error_deg\"\x82\x01\n" +
	"\x06Spokes\x12\x1b\n" +
	"\tnum_spoke\x18\x01 \x01(\rR\bnumSpoke\x12*\n" +
	"\x11first_spoke_index\x18\x02 \x01(\rR\x0ffirstSpokeIndex\x12\x17\n" +
	"\arange_m\x18\x03 \x01(\x02R\x06rangeM\x12\x16\n" +
	"\x06spokes\x18\x04 \x03(\fR\x06spokes\"\x92\x03\n" +
	"\x04Info\x128\n" +
	"\ttimestamp\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12$\n" +
	"\x0ettag_steady_ns\x18\t \x01(\x04R\fttagSteadyNs\x123\n" +
	"\x05state\x18\x02 \x01(\x0e2\x1d.magothy.protobuf.radar.StateR\x05state\x12\x12\n" +
	"\x04gain\x18\x03 \x01(\x05R\x04gain\x12\x12\n" +
	"\x04rain\x18\x04 \x01(\x05R\x04rain\x12\x10\n" +
	"\x03sea\x18\x05 \x01(\x05R\x03sea\x12\x14\n" +
	"\x05range\x18\x06 \x01(\x05R\x05range\x12\x1d\n" +
	"\n" +
	"scan_speed\x18\a \x01(\x05R\tscanSpeed\x12;\n" +
	"\x06spokes\x18\b \x01(\v2\x1e.magothy.protobuf.radar.SpokesH\x00R\x06spokes\x88\x01\x01\x125\n" +
	"\x04pose\x18\n" +
	" \x01(\v2\x1c.magothy.protobuf.radar.PoseH\x01R\x04pose\x88\x01\x01B\t\n" +
	"\a_spokesB\a\n" +
	"\x05_pose\"^\n" +
	"\aCommand\x12=\n" +
	"\acommand\x18\x01 \x01(\x0e2#.magothy.protobuf.radar.CommandTypeR\acommand\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x03R\x05value*\xc1\x01\n" +
	"\x05State\x12\r\n" +
	"\tSTATE_OFF\x10\x00\x12\x11\n" +
	"\rSTATE_STANDBY\x10\x01\x12\x14\n" +
	"\x10STATE_WARMING_UP\x10\x02\x12\x14\n" +
	"\x10STATE_TIMED_IDLE\x10\x03\x12\x12\n" +
	"\x0eSTATE_STOPPING\x10\x04\x12\x17\n" +
	"\x13STATE_SPINNING_DOWN\x10\x05\x12\x12\n" +
	"\x0eSTATE_STARTING\x10\x06\x12\x15\n" +
	"\x11STATE_SPINNING_UP\x10\a\x12\x12\n" +
	"\x0eSTATE_TRANSMIT\x10\b*\xa2\x01\n" +
	"\vCommandType\x12\x10\n" +
	"\fCOMMAND_NONE\x10\x00\x12\x13\n" +
	"\x0fCOMMAND_TURN_ON\x10\x01\x12\x14\n" +
	"\x10COMMAND_TURN_OFF\x10\x02\x12\x15\n" +
	"\x11COMMAND_SET_RANGE\x10\x03\x12\x14\n" +
	"\x10COMMAND_SET_GAIN\x10\x04\x12\x14\n" +
	"\x10COMMAND_SET_RAIN\x10\x05\x12\x13\n" +
	"\x0fCOMMAND_SET_SEA\x10\x06b\x06proto3"

var (
	file_radar_proto_rawDescOnce sync.Once
	file_radar_proto_rawDescData []byte
)

func file_radar_proto_rawDescGZIP() []byte {
	file_radar_proto_rawDescOnce.Do(func() {
		file_radar_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_radar_proto_rawDesc), len(file_radar_proto_rawDesc)))
	})
	return file_radar_proto_rawDescData
}

var file_radar_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_radar_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_radar_proto_goTypes = []any{
	(State)(0),                    // 0: magothy.protobuf.radar.State
	(CommandType)(0),              // 1: magothy.protobuf.radar.CommandType
	(*Pose)(nil),                  // 2: magothy.protobuf.radar.Pose
	(*Spokes)(nil),                // 3: magothy.protobuf.radar.Spokes
	(*Info)(nil),                  // 4: magothy.protobuf.radar.Info
	(*Command)(nil),               // 5: magothy.protobuf.radar.Command
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_radar_proto_depIdxs = []int32{
	6, // 0: magothy.protobuf.radar.Info.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: magothy.protobuf.radar.Info.state:type_name -> magothy.protobuf.radar.State
	3, // 2: magothy.protobuf.radar.Info.spokes:type_name -> magothy.protobuf.radar.Spokes
	2, // 3: magothy.protobuf.radar.Info.pose:type_name -> magothy.protobuf.radar.Pose
	1, // 4: magothy.protobuf.radar.Command.command:type_name -> magothy.protobuf.radar.CommandType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_radar_proto_init() }
func file_radar_proto_init() {
	if File_radar_proto != nil {
		return
	}
	file_radar_proto_msgTypes[0].OneofWrappers = []any{}
	file_radar_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_radar_proto_rawDesc), len(file_radar_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_radar_proto_goTypes,
		DependencyIndexes: file_radar_proto_depIdxs,
		EnumInfos:         file_radar_proto_enumTypes,
		MessageInfos:      file_radar_proto_msgTypes,
	}.Build()
	File_radar_proto = out.File
	file_radar_proto_goTypes = nil
	file_radar_proto_depIdxs = nil
}
