package protocol

import (
	fmt "fmt"

	proto1 "github.com/gogo/protobuf/proto"

	math "math"

	binary "encoding/binary"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type BenchmarkMessage struct {
	Field1   string   `protobuf:"bytes,1,req,name=field1" json:"field1"`
	Field9   string   `protobuf:"bytes,9,opt,name=field9" json:"field9"`
	Field18  string   `protobuf:"bytes,18,opt,name=field18" json:"field18"`
	Field80  *bool    `protobuf:"varint,80,opt,name=field80,def=0" json:"field80,omitempty"`
	Field81  *bool    `protobuf:"varint,81,opt,name=field81,def=1" json:"field81,omitempty"`
	Field2   int32    `protobuf:"varint,2,req,name=field2" json:"field2"`
	Field3   int32    `protobuf:"varint,3,req,name=field3" json:"field3"`
	Field280 int32    `protobuf:"varint,280,opt,name=field280" json:"field280"`
	Field6   *int32   `protobuf:"varint,6,opt,name=field6,def=0" json:"field6,omitempty"`
	Field22  int64    `protobuf:"varint,22,opt,name=field22" json:"field22"`
	Field4   string   `protobuf:"bytes,4,opt,name=field4" json:"field4"`
	Field5   []uint64 `protobuf:"fixed64,5,rep,name=field5" json:"field5,omitempty"`
	Field59  *bool    `protobuf:"varint,59,opt,name=field59,def=0" json:"field59,omitempty"`
	Field7   string   `protobuf:"bytes,7,opt,name=field7" json:"field7"`
	Field16  int32    `protobuf:"varint,16,opt,name=field16" json:"field16"`
	Field130 *int32   `protobuf:"varint,130,opt,name=field130,def=0" json:"field130,omitempty"`
	Field12  *bool    `protobuf:"varint,12,opt,name=field12,def=1" json:"field12,omitempty"`
	Field17  *bool    `protobuf:"varint,17,opt,name=field17,def=1" json:"field17,omitempty"`
	Field13  *bool    `protobuf:"varint,13,opt,name=field13,def=1" json:"field13,omitempty"`
	Field14  *bool    `protobuf:"varint,14,opt,name=field14,def=1" json:"field14,omitempty"`
	Field104 *int32   `protobuf:"varint,104,opt,name=field104,def=0" json:"field104,omitempty"`
	Field100 *int32   `protobuf:"varint,100,opt,name=field100,def=0" json:"field100,omitempty"`
	Field101 *int32   `protobuf:"varint,101,opt,name=field101,def=0" json:"field101,omitempty"`
	Field102 string   `protobuf:"bytes,102,opt,name=field102" json:"field102"`
	Field103 string   `protobuf:"bytes,103,opt,name=field103" json:"field103"`
	Field29  *int32   `protobuf:"varint,29,opt,name=field29,def=0" json:"field29,omitempty"`
	Field30  *bool    `protobuf:"varint,30,opt,name=field30,def=0" json:"field30,omitempty"`
	Field60  *int32   `protobuf:"varint,60,opt,name=field60,def=-1" json:"field60,omitempty"`
	Field271 *int32   `protobuf:"varint,271,opt,name=field271,def=-1" json:"field271,omitempty"`
	Field272 *int32   `protobuf:"varint,272,opt,name=field272,def=-1" json:"field272,omitempty"`
	Field150 int32    `protobuf:"varint,150,opt,name=field150" json:"field150"`
	Field23  *int32   `protobuf:"varint,23,opt,name=field23,def=0" json:"field23,omitempty"`
	Field24  *bool    `protobuf:"varint,24,opt,name=field24,def=0" json:"field24,omitempty"`
	Field25  *int32   `protobuf:"varint,25,opt,name=field25,def=0" json:"field25,omitempty"`
	Field78  bool     `protobuf:"varint,78,opt,name=field78" json:"field78"`
	Field67  *int32   `protobuf:"varint,67,opt,name=field67,def=0" json:"field67,omitempty"`
	Field68  int32    `protobuf:"varint,68,opt,name=field68" json:"field68"`
	Field128 *int32   `protobuf:"varint,128,opt,name=field128,def=0" json:"field128,omitempty"`
	Field129 *string  `protobuf:"bytes,129,opt,name=field129,def=xxxxxxxxxxxxxxxxxxxxx" json:"field129,omitempty"`
	Field131 *int32   `protobuf:"varint,131,opt,name=field131,def=0" json:"field131,omitempty"`
}

func (m *BenchmarkMessage) Reset()                    { *m = BenchmarkMessage{} }
func (m *BenchmarkMessage) String() string            { return proto1.CompactTextString(m) }
func (*BenchmarkMessage) ProtoMessage()               {}
func (*BenchmarkMessage) Descriptor() ([]byte, []int) { return fileDescriptorBenchmark, []int{0} }

const Default_BenchmarkMessage_Field80 bool = false
const Default_BenchmarkMessage_Field81 bool = true
const Default_BenchmarkMessage_Field6 int32 = 0
const Default_BenchmarkMessage_Field59 bool = false
const Default_BenchmarkMessage_Field130 int32 = 0
const Default_BenchmarkMessage_Field12 bool = true
const Default_BenchmarkMessage_Field17 bool = true
const Default_BenchmarkMessage_Field13 bool = true
const Default_BenchmarkMessage_Field14 bool = true
const Default_BenchmarkMessage_Field104 int32 = 0
const Default_BenchmarkMessage_Field100 int32 = 0
const Default_BenchmarkMessage_Field101 int32 = 0
const Default_BenchmarkMessage_Field29 int32 = 0
const Default_BenchmarkMessage_Field30 bool = false
const Default_BenchmarkMessage_Field60 int32 = -1
const Default_BenchmarkMessage_Field271 int32 = -1
const Default_BenchmarkMessage_Field272 int32 = -1
const Default_BenchmarkMessage_Field23 int32 = 0
const Default_BenchmarkMessage_Field24 bool = false
const Default_BenchmarkMessage_Field25 int32 = 0
const Default_BenchmarkMessage_Field67 int32 = 0
const Default_BenchmarkMessage_Field128 int32 = 0
const Default_BenchmarkMessage_Field129 string = "xxxxxxxxxxxxxxxxxxxxx"
const Default_BenchmarkMessage_Field131 int32 = 0

func (m *BenchmarkMessage) GetField1() string {
	if m != nil {
		return m.Field1
	}
	return ""
}

func (m *BenchmarkMessage) GetField9() string {
	if m != nil {
		return m.Field9
	}
	return ""
}

func (m *BenchmarkMessage) GetField18() string {
	if m != nil {
		return m.Field18
	}
	return ""
}

func (m *BenchmarkMessage) GetField80() bool {
	if m != nil && m.Field80 != nil {
		return *m.Field80
	}
	return Default_BenchmarkMessage_Field80
}

func (m *BenchmarkMessage) GetField81() bool {
	if m != nil && m.Field81 != nil {
		return *m.Field81
	}
	return Default_BenchmarkMessage_Field81
}

func (m *BenchmarkMessage) GetField2() int32 {
	if m != nil {
		return m.Field2
	}
	return 0
}

func (m *BenchmarkMessage) GetField3() int32 {
	if m != nil {
		return m.Field3
	}
	return 0
}

func (m *BenchmarkMessage) GetField280() int32 {
	if m != nil {
		return m.Field280
	}
	return 0
}

func (m *BenchmarkMessage) GetField6() int32 {
	if m != nil && m.Field6 != nil {
		return *m.Field6
	}
	return Default_BenchmarkMessage_Field6
}

func (m *BenchmarkMessage) GetField22() int64 {
	if m != nil {
		return m.Field22
	}
	return 0
}

func (m *BenchmarkMessage) GetField4() string {
	if m != nil {
		return m.Field4
	}
	return ""
}

func (m *BenchmarkMessage) GetField5() []uint64 {
	if m != nil {
		return m.Field5
	}
	return nil
}

func (m *BenchmarkMessage) GetField59() bool {
	if m != nil && m.Field59 != nil {
		return *m.Field59
	}
	return Default_BenchmarkMessage_Field59
}

func (m *BenchmarkMessage) GetField7() string {
	if m != nil {
		return m.Field7
	}
	return ""
}

func (m *BenchmarkMessage) GetField16() int32 {
	if m != nil {
		return m.Field16
	}
	return 0
}

func (m *BenchmarkMessage) GetField130() int32 {
	if m != nil && m.Field130 != nil {
		return *m.Field130
	}
	return Default_BenchmarkMessage_Field130
}

func (m *BenchmarkMessage) GetField12() bool {
	if m != nil && m.Field12 != nil {
		return *m.Field12
	}
	return Default_BenchmarkMessage_Field12
}

func (m *BenchmarkMessage) GetField17() bool {
	if m != nil && m.Field17 != nil {
		return *m.Field17
	}
	return Default_BenchmarkMessage_Field17
}

func (m *BenchmarkMessage) GetField13() bool {
	if m != nil && m.Field13 != nil {
		return *m.Field13
	}
	return Default_BenchmarkMessage_Field13
}

func (m *BenchmarkMessage) GetField14() bool {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return Default_BenchmarkMessage_Field14
}

func (m *BenchmarkMessage) GetField104() int32 {
	if m != nil && m.Field104 != nil {
		return *m.Field104
	}
	return Default_BenchmarkMessage_Field104
}

func (m *BenchmarkMessage) GetField100() int32 {
	if m != nil && m.Field100 != nil {
		return *m.Field100
	}
	return Default_BenchmarkMessage_Field100
}

func (m *BenchmarkMessage) GetField101() int32 {
	if m != nil && m.Field101 != nil {
		return *m.Field101
	}
	return Default_BenchmarkMessage_Field101
}

func (m *BenchmarkMessage) GetField102() string {
	if m != nil {
		return m.Field102
	}
	return ""
}

func (m *BenchmarkMessage) GetField103() string {
	if m != nil {
		return m.Field103
	}
	return ""
}

func (m *BenchmarkMessage) GetField29() int32 {
	if m != nil && m.Field29 != nil {
		return *m.Field29
	}
	return Default_BenchmarkMessage_Field29
}

func (m *BenchmarkMessage) GetField30() bool {
	if m != nil && m.Field30 != nil {
		return *m.Field30
	}
	return Default_BenchmarkMessage_Field30
}

func (m *BenchmarkMessage) GetField60() int32 {
	if m != nil && m.Field60 != nil {
		return *m.Field60
	}
	return Default_BenchmarkMessage_Field60
}

func (m *BenchmarkMessage) GetField271() int32 {
	if m != nil && m.Field271 != nil {
		return *m.Field271
	}
	return Default_BenchmarkMessage_Field271
}

func (m *BenchmarkMessage) GetField272() int32 {
	if m != nil && m.Field272 != nil {
		return *m.Field272
	}
	return Default_BenchmarkMessage_Field272
}

func (m *BenchmarkMessage) GetField150() int32 {
	if m != nil {
		return m.Field150
	}
	return 0
}

func (m *BenchmarkMessage) GetField23() int32 {
	if m != nil && m.Field23 != nil {
		return *m.Field23
	}
	return Default_BenchmarkMessage_Field23
}

func (m *BenchmarkMessage) GetField24() bool {
	if m != nil && m.Field24 != nil {
		return *m.Field24
	}
	return Default_BenchmarkMessage_Field24
}

func (m *BenchmarkMessage) GetField25() int32 {
	if m != nil && m.Field25 != nil {
		return *m.Field25
	}
	return Default_BenchmarkMessage_Field25
}

func (m *BenchmarkMessage) GetField78() bool {
	if m != nil {
		return m.Field78
	}
	return false
}

func (m *BenchmarkMessage) GetField67() int32 {
	if m != nil && m.Field67 != nil {
		return *m.Field67
	}
	return Default_BenchmarkMessage_Field67
}

func (m *BenchmarkMessage) GetField68() int32 {
	if m != nil {
		return m.Field68
	}
	return 0
}

func (m *BenchmarkMessage) GetField128() int32 {
	if m != nil && m.Field128 != nil {
		return *m.Field128
	}
	return Default_BenchmarkMessage_Field128
}

func (m *BenchmarkMessage) GetField129() string {
	if m != nil && m.Field129 != nil {
		return *m.Field129
	}
	return Default_BenchmarkMessage_Field129
}

func (m *BenchmarkMessage) GetField131() int32 {
	if m != nil && m.Field131 != nil {
		return *m.Field131
	}
	return Default_BenchmarkMessage_Field131
}

func init() {
	proto1.RegisterType((*BenchmarkMessage)(nil), "proto.BenchmarkMessage")
}
func (m *BenchmarkMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BenchmarkMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(len(m.Field1)))
	i += copy(dAtA[i:], m.Field1)
	dAtA[i] = 0x10
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(m.Field2))
	dAtA[i] = 0x18
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(m.Field3))
	dAtA[i] = 0x22
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(len(m.Field4)))
	i += copy(dAtA[i:], m.Field4)
	if len(m.Field5) > 0 {
		for _, num := range m.Field5 {
			dAtA[i] = 0x29
			i++
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(num))
			i += 8
		}
	}
	if m.Field6 != nil {
		dAtA[i] = 0x30
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field6))
	}
	dAtA[i] = 0x3a
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(len(m.Field7)))
	i += copy(dAtA[i:], m.Field7)
	dAtA[i] = 0x4a
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(len(m.Field9)))
	i += copy(dAtA[i:], m.Field9)
	if m.Field12 != nil {
		dAtA[i] = 0x60
		i++
		if *m.Field12 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Field13 != nil {
		dAtA[i] = 0x68
		i++
		if *m.Field13 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Field14 != nil {
		dAtA[i] = 0x70
		i++
		if *m.Field14 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x80
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(m.Field16))
	if m.Field17 != nil {
		dAtA[i] = 0x88
		i++
		dAtA[i] = 0x1
		i++
		if *m.Field17 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x92
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(len(m.Field18)))
	i += copy(dAtA[i:], m.Field18)
	dAtA[i] = 0xb0
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(m.Field22))
	if m.Field23 != nil {
		dAtA[i] = 0xb8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field23))
	}
	if m.Field24 != nil {
		dAtA[i] = 0xc0
		i++
		dAtA[i] = 0x1
		i++
		if *m.Field24 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Field25 != nil {
		dAtA[i] = 0xc8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field25))
	}
	if m.Field29 != nil {
		dAtA[i] = 0xe8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field29))
	}
	if m.Field30 != nil {
		dAtA[i] = 0xf0
		i++
		dAtA[i] = 0x1
		i++
		if *m.Field30 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Field59 != nil {
		dAtA[i] = 0xd8
		i++
		dAtA[i] = 0x3
		i++
		if *m.Field59 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Field60 != nil {
		dAtA[i] = 0xe0
		i++
		dAtA[i] = 0x3
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field60))
	}
	if m.Field67 != nil {
		dAtA[i] = 0x98
		i++
		dAtA[i] = 0x4
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field67))
	}
	dAtA[i] = 0xa0
	i++
	dAtA[i] = 0x4
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(m.Field68))
	dAtA[i] = 0xf0
	i++
	dAtA[i] = 0x4
	i++
	if m.Field78 {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if m.Field80 != nil {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x5
		i++
		if *m.Field80 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Field81 != nil {
		dAtA[i] = 0x88
		i++
		dAtA[i] = 0x5
		i++
		if *m.Field81 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Field100 != nil {
		dAtA[i] = 0xa0
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field100))
	}
	if m.Field101 != nil {
		dAtA[i] = 0xa8
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field101))
	}
	dAtA[i] = 0xb2
	i++
	dAtA[i] = 0x6
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(len(m.Field102)))
	i += copy(dAtA[i:], m.Field102)
	dAtA[i] = 0xba
	i++
	dAtA[i] = 0x6
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(len(m.Field103)))
	i += copy(dAtA[i:], m.Field103)
	if m.Field104 != nil {
		dAtA[i] = 0xc0
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field104))
	}
	if m.Field128 != nil {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x8
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field128))
	}
	if m.Field129 != nil {
		dAtA[i] = 0x8a
		i++
		dAtA[i] = 0x8
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(len(*m.Field129)))
		i += copy(dAtA[i:], *m.Field129)
	}
	if m.Field130 != nil {
		dAtA[i] = 0x90
		i++
		dAtA[i] = 0x8
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field130))
	}
	if m.Field131 != nil {
		dAtA[i] = 0x98
		i++
		dAtA[i] = 0x8
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field131))
	}
	dAtA[i] = 0xb0
	i++
	dAtA[i] = 0x9
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(m.Field150))
	if m.Field271 != nil {
		dAtA[i] = 0xf8
		i++
		dAtA[i] = 0x10
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field271))
	}
	if m.Field272 != nil {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x11
		i++
		i = encodeVarintBenchmark(dAtA, i, uint64(*m.Field272))
	}
	dAtA[i] = 0xc0
	i++
	dAtA[i] = 0x11
	i++
	i = encodeVarintBenchmark(dAtA, i, uint64(m.Field280))
	return i, nil
}

