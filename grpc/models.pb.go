// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: gotbitoriginal/common/models.proto

package grpc

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

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  string `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Quote string `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotbitoriginal_common_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_gotbitoriginal_common_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_gotbitoriginal_common_models_proto_rawDescGZIP(), []int{0}
}

func (x *Pair) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *Pair) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

type FeeRatios struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maker string `protobuf:"bytes,1,opt,name=maker,proto3" json:"maker,omitempty"`
	Taker string `protobuf:"bytes,2,opt,name=taker,proto3" json:"taker,omitempty"`
}

func (x *FeeRatios) Reset() {
	*x = FeeRatios{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotbitoriginal_common_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeRatios) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeRatios) ProtoMessage() {}

func (x *FeeRatios) ProtoReflect() protoreflect.Message {
	mi := &file_gotbitoriginal_common_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeRatios.ProtoReflect.Descriptor instead.
func (*FeeRatios) Descriptor() ([]byte, []int) {
	return file_gotbitoriginal_common_models_proto_rawDescGZIP(), []int{1}
}

func (x *FeeRatios) GetMaker() string {
	if x != nil {
		return x.Maker
	}
	return ""
}

func (x *FeeRatios) GetTaker() string {
	if x != nil {
		return x.Taker
	}
	return ""
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Free   string `protobuf:"bytes,1,opt,name=free,proto3" json:"free,omitempty"`
	Locked string `protobuf:"bytes,2,opt,name=locked,proto3" json:"locked,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotbitoriginal_common_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_gotbitoriginal_common_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_gotbitoriginal_common_models_proto_rawDescGZIP(), []int{2}
}

func (x *Balance) GetFree() string {
	if x != nil {
		return x.Free
	}
	return ""
}

func (x *Balance) GetLocked() string {
	if x != nil {
		return x.Locked
	}
	return ""
}

type Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Price  string `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Offer) Reset() {
	*x = Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotbitoriginal_common_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offer) ProtoMessage() {}

func (x *Offer) ProtoReflect() protoreflect.Message {
	mi := &file_gotbitoriginal_common_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offer.ProtoReflect.Descriptor instead.
func (*Offer) Descriptor() ([]byte, []int) {
	return file_gotbitoriginal_common_models_proto_rawDescGZIP(), []int{3}
}

func (x *Offer) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Offer) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Side             string `protobuf:"bytes,2,opt,name=side,proto3" json:"side,omitempty"`
	Pair             *Pair  `protobuf:"bytes,3,opt,name=pair,proto3" json:"pair,omitempty"`
	Amount           string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountFilled     string `protobuf:"bytes,5,opt,name=amountFilled,proto3" json:"amountFilled,omitempty"`
	Price            string `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	PriceFillAgerage string `protobuf:"bytes,7,opt,name=priceFillAgerage,proto3" json:"priceFillAgerage,omitempty"`
	Fee              string `protobuf:"bytes,8,opt,name=fee,proto3" json:"fee,omitempty"`
	Status           string `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Time             int64  `protobuf:"varint,10,opt,name=time,proto3" json:"time,omitempty"`
	FillOrKill       bool   `protobuf:"varint,11,opt,name=fillOrKill,proto3" json:"fillOrKill,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotbitoriginal_common_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_gotbitoriginal_common_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_gotbitoriginal_common_models_proto_rawDescGZIP(), []int{4}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetSide() string {
	if x != nil {
		return x.Side
	}
	return ""
}

func (x *Order) GetPair() *Pair {
	if x != nil {
		return x.Pair
	}
	return nil
}

func (x *Order) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Order) GetAmountFilled() string {
	if x != nil {
		return x.AmountFilled
	}
	return ""
}

func (x *Order) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Order) GetPriceFillAgerage() string {
	if x != nil {
		return x.PriceFillAgerage
	}
	return ""
}

func (x *Order) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Order) GetFillOrKill() bool {
	if x != nil {
		return x.FillOrKill
	}
	return false
}

type Candle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time   int64  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Open   string `protobuf:"bytes,2,opt,name=open,proto3" json:"open,omitempty"`
	Close  string `protobuf:"bytes,3,opt,name=close,proto3" json:"close,omitempty"`
	High   string `protobuf:"bytes,4,opt,name=high,proto3" json:"high,omitempty"`
	Low    string `protobuf:"bytes,5,opt,name=low,proto3" json:"low,omitempty"`
	Volume string `protobuf:"bytes,6,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *Candle) Reset() {
	*x = Candle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotbitoriginal_common_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candle) ProtoMessage() {}

func (x *Candle) ProtoReflect() protoreflect.Message {
	mi := &file_gotbitoriginal_common_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candle.ProtoReflect.Descriptor instead.
func (*Candle) Descriptor() ([]byte, []int) {
	return file_gotbitoriginal_common_models_proto_rawDescGZIP(), []int{5}
}

func (x *Candle) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Candle) GetOpen() string {
	if x != nil {
		return x.Open
	}
	return ""
}

func (x *Candle) GetClose() string {
	if x != nil {
		return x.Close
	}
	return ""
}

func (x *Candle) GetHigh() string {
	if x != nil {
		return x.High
	}
	return ""
}

func (x *Candle) GetLow() string {
	if x != nil {
		return x.Low
	}
	return ""
}

func (x *Candle) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

type Trade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time   int64  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Side   string `protobuf:"bytes,2,opt,name=side,proto3" json:"side,omitempty"`
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Price  string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Trade) Reset() {
	*x = Trade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotbitoriginal_common_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trade) ProtoMessage() {}

