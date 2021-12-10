package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)


type TBBoard struct {
	Timestamp int64      `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Broad     []*TBArray `protobuf:"bytes,2,rep,name=broad" json:"broad,omitempty"`
}

func (m *TBBoard) Reset()                    { *m = TBBoard{} }
func (m *TBBoard) String() string            { return proto.CompactTextString(m) }
func (*TBBoard) ProtoMessage()               {}
func (*TBBoard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *TBBoard) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TBBoard) GetBroad() []*TBArray {
	if m != nil {
		return m.Broad
	}
	return nil
}

type TBArray struct {
	Val []int32 `protobuf:"varint,1,rep,packed,name=val" json:"val,omitempty"`
}

func (m *TBArray) Reset()                    { *m = TBArray{} }
func (m *TBArray) String() string            { return proto.CompactTextString(m) }
func (*TBArray) ProtoMessage()               {}
func (*TBArray) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TBArray) GetVal() []int32 {
	if m != nil {
		return m.Val
	}
	return nil
}

var fileDescriptor0 = []byte{
	// 1239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0xcb, 0x6e, 0x1b, 0x37,
	0x17, 0xce, 0x68, 0x2c, 0x59, 0x3a, 0x92, 0xed, 0x11, 0xe3, 0x24, 0x8a, 0x9d, 0x8b, 0xa3, 0xfc,
	0x41, 0x0c, 0xff, 0x85, 0xdd, 0x38, 0x8b, 0x16, 0xbd, 0x20, 0x95, 0x06, 0x4a, 0x20, 0xc0, 0x91,
	0xd2, 0x91, 0x83, 0x2e, 0x07, 0x94, 0x86, 0x91, 0x27, 0x9e, 0x8b, 0x4c, 0x52, 0x41, 0x03, 0x14,
	0xe8, 0xbe, 0x0f, 0xd0, 0x37, 0xe9, 0x3b, 0x74, 0x5b, 0xb4, 0x40, 0xa1, 0x5e, 0xde, 0xa4, 0x8b,
	0x82, 0xe4, 0x68, 0x34, 0x17, 0x45, 0x71, 0xba, 0xd2, 0xf0, 0x3b, 0x17, 0xf2, 0x7c, 0xe4, 0xf9,
	0x48, 0x81, 0x39, 0xe5, 0xae, 0x77, 0x34, 0xa1, 0x21, 0x0f, 0x47, 0xa1, 0x77, 0x34, 0xc6, 0x3e,
	0xb1, 0x27, 0x43, 0xf5, 0x4b, 0x02, 0x4e, 0xdf, 0xb2, 0x23, 0x4e, 0x38, 0x75, 0x99, 0x3d, 0xc4,
	0x9c, 0x7b, 0x24, 0x3d, 0x3a, 0x94, 0x81, 0x68, 0x23, 0x05, 0xee, 0x3c, 0x5c, 0x91, 0x73, 0x14,
	0xfa, 0x7e, 0x18, 0xa8, 0xb8, 0x9d, 0xff, 0xaf, 0x70, 0x64, 0x84, 0xbe, 0x21, 0xf4, 0x90, 0x7c,
	0xcb, 0x95, 0x73, 0xf3, 0x27, 0x1d, 0xca, 0x2f, 0x19, 0xa1, 0xdd, 0xe0, 0x55, 0x88, 0x1e, 0xc2,
	0xda, 0x94, 0x11, 0xda, 0xd0, 0xf6, 0xb4, 0xfd, 0xea, 0xf1, 0xd5, 0xc3, 0x28, 0xf4, 0xf0, 0x19,
	0xf6, 0xc9, 0x0b, 0x0f, 0xbf, 0x25, 0xd4, 0x92, 0x0e, 0xe8, 0x11, 0x94, 0x18, 0xc7, 0x7c, 0xca,
	0x1a, 0x85, 0x3d, 0x6d, 0x7f, 0xf3, 0xf8, 0xe6, 0x61, 0xba, 0x00, 0x91, 0x71, 0x20, 0x1d, 0xac,
	0xc8, 0x11, 0x5d, 0x83, 0x92, 0xcb, 0xec, 0x70, 0xca, 0x1b, 0xfa, 0x9e, 0xb6, 0x5f, 0xb6, 0x8a,
	0x2e, 0xeb, 0x4f, 0x39, 0xba, 0x0d, 0xc0, 0x08, 0x75, 0xb1, 0x67, 0x07, 0x53, 0xbf, 0xb1, 0xb6,
	0xa7, 0xed, 0x17, 0xad, 0x8a, 0x42, 0x7a, 0x53, 0x5f, 0x98, 0xcf, 0x5d, 0xcf, 0x23, 0x54, 0x9a,
	0x8b, 0xca, 0xac, 0x10, 0x61, 0xfe, 0x08, 0x8a, 0x43, 0x1a, 0x62, 0xa7, 0x51, 0xda, 0xd3, 0xf7,
	0xab, 0xc7, 0xd7, 0x33, 0xcb, 0x38, 0x6d, 0xb7, 0x28, 0xc5, 0x6f, 0x2d, 0xe5, 0x84, 0x9e, 0x40,
	0x95, 0x53, 0xcc, 0xce, 0x6c, 0xcf, 0x0d, 0x08, 0x6b, 0xac, 0xcb, 0x98, 0x3b, 0xf9, 0x18, 0xc7,
	0x39, 0x15, 0x6e, 0x27, 0xc2, 0xcb, 0x02, 0x1e, 0x7f, 0xa3, 0x7b, 0x50, 0xc3, 0x9c, 0xe3, 0xd1,
	0x79, 0x94, 0xa1, 0x2c, 0xd7, 0x53, 0x55, 0x98, 0x72, 0xd9, 0x86, 0xe2, 0x28, 0xf4, 0x87, 0x61,
	0xa3, 0x22, 0x6d, 0x6a, 0x20, 0x02, 0x29, 0xe1, 0xd8, 0x0d, 0xec, 0x89, 0x4b, 0x46, 0xa4, 0x01,
	0x2a, 0x50, 0x61, 0x2f, 0x04, 0x84, 0x6e, 0x42, 0xf9, 0x15, 0xf6, 0x14, 0x0d, 0x55, 0x69, 0x5e,
	0x17, 0x63, 0x51, 0x25, 0x82, 0x35, 0x8a, 0x83, 0xf3, 0x46, 0x4d, 0xc2, 0xf2, 0xbb, 0xf9, 0x63,
	0x01, 0x8c, 0xb6, 0x79, 0x2a, 0x97, 0xde, 0x0e, 0x31, 0x75, 0x2c, 0x72, 0xb1, 0xa0, 0x43, 0xbb,
	0x0c, 0x1d, 0xd9, 0x6a, 0x0a, 0xf9, 0x6a, 0x32, 0x8c, 0xe9, 0x1f, 0xcc, 0xd8, 0x2d, 0xa8, 0x70,
	0xd7, 0x27, 0x8c, 0x63, 0x7f, 0x22, 0x77, 0x57, 0xb7, 0x16, 0xc0, 0x82, 0xac, 0xe2, 0x2a, 0xb2,
	0x4a, 0xab, 0xc9, 0x5a, 0x4f, 0x91, 0xd5, 0xdc, 0x85, 0xf5, 0xa8, 0x4e, 0x64, 0x80, 0xfe, 0x06,
	0x7b, 0x92, 0x8c, 0xa2, 0x25, 0x3e, 0x9b, 0x04, 0xb6, 0xdb, 0xe6, 0x80, 0xf0, 0x96, 0xac, 0xb1,
	0x3f, 0x7c, 0x4d, 0x46, 0x5c, 0x10, 0xf7, 0x31, 0xac, 0xf9, 0xa1, 0x43, 0xe4, 0xc1, 0xdf, 0x3c,
	0xbe, 0x95, 0x2b, 0xd0, 0x56, 0x01, 0xcf, 0x43, 0x87, 0x58, 0xd2, 0x13, 0xed, 0x42, 0x85, 0x63,
	0x3a, 0x26, 0xdc, 0x76, 0x1d, 0xc9, 0x9c, 0x6e, 0x95, 0x15, 0xd0, 0x75, 0x9a, 0x13, 0x40, 0x6d,
	0x33, 0x4d, 0x0a, 0xb9, 0x48, 0x87, 0x68, 0xe9, 0x10, 0x74, 0x37, 0xcd, 0xb4, 0xda, 0x8b, 0x77,
	0x32, 0xa9, 0x67, 0x98, 0x6c, 0x86, 0xb0, 0xd5, 0x36, 0x55, 0x8b, 0x9e, 0x84, 0x8c, 0x88, 0xe9,
	0xb2, 0xdb, 0xab, 0xe5, 0xb7, 0x77, 0x17, 0xa2, 0x5e, 0x4a, 0x14, 0xa1, 0x80, 0xae, 0xf3, 0x9e,
	0x09, 0x1d, 0xc1, 0xa4, 0x9a, 0xb0, 0xe3, 0xb9, 0xbe, 0x1b, 0x60, 0x4e, 0x2c, 0xc2, 0xd0, 0x11,
	0x94, 0x26, 0x12, 0x8d, 0x44, 0xe4, 0xc6, 0x12, 0x65, 0x10, 0x5a, 0x63, 0x45, 0x6e, 0x2b, 0xd7,
	0xd0, 0xfc, 0x02, 0x2a, 0x6d, 0xd3, 0xc2, 0xc1, 0xb9, 0x4a, 0xbd, 0xe6, 0xb9, 0x8c, 0x47, 0x87,
	0x7b, 0x37, 0xb7, 0x49, 0x91, 0x42, 0x09, 0x6f, 0xe9, 0xd8, 0xfc, 0x1e, 0x6a, 0x49, 0x14, 0x7d,
	0x0a, 0x55, 0xa5, 0x68, 0x97, 0x5a, 0x20, 0x8c, 0x63, 0xd5, 0xbb, 0x4c, 0xab, 0xcc, 0x9b, 0x54,
	0x4f, 0x34, 0xe9, 0x09, 0xd4, 0x53, 0x3d, 0xca, 0x44, 0x19, 0x9f, 0xa4, 0xca, 0xb8, 0x9f, 0x2d,
	0x23, 0xe1, 0x3d, 0x20, 0x17, 0x53, 0x12, 0x8c, 0x48, 0x54, 0xce, 0x19, 0x6c, 0x2f, 0xb3, 0x0a,
	0x06, 0x55, 0x45, 0x89, 0x73, 0xa5, 0x80, 0xae, 0x83, 0x1e, 0x43, 0x85, 0x0f, 0xed, 0xa1, 0x8c,
	0x68, 0x14, 0xde, 0x21, 0x0b, 0x4a, 0x40, 0xca, 0x7c, 0xa8, 0x32, 0x37, 0x5f, 0x8a, 0x1e, 0x92,
	0xdf, 0xe9, 0x53, 0xa0, 0x65, 0x1b, 0x38, 0x16, 0x9c, 0xc2, 0x25, 0x04, 0xa7, 0xf9, 0x54, 0xd0,
	0x91, 0x6e, 0x3d, 0x86, 0x1e, 0xa5, 0xe8, 0xb8, 0x9d, 0xcf, 0x90, 0xf4, 0x57, 0x44, 0x7c, 0x07,
	0x9b, 0x69, 0xfc, 0xbf, 0xf5, 0xef, 0x82, 0xb4, 0x42, 0x86, 0xb4, 0x54, 0xa7, 0xea, 0x99, 0xe6,
	0xee, 0x88, 0x56, 0x4b, 0x35, 0x37, 0x3a, 0x4e, 0xd5, 0xf0, 0x3e, 0x7d, 0x54, 0x45, 0xfc, 0xa0,
	0xc1, 0x56, 0xc6, 0xb2, 0x7a, 0x27, 0x57, 0x29, 0x4e, 0x56, 0x3e, 0xf4, 0xd5, 0xf2, 0x91, 0x15,
	0xe2, 0xe6, 0x97, 0x00, 0x6d, 0x53, 0xdc, 0xf2, 0xfd, 0x37, 0x84, 0x7e, 0x78, 0xa3, 0xfd, 0xa3,
	0x09, 0x35, 0x88, 0x0e, 0xa7, 0xf4, 0x52, 0x97, 0x3f, 0xfa, 0x1c, 0x8a, 0xe2, 0xfa, 0x9f, 0x6f,
	0xcc, 0x83, 0xe5, 0x87, 0x5d, 0x0e, 0x6c, 0x31, 0xbd, 0x08, 0x23, 0x96, 0x8a, 0x41, 0x9f, 0x41,
	0x2d, 0xd1, 0xae, 0xf3, 0xd3, 0xfb, 0xce, 0x7e, 0xad, 0x2e, 0xfa, 0x95, 0xa1, 0x07, 0xb0, 0x49,
	0x89, 0x2f, 0xef, 0x90, 0x28, 0x5a, 0x51, 0xb2, 0xa1, 0xd0, 0xb9, 0xdb, 0x7d, 0xd8, 0xe0, 0x21,
	0xc7, 0x5e, 0xec, 0xa5, 0x1e, 0x20, 0x35, 0x09, 0xce, 0x9d, 0xee, 0x42, 0x95, 0xe2, 0xc0, 0x09,
	0x7d, 0x9b, 0x11, 0xe2, 0xc8, 0xbb, 0x4a, 0xb7, 0x40, 0x41, 0x03, 0x42, 0x9c, 0x03, 0x07, 0xae,
	0x2f, 0xaf, 0x04, 0x6d, 0x40, 0xe5, 0xb4, 0xfd, 0x6c, 0x60, 0xf7, 0xc2, 0x80, 0x18, 0x57, 0xe2,
	0xe1, 0x37, 0xd8, 0xe5, 0x86, 0x86, 0x36, 0x01, 0xe4, 0x70, 0xc0, 0x31, 0xe5, 0x46, 0x21, 0x1e,
	0x3f, 0x27, 0x74, 0x4c, 0x0c, 0x1d, 0xd5, 0xa0, 0x2c, 0xc7, 0x9d, 0xc0, 0x31, 0xd6, 0x0e, 0xfe,
	0x07, 0xb0, 0x78, 0x56, 0x21, 0x80, 0x52, 0x67, 0x60, 0xb6, 0x5e, 0x74, 0x8c, 0x2b, 0xe2, 0xbb,
	0xdf, 0x3b, 0xe9, 0xf6, 0x3a, 0x86, 0x76, 0xf0, 0x15, 0x6c, 0xa4, 0x8e, 0xbb, 0x9a, 0xb3, 0xf5,
	0xdc, 0xee, 0xf5, 0x7b, 0xc2, 0x77, 0x0b, 0xaa, 0x72, 0xd8, 0x62, 0xcc, 0x1d, 0x07, 0x86, 0x16,
	0x03, 0x96, 0xac, 0xc7, 0x28, 0x1c, 0xfc, 0xac, 0x01, 0x4a, 0x95, 0x63, 0xfa, 0x4e, 0xd7, 0x51,
	0x79, 0xcc, 0xee, 0xbc, 0x94, 0x6b, 0xe2, 0xf4, 0x9a, 0xdd, 0xb8, 0xd6, 0x29, 0x33, 0x7e, 0x99,
	0xd5, 0xd1, 0x0e, 0x6c, 0x4b, 0x38, 0x73, 0x33, 0x18, 0xbf, 0xce, 0xea, 0x68, 0x2b, 0xca, 0x20,
	0x0e, 0x8e, 0xf1, 0xdb, 0xac, 0x8e, 0xae, 0x8a, 0xb5, 0x46, 0x39, 0xc4, 0xc1, 0x33, 0xfe, 0x98,
	0xd5, 0xd1, 0x0d, 0xa8, 0x4b, 0x30, 0x29, 0x75, 0xc6, 0x9f, 0x09, 0x43, 0xb2, 0xf1, 0x8d, 0xbf,
	0x66, 0x75, 0xd4, 0x00, 0xa4, 0x0c, 0xc9, 0x5e, 0x32, 0xfe, 0x9e, 0xd5, 0x8f, 0x7f, 0xd7, 0xd3,
	0xa5, 0x74, 0x03, 0xde, 0x1f, 0xbe, 0x16, 0x9a, 0x28, 0x29, 0x17, 0xf3, 0xa2, 0x6b, 0xf1, 0x2b,
	0x37, 0xc6, 0x2c, 0x72, 0xb1, 0xb3, 0x15, 0xc3, 0xbd, 0xb0, 0x45, 0xc7, 0x42, 0xa7, 0xca, 0x5f,
	0x4f, 0x5d, 0x15, 0xb3, 0x1d, 0x1b, 0xe7, 0xd0, 0xd2, 0x10, 0x13, 0xea, 0x4f, 0x09, 0x1f, 0x9d,
	0x49, 0x8e, 0x02, 0x3c, 0x61, 0x67, 0x21, 0x47, 0xb7, 0x63, 0xaf, 0x9c, 0x6d, 0x69, 0x92, 0x16,
	0x6c, 0xa4, 0xee, 0x10, 0x74, 0x37, 0xd3, 0x00, 0xd9, 0x57, 0x60, 0x3e, 0x45, 0x57, 0xe8, 0x6e,
	0xe6, 0xd5, 0x83, 0xee, 0xe7, 0xd2, 0xe4, 0xdf, 0x45, 0xf9, 0x54, 0x4b, 0xc4, 0xef, 0x5e, 0x2e,
	0x51, 0xf6, 0xe5, 0x93, 0x4f, 0xf3, 0x04, 0x6a, 0xc9, 0xe7, 0x0a, 0xba, 0x93, 0xcb, 0x91, 0x7a,
	0xcb, 0xe4, 0x12, 0x0c, 0x4b, 0xf2, 0xdf, 0xcb, 0xe3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe1,
	0x9e, 0x2a, 0x15, 0x69, 0x0d, 0x00, 0x00,
}


func main() {
	var BoardFrame = [][]*TBArray{
		{
			&TBArray{Val: []int32{0, 0, 0, 0, 2, 0, 0, 0, 0, 0}},
			&TBArray{Val: []int32{0, 0, 0, 0, 2, 0, 0, 0, 0, 0}},
			&TBArray{Val: []int32{0, 0, 0, 0, 2, 0, 0, 0, 0, 0}},
			&TBArray{Val: []int32{0, 0, 0, 0, 2, 0, 0, 0, 0, 0}},
		},
		{
			&TBArray{Val: []int32{0, 0, 0, 0, 2, 0, 0, 0, 0, 0}},
			&TBArray{Val: []int32{0, 0, 0, 0, 2, 0, 0, 0, 0, 0}},
			&TBArray{Val: []int32{0, 0, 0, 0, 2, 0, 0, 0, 0, 0}},
			&TBArray{Val: []int32{0, 0, 0, 0, 2, 0, 0, 0, 0, 0}},
		},
	}
	tbArray := BoardFrame[0]

	/*S
	tbArray := make([]*TBArray, 0, 12)
	//tbArray[0] = &TBArray{Val: Vals} // panic: runtime error: index out of range
	tbArray = append(tbArray,
		&TBArray{Val: []int32{0,0,0,0,2,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,2,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,2,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,2,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,0,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,0,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,0,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,0,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,0,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,0,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,0,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,0,0,0,0,0,0}},
		&TBArray{Val: []int32{0,0,0,0,0,0,0,0,0,0}})
	 */

	tbBoard := TBBoard{
		Timestamp:1234567890,
		Broad: tbArray,
	}

	ret, err := proto.Marshal(&tbBoard)
	if err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println("before ret=", ret)
	}

	tbArray1 := new(TBBoard)
	err = proto.Unmarshal(ret, tbArray1)
	if err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println("after ret=", tbArray1.Broad)
	}
}