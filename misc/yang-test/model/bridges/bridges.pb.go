// Code generated by protoc-gen-go.
// source: bridges.proto
// DO NOT EDIT!

/*
Package bridges is a generated protocol buffer package.

It is generated from these files:
	bridges.proto

It has these top-level messages:
	Bridges
*/
package bridges

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// NLAN bridges module
type Bridges struct {
	OvsBridges       *bool  `protobuf:"varint,1,opt,name=OvsBridges" json:"OvsBridges,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Bridges) Reset()         { *m = Bridges{} }
func (m *Bridges) String() string { return proto.CompactTextString(m) }
func (*Bridges) ProtoMessage()    {}

func (m *Bridges) GetOvsBridges() bool {
	if m != nil && m.OvsBridges != nil {
		return *m.OvsBridges
	}
	return false
}