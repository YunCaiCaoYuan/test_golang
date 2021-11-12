package main

import "fmt"

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
	var req StartGameReq
	list1 := make([]*GamePlayer, 0, 20)
	for i:=1;  i<= 10; i++ {
		list1 = append(list1, &GamePlayer{Id: int64(i), IsRobot: false})
	}

	for i:=11;  i< 20; i++ {
		list1 = append(list1, &GamePlayer{Id: int64(i), IsRobot: true})
	}
	req.Players = list1
	fmt.Println("before:")
	for _, tmp := range req.Players {
		fmt.Println(tmp)
	}

	robotNum := 0
	total := len(req.Players)
	for i:=0; i<total; i++ {
		if req.Players[i].IsRobot {
			robotNum++
		}
	}

	for i:=0; i<robotNum/2; i++ {
		tmpPlayer := req.Players[i]
		req.Players[i] = req.Players[total-i-1]
		req.Players[total-i-1] = tmpPlayer
	}
	fmt.Println("after :")
	for _, tmp := range req.Players {
		fmt.Println(tmp)
	}
}
