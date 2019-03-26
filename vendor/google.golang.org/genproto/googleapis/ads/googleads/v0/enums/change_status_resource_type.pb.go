// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/change_status_resource_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum listing the resource types support by the ChangeStatus resource.
type ChangeStatusResourceTypeEnum_ChangeStatusResourceType int32

const (
	// No value has been specified.
	ChangeStatusResourceTypeEnum_UNSPECIFIED ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 0
	// Used for return value only. Represents an unclassified resource unknown
	// in this version.
	ChangeStatusResourceTypeEnum_UNKNOWN ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 1
	// An AdGroup resource change.
	ChangeStatusResourceTypeEnum_AD_GROUP ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 3
	// An AdGroupAd resource change.
	ChangeStatusResourceTypeEnum_AD_GROUP_AD ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 4
	// An AdGroupCriterion resource change.
	ChangeStatusResourceTypeEnum_AD_GROUP_CRITERION ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 5
	// A Campaign resource change.
	ChangeStatusResourceTypeEnum_CAMPAIGN ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 6
	// A CampaignCriterion resource change.
	ChangeStatusResourceTypeEnum_CAMPAIGN_CRITERION ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 7
	// A Feed resource change.
	ChangeStatusResourceTypeEnum_FEED ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 9
	// A FeedItem resource change.
	ChangeStatusResourceTypeEnum_FEED_ITEM ChangeStatusResourceTypeEnum_ChangeStatusResourceType = 10
)

var ChangeStatusResourceTypeEnum_ChangeStatusResourceType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "AD_GROUP",
	4:  "AD_GROUP_AD",
	5:  "AD_GROUP_CRITERION",
	6:  "CAMPAIGN",
	7:  "CAMPAIGN_CRITERION",
	9:  "FEED",
	10: "FEED_ITEM",
}
var ChangeStatusResourceTypeEnum_ChangeStatusResourceType_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"AD_GROUP":           3,
	"AD_GROUP_AD":        4,
	"AD_GROUP_CRITERION": 5,
	"CAMPAIGN":           6,
	"CAMPAIGN_CRITERION": 7,
	"FEED":               9,
	"FEED_ITEM":          10,
}

func (x ChangeStatusResourceTypeEnum_ChangeStatusResourceType) String() string {
	return proto.EnumName(ChangeStatusResourceTypeEnum_ChangeStatusResourceType_name, int32(x))
}
func (ChangeStatusResourceTypeEnum_ChangeStatusResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_change_status_resource_type_b24c7a466a3225d8, []int{0, 0}
}

// Container for enum describing supported resource types for the ChangeStatus
// resource.
type ChangeStatusResourceTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeStatusResourceTypeEnum) Reset()         { *m = ChangeStatusResourceTypeEnum{} }
func (m *ChangeStatusResourceTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusResourceTypeEnum) ProtoMessage()    {}
func (*ChangeStatusResourceTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_change_status_resource_type_b24c7a466a3225d8, []int{0}
}
func (m *ChangeStatusResourceTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusResourceTypeEnum.Unmarshal(m, b)
}
func (m *ChangeStatusResourceTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusResourceTypeEnum.Marshal(b, m, deterministic)
}
func (dst *ChangeStatusResourceTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusResourceTypeEnum.Merge(dst, src)
}
func (m *ChangeStatusResourceTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusResourceTypeEnum.Size(m)
}
func (m *ChangeStatusResourceTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusResourceTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusResourceTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChangeStatusResourceTypeEnum)(nil), "google.ads.googleads.v0.enums.ChangeStatusResourceTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType", ChangeStatusResourceTypeEnum_ChangeStatusResourceType_name, ChangeStatusResourceTypeEnum_ChangeStatusResourceType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/change_status_resource_type.proto", fileDescriptor_change_status_resource_type_b24c7a466a3225d8)
}

