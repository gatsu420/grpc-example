// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.21.12
// source: common/model/garage.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type GarageCoordinate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Latitude      float32                `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     float32                `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GarageCoordinate) Reset() {
	*x = GarageCoordinate{}
	mi := &file_common_model_garage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GarageCoordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageCoordinate) ProtoMessage() {}

func (x *GarageCoordinate) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_garage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageCoordinate.ProtoReflect.Descriptor instead.
func (*GarageCoordinate) Descriptor() ([]byte, []int) {
	return file_common_model_garage_proto_rawDescGZIP(), []int{0}
}

func (x *GarageCoordinate) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GarageCoordinate) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Garage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Coordinate    *GarageCoordinate      `protobuf:"bytes,3,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Garage) Reset() {
	*x = Garage{}
	mi := &file_common_model_garage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Garage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Garage) ProtoMessage() {}

func (x *Garage) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_garage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Garage.ProtoReflect.Descriptor instead.
func (*Garage) Descriptor() ([]byte, []int) {
	return file_common_model_garage_proto_rawDescGZIP(), []int{1}
}

func (x *Garage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Garage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Garage) GetCoordinate() *GarageCoordinate {
	if x != nil {
		return x.Coordinate
	}
	return nil
}

type GarageList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Garage        []*Garage              `protobuf:"bytes,1,rep,name=garage,proto3" json:"garage,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GarageList) Reset() {
	*x = GarageList{}
	mi := &file_common_model_garage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GarageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageList) ProtoMessage() {}

func (x *GarageList) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_garage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageList.ProtoReflect.Descriptor instead.
func (*GarageList) Descriptor() ([]byte, []int) {
	return file_common_model_garage_proto_rawDescGZIP(), []int{2}
}

func (x *GarageList) GetGarage() []*Garage {
	if x != nil {
		return x.Garage
	}
	return nil
}

type GarageListByUser struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          map[string]*GarageList `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GarageListByUser) Reset() {
	*x = GarageListByUser{}
	mi := &file_common_model_garage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GarageListByUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageListByUser) ProtoMessage() {}

func (x *GarageListByUser) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_garage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageListByUser.ProtoReflect.Descriptor instead.
func (*GarageListByUser) Descriptor() ([]byte, []int) {
	return file_common_model_garage_proto_rawDescGZIP(), []int{3}
}

func (x *GarageListByUser) GetUser() map[string]*GarageList {
	if x != nil {
		return x.User
	}
	return nil
}

type GarageUserID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserID        string                 `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GarageUserID) Reset() {
	*x = GarageUserID{}
	mi := &file_common_model_garage_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GarageUserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageUserID) ProtoMessage() {}

func (x *GarageUserID) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_garage_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageUserID.ProtoReflect.Descriptor instead.
func (*GarageUserID) Descriptor() ([]byte, []int) {
	return file_common_model_garage_proto_rawDescGZIP(), []int{4}
}

func (x *GarageUserID) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GarageAndUserID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserID        string                 `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Garage        *Garage                `protobuf:"bytes,2,opt,name=garage,proto3" json:"garage,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GarageAndUserID) Reset() {
	*x = GarageAndUserID{}
	mi := &file_common_model_garage_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GarageAndUserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageAndUserID) ProtoMessage() {}

func (x *GarageAndUserID) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_garage_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageAndUserID.ProtoReflect.Descriptor instead.
func (*GarageAndUserID) Descriptor() ([]byte, []int) {
	return file_common_model_garage_proto_rawDescGZIP(), []int{5}
}

func (x *GarageAndUserID) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GarageAndUserID) GetGarage() *Garage {
	if x != nil {
		return x.Garage
	}
	return nil
}

var File_common_model_garage_proto protoreflect.FileDescriptor

var file_common_model_garage_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67,
	0x61, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4c, 0x0a, 0x10, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x65, 0x0a,
	0x06, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x22, 0x33, 0x0a, 0x0a, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x06, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x10, 0x47, 0x61,
	0x72, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x35,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x4a, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x26, 0x0a, 0x0c, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x50, 0x0a, 0x0f, 0x47, 0x61, 0x72,
	0x61, 0x67, 0x65, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x25, 0x0a, 0x06, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x06, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x32, 0x74, 0x0a, 0x07, 0x47,
	0x61, 0x72, 0x61, 0x67, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x41, 0x6e,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_common_model_garage_proto_rawDescOnce sync.Once
	file_common_model_garage_proto_rawDescData []byte
)

func file_common_model_garage_proto_rawDescGZIP() []byte {
	file_common_model_garage_proto_rawDescOnce.Do(func() {
		file_common_model_garage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_model_garage_proto_rawDesc), len(file_common_model_garage_proto_rawDesc)))
	})
	return file_common_model_garage_proto_rawDescData
}

var file_common_model_garage_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_common_model_garage_proto_goTypes = []any{
	(*GarageCoordinate)(nil), // 0: model.GarageCoordinate
	(*Garage)(nil),           // 1: model.Garage
	(*GarageList)(nil),       // 2: model.GarageList
	(*GarageListByUser)(nil), // 3: model.GarageListByUser
	(*GarageUserID)(nil),     // 4: model.GarageUserID
	(*GarageAndUserID)(nil),  // 5: model.GarageAndUserID
	nil,                      // 6: model.GarageListByUser.UserEntry
	(*emptypb.Empty)(nil),    // 7: google.protobuf.Empty
}
var file_common_model_garage_proto_depIdxs = []int32{
	0, // 0: model.Garage.coordinate:type_name -> model.GarageCoordinate
	1, // 1: model.GarageList.garage:type_name -> model.Garage
	6, // 2: model.GarageListByUser.user:type_name -> model.GarageListByUser.UserEntry
	1, // 3: model.GarageAndUserID.garage:type_name -> model.Garage
	2, // 4: model.GarageListByUser.UserEntry.value:type_name -> model.GarageList
	4, // 5: model.Garages.List:input_type -> model.GarageUserID
	5, // 6: model.Garages.Add:input_type -> model.GarageAndUserID
	2, // 7: model.Garages.List:output_type -> model.GarageList
	7, // 8: model.Garages.Add:output_type -> google.protobuf.Empty
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_model_garage_proto_init() }
func file_common_model_garage_proto_init() {
	if File_common_model_garage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_model_garage_proto_rawDesc), len(file_common_model_garage_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_model_garage_proto_goTypes,
		DependencyIndexes: file_common_model_garage_proto_depIdxs,
		MessageInfos:      file_common_model_garage_proto_msgTypes,
	}.Build()
	File_common_model_garage_proto = out.File
	file_common_model_garage_proto_goTypes = nil
	file_common_model_garage_proto_depIdxs = nil
}
