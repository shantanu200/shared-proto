// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: shop/shop.proto

package shop

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

type ShopReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *ShopReq) Reset() {
	*x = ShopReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopReq) ProtoMessage() {}

func (x *ShopReq) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopReq.ProtoReflect.Descriptor instead.
func (*ShopReq) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{0}
}

func (x *ShopReq) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

type Shop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ShopId      int64   `protobuf:"varint,3,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	PhoneNumber string  `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	IsActive    bool    `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Latitude    float64 `protobuf:"fixed64,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float64 `protobuf:"fixed64,7,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Address     string  `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Shop) Reset() {
	*x = Shop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_shop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{1}
}

func (x *Shop) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Shop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shop) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *Shop) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Shop) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Shop) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Shop) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Shop) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_shop_shop_proto protoreflect.FileDescriptor

var file_shop_shop_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x22, 0x22, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x70, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0xd7, 0x01, 0x0a, 0x04,
	0x53, 0x68, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x35, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x12,
	0x0d, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0a,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x6e, 0x74,
	0x61, 0x6e, 0x75, 0x32, 0x30, 0x30, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_shop_proto_rawDescOnce sync.Once
	file_shop_shop_proto_rawDescData = file_shop_shop_proto_rawDesc
)

func file_shop_shop_proto_rawDescGZIP() []byte {
	file_shop_shop_proto_rawDescOnce.Do(func() {
		file_shop_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_shop_proto_rawDescData)
	})
	return file_shop_shop_proto_rawDescData
}

var file_shop_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shop_shop_proto_goTypes = []interface{}{
	(*ShopReq)(nil), // 0: shop.ShopReq
	(*Shop)(nil),    // 1: shop.Shop
}
var file_shop_shop_proto_depIdxs = []int32{
	0, // 0: shop.ShopService.GetShop:input_type -> shop.ShopReq
	1, // 1: shop.ShopService.GetShop:output_type -> shop.Shop
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shop_shop_proto_init() }
func file_shop_shop_proto_init() {
	if File_shop_shop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopReq); i {
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
		file_shop_shop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shop); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shop_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_shop_proto_goTypes,
		DependencyIndexes: file_shop_shop_proto_depIdxs,
		MessageInfos:      file_shop_shop_proto_msgTypes,
	}.Build()
	File_shop_shop_proto = out.File
	file_shop_shop_proto_rawDesc = nil
	file_shop_shop_proto_goTypes = nil
	file_shop_shop_proto_depIdxs = nil
}