func (x *Trade) ProtoReflect() protoreflect.Message {
	mi := &file_gotbitoriginal_common_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trade.ProtoReflect.Descriptor instead.
func (*Trade) Descriptor() ([]byte, []int) {
	return file_gotbitoriginal_common_models_proto_rawDescGZIP(), []int{6}
}

func (x *Trade) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Trade) GetSide() string {
	if x != nil {
		return x.Side
	}
	return ""
}

func (x *Trade) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Trade) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

type Precision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int32 `protobuf:"varint,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Price  int32 `protobuf:"varint,2,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *Precision) Reset() {
	*x = Precision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotbitoriginal_common_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Precision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Precision) ProtoMessage() {}

func (x *Precision) ProtoReflect() protoreflect.Message {
	mi := &file_gotbitoriginal_common_models_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Precision.ProtoReflect.Descriptor instead.
func (*Precision) Descriptor() ([]byte, []int) {
	return file_gotbitoriginal_common_models_proto_rawDescGZIP(), []int{7}
}

func (x *Precision) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Precision) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type OrderBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asks []*Offer `protobuf:"bytes,1,rep,name=Asks,proto3" json:"Asks,omitempty"`
	Bids []*Offer `protobuf:"bytes,2,rep,name=Bids,proto3" json:"Bids,omitempty"`
}

func (x *OrderBook) Reset() {
	*x = OrderBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotbitoriginal_common_models_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBook) ProtoMessage() {}

func (x *OrderBook) ProtoReflect() protoreflect.Message {
	mi := &file_gotbitoriginal_common_models_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBook.ProtoReflect.Descriptor instead.
func (*OrderBook) Descriptor() ([]byte, []int) {
	return file_gotbitoriginal_common_models_proto_rawDescGZIP(), []int{8}
}

func (x *OrderBook) GetAsks() []*Offer {
	if x != nil {
		return x.Asks
	}
	return nil
}

func (x *OrderBook) GetBids() []*Offer {
	if x != nil {
		return x.Bids
	}
	return nil
}

var File_gotbitoriginal_common_models_proto protoreflect.FileDescriptor

var file_gotbitoriginal_common_models_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x74, 0x62, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x22, 0x37, 0x0a, 0x09, 0x46, 0x65, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x6b,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x22,
	0x35, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x65, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x35, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0xa2, 0x02,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x70,
	0x61, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x04, 0x70, 0x61, 0x69, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x6c,
	0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x46, 0x69, 0x6c, 0x6c, 0x41, 0x67, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x69, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x41, 0x67, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x6c, 0x4f, 0x72, 0x4b, 0x69, 0x6c, 0x6c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x6c, 0x4f, 0x72, 0x4b, 0x69,
	0x6c, 0x6c, 0x22, 0x84, 0x01, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x69, 0x67, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f,
	0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x05, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x39, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x22, 0x43, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x1a, 0x0a, 0x04, 0x41, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x04, 0x41, 0x73, 0x6b, 0x73, 0x12, 0x1a, 0x0a, 0x04,
	0x42, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x52, 0x04, 0x42, 0x69, 0x64, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x74, 0x62, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gotbitoriginal_common_models_proto_rawDescOnce sync.Once
	file_gotbitoriginal_common_models_proto_rawDescData = file_gotbitoriginal_common_models_proto_rawDesc
)

func file_gotbitoriginal_common_models_proto_rawDescGZIP() []byte {
	file_gotbitoriginal_common_models_proto_rawDescOnce.Do(func() {
		file_gotbitoriginal_common_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_gotbitoriginal_common_models_proto_rawDescData)
	})
	return file_gotbitoriginal_common_models_proto_rawDescData
}

var file_gotbitoriginal_common_models_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gotbitoriginal_common_models_proto_goTypes = []interface{}{
	(*Pair)(nil),      // 0: Pair
	(*FeeRatios)(nil), // 1: FeeRatios
	(*Balance)(nil),   // 2: Balance
	(*Offer)(nil),     // 3: Offer
	(*Order)(nil),     // 4: Order
	(*Candle)(nil),    // 5: Candle
	(*Trade)(nil),     // 6: Trade
	(*Precision)(nil), // 7: Precision
	(*OrderBook)(nil), // 8: OrderBook
}
var file_gotbitoriginal_common_models_proto_depIdxs = []int32{
	0, // 0: Order.pair:type_name -> Pair
	3, // 1: OrderBook.Asks:type_name -> Offer
	3, // 2: OrderBook.Bids:type_name -> Offer
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gotbitoriginal_common_models_proto_init() }
func file_gotbitoriginal_common_models_proto_init() {
	if File_gotbitoriginal_common_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gotbitoriginal_common_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_gotbitoriginal_common_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeRatios); i {
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
		file_gotbitoriginal_common_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_gotbitoriginal_common_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offer); i {
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
		file_gotbitoriginal_common_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_gotbitoriginal_common_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candle); i {
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
		file_gotbitoriginal_common_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trade); i {
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
		file_gotbitoriginal_common_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Precision); i {
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
		file_gotbitoriginal_common_models_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderBook); i {
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
			RawDescriptor: file_gotbitoriginal_common_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gotbitoriginal_common_models_proto_goTypes,
		DependencyIndexes: file_gotbitoriginal_common_models_proto_depIdxs,
		MessageInfos:      file_gotbitoriginal_common_models_proto_msgTypes,
	}.Build()
	File_gotbitoriginal_common_models_proto = out.File
	file_gotbitoriginal_common_models_proto_rawDesc = nil
	file_gotbitoriginal_common_models_proto_goTypes = nil
	file_gotbitoriginal_common_models_proto_depIdxs = nil
}