var fileDescriptor_change_status_resource_type_b24c7a466a3225d8 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xc7, 0x2b, 0xdb, 0xf5, 0xc7, 0xba, 0xa5, 0xcb, 0x1e, 0x4a, 0x0f, 0xf5, 0xc1, 0x7e, 0x80,
	0x95, 0xa0, 0xb7, 0xed, 0xa1, 0xac, 0x25, 0x59, 0x88, 0x62, 0x59, 0xc8, 0x1f, 0x85, 0x22, 0x10,
	0x8a, 0xb5, 0x28, 0x01, 0x5b, 0x2b, 0xb4, 0x92, 0xc1, 0xaf, 0x93, 0x63, 0x0e, 0x79, 0x90, 0x5c,
	0xf2, 0x1e, 0x39, 0xe6, 0x09, 0xc2, 0xae, 0x2c, 0x91, 0x8b, 0x73, 0x59, 0xfe, 0x33, 0xf3, 0x9f,
	0x1f, 0x3b, 0x33, 0xe0, 0x4f, 0xca, 0x79, 0x7a, 0x60, 0x7a, 0x9c, 0x08, 0xbd, 0x96, 0x52, 0x9d,
	0x0c, 0x9d, 0x65, 0xd5, 0x51, 0xe8, 0xfb, 0xdb, 0x38, 0x4b, 0x59, 0x24, 0xca, 0xb8, 0xac, 0x44,
	0x54, 0x30, 0xc1, 0xab, 0x62, 0xcf, 0xa2, 0xf2, 0x9c, 0x33, 0x9c, 0x17, 0xbc, 0xe4, 0x68, 0x52,
	0x77, 0xe1, 0x38, 0x11, 0xb8, 0x05, 0xe0, 0x93, 0x81, 0x15, 0x60, 0xf6, 0xac, 0x81, 0x9f, 0xa6,
	0x82, 0xac, 0x15, 0x23, 0xb8, 0x20, 0x36, 0xe7, 0x9c, 0xd9, 0x59, 0x75, 0x9c, 0x3d, 0x6a, 0xe0,
	0xc7, 0x35, 0x03, 0xfa, 0x06, 0xc6, 0x5b, 0x6f, 0xed, 0xdb, 0xa6, 0xbb, 0x70, 0x6d, 0x0b, 0x7e,
	0x42, 0x63, 0x30, 0xd8, 0x7a, 0x7f, 0xbd, 0xd5, 0x3f, 0x0f, 0x6a, 0xe8, 0x0b, 0x18, 0x52, 0x2b,
	0x72, 0x82, 0xd5, 0xd6, 0x87, 0x5d, 0xe9, 0x6d, 0xa2, 0x88, 0x5a, 0xb0, 0x87, 0xbe, 0x03, 0xd4,
	0x26, 0xcc, 0xc0, 0xdd, 0xd8, 0x81, 0xbb, 0xf2, 0xe0, 0x67, 0xd9, 0x66, 0xd2, 0xa5, 0x4f, 0x5d,
	0xc7, 0x83, 0x7d, 0xe9, 0x6a, 0xa2, 0x77, 0xae, 0x01, 0x1a, 0x82, 0xde, 0xc2, 0xb6, 0x2d, 0x38,
	0x42, 0x5f, 0xc1, 0x48, 0xaa, 0xc8, 0xdd, 0xd8, 0x4b, 0x08, 0xe6, 0xaf, 0x1a, 0x98, 0xee, 0xf9,
	0x11, 0x7f, 0x38, 0xf7, 0x7c, 0x72, 0x6d, 0x26, 0x5f, 0x6e, 0xcd, 0xd7, 0xfe, 0xcf, 0x2f, 0xfd,
	0x29, 0x3f, 0xc4, 0x59, 0x8a, 0x79, 0x91, 0xea, 0x29, 0xcb, 0xd4, 0x4e, 0x9b, 0x43, 0xe4, 0x77,
	0xe2, 0xca, 0x5d, 0x7e, 0xab, 0xf7, 0xbe, 0xd3, 0x75, 0x28, 0x7d, 0xe8, 0x4c, 0x9c, 0x1a, 0x45,
	0x13, 0x81, 0x6b, 0x29, 0xd5, 0xce, 0xc0, 0x72, 0xc1, 0xe2, 0xa9, 0xa9, 0x87, 0x34, 0x11, 0x61,
	0x5b, 0x0f, 0x77, 0x46, 0xa8, 0xea, 0x2f, 0x9d, 0x69, 0x9d, 0x24, 0x84, 0x26, 0x82, 0x90, 0xd6,
	0x41, 0xc8, 0xce, 0x20, 0x44, 0x79, 0x6e, 0xfa, 0xea, 0x63, 0xbf, 0xde, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x08, 0x5b, 0x5b, 0xd0, 0x2f, 0x02, 0x00, 0x00,
}