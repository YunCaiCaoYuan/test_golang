package main


type StartGameReq struct {
	Players []*GamePlayer `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
	Uuid    string        `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
}

type GamePlayer struct {
	Id       int64    `protobuf:"zigzag64,1,opt,name=id" json:"id,omitempty"`
	Nickname string   `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	Sex      GSexType `protobuf:"varint,3,opt,name=sex,enum=game_pb.GSexType" json:"sex,omitempty"`
	Id2      int64    `protobuf:"zigzag64,4,opt,name=id2" json:"id2,omitempty"`
	Icon     string   `protobuf:"bytes,5,opt,name=icon" json:"icon,omitempty"`
	IsRobot  bool     `protobuf:"varint,6,opt,name=is_robot,json=isRobot" json:"is_robot,omitempty"`
}

type GSexType int32


func main() {

}
