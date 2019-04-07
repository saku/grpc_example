// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pancake.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// メニュー表
type Pancake_Menu int32

const (
	Pancake_UNKNOWN           Pancake_Menu = 0
	Pancake_CLASSIC           Pancake_Menu = 1
	Pancake_BANANA_AND_WHIP   Pancake_Menu = 2
	Pancake_BACON_AND_CHEESE  Pancake_Menu = 3
	Pancake_MIX_BERRY         Pancake_Menu = 4
	Pancake_BAKED_MARSHMALLOW Pancake_Menu = 5
	Pancake_SPICY_CURRY       Pancake_Menu = 6
)

var Pancake_Menu_name = map[int32]string{
	0: "UNKNOWN",
	1: "CLASSIC",
	2: "BANANA_AND_WHIP",
	3: "BACON_AND_CHEESE",
	4: "MIX_BERRY",
	5: "BAKED_MARSHMALLOW",
	6: "SPICY_CURRY",
}

var Pancake_Menu_value = map[string]int32{
	"UNKNOWN":           0,
	"CLASSIC":           1,
	"BANANA_AND_WHIP":   2,
	"BACON_AND_CHEESE":  3,
	"MIX_BERRY":         4,
	"BAKED_MARSHMALLOW": 5,
	"SPICY_CURRY":       6,
}

func (x Pancake_Menu) String() string {
	return proto.EnumName(Pancake_Menu_name, int32(x))
}

func (Pancake_Menu) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd84accc06629c9d, []int{0, 0}
}

// Pancakeは一枚一枚の焼かれたパンケーキを表します。
type Pancake struct {
	// シェフの名前
	ChefName string `protobuf:"bytes,1,opt,name=chef_name,json=chefName,proto3" json:"chef_name,omitempty"`
	// メニュー
	Menu Pancake_Menu `protobuf:"varint,2,opt,name=menu,proto3,enum=pancake.baker.Pancake_Menu" json:"menu,omitempty"`
	// 焼き具合を表すスコアです(0-0.9)
	TechnicalScore float32 `protobuf:"fixed32,3,opt,name=technical_score,json=technicalScore,proto3" json:"technical_score,omitempty"`
	// 焼いた日時
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,15,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pancake) Reset()         { *m = Pancake{} }
func (m *Pancake) String() string { return proto.CompactTextString(m) }
func (*Pancake) ProtoMessage()    {}
func (*Pancake) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd84accc06629c9d, []int{0}
}

func (m *Pancake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pancake.Unmarshal(m, b)
}
func (m *Pancake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pancake.Marshal(b, m, deterministic)
}
func (m *Pancake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pancake.Merge(m, src)
}
func (m *Pancake) XXX_Size() int {
	return xxx_messageInfo_Pancake.Size(m)
}
func (m *Pancake) XXX_DiscardUnknown() {
	xxx_messageInfo_Pancake.DiscardUnknown(m)
}

var xxx_messageInfo_Pancake proto.InternalMessageInfo

func (m *Pancake) GetChefName() string {
	if m != nil {
		return m.ChefName
	}
	return ""
}

func (m *Pancake) GetMenu() Pancake_Menu {
	if m != nil {
		return m.Menu
	}
	return Pancake_UNKNOWN
}

func (m *Pancake) GetTechnicalScore() float32 {
	if m != nil {
		return m.TechnicalScore
	}
	return 0
}

func (m *Pancake) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

