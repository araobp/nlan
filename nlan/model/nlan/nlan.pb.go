// Code generated by protoc-gen-go.
// source: nlan.proto
// DO NOT EDIT!

/*
Package nlan is a generated protocol buffer package.

It is generated from these files:
	nlan.proto

It has these top-level messages:
	Nlan
	Request
	Model
	Dvr
	Subnets
	IpDvr
	Vxlan
	Ptn
	PtnL2Vpn
	PtnLinks
	PtnNodes
	Nodes
	Response
*/
package nlan

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// NLAN module
type Nlan struct {
	Request  *Request  `protobuf:"bytes,1,opt,name=Request" json:"Request,omitempty"`
	Response *Response `protobuf:"bytes,2,opt,name=Response" json:"Response,omitempty"`
}

func (m *Nlan) Reset()         { *m = Nlan{} }
func (m *Nlan) String() string { return proto.CompactTextString(m) }
func (*Nlan) ProtoMessage()    {}

func (m *Nlan) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Nlan) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type Request struct {
	Model *Model `protobuf:"bytes,1,opt,name=Model" json:"Model,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetModel() *Model {
	if m != nil {
		return m.Model
	}
	return nil
}

type Model struct {
	Dvr *Dvr `protobuf:"bytes,1,opt,name=Dvr" json:"Dvr,omitempty"`
	Ptn *Ptn `protobuf:"bytes,2,opt,name=Ptn" json:"Ptn,omitempty"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}

func (m *Model) GetDvr() *Dvr {
	if m != nil {
		return m.Dvr
	}
	return nil
}

func (m *Model) GetPtn() *Ptn {
	if m != nil {
		return m.Ptn
	}
	return nil
}

type Dvr struct {
	OvsBridges bool       `protobuf:"varint,1,opt,name=OvsBridges" json:"OvsBridges,omitempty"`
	Subnets    []*Subnets `protobuf:"bytes,2,rep,name=Subnets" json:"Subnets,omitempty"`
	Vxlan      []*Vxlan   `protobuf:"bytes,3,rep,name=Vxlan" json:"Vxlan,omitempty"`
}

func (m *Dvr) Reset()         { *m = Dvr{} }
func (m *Dvr) String() string { return proto.CompactTextString(m) }
func (*Dvr) ProtoMessage()    {}

