// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: product/product.proto

package product

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

type GetProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type GetProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductIds []string `protobuf:"bytes,1,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
}

func (x *GetProductsRequest) Reset() {
	*x = GetProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsRequest) ProtoMessage() {}

func (x *GetProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsRequest.ProtoReflect.Descriptor instead.
func (*GetProductsRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductsRequest) GetProductIds() []string {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

type GetProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products   []*MasterProduct `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	TotalCount int32            `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *GetProductsResponse) Reset() {
	*x = GetProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsResponse) ProtoMessage() {}

func (x *GetProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsResponse.ProtoReflect.Descriptor instead.
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductsResponse) GetProducts() []*MasterProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *GetProductsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type GetCartProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products   []*CartProduct `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	TotalCount int32          `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *GetCartProductResponse) Reset() {
	*x = GetCartProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartProductResponse) ProtoMessage() {}

func (x *GetCartProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartProductResponse.ProtoReflect.Descriptor instead.
func (*GetCartProductResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetCartProductResponse) GetProducts() []*CartProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *GetCartProductResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CartProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KId              string   `protobuf:"bytes,1,opt,name=K_id,json=KId,proto3" json:"K_id,omitempty"`
	KDisplayName     string   `protobuf:"bytes,2,opt,name=K_display_name,json=KDisplayName,proto3" json:"K_display_name,omitempty"`
	KDescription     string   `protobuf:"bytes,3,opt,name=K_description,json=KDescription,proto3" json:"K_description,omitempty"`
	KPreviewImage    string   `protobuf:"bytes,4,opt,name=K_preview_image,json=KPreviewImage,proto3" json:"K_preview_image,omitempty"`
	KShelfLifeInDays int32    `protobuf:"varint,5,opt,name=K_shelf_life_in_days,json=KShelfLifeInDays,proto3" json:"K_shelf_life_in_days,omitempty"`
	KDisplayPrice    float64  `protobuf:"fixed64,6,opt,name=K_display_price,json=KDisplayPrice,proto3" json:"K_display_price,omitempty"`
	KPricePerKg      int64    `protobuf:"varint,7,opt,name=K_price_per_kg,json=KPricePerKg,proto3" json:"K_price_per_kg,omitempty"`
	KRating          float64  `protobuf:"fixed64,8,opt,name=K_rating,json=KRating,proto3" json:"K_rating,omitempty"`
	KProductId       int64    `protobuf:"varint,9,opt,name=K_product_id,json=KProductId,proto3" json:"K_product_id,omitempty"`
	KCategory        []string `protobuf:"bytes,10,rep,name=K_category,json=KCategory,proto3" json:"K_category,omitempty"`
	KPortionsGm      []string `protobuf:"bytes,11,rep,name=K_portions_gm,json=KPortionsGm,proto3" json:"K_portions_gm,omitempty"`
	KVegNonVegTypeId string   `protobuf:"bytes,12,opt,name=K_veg_non_veg_type_id,json=KVegNonVegTypeId,proto3" json:"K_veg_non_veg_type_id,omitempty"`
	RW_GSTPer        float64  `protobuf:"fixed64,13,opt,name=RW_GSTPer,json=RWGSTPer,proto3" json:"RW_GSTPer,omitempty"`
}

func (x *CartProduct) Reset() {
	*x = CartProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartProduct) ProtoMessage() {}

func (x *CartProduct) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartProduct.ProtoReflect.Descriptor instead.
func (*CartProduct) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{4}
}

func (x *CartProduct) GetKId() string {
	if x != nil {
		return x.KId
	}
	return ""
}

func (x *CartProduct) GetKDisplayName() string {
	if x != nil {
		return x.KDisplayName
	}
	return ""
}

func (x *CartProduct) GetKDescription() string {
	if x != nil {
		return x.KDescription
	}
	return ""
}

func (x *CartProduct) GetKPreviewImage() string {
	if x != nil {
		return x.KPreviewImage
	}
	return ""
}

func (x *CartProduct) GetKShelfLifeInDays() int32 {
	if x != nil {
		return x.KShelfLifeInDays
	}
	return 0
}

func (x *CartProduct) GetKDisplayPrice() float64 {
	if x != nil {
		return x.KDisplayPrice
	}
	return 0
}

func (x *CartProduct) GetKPricePerKg() int64 {
	if x != nil {
		return x.KPricePerKg
	}
	return 0
}

func (x *CartProduct) GetKRating() float64 {
	if x != nil {
		return x.KRating
	}
	return 0
}

func (x *CartProduct) GetKProductId() int64 {
	if x != nil {
		return x.KProductId
	}
	return 0
}

func (x *CartProduct) GetKCategory() []string {
	if x != nil {
		return x.KCategory
	}
	return nil
}

func (x *CartProduct) GetKPortionsGm() []string {
	if x != nil {
		return x.KPortionsGm
	}
	return nil
}

func (x *CartProduct) GetKVegNonVegTypeId() string {
	if x != nil {
		return x.KVegNonVegTypeId
	}
	return ""
}

func (x *CartProduct) GetRW_GSTPer() float64 {
	if x != nil {
		return x.RW_GSTPer
	}
	return 0
}

type MasterProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KId              string   `protobuf:"bytes,1,opt,name=K_id,json=KId,proto3" json:"K_id,omitempty"`
	KProductId       int64    `protobuf:"varint,2,opt,name=K_product_id,json=KProductId,proto3" json:"K_product_id,omitempty"`
	KVisibility      int64    `protobuf:"varint,3,opt,name=K_visibility,json=KVisibility,proto3" json:"K_visibility,omitempty"`
	KDisplayName     string   `protobuf:"bytes,4,opt,name=K_display_name,json=KDisplayName,proto3" json:"K_display_name,omitempty"`
	KPreviewImage    string   `protobuf:"bytes,6,opt,name=K_preview_image,json=KPreviewImage,proto3" json:"K_preview_image,omitempty"`
	KShelfLifeInDays int32    `protobuf:"varint,7,opt,name=K_shelf_life_in_days,json=KShelfLifeInDays,proto3" json:"K_shelf_life_in_days,omitempty"`
	KDisplayPrice    float64  `protobuf:"fixed64,8,opt,name=K_display_price,json=KDisplayPrice,proto3" json:"K_display_price,omitempty"`
	KPricePerKg      int64    `protobuf:"varint,9,opt,name=K_price_per_kg,json=KPricePerKg,proto3" json:"K_price_per_kg,omitempty"`
	KRating          float64  `protobuf:"fixed64,10,opt,name=K_rating,json=KRating,proto3" json:"K_rating,omitempty"`
	KCategory        []string `protobuf:"bytes,11,rep,name=K_category,json=KCategory,proto3" json:"K_category,omitempty"`
	KPortionsGm      []string `protobuf:"bytes,12,rep,name=K_portions_gm,json=KPortionsGm,proto3" json:"K_portions_gm,omitempty"`
	KVegNonVegTypeId string   `protobuf:"bytes,13,opt,name=K_veg_non_veg_type_id,json=KVegNonVegTypeId,proto3" json:"K_veg_non_veg_type_id,omitempty"`
}

func (x *MasterProduct) Reset() {
	*x = MasterProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterProduct) ProtoMessage() {}

func (x *MasterProduct) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterProduct.ProtoReflect.Descriptor instead.
func (*MasterProduct) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{5}
}

func (x *MasterProduct) GetKId() string {
	if x != nil {
		return x.KId
	}
	return ""
}

func (x *MasterProduct) GetKProductId() int64 {
	if x != nil {
		return x.KProductId
	}
	return 0
}

func (x *MasterProduct) GetKVisibility() int64 {
	if x != nil {
		return x.KVisibility
	}
	return 0
}

func (x *MasterProduct) GetKDisplayName() string {
	if x != nil {
		return x.KDisplayName
	}
	return ""
}

func (x *MasterProduct) GetKPreviewImage() string {
	if x != nil {
		return x.KPreviewImage
	}
	return ""
}

func (x *MasterProduct) GetKShelfLifeInDays() int32 {
	if x != nil {
		return x.KShelfLifeInDays
	}
	return 0
}

func (x *MasterProduct) GetKDisplayPrice() float64 {
	if x != nil {
		return x.KDisplayPrice
	}
	return 0
}

func (x *MasterProduct) GetKPricePerKg() int64 {
	if x != nil {
		return x.KPricePerKg
	}
	return 0
}

func (x *MasterProduct) GetKRating() float64 {
	if x != nil {
		return x.KRating
	}
	return 0
}

func (x *MasterProduct) GetKCategory() []string {
	if x != nil {
		return x.KCategory
	}
	return nil
}

func (x *MasterProduct) GetKPortionsGm() []string {
	if x != nil {
		return x.KPortionsGm
	}
	return nil
}

func (x *MasterProduct) GetKVegNonVegTypeId() string {
	if x != nil {
		return x.KVegNonVegTypeId
	}
	return ""
}

type ProductsMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products   map[string]*MasterProduct `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TotalCount int64                     `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ProductsMap) Reset() {
	*x = ProductsMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsMap) ProtoMessage() {}

func (x *ProductsMap) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsMap.ProtoReflect.Descriptor instead.
func (*ProductsMap) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{6}
}

func (x *ProductsMap) GetProducts() map[string]*MasterProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ProductsMap) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_product_product_proto protoreflect.FileDescriptor

var file_product_product_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x22, 0x6a, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xde, 0x03, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x11, 0x0a, 0x04, 0x4b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x4b, 0x5f, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x4b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x4b, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x4b, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4b, 0x50,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x14, 0x4b,
	0x5f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x64,
	0x61, 0x79, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x4b, 0x53, 0x68, 0x65, 0x6c,
	0x66, 0x4c, 0x69, 0x66, 0x65, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x4b,
	0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x4b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0e, 0x4b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x6b, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x4b, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x4b, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x4b, 0x5f, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x4b, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0c, 0x4b, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4b, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x4b, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x4b, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x4b, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x5f, 0x67, 0x6d, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x4b, 0x50, 0x6f,
	0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x47, 0x6d, 0x12, 0x2f, 0x0a, 0x15, 0x4b, 0x5f, 0x76, 0x65,
	0x67, 0x5f, 0x6e, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4b, 0x56, 0x65, 0x67, 0x4e, 0x6f, 0x6e,
	0x56, 0x65, 0x67, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x52, 0x57, 0x5f,
	0x47, 0x53, 0x54, 0x50, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x52, 0x57,
	0x47, 0x53, 0x54, 0x50, 0x65, 0x72, 0x22, 0xc1, 0x03, 0x0a, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x11, 0x0a, 0x04, 0x4b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x4b,
	0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x4b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x4b, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x4b, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x24, 0x0a, 0x0e, 0x4b, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4b, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x4b, 0x5f, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x4b, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2e,
	0x0a, 0x14, 0x4b, 0x5f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x69,
	0x6e, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x4b, 0x53,
	0x68, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x66, 0x65, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x4b, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x4b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0e, 0x4b, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6b, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x4b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x4b, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x4b,
	0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x4b,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x4b, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x4b, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x4b, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x67, 0x6d, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x4b, 0x50,
	0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x47, 0x6d, 0x12, 0x2f, 0x0a, 0x15, 0x4b, 0x5f, 0x76,
	0x65, 0x67, 0x5f, 0x6e, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4b, 0x56, 0x65, 0x67, 0x4e, 0x6f,
	0x6e, 0x56, 0x65, 0x67, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0xc3, 0x01, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4d, 0x61, 0x70, 0x12, 0x3e, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4d,
	0x61, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x53, 0x0a, 0x0d, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x32, 0xbc, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4d, 0x61, 0x70, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4d, 0x61, 0x70, 0x22, 0x00, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68,
	0x61, 0x6e, 0x74, 0x61, 0x6e, 0x75, 0x32, 0x30, 0x30, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_product_proto_rawDescOnce sync.Once
	file_product_product_proto_rawDescData = file_product_product_proto_rawDesc
)

func file_product_product_proto_rawDescGZIP() []byte {
	file_product_product_proto_rawDescOnce.Do(func() {
		file_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_product_proto_rawDescData)
	})
	return file_product_product_proto_rawDescData
}

var file_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_product_product_proto_goTypes = []interface{}{
	(*GetProductRequest)(nil),      // 0: product.GetProductRequest
	(*GetProductsRequest)(nil),     // 1: product.GetProductsRequest
	(*GetProductsResponse)(nil),    // 2: product.GetProductsResponse
	(*GetCartProductResponse)(nil), // 3: product.GetCartProductResponse
	(*CartProduct)(nil),            // 4: product.CartProduct
	(*MasterProduct)(nil),          // 5: product.MasterProduct
	(*ProductsMap)(nil),            // 6: product.ProductsMap
	nil,                            // 7: product.ProductsMap.ProductsEntry
}
var file_product_product_proto_depIdxs = []int32{
	5, // 0: product.GetProductsResponse.products:type_name -> product.MasterProduct
	4, // 1: product.GetCartProductResponse.products:type_name -> product.CartProduct
	7, // 2: product.ProductsMap.products:type_name -> product.ProductsMap.ProductsEntry
	5, // 3: product.ProductsMap.ProductsEntry.value:type_name -> product.MasterProduct
	1, // 4: product.ProductService.GetProducts:input_type -> product.GetProductsRequest
	1, // 5: product.ProductService.GetCartProducts:input_type -> product.GetProductsRequest
	0, // 6: product.ProductService.GetCartProduct:input_type -> product.GetProductRequest
	1, // 7: product.ProductService.GetProductsMap:input_type -> product.GetProductsRequest
	2, // 8: product.ProductService.GetProducts:output_type -> product.GetProductsResponse
	3, // 9: product.ProductService.GetCartProducts:output_type -> product.GetCartProductResponse
	4, // 10: product.ProductService.GetCartProduct:output_type -> product.CartProduct
	6, // 11: product.ProductService.GetProductsMap:output_type -> product.ProductsMap
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_product_product_proto_init() }
func file_product_product_proto_init() {
	if File_product_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductRequest); i {
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
		file_product_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductsRequest); i {
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
		file_product_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductsResponse); i {
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
		file_product_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartProductResponse); i {
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
		file_product_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartProduct); i {
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
		file_product_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterProduct); i {
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
		file_product_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsMap); i {
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
			RawDescriptor: file_product_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_product_proto_goTypes,
		DependencyIndexes: file_product_product_proto_depIdxs,
		MessageInfos:      file_product_product_proto_msgTypes,
	}.Build()
	File_product_product_proto = out.File
	file_product_product_proto_rawDesc = nil
	file_product_product_proto_goTypes = nil
	file_product_product_proto_depIdxs = nil
}
