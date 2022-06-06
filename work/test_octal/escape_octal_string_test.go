package test_octal

import (
	"github.com/golang/protobuf/proto"
	"testing"
)

type MailMsg struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	CmdId    int32  `protobuf:"varint,2,opt,name=cmd_id,json=cmdId" json:"cmd_id,omitempty"`
	CreateAt int32  `protobuf:"varint,3,opt,name=create_at,json=createAt" json:"create_at,omitempty"`
	Data     []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	MailType int32  `protobuf:"varint,5,opt,name=mail_type,json=mailType" json:"mail_type,omitempty"`
}

func (m *MailMsg) Reset()         { *m = MailMsg{} }
func (m *MailMsg) String() string { return proto.CompactTextString(m) }
func (*MailMsg) ProtoMessage()    {}

type BroadcastMailMsg struct {
	Mail *MailMsg `protobuf:"bytes,1,opt,name=mail" json:"mail,omitempty"`
}

func (m *BroadcastMailMsg) Reset()         { *m = BroadcastMailMsg{} }
func (m *BroadcastMailMsg) String() string { return proto.CompactTextString(m) }
func (*BroadcastMailMsg) ProtoMessage()    {}

func TestEscapeOctalString(t *testing.T) {
	//s := "\\344\\275\\240\\345\\245\\275"
	//s := "\\344\\275\\240\\345\\245\\275123"
	s := "\\350\\216\\267\\345\\217\\226\\345\\257\\206\\351\\222\\245\\345\\244\\261\\350\\264\\245"
	ret := EscapeOctalString(s)
	t.Log("原转义\n", []byte(ret))
	t.Log("转义\n", ret)

	BroadcastMailMsg := new(BroadcastMailMsg)
	err := proto.Unmarshal([]byte(ret), BroadcastMailMsg)
	t.Log("err=", err)

	t.Log("BroadcastMailMsg=", BroadcastMailMsg)
}
