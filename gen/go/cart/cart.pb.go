// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: cart/cart.proto

package cart

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string  `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int64   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Portion   int64   `protobuf:"varint,3,opt,name=portion,proto3" json:"portion,omitempty"`
	Price     float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	GstPer    float64 `protobuf:"fixed64,5,opt,name=gst_per,json=gstPer,proto3" json:"gst_per,omitempty"`
}

func (x *OrderProduct) Reset() {
	*x = OrderProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProduct) ProtoMessage() {}

func (x *OrderProduct) ProtoReflect() protoreflect.Message {
	mi := &file_cart_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProduct.ProtoReflect.Descriptor instead.
func (*OrderProduct) Descriptor() ([]byte, []int) {
	return file_cart_cart_proto_rawDescGZIP(), []int{0}
}

func (x *OrderProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderProduct) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderProduct) GetPortion() int64 {
	if x != nil {
		return x.Portion
	}
	return 0
}

func (x *OrderProduct) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderProduct) GetGstPer() float64 {
	if x != nil {
		return x.GstPer
	}
	return 0
}

type OrderCart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId             string          `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AddressId          string          `protobuf:"bytes,3,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	StoreId            int64           `protobuf:"varint,4,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	IsServicable       bool            `protobuf:"varint,5,opt,name=is_servicable,json=isServicable,proto3" json:"is_servicable,omitempty"`
	IsFulfilled        bool            `protobuf:"varint,6,opt,name=is_fulfilled,json=isFulfilled,proto3" json:"is_fulfilled,omitempty"`
	TotalProductsPrice float64         `protobuf:"fixed64,7,opt,name=total_products_price,json=totalProductsPrice,proto3" json:"total_products_price,omitempty"`
	DeliveryPrice      float64         `protobuf:"fixed64,8,opt,name=delivery_price,json=deliveryPrice,proto3" json:"delivery_price,omitempty"`
	TotalTax           float64         `protobuf:"fixed64,9,opt,name=total_tax,json=totalTax,proto3" json:"total_tax,omitempty"`
	DiscountPrice      float64         `protobuf:"fixed64,10,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
	TotalCartPrice     float64         `protobuf:"fixed64,11,opt,name=total_cart_price,json=totalCartPrice,proto3" json:"total_cart_price,omitempty"`
	Products           []*OrderProduct `protobuf:"bytes,12,rep,name=products,proto3" json:"products,omitempty"`
	Type               string          `protobuf:"bytes,13,opt,name=type,proto3" json:"type,omitempty"`
	TaxType            string          `protobuf:"bytes,14,opt,name=tax_type,json=taxType,proto3" json:"tax_type,omitempty"`
}

func (x *OrderCart) Reset() {
	*x = OrderCart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCart) ProtoMessage() {}

func (x *OrderCart) ProtoReflect() protoreflect.Message {
	mi := &file_cart_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCart.ProtoReflect.Descriptor instead.
func (*OrderCart) Descriptor() ([]byte, []int) {
	return file_cart_cart_proto_rawDescGZIP(), []int{1}
}

func (x *OrderCart) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderCart) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderCart) GetAddressId() string {
	if x != nil {
		return x.AddressId
	}
	return ""
}

func (x *OrderCart) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *OrderCart) GetIsServicable() bool {
	if x != nil {
		return x.IsServicable
	}
	return false
}

func (x *OrderCart) GetIsFulfilled() bool {
	if x != nil {
		return x.IsFulfilled
	}
	return false
}

func (x *OrderCart) GetTotalProductsPrice() float64 {
	if x != nil {
		return x.TotalProductsPrice
	}
	return 0
}

func (x *OrderCart) GetDeliveryPrice() float64 {
	if x != nil {
		return x.DeliveryPrice
	}
	return 0
}

func (x *OrderCart) GetTotalTax() float64 {
	if x != nil {
		return x.TotalTax
	}
	return 0
}

func (x *OrderCart) GetDiscountPrice() float64 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *OrderCart) GetTotalCartPrice() float64 {
	if x != nil {
		return x.TotalCartPrice
	}
	return 0
}

func (x *OrderCart) GetProducts() []*OrderProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *OrderCart) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderCart) GetTaxType() string {
	if x != nil {
		return x.TaxType
	}
	return ""
}

type OrderCartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId string `protobuf:"bytes,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
}

func (x *OrderCartReq) Reset() {
	*x = OrderCartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCartReq) ProtoMessage() {}

