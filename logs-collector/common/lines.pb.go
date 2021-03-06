// Code generated by protoc-gen-go.
// source: common/lines.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/lines.proto

It has these top-level messages:
	LogLines
*/
package common

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type LogLines struct {
	FileName         *string  `protobuf:"bytes,1,req" json:"FileName,omitempty"`
	Offset           *int64   `protobuf:"varint,2,req" json:"Offset,omitempty"`
	Inode            *uint64  `protobuf:"varint,3,req" json:"Inode,omitempty"`
	Lines            []string `protobuf:"bytes,4,rep" json:"Lines,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *LogLines) Reset()         { *m = LogLines{} }
func (m *LogLines) String() string { return proto.CompactTextString(m) }
func (*LogLines) ProtoMessage()    {}

func (m *LogLines) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *LogLines) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *LogLines) GetInode() uint64 {
	if m != nil && m.Inode != nil {
		return *m.Inode
	}
	return 0
}

func (m *LogLines) GetLines() []string {
	if m != nil {
		return m.Lines
	}
	return nil
}

func init() {
}