// Reportはどのくらいパンケーキを焼いたかについての報告書を表します
type Report struct {
	BakeCounts           []*Report_BakeCount `protobuf:"bytes,1,rep,name=bake_counts,json=bakeCounts,proto3" json:"bake_counts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd84accc06629c9d, []int{1}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetBakeCounts() []*Report_BakeCount {
	if m != nil {
		return m.BakeCounts
	}
	return nil
}

type Report_BakeCount struct {
	Menu                 Pancake_Menu `protobuf:"varint,1,opt,name=menu,proto3,enum=pancake.baker.Pancake_Menu" json:"menu,omitempty"`
	Count                int32        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Report_BakeCount) Reset()         { *m = Report_BakeCount{} }
func (m *Report_BakeCount) String() string { return proto.CompactTextString(m) }
func (*Report_BakeCount) ProtoMessage()    {}
func (*Report_BakeCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd84accc06629c9d, []int{1, 0}
}

func (m *Report_BakeCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report_BakeCount.Unmarshal(m, b)
}
func (m *Report_BakeCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report_BakeCount.Marshal(b, m, deterministic)
}
func (m *Report_BakeCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report_BakeCount.Merge(m, src)
}
func (m *Report_BakeCount) XXX_Size() int {
	return xxx_messageInfo_Report_BakeCount.Size(m)
}
func (m *Report_BakeCount) XXX_DiscardUnknown() {
	xxx_messageInfo_Report_BakeCount.DiscardUnknown(m)
}

var xxx_messageInfo_Report_BakeCount proto.InternalMessageInfo

func (m *Report_BakeCount) GetMenu() Pancake_Menu {
	if m != nil {
		return m.Menu
	}
	return Pancake_UNKNOWN
}

func (m *Report_BakeCount) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type BakeRequest struct {
	Menu                 Pancake_Menu `protobuf:"varint,1,opt,name=menu,proto3,enum=pancake.baker.Pancake_Menu" json:"menu,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BakeRequest) Reset()         { *m = BakeRequest{} }
func (m *BakeRequest) String() string { return proto.CompactTextString(m) }
func (*BakeRequest) ProtoMessage()    {}
func (*BakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd84accc06629c9d, []int{2}
}

func (m *BakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BakeRequest.Unmarshal(m, b)
}
func (m *BakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BakeRequest.Marshal(b, m, deterministic)
}
func (m *BakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BakeRequest.Merge(m, src)
}
func (m *BakeRequest) XXX_Size() int {
	return xxx_messageInfo_BakeRequest.Size(m)
}
func (m *BakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BakeRequest proto.InternalMessageInfo

func (m *BakeRequest) GetMenu() Pancake_Menu {
	if m != nil {
		return m.Menu
	}
	return Pancake_UNKNOWN
}

type BakeResponse struct {
	Pancake              *Pancake `protobuf:"bytes,1,opt,name=pancake,proto3" json:"pancake,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BakeResponse) Reset()         { *m = BakeResponse{} }
func (m *BakeResponse) String() string { return proto.CompactTextString(m) }
func (*BakeResponse) ProtoMessage()    {}
func (*BakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd84accc06629c9d, []int{3}
}

func (m *BakeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BakeResponse.Unmarshal(m, b)
}
func (m *BakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BakeResponse.Marshal(b, m, deterministic)
}
func (m *BakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BakeResponse.Merge(m, src)
}
func (m *BakeResponse) XXX_Size() int {
	return xxx_messageInfo_BakeResponse.Size(m)
}
func (m *BakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BakeResponse proto.InternalMessageInfo

func (m *BakeResponse) GetPancake() *Pancake {
	if m != nil {
		return m.Pancake
	}
	return nil
}

type ReportRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportRequest) Reset()         { *m = ReportRequest{} }
func (m *ReportRequest) String() string { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()    {}
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd84accc06629c9d, []int{4}
}

func (m *ReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportRequest.Unmarshal(m, b)
}
func (m *ReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportRequest.Marshal(b, m, deterministic)
}
func (m *ReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportRequest.Merge(m, src)
}
func (m *ReportRequest) XXX_Size() int {
	return xxx_messageInfo_ReportRequest.Size(m)
}
func (m *ReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportRequest proto.InternalMessageInfo

type ReportResponse struct {
	Report               *Report  `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportResponse) Reset()         { *m = ReportResponse{} }
func (m *ReportResponse) String() string { return proto.CompactTextString(m) }
func (*ReportResponse) ProtoMessage()    {}
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd84accc06629c9d, []int{5}
}

func (m *ReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportResponse.Unmarshal(m, b)
}
func (m *ReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportResponse.Marshal(b, m, deterministic)
}
func (m *ReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportResponse.Merge(m, src)
}
func (m *ReportResponse) XXX_Size() int {
	return xxx_messageInfo_ReportResponse.Size(m)
}
func (m *ReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportResponse proto.InternalMessageInfo

func (m *ReportResponse) GetReport() *Report {
	if m != nil {
		return m.Report
	}
	return nil
}

func init() {
	proto.RegisterEnum("pancake.baker.Pancake_Menu", Pancake_Menu_name, Pancake_Menu_value)
	proto.RegisterType((*Pancake)(nil), "pancake.baker.Pancake")
	proto.RegisterType((*Report)(nil), "pancake.baker.Report")
	proto.RegisterType((*Report_BakeCount)(nil), "pancake.baker.Report.BakeCount")
	proto.RegisterType((*BakeRequest)(nil), "pancake.baker.BakeRequest")
	proto.RegisterType((*BakeResponse)(nil), "pancake.baker.BakeResponse")
	proto.RegisterType((*ReportRequest)(nil), "pancake.baker.ReportRequest")
	proto.RegisterType((*ReportResponse)(nil), "pancake.baker.ReportResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PancakeBakerServiceClient is the client API for PancakeBakerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PancakeBakerServiceClient interface {
	// Bakeは指定されたメニューのパンケーキを焼く関数です
	Bake(ctx context.Context, in *BakeRequest, opts ...grpc.CallOption) (*BakeResponse, error)
	// Reportはメニューごとに焼いたパンケーキの数を返します
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type pancakeBakerServiceClient struct {
	cc *grpc.ClientConn
}

func NewPancakeBakerServiceClient(cc *grpc.ClientConn) PancakeBakerServiceClient {
	return &pancakeBakerServiceClient{cc}
}

func (c *pancakeBakerServiceClient) Bake(ctx context.Context, in *BakeRequest, opts ...grpc.CallOption) (*BakeResponse, error) {
	out := new(BakeResponse)
	err := c.cc.Invoke(ctx, "/pancake.baker.PancakeBakerService/Bake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pancakeBakerServiceClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, "/pancake.baker.PancakeBakerService/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PancakeBakerServiceServer is the server API for PancakeBakerService service.
type PancakeBakerServiceServer interface {
	// Bakeは指定されたメニューのパンケーキを焼く関数です
	Bake(context.Context, *BakeRequest) (*BakeResponse, error)
	// Reportはメニューごとに焼いたパンケーキの数を返します
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
}

func RegisterPancakeBakerServiceServer(s *grpc.Server, srv PancakeBakerServiceServer) {
	s.RegisterService(&_PancakeBakerService_serviceDesc, srv)
}

func _PancakeBakerService_Bake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PancakeBakerServiceServer).Bake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pancake.baker.PancakeBakerService/Bake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PancakeBakerServiceServer).Bake(ctx, req.(*BakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PancakeBakerService_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PancakeBakerServiceServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pancake.baker.PancakeBakerService/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PancakeBakerServiceServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PancakeBakerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pancake.baker.PancakeBakerService",
	HandlerType: (*PancakeBakerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bake",
			Handler:    _PancakeBakerService_Bake_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _PancakeBakerService_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pancake.proto",
}

func init() { proto.RegisterFile("pancake.proto", fileDescriptor_bd84accc06629c9d) }

var fileDescriptor_bd84accc06629c9d = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xed, 0xe6, 0xf7, 0xcb, 0xf8, 0x4b, 0x62, 0xb6, 0x2d, 0x8a, 0x12, 0x50, 0x23, 0xdf, 0x90,
	0x1b, 0x1c, 0x14, 0x2e, 0x91, 0xa0, 0x6b, 0x37, 0x22, 0x51, 0x13, 0x27, 0x5a, 0x53, 0x85, 0x72,
	0x63, 0x6d, 0xac, 0x69, 0x1a, 0xa5, 0xb1, 0x8d, 0xed, 0xf0, 0x00, 0x3c, 0x08, 0x12, 0xaf, 0xc2,
	0x93, 0xa1, 0xf5, 0x9f, 0xa0, 0x0a, 0x12, 0xdc, 0x79, 0xce, 0x9c, 0x39, 0x73, 0x7c, 0x66, 0xa1,
	0x19, 0x08, 0xcf, 0x15, 0x3b, 0xd4, 0x83, 0xd0, 0x8f, 0x7d, 0x5a, 0x94, 0x6b, 0xb1, 0xc3, 0xb0,
	0x7b, 0xb1, 0xf1, 0xfd, 0xcd, 0x03, 0x0e, 0x93, 0xe6, 0xfa, 0x70, 0x37, 0x8c, 0xb7, 0x7b, 0x8c,
	0x62, 0xb1, 0x0f, 0x52, 0xbe, 0xf6, 0xa3, 0x04, 0xf5, 0x65, 0x3a, 0x42, 0x7b, 0xd0, 0x70, 0xef,
	0xf1, 0xce, 0xf1, 0xc4, 0x1e, 0x3b, 0xa4, 0x4f, 0x06, 0x0d, 0xfe, 0x9f, 0x04, 0x2c, 0xb1, 0x47,
	0x3a, 0x84, 0xca, 0x1e, 0xbd, 0x43, 0xa7, 0xd4, 0x27, 0x83, 0xd6, 0xa8, 0xa7, 0xff, 0xb6, 0x47,
	0xcf, 0x24, 0xf4, 0x39, 0x7a, 0x07, 0x9e, 0x10, 0xe9, 0x0b, 0x68, 0xc7, 0xe8, 0xde, 0x7b, 0x5b,
	0x57, 0x3c, 0x38, 0x91, 0xeb, 0x87, 0xd8, 0x29, 0xf7, 0xc9, 0xa0, 0xc4, 0x5b, 0x05, 0x6c, 0x4b,
	0x94, 0xbe, 0x01, 0xc5, 0x0d, 0x51, 0xc4, 0xe8, 0x48, 0x73, 0x9d, 0x76, 0x9f, 0x0c, 0x94, 0x51,
	0x57, 0x4f, 0x9d, 0xeb, 0xb9, 0x73, 0xfd, 0x43, 0xee, 0x9c, 0x43, 0x4a, 0x97, 0x80, 0xf6, 0x95,
	0x40, 0x45, 0x2e, 0xa5, 0x0a, 0xd4, 0x6f, 0xac, 0x6b, 0x6b, 0xb1, 0xb2, 0xd4, 0x13, 0x59, 0x98,
	0x33, 0x66, 0xdb, 0x53, 0x53, 0x25, 0xf4, 0x14, 0xda, 0x06, 0xb3, 0x98, 0xc5, 0x1c, 0x66, 0x5d,
	0x39, 0xab, 0xc9, 0x74, 0xa9, 0x96, 0xe8, 0x19, 0xa8, 0x06, 0x33, 0x17, 0x56, 0x82, 0x99, 0x93,
	0xf1, 0xd8, 0x1e, 0xab, 0x65, 0xda, 0x84, 0xc6, 0x7c, 0xfa, 0xd1, 0x31, 0xc6, 0x9c, 0xdf, 0xaa,
	0x15, 0x7a, 0x0e, 0x4f, 0x0c, 0x76, 0x3d, 0xbe, 0x72, 0xe6, 0x8c, 0xdb, 0x93, 0x39, 0x9b, 0xcd,
	0x16, 0x2b, 0xb5, 0x4a, 0xdb, 0xa0, 0xd8, 0xcb, 0xa9, 0x79, 0xeb, 0x98, 0x37, 0x92, 0x57, 0xd3,
	0xbe, 0x11, 0xa8, 0x71, 0x0c, 0xfc, 0x30, 0xa6, 0x97, 0xa0, 0xc8, 0x44, 0x1c, 0xd7, 0x3f, 0x78,
	0x71, 0xd4, 0x21, 0xfd, 0xf2, 0x40, 0x19, 0x5d, 0x3c, 0x4a, 0x2b, 0xe5, 0xea, 0x86, 0xd8, 0xa1,
	0x29, 0x79, 0x1c, 0xd6, 0xf9, 0x67, 0xd4, 0xe5, 0xd0, 0x28, 0x1a, 0x45, 0xea, 0xe4, 0x6f, 0x53,
	0x3f, 0x83, 0x6a, 0xb2, 0x3a, 0xb9, 0x53, 0x95, 0xa7, 0x85, 0xf6, 0x16, 0x14, 0xa9, 0xc9, 0xf1,
	0xf3, 0x01, 0xa3, 0x7f, 0x57, 0xd5, 0x2e, 0xe1, 0xff, 0x74, 0x3e, 0x0a, 0x7c, 0x2f, 0x42, 0xfa,
	0x0a, 0xea, 0xd9, 0x4c, 0xa2, 0xa1, 0x8c, 0x9e, 0x1e, 0xd7, 0xe0, 0x39, 0x4d, 0x6b, 0x43, 0x33,
	0xfd, 0xeb, 0xcc, 0x83, 0xf6, 0x0e, 0x5a, 0x39, 0x90, 0x89, 0xbe, 0x84, 0x5a, 0x98, 0x20, 0x99,
	0xe6, 0xf9, 0xd1, 0xd4, 0x78, 0x46, 0x1a, 0x7d, 0x27, 0x70, 0x9a, 0xad, 0x91, 0xde, 0x42, 0x1b,
	0xc3, 0x2f, 0x5b, 0x17, 0x29, 0x83, 0x8a, 0xac, 0x69, 0xf7, 0xd1, 0xf8, 0x2f, 0x01, 0x74, 0x7b,
	0x47, 0x7b, 0xa9, 0x0f, 0xed, 0x84, 0xbe, 0x2f, 0xce, 0xf9, 0xec, 0xb8, 0x87, 0x4c, 0xe6, 0xf9,
	0x1f, 0xba, 0xb9, 0x90, 0xd1, 0xf8, 0x54, 0xdf, 0xa0, 0x37, 0x14, 0xc1, 0x76, 0x5d, 0x4b, 0x1e,
	0xf2, 0xeb, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x59, 0x9d, 0x12, 0xb0, 0x03, 0x00, 0x00,
}
