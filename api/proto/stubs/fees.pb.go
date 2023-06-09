// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: fees.proto

package stubs

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

type GetFeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId     int32    `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	BuyerCurrency Currency `protobuf:"varint,2,opt,name=buyerCurrency,proto3,enum=proto.Currency" json:"buyerCurrency,omitempty"`
}

func (x *GetFeeRequest) Reset() {
	*x = GetFeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fees_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeeRequest) ProtoMessage() {}

func (x *GetFeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fees_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeeRequest.ProtoReflect.Descriptor instead.
func (*GetFeeRequest) Descriptor() ([]byte, []int) {
	return file_fees_proto_rawDescGZIP(), []int{0}
}

func (x *GetFeeRequest) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *GetFeeRequest) GetBuyerCurrency() Currency {
	if x != nil {
		return x.BuyerCurrency
	}
	return Currency_EUR
}

type GetFeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fee   *Fee   `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetFeeResponse) Reset() {
	*x = GetFeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fees_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeeResponse) ProtoMessage() {}

func (x *GetFeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fees_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeeResponse.ProtoReflect.Descriptor instead.
func (*GetFeeResponse) Descriptor() ([]byte, []int) {
	return file_fees_proto_rawDescGZIP(), []int{1}
}

func (x *GetFeeResponse) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *GetFeeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FeeResponse) Reset() {
	*x = FeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fees_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeResponse) ProtoMessage() {}

func (x *FeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fees_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeResponse.ProtoReflect.Descriptor instead.
func (*FeeResponse) Descriptor() ([]byte, []int) {
	return file_fees_proto_rawDescGZIP(), []int{2}
}

func (x *FeeResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *FeeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId     int32    `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	BuyerCurrency Currency `protobuf:"varint,2,opt,name=buyerCurrency,proto3,enum=proto.Currency" json:"buyerCurrency,omitempty"`
	Fee           *Fee     `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *FeeRequest) Reset() {
	*x = FeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fees_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeRequest) ProtoMessage() {}

func (x *FeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fees_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeRequest.ProtoReflect.Descriptor instead.
func (*FeeRequest) Descriptor() ([]byte, []int) {
	return file_fees_proto_rawDescGZIP(), []int{3}
}

func (x *FeeRequest) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *FeeRequest) GetBuyerCurrency() Currency {
	if x != nil {
		return x.BuyerCurrency
	}
	return Currency_EUR
}

func (x *FeeRequest) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

var File_fees_proto protoreflect.FileDescriptor

var file_fees_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x65, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0d, 0x62, 0x75, 0x79, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0d, 0x62, 0x75, 0x79,
	0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x44, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x03,
	0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x3b, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x7f, 0x0a,
	0x0a, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0d, 0x62, 0x75, 0x79,
	0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x0d, 0x62, 0x75, 0x79, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x1c, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x03, 0x66, 0x65, 0x65, 0x32, 0xd2,
	0x01, 0x0a, 0x0a, 0x46, 0x65, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x46, 0x65, 0x65, 0x12, 0x18, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x65, 0x65, 0x12, 0x18,
	0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x46, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fees_proto_rawDescOnce sync.Once
	file_fees_proto_rawDescData = file_fees_proto_rawDesc
)

func file_fees_proto_rawDescGZIP() []byte {
	file_fees_proto_rawDescOnce.Do(func() {
		file_fees_proto_rawDescData = protoimpl.X.CompressGZIP(file_fees_proto_rawDescData)
	})
	return file_fees_proto_rawDescData
}

var file_fees_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_fees_proto_goTypes = []interface{}{
	(*GetFeeRequest)(nil),  // 0: repositories.GetFeeRequest
	(*GetFeeResponse)(nil), // 1: repositories.GetFeeResponse
	(*FeeResponse)(nil),    // 2: repositories.FeeResponse
	(*FeeRequest)(nil),     // 3: repositories.FeeRequest
	(Currency)(0),          // 4: proto.Currency
	(*Fee)(nil),            // 5: proto.Fee
}
var file_fees_proto_depIdxs = []int32{
	4, // 0: repositories.GetFeeRequest.buyerCurrency:type_name -> proto.Currency
	5, // 1: repositories.GetFeeResponse.fee:type_name -> proto.Fee
	4, // 2: repositories.FeeRequest.buyerCurrency:type_name -> proto.Currency
	5, // 3: repositories.FeeRequest.fee:type_name -> proto.Fee
	0, // 4: repositories.FeeService.GetFee:input_type -> repositories.GetFeeRequest
	3, // 5: repositories.FeeService.AddFee:input_type -> repositories.FeeRequest
	3, // 6: repositories.FeeService.RemoveFee:input_type -> repositories.FeeRequest
	1, // 7: repositories.FeeService.GetFee:output_type -> repositories.GetFeeResponse
	2, // 8: repositories.FeeService.AddFee:output_type -> repositories.FeeResponse
	2, // 9: repositories.FeeService.RemoveFee:output_type -> repositories.FeeResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_fees_proto_init() }
func file_fees_proto_init() {
	if File_fees_proto != nil {
		return
	}
	file_product_price_fees_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fees_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeeRequest); i {
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
		file_fees_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeeResponse); i {
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
		file_fees_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeResponse); i {
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
		file_fees_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeRequest); i {
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
			RawDescriptor: file_fees_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fees_proto_goTypes,
		DependencyIndexes: file_fees_proto_depIdxs,
		MessageInfos:      file_fees_proto_msgTypes,
	}.Build()
	File_fees_proto = out.File
	file_fees_proto_rawDesc = nil
	file_fees_proto_goTypes = nil
	file_fees_proto_depIdxs = nil
}
