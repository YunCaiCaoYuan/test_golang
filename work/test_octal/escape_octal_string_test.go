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
func (m *MailMsg) Reset()                    { *m = MailMsg{} }
func (m *MailMsg) String() string            { return proto.CompactTextString(m) }
func (*MailMsg) ProtoMessage()               {}

type BroadcastMailMsg struct {
	Mail *MailMsg `protobuf:"bytes,1,opt,name=mail" json:"mail,omitempty"`
}
func (m *BroadcastMailMsg) Reset()                    { *m = BroadcastMailMsg{} }
func (m *BroadcastMailMsg) String() string            { return proto.CompactTextString(m) }
func (*BroadcastMailMsg) ProtoMessage()               {}

func TestEscapeOctalString(t *testing.T) {
	s := "\\nh\\010\\242\\376)\\020\\211\\225\\006\\030\\255\\367\\337\\202\\006\\\"X\\nP\\346\\202\\250\\344\\272\\2162021-03-22 10:03:58\\346\\217\\220\\344\\272\\244\\347\\232\\204\\347\\224\\250\\346\\210\\267\\345\\244\\264\\345\\203\\217\\345\\256\\241\\346\\240\\270\\346\\234\\252\\351\\200\\232\\350\\277\\207,\\350\\257\\267\\344\\277\\256\\346\\224\\271\\345\\220\\216\\346\\217\\220\\344\\272\\244\\020\\255\\367\\337\\202\\006"

	ret := EscapeOctalString(s)
	t.Log("原转义\n", []byte(ret))
	t.Log("转义\n", ret)

	BroadcastMailMsg := new(BroadcastMailMsg)
	err := proto.Unmarshal([]byte(ret), BroadcastMailMsg)
	t.Log("err=", err)

	t.Log("BroadcastMailMsg=", BroadcastMailMsg)
}
