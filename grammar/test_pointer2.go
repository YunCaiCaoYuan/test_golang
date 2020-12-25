package main

import (
	"fmt"
	"time"
)

type ChannelTagConf struct {
	Id          int32     `json:"id" gorm:"column:id;PRIMARY_KEY"`
	Name        string    `json:"name" gorm:"column:name;PRIMARY_KEY"`
	TagImage    string    `json:"tagImage" gorm:"column:tag_image"`
	DefaultIcon string    `json:"defaultIcon" gorm:"column:default_icon"`
	Operator    string    `json:"operator" gorm:"column:operator"`
	Weight      int       `json:"weight" gorm:"column:weight"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

func GetChannelTag() *ChannelTagConf {
	return nil
}

type ChannelTag struct {
	Id          int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	TagImage    string `protobuf:"bytes,3,opt,name=tag_image,json=tagImage" json:"tag_image,omitempty"`
	DefaultIcon string `protobuf:"bytes,4,opt,name=default_icon,json=defaultIcon" json:"default_icon,omitempty"`
}

func (c *ChannelTagConf) ToPBChannelTag() *ChannelTag {
	if c == nil {
		return nil
	}
	pbChannelTag := &ChannelTag{
		Id:          c.Id,
		Name:        c.Name,
		TagImage:    c.TagImage,
		DefaultIcon: c.DefaultIcon,
	}
	return pbChannelTag
}

func main() {
	channelTagConf := GetChannelTag()
	fmt.Printf("channelTagConf=%T\n", channelTagConf) // *main.ChannelTagConf
	pbChannelTag := channelTagConf.ToPBChannelTag()
	fmt.Println("pbChannelTag=", pbChannelTag)
}