func (m *Dvr) GetSubnets() []*Subnets {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func (m *Dvr) GetVxlan() []*Vxlan {
	if m != nil {
		return m.Vxlan
	}
	return nil
}

type Subnets struct {
	IpDvr []*IpDvr `protobuf:"bytes,1,rep,name=IpDvr" json:"IpDvr,omitempty"`
	Peers []string `protobuf:"bytes,2,rep,name=Peers" json:"Peers,omitempty"`
	Ports []string `protobuf:"bytes,3,rep,name=Ports" json:"Ports,omitempty"`
	Vid   uint32   `protobuf:"varint,4,opt,name=Vid" json:"Vid,omitempty"`
	Vni   uint32   `protobuf:"varint,5,opt,name=Vni" json:"Vni,omitempty"`
}

func (m *Subnets) Reset()         { *m = Subnets{} }
func (m *Subnets) String() string { return proto.CompactTextString(m) }
func (*Subnets) ProtoMessage()    {}

func (m *Subnets) GetIpDvr() []*IpDvr {
	if m != nil {
		return m.IpDvr
	}
	return nil
}

type IpDvr struct {
	Addr string `protobuf:"bytes,1,opt,name=Addr" json:"Addr,omitempty"`
	Dhcp string `protobuf:"bytes,2,opt,name=Dhcp" json:"Dhcp,omitempty"`
	Mode string `protobuf:"bytes,3,opt,name=Mode" json:"Mode,omitempty"`
}

func (m *IpDvr) Reset()         { *m = IpDvr{} }
func (m *IpDvr) String() string { return proto.CompactTextString(m) }
func (*IpDvr) ProtoMessage()    {}

type Vxlan struct {
	LocalIp   string   `protobuf:"bytes,1,opt,name=LocalIp" json:"LocalIp,omitempty"`
	RemoteIps []string `protobuf:"bytes,2,rep,name=RemoteIps" json:"RemoteIps,omitempty"`
}

func (m *Vxlan) Reset()         { *m = Vxlan{} }
func (m *Vxlan) String() string { return proto.CompactTextString(m) }
func (*Vxlan) ProtoMessage()    {}

type Ptn struct {
	PtnL2Vpn []*PtnL2Vpn `protobuf:"bytes,1,rep,name=PtnL2Vpn" json:"PtnL2Vpn,omitempty"`
	PtnLinks []*PtnLinks `protobuf:"bytes,2,rep,name=PtnLinks" json:"PtnLinks,omitempty"`
	PtnNodes []*PtnNodes `protobuf:"bytes,3,rep,name=PtnNodes" json:"PtnNodes,omitempty"`
}

func (m *Ptn) Reset()         { *m = Ptn{} }
func (m *Ptn) String() string { return proto.CompactTextString(m) }
func (*Ptn) ProtoMessage()    {}

func (m *Ptn) GetPtnL2Vpn() []*PtnL2Vpn {
	if m != nil {
		return m.PtnL2Vpn
	}
	return nil
}

func (m *Ptn) GetPtnLinks() []*PtnLinks {
	if m != nil {
		return m.PtnLinks
	}
	return nil
}

func (m *Ptn) GetPtnNodes() []*PtnNodes {
	if m != nil {
		return m.PtnNodes
	}
	return nil
}

type PtnL2Vpn struct {
	Id    string   `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Ip    string   `protobuf:"bytes,2,opt,name=Ip" json:"Ip,omitempty"`
	Peers []string `protobuf:"bytes,3,rep,name=Peers" json:"Peers,omitempty"`
	Vid   uint32   `protobuf:"varint,4,opt,name=Vid" json:"Vid,omitempty"`
	Vni   uint32   `protobuf:"varint,5,opt,name=Vni" json:"Vni,omitempty"`
}

func (m *PtnL2Vpn) Reset()         { *m = PtnL2Vpn{} }
func (m *PtnL2Vpn) String() string { return proto.CompactTextString(m) }
func (*PtnL2Vpn) ProtoMessage()    {}

type PtnLinks struct {
	Id        string   `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	LocalIp   string   `protobuf:"bytes,2,opt,name=LocalIp" json:"LocalIp,omitempty"`
	RemoteIps []string `protobuf:"bytes,3,rep,name=RemoteIps" json:"RemoteIps,omitempty"`
}

func (m *PtnLinks) Reset()         { *m = PtnLinks{} }
func (m *PtnLinks) String() string { return proto.CompactTextString(m) }
func (*PtnLinks) ProtoMessage()    {}

type PtnNodes struct {
	Id    string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Nodes *Nodes `protobuf:"bytes,2,opt,name=Nodes" json:"Nodes,omitempty"`
}

func (m *PtnNodes) Reset()         { *m = PtnNodes{} }
func (m *PtnNodes) String() string { return proto.CompactTextString(m) }
func (*PtnNodes) ProtoMessage()    {}

func (m *PtnNodes) GetNodes() *Nodes {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Nodes struct {
	L2Sw string `protobuf:"bytes,1,opt,name=L2Sw" json:"L2Sw,omitempty"`
	Ptn  string `protobuf:"bytes,2,opt,name=Ptn" json:"Ptn,omitempty"`
}

func (m *Nodes) Reset()         { *m = Nodes{} }
func (m *Nodes) String() string { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()    {}

type Response struct {
	Exit       uint32 `protobuf:"varint,1,opt,name=Exit" json:"Exit,omitempty"`
	LogMessage string `protobuf:"bytes,2,opt,name=LogMessage" json:"LogMessage,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for NlanAgent service

type NlanAgentClient interface {
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type nlanAgentClient struct {
	cc *grpc.ClientConn
}

func NewNlanAgentClient(cc *grpc.ClientConn) NlanAgentClient {
	return &nlanAgentClient{cc}
}

func (c *nlanAgentClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/nlan.NlanAgent/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nlanAgentClient) Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/nlan.NlanAgent/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nlanAgentClient) Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/nlan.NlanAgent/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NlanAgent service

type NlanAgentServer interface {
	Add(context.Context, *Request) (*Response, error)
	Update(context.Context, *Request) (*Response, error)
	Delete(context.Context, *Request) (*Response, error)
}

func RegisterNlanAgentServer(s *grpc.Server, srv NlanAgentServer) {
	s.RegisterService(&_NlanAgent_serviceDesc, srv)
}

func _NlanAgent_Add_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Request)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(NlanAgentServer).Add(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _NlanAgent_Update_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Request)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(NlanAgentServer).Update(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _NlanAgent_Delete_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Request)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(NlanAgentServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _NlanAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nlan.NlanAgent",
	HandlerType: (*NlanAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _NlanAgent_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NlanAgent_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NlanAgent_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