func (x *OrderCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCartReq.ProtoReflect.Descriptor instead.
func (*OrderCartReq) Descriptor() ([]byte, []int) {
	return file_cart_cart_proto_rawDescGZIP(), []int{2}
}

func (x *OrderCartReq) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

type UpdateCartFulfilReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId      string `protobuf:"bytes,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	IsFulfilled bool   `protobuf:"varint,2,opt,name=is_fulfilled,json=isFulfilled,proto3" json:"is_fulfilled,omitempty"`
}

func (x *UpdateCartFulfilReq) Reset() {
	*x = UpdateCartFulfilReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartFulfilReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartFulfilReq) ProtoMessage() {}

func (x *UpdateCartFulfilReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartFulfilReq.ProtoReflect.Descriptor instead.
func (*UpdateCartFulfilReq) Descriptor() ([]byte, []int) {
	return file_cart_cart_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCartFulfilReq) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

func (x *UpdateCartFulfilReq) GetIsFulfilled() bool {
	if x != nil {
		return x.IsFulfilled
	}
	return false
}

type UpdateCartFulfilRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId      string `protobuf:"bytes,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	IsFulfilled bool   `protobuf:"varint,2,opt,name=is_fulfilled,json=isFulfilled,proto3" json:"is_fulfilled,omitempty"`
}

func (x *UpdateCartFulfilRes) Reset() {
	*x = UpdateCartFulfilRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartFulfilRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartFulfilRes) ProtoMessage() {}

func (x *UpdateCartFulfilRes) ProtoReflect() protoreflect.Message {
	mi := &file_cart_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartFulfilRes.ProtoReflect.Descriptor instead.
func (*UpdateCartFulfilRes) Descriptor() ([]byte, []int) {
	return file_cart_cart_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCartFulfilRes) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

func (x *UpdateCartFulfilRes) GetIsFulfilled() bool {
	if x != nil {
		return x.IsFulfilled
	}
	return false
}

var File_cart_cart_proto protoreflect.FileDescriptor

var file_cart_cart_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x61, 0x72, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x67, 0x73, 0x74, 0x50, 0x65, 0x72, 0x22, 0xdc, 0x03,
	0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c,
	0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46, 0x75, 0x6c,
	0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x78, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x61, 0x78, 0x12, 0x25, 0x0a, 0x0e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x61, 0x72,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x78, 0x54, 0x79, 0x70, 0x65, 0x22, 0x27, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x74, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x66, 0x75, 0x6c, 0x66,
	0x69, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46,
	0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x22, 0x51, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x66,
	0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x32, 0x90, 0x01, 0x0a, 0x0b,
	0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x72, 0x74, 0x12, 0x12, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x72, 0x74,
	0x22, 0x00, 0x12, 0x4a, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x61,
	0x6e, 0x74, 0x61, 0x6e, 0x75, 0x32, 0x30, 0x30, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x61, 0x72,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_cart_proto_rawDescOnce sync.Once
	file_cart_cart_proto_rawDescData = file_cart_cart_proto_rawDesc
)

func file_cart_cart_proto_rawDescGZIP() []byte {
	file_cart_cart_proto_rawDescOnce.Do(func() {
		file_cart_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_cart_proto_rawDescData)
	})
	return file_cart_cart_proto_rawDescData
}

var file_cart_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cart_cart_proto_goTypes = []interface{}{
	(*OrderProduct)(nil),        // 0: cart.OrderProduct
	(*OrderCart)(nil),           // 1: cart.OrderCart
	(*OrderCartReq)(nil),        // 2: cart.OrderCartReq
	(*UpdateCartFulfilReq)(nil), // 3: cart.UpdateCartFulfilReq
	(*UpdateCartFulfilRes)(nil), // 4: cart.UpdateCartFulfilRes
}
var file_cart_cart_proto_depIdxs = []int32{
	0, // 0: cart.OrderCart.products:type_name -> cart.OrderProduct
	2, // 1: cart.CartService.GetOrderCart:input_type -> cart.OrderCartReq
	3, // 2: cart.CartService.UpdateCartFulfil:input_type -> cart.UpdateCartFulfilReq
	1, // 3: cart.CartService.GetOrderCart:output_type -> cart.OrderCart
	4, // 4: cart.CartService.UpdateCartFulfil:output_type -> cart.UpdateCartFulfilRes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cart_cart_proto_init() }
func file_cart_cart_proto_init() {
	if File_cart_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderProduct); i {
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
		file_cart_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCart); i {
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
		file_cart_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCartReq); i {
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
		file_cart_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCartFulfilReq); i {
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
		file_cart_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCartFulfilRes); i {
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
			RawDescriptor: file_cart_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_cart_proto_goTypes,
		DependencyIndexes: file_cart_cart_proto_depIdxs,
		MessageInfos:      file_cart_cart_proto_msgTypes,
	}.Build()
	File_cart_cart_proto = out.File
	file_cart_cart_proto_rawDesc = nil
	file_cart_cart_proto_goTypes = nil
	file_cart_cart_proto_depIdxs = nil
}
