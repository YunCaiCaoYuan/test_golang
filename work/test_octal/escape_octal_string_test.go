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
	s := "\\032\\016\\n\\nno_auth_id\\022\\000\\032\\032\\n\\tclient-ip\\022\\r192.168.0.200\\032 \\n\\014wired-connid\\022\\0202050000000459643\\032\\031\\n\\007version\\022\\0162.0.1-SNAPSHOT\\032\\016\\n\\003uid\\022\\0071001010\\032Q\\n\\006client\\022Gweb;86c96ab8-c805-11eb-9234-0a44df474c03;2.0.1-SNAPSHOT;2.0.1-SNAPSHOT;\\032\\033\\n\\023heartbeat-threshold\\022\\0045m0s\\032!\\n\\023user-heart-interval\\022\\n1623124397\\032\\024\\n\\006connid\\022\\n1767675782\\032\\n\\n\\005appid\\022\\0018\\032\\031\\n\\016consisthashkey\\022\\0071001010\\032 \\n\\017access-point-ip\\022\\r192.168.0.205\\032\\022\\n\\013application\\022\\003yes\\032\\024\\n\\013client-port\\022\\00558098\\032\\031\\n\\021access-point-port\\022\\0048091\\032\\035\\n\\006uni-id\\022\\0231623124407806112097\\032\\013\\n\\007X-Token\\022\\000\\032\\013\\n\\003env\\022\\004beta\\032\\014\\n\\006roomid\\022\\002-1*\\023yes.user.UserExtObj2\\020saveOssImageList\""

	ret := EscapeOctalString(s)
	t.Log("原转义\n", []byte(ret))
	t.Log("转义\n", ret)

	BroadcastMailMsg := new(BroadcastMailMsg)
	err := proto.Unmarshal([]byte(ret), BroadcastMailMsg)
	t.Log("err=", err)

	t.Log("BroadcastMailMsg=", BroadcastMailMsg)
}
