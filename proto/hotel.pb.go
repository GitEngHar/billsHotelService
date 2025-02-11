// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/hotel.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type HotelRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PricePernight  int32                  `protobuf:"varint,3,opt,name=price_pernight,json=pricePernight,proto3" json:"price_pernight,omitempty"`
	RoomsAvailable int32                  `protobuf:"varint,4,opt,name=rooms_available,json=roomsAvailable,proto3" json:"rooms_available,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *HotelRequest) Reset() {
	*x = HotelRequest{}
	mi := &file_proto_hotel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HotelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelRequest) ProtoMessage() {}

func (x *HotelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelRequest.ProtoReflect.Descriptor instead.
func (*HotelRequest) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{0}
}

func (x *HotelRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HotelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HotelRequest) GetPricePernight() int32 {
	if x != nil {
		return x.PricePernight
	}
	return 0
}

func (x *HotelRequest) GetRoomsAvailable() int32 {
	if x != nil {
		return x.RoomsAvailable
	}
	return 0
}

type Hotel struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PricePernight  int32                  `protobuf:"varint,3,opt,name=price_pernight,json=pricePernight,proto3" json:"price_pernight,omitempty"`
	RoomsAvailable int32                  `protobuf:"varint,4,opt,name=rooms_available,json=roomsAvailable,proto3" json:"rooms_available,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Hotel) Reset() {
	*x = Hotel{}
	mi := &file_proto_hotel_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hotel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hotel) ProtoMessage() {}

func (x *Hotel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hotel.ProtoReflect.Descriptor instead.
func (*Hotel) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{1}
}

func (x *Hotel) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Hotel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hotel) GetPricePernight() int32 {
	if x != nil {
		return x.PricePernight
	}
	return 0
}

func (x *Hotel) GetRoomsAvailable() int32 {
	if x != nil {
		return x.RoomsAvailable
	}
	return 0
}

type HotelResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Hotel         *Hotel                 `protobuf:"bytes,1,opt,name=hotel,proto3" json:"hotel,omitempty"` //TODO: 設定内容を詳細にする
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HotelResponse) Reset() {
	*x = HotelResponse{}
	mi := &file_proto_hotel_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HotelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelResponse) ProtoMessage() {}

func (x *HotelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelResponse.ProtoReflect.Descriptor instead.
func (*HotelResponse) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{2}
}

func (x *HotelResponse) GetHotel() *Hotel {
	if x != nil {
		return x.Hotel
	}
	return nil
}

var File_proto_hotel_proto protoreflect.FileDescriptor

var file_proto_hotel_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6e, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65,
	0x72, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x5f,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0x7b, 0x0a, 0x05, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6e, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x5f, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x6f,
	0x6f, 0x6d, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x33, 0x0a, 0x0d,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x05, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x05, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x32, 0x7f, 0x0a, 0x0c, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x13, 0x2e,
	0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x13, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x48, 0x6f, 0x74, 0x65, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_hotel_proto_rawDescOnce sync.Once
	file_proto_hotel_proto_rawDescData []byte
)

func file_proto_hotel_proto_rawDescGZIP() []byte {
	file_proto_hotel_proto_rawDescOnce.Do(func() {
		file_proto_hotel_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_hotel_proto_rawDesc), len(file_proto_hotel_proto_rawDesc)))
	})
	return file_proto_hotel_proto_rawDescData
}

var file_proto_hotel_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_hotel_proto_goTypes = []any{
	(*HotelRequest)(nil),  // 0: hotel.HotelRequest
	(*Hotel)(nil),         // 1: hotel.Hotel
	(*HotelResponse)(nil), // 2: hotel.HotelResponse
}
var file_proto_hotel_proto_depIdxs = []int32{
	1, // 0: hotel.HotelResponse.hotel:type_name -> hotel.Hotel
	0, // 1: hotel.HotelService.GetHotel:input_type -> hotel.HotelRequest
	0, // 2: hotel.HotelService.CreateHotel:input_type -> hotel.HotelRequest
	2, // 3: hotel.HotelService.GetHotel:output_type -> hotel.HotelResponse
	2, // 4: hotel.HotelService.CreateHotel:output_type -> hotel.HotelResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_hotel_proto_init() }
func file_proto_hotel_proto_init() {
	if File_proto_hotel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_hotel_proto_rawDesc), len(file_proto_hotel_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hotel_proto_goTypes,
		DependencyIndexes: file_proto_hotel_proto_depIdxs,
		MessageInfos:      file_proto_hotel_proto_msgTypes,
	}.Build()
	File_proto_hotel_proto = out.File
	file_proto_hotel_proto_goTypes = nil
	file_proto_hotel_proto_depIdxs = nil
}
