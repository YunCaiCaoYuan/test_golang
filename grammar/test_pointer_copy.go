package main

import "fmt"

type ConfChannelStar struct {
	Id           int32                     `json:"id,omitempty" gorm:"column:id"`
	Name         string                    `json:"name,omitempty" gorm:"column:name"`
	Weight       int32                     `json:"weight,omitempty" gorm:"column:weight"`
	TaskType     int32     				   `json:"task_type,omitempty" gorm:"column:task_type"`
	Type         int32 					   `json:"type,omitempty" gorm:"column:type"`
	PotentialVal int32                     `json:"potential_val,omitempty" gorm:"column:potential_val"`
	OneVal       int32                     `json:"one_val,omitempty" gorm:"column:one_val"`
	TwoVal       int32                     `json:"two_val,omitempty" gorm:"column:two_val"`
	ThreeVal     int32                     `json:"three_val,omitempty" gorm:"column:three_val"`
}

// 星级 任务配置
type ConfChannelStarTask struct {
	Id           int32                   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name         string                  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Weight       int32                   `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	TaskType     int32     				 `protobuf:"varint,4,opt,name=task_type,json=taskType,enum=pb.ChannelStarTaskType" json:"task_type,omitempty"`
	PriorityType int32 					 `protobuf:"varint,5,opt,name=priority_type,json=priorityType,enum=pb.ChannelStarPriorityType" json:"priority_type,omitempty"`
	PotentialVal int32                   `protobuf:"varint,6,opt,name=potential_val,json=potentialVal" json:"potential_val,omitempty"`
	OneVal       int32                   `protobuf:"varint,7,opt,name=one_val,json=oneVal" json:"one_val,omitempty"`
	TwoVal       int32                   `protobuf:"varint,8,opt,name=two_val,json=twoVal" json:"two_val,omitempty"`
	ThreeVal     int32                   `protobuf:"varint,9,opt,name=three_val,json=threeVal" json:"three_val,omitempty"`
}

// 返回值 vs 返回引用
//func (c *ConfChannelStar) ToPB() ConfChannelStarTask {
//	ret := ConfChannelStarTask{Id : c.Id, Name: "ToPB"}
//	fmt.Printf("ret=%p\n", &ret)
//	return ret
//}
func (c *ConfChannelStar) ToPB() *ConfChannelStarTask {
	ret := &ConfChannelStarTask{Id : c.Id, Name: "ToPB"}
	fmt.Printf("ret=%p\n", ret)
	return ret
}

func main() {
	confList := make([]*ConfChannelStar, 0)
	confList = append(confList,
		&ConfChannelStar{Id:1,Name:"1"},
		//&ConfChannelStar{Id:2,Name:"2"}
	)

	pbTaskList := make([]*ConfChannelStarTask, 0)
	for _, item := range confList {
		pbTask := item.ToPB()
		fmt.Printf("pbTask=%p\n", pbTask)
		pbTaskList = append(pbTaskList, pbTask)
	}
}