func encodeVarintBenchmark(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BenchmarkMessage) Size() (n int) {
	var l int
	_ = l
	l = len(m.Field1)
	n += 1 + l + sovBenchmark(uint64(l))
	n += 1 + sovBenchmark(uint64(m.Field2))
	n += 1 + sovBenchmark(uint64(m.Field3))
	l = len(m.Field4)
	n += 1 + l + sovBenchmark(uint64(l))
	if len(m.Field5) > 0 {
		n += 9 * len(m.Field5)
	}
	if m.Field6 != nil {
		n += 1 + sovBenchmark(uint64(*m.Field6))
	}
	l = len(m.Field7)
	n += 1 + l + sovBenchmark(uint64(l))
	l = len(m.Field9)
	n += 1 + l + sovBenchmark(uint64(l))
	if m.Field12 != nil {
		n += 2
	}
	if m.Field13 != nil {
		n += 2
	}
	if m.Field14 != nil {
		n += 2
	}
	n += 2 + sovBenchmark(uint64(m.Field16))
	if m.Field17 != nil {
		n += 3
	}
	l = len(m.Field18)
	n += 2 + l + sovBenchmark(uint64(l))
	n += 2 + sovBenchmark(uint64(m.Field22))
	if m.Field23 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field23))
	}
	if m.Field24 != nil {
		n += 3
	}
	if m.Field25 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field25))
	}
	if m.Field29 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field29))
	}
	if m.Field30 != nil {
		n += 3
	}
	if m.Field59 != nil {
		n += 3
	}
	if m.Field60 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field60))
	}
	if m.Field67 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field67))
	}
	n += 2 + sovBenchmark(uint64(m.Field68))
	n += 3
	if m.Field80 != nil {
		n += 3
	}
	if m.Field81 != nil {
		n += 3
	}
	if m.Field100 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field100))
	}
	if m.Field101 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field101))
	}
	l = len(m.Field102)
	n += 2 + l + sovBenchmark(uint64(l))
	l = len(m.Field103)
	n += 2 + l + sovBenchmark(uint64(l))
	if m.Field104 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field104))
	}
	if m.Field128 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field128))
	}
	if m.Field129 != nil {
		l = len(*m.Field129)
		n += 2 + l + sovBenchmark(uint64(l))
	}
	if m.Field130 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field130))
	}
	if m.Field131 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field131))
	}
	n += 2 + sovBenchmark(uint64(m.Field150))
	if m.Field271 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field271))
	}
	if m.Field272 != nil {
		n += 2 + sovBenchmark(uint64(*m.Field272))
	}
	n += 2 + sovBenchmark(uint64(m.Field280))
	return n
}

func sovBenchmark(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBenchmark(x uint64) (n int) {
	return sovBenchmark(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BenchmarkMessage) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBenchmark
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BenchmarkMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BenchmarkMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBenchmark
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field2", wireType)
			}
			m.Field2 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Field2 |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field3", wireType)
			}
			m.Field3 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Field3 |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field4", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBenchmark
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field4 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				m.Field5 = append(m.Field5, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBenchmark
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBenchmark
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					m.Field5 = append(m.Field5, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Field5", wireType)
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field6", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field6 = &v
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field7", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBenchmark
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field7 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field9", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBenchmark
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field9 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field12", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Field12 = &b
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field13", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Field13 = &b
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field14", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Field14 = &b
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field16", wireType)
			}
			m.Field16 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Field16 |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field17", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Field17 = &b
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field18", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBenchmark
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field18 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field22", wireType)
			}
			m.Field22 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Field22 |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 23:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field23", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field23 = &v
		case 24:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field24", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Field24 = &b
		case 25:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field25", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field25 = &v
		case 29:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field29", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field29 = &v
		case 30:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field30", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Field30 = &b
		case 59:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field59", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Field59 = &b
		case 60:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field60", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field60 = &v
		case 67:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field67", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field67 = &v
		case 68:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field68", wireType)
			}
			m.Field68 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Field68 |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 78:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field78", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field78 = bool(v != 0)
		case 80:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field80", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Field80 = &b
		case 81:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field81", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Field81 = &b
		case 100:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field100", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field100 = &v
		case 101:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field101", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field101 = &v
		case 102:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field102", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBenchmark
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field102 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 103:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field103", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBenchmark
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field103 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 104:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field104", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field104 = &v
		case 128:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field128", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field128 = &v
		case 129:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field129", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBenchmark
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Field129 = &s
			iNdEx = postIndex
		case 130:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field130", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field130 = &v
		case 131:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field131", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field131 = &v
		case 150:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field150", wireType)
			}
			m.Field150 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Field150 |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 271:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field271", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field271 = &v
		case 272:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field272", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Field272 = &v
		case 280:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field280", wireType)
			}
			m.Field280 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Field280 |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBenchmark(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBenchmark
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return proto1.NewRequiredNotSetError("field1")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return proto1.NewRequiredNotSetError("field2")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return proto1.NewRequiredNotSetError("field3")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBenchmark(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBenchmark
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBenchmark
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBenchmark
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBenchmark
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBenchmark(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBenchmark = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBenchmark   = fmt.Errorf("proto: integer overflow")
)

func init() { proto1.RegisterFile("benchmark.proto", fileDescriptorBenchmark) }

var fileDescriptorBenchmark = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcb, 0x6e, 0xd3, 0x40,
	0x18, 0x85, 0x3b, 0x4e, 0x9c, 0x36, 0x16, 0x97, 0x32, 0x12, 0xe5, 0x54, 0xb4, 0x8e, 0x61, 0xe5,
	0x0d, 0x68, 0x3c, 0xe3, 0xcb, 0x38, 0xb0, 0x0a, 0x2c, 0xd8, 0x80, 0x20, 0x6f, 0x10, 0xa8, 0xd3,
	0x22, 0x0a, 0x45, 0x49, 0x91, 0x58, 0x72, 0xd9, 0x22, 0xd1, 0x15, 0xe2, 0x91, 0xba, 0xe4, 0x09,
	0x10, 0x0a, 0x2f, 0x52, 0x25, 0xf5, 0xfc, 0xe3, 0x78, 0x9a, 0x4d, 0xa2, 0xf3, 0x9d, 0x39, 0xe7,
	0xcf, 0xfc, 0x13, 0xdc, 0x7c, 0x5d, 0x7d, 0x78, 0x73, 0xf4, 0x7e, 0x32, 0x7b, 0xf7, 0xf0, 0xe3,
	0xec, 0xe4, 0xf4, 0x84, 0xfb, 0xab, 0xaf, 0xfb, 0x3f, 0x82, 0x60, 0x7b, 0x64, 0xd0, 0xf3, 0x6a,
	0x3e, 0x9f, 0x1c, 0x56, 0x7c, 0x2f, 0xe8, 0x4d, 0xdf, 0x56, 0xc7, 0x07, 0x09, 0x58, 0xe4, 0xc5,
	0xfd, 0x51, 0xf7, 0xfc, 0xef, 0x60, 0x63, 0x5c, 0x6b, 0x44, 0x25, 0xbc, 0xc8, 0x8b, 0xfd, 0x35,
	0x2a, 0x89, 0x2a, 0x74, 0x1c, 0xaa, 0x88, 0xa6, 0xe8, 0x46, 0xac, 0x95, 0x9c, 0xf2, 0x9d, 0x9a,
	0x66, 0xf0, 0xa3, 0x4e, 0xdc, 0xab, 0xf5, 0x8c, 0xef, 0xd6, 0x7a, 0x8e, 0x5e, 0xc4, 0x62, 0x7f,
	0xc8, 0x44, 0x8d, 0x72, 0x0a, 0x2c, 0xb0, 0xe9, 0x04, 0x16, 0x44, 0x4b, 0xf4, 0x1d, 0x5a, 0xf2,
	0x30, 0xd8, 0xbc, 0xfc, 0x4b, 0x12, 0xd7, 0x22, 0x16, 0x6f, 0x0d, 0xbb, 0xa7, 0xb3, 0x4f, 0xd5,
	0xd8, 0x88, 0x96, 0x2b, 0x5c, 0x77, 0xb9, 0xb2, 0x3c, 0xc5, 0x0d, 0x97, 0xa7, 0x96, 0xe7, 0xd8,
	0x5e, 0xce, 0x5d, 0xd7, 0x1b, 0xd1, 0xf2, 0x02, 0xb7, 0xdc, 0xf3, 0x85, 0xe5, 0x1a, 0xbc, 0x31,
	0xbe, 0x11, 0x89, 0x4b, 0x89, 0x9d, 0x88, 0xc5, 0x9d, 0x35, 0x2e, 0x25, 0xbf, 0x6b, 0xb8, 0xc2,
	0x1d, 0x73, 0x6f, 0x46, 0xe1, 0x03, 0x03, 0x53, 0x60, 0x55, 0xee, 0x4f, 0x27, 0xc7, 0x73, 0xd3,
	0x2e, 0x53, 0x7b, 0x3a, 0xc3, 0x6e, 0xeb, 0x74, 0x66, 0x61, 0x89, 0xfd, 0x16, 0x2c, 0x29, 0x5a,
	0x09, 0x84, 0x6e, 0xb4, 0x12, 0x64, 0xc8, 0x4a, 0x3c, 0x72, 0x0d, 0x59, 0xc9, 0xf7, 0x6a, 0x43,
	0x2e, 0xf0, 0x78, 0x15, 0xef, 0x3d, 0x48, 0xc6, 0x46, 0xa2, 0xf2, 0xbc, 0xc0, 0x93, 0xf5, 0xf2,
	0xdc, 0x5e, 0x5a, 0xae, 0xf1, 0xd4, 0xb9, 0xf4, 0xdc, 0x5e, 0x5a, 0xa1, 0xf1, 0x62, 0xd9, 0xbd,
	0xc6, 0x0b, 0x4d, 0xb3, 0x69, 0x81, 0x97, 0xee, 0x6c, 0x5a, 0x50, 0x80, 0x4e, 0xf0, 0xca, 0xd9,
	0x9a, 0x4e, 0xf8, 0x7e, 0xb0, 0x75, 0xb9, 0x20, 0x21, 0x70, 0x60, 0xc6, 0x23, 0xa9, 0x81, 0x13,
	0x54, 0x6d, 0x9c, 0xf0, 0x88, 0xb0, 0xc4, 0xb4, 0xb1, 0x74, 0x52, 0x1b, 0x0e, 0x85, 0xc3, 0x2b,
	0x1c, 0xaa, 0x51, 0x91, 0xe2, 0xa8, 0x5d, 0xb1, 0x7c, 0x96, 0xf5, 0x6f, 0xa9, 0xf1, 0x85, 0xb5,
	0xb8, 0xd4, 0x5c, 0x12, 0x2f, 0xf1, 0x75, 0xc9, 0xfb, 0xc3, 0xdb, 0x9f, 0xaf, 0xfa, 0xd0, 0x99,
	0xd2, 0x66, 0x2a, 0x81, 0x6f, 0xed, 0x4c, 0x25, 0x1a, 0x3c, 0xc1, 0x77, 0x87, 0x27, 0xfc, 0x9e,
	0xe1, 0x99, 0xc0, 0x2f, 0xd6, 0xd8, 0x1b, 0xc9, 0x7c, 0x50, 0x5b, 0x64, 0x91, 0xe0, 0xa7, 0x47,
	0xaf, 0x82, 0xc4, 0x86, 0x41, 0xe2, 0xcc, 0x35, 0x48, 0x2a, 0x91, 0x5a, 0xe0, 0xb7, 0xe7, 0x94,
	0x48, 0x2d, 0x46, 0xfc, 0x7c, 0x11, 0xb2, 0x3f, 0x8b, 0x90, 0xfd, 0x5b, 0x84, 0xec, 0xec, 0x7f,
	0xb8, 0xf1, 0x8c, 0x5d, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x48, 0xff, 0x6c, 0x3b, 0x05, 0x00,
	0x00,
}
