// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: funcpb.proto

/*
	Package funcpb is a generated protocol buffer package.

	It is generated from these files:
		funcpb.proto

	It has these top-level messages:
*/
package funcpb

import (
	"fmt"
	"math"

	proto "github.com/golang/protobuf/proto"

	_ "github.com/gogo/protobuf/gogoproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FunctionID int32

const (
	FunctionID_kFuncHeartbeat           FunctionID = 0
	FunctionID_kFuncRawGet              FunctionID = 1
	FunctionID_kFuncRawPut              FunctionID = 2
	FunctionID_kFuncRawDelete           FunctionID = 3
	FunctionID_kFuncRawExecute          FunctionID = 4
	FunctionID_kFuncSelect              FunctionID = 10
	FunctionID_kFuncInsert              FunctionID = 11
	FunctionID_kFuncDelete              FunctionID = 12
	FunctionID_kFuncUpdate              FunctionID = 13
	FunctionID_KFuncReplace             FunctionID = 14
	FunctionID_kFuncWatchGet            FunctionID = 50
	FunctionID_kFuncPureGet             FunctionID = 51
	FunctionID_kFuncWatchPut            FunctionID = 52
	FunctionID_kFuncWatchDel            FunctionID = 53
	FunctionID_kFuncKvSet               FunctionID = 100
	FunctionID_kFuncKvGet               FunctionID = 101
	FunctionID_kFuncKvBatchSet          FunctionID = 102
	FunctionID_kFuncKvBatchGet          FunctionID = 103
	FunctionID_kFuncKvDel               FunctionID = 104
	FunctionID_kFuncKvBatchDel          FunctionID = 105
	FunctionID_kFuncKvRangeDel          FunctionID = 106
	FunctionID_kFuncKvScan              FunctionID = 107
	FunctionID_kFuncLock                FunctionID = 200
	FunctionID_kFuncLockUpdate          FunctionID = 201
	FunctionID_kFuncUnlock              FunctionID = 202
	FunctionID_kFuncUnlockForce         FunctionID = 203
	FunctionID_kFuncLockWatch           FunctionID = 204
	FunctionID_kFuncCreateRange         FunctionID = 1001
	FunctionID_kFuncDeleteRange         FunctionID = 1002
	FunctionID_kFuncRangeTransferLeader FunctionID = 1003
	FunctionID_kFuncUpdateRange         FunctionID = 1004
	FunctionID_kFuncGetPeerInfo         FunctionID = 1005
	FunctionID_kFuncSetNodeLogLevel     FunctionID = 1006
	FunctionID_kFuncOfflineRange        FunctionID = 1007
	FunctionID_kFuncReplaceRange        FunctionID = 1008
)

var FunctionID_name = map[int32]string{
	0:    "kFuncHeartbeat",
	1:    "kFuncRawGet",
	2:    "kFuncRawPut",
	3:    "kFuncRawDelete",
	4:    "kFuncRawExecute",
	10:   "kFuncSelect",
	11:   "kFuncInsert",
	12:   "kFuncDelete",
	13:   "kFuncUpdate",
	14:   "KFuncReplace",
	50:   "kFuncWatchGet",
	51:   "kFuncPureGet",
	52:   "kFuncWatchPut",
	53:   "kFuncWatchDel",
	100:  "kFuncKvSet",
	101:  "kFuncKvGet",
	102:  "kFuncKvBatchSet",
	103:  "kFuncKvBatchGet",
	104:  "kFuncKvDel",
	105:  "kFuncKvBatchDel",
	106:  "kFuncKvRangeDel",
	107:  "kFuncKvScan",
	200:  "kFuncLock",
	201:  "kFuncLockUpdate",
	202:  "kFuncUnlock",
	203:  "kFuncUnlockForce",
	204:  "kFuncLockWatch",
	1001: "kFuncCreateRange",
	1002: "kFuncDeleteRange",
	1003: "kFuncRangeTransferLeader",
	1004: "kFuncUpdateRange",
	1005: "kFuncGetPeerInfo",
	1006: "kFuncSetNodeLogLevel",
	1007: "kFuncOfflineRange",
	1008: "kFuncReplaceRange",
}
var FunctionID_value = map[string]int32{
	"kFuncHeartbeat":           0,
	"kFuncRawGet":              1,
	"kFuncRawPut":              2,
	"kFuncRawDelete":           3,
	"kFuncRawExecute":          4,
	"kFuncSelect":              10,
	"kFuncInsert":              11,
	"kFuncDelete":              12,
	"kFuncUpdate":              13,
	"KFuncReplace":             14,
	"kFuncWatchGet":            50,
	"kFuncPureGet":             51,
	"kFuncWatchPut":            52,
	"kFuncWatchDel":            53,
	"kFuncKvSet":               100,
	"kFuncKvGet":               101,
	"kFuncKvBatchSet":          102,
	"kFuncKvBatchGet":          103,
	"kFuncKvDel":               104,
	"kFuncKvBatchDel":          105,
	"kFuncKvRangeDel":          106,
	"kFuncKvScan":              107,
	"kFuncLock":                200,
	"kFuncLockUpdate":          201,
	"kFuncUnlock":              202,
	"kFuncUnlockForce":         203,
	"kFuncLockWatch":           204,
	"kFuncCreateRange":         1001,
	"kFuncDeleteRange":         1002,
	"kFuncRangeTransferLeader": 1003,
	"kFuncUpdateRange":         1004,
	"kFuncGetPeerInfo":         1005,
	"kFuncSetNodeLogLevel":     1006,
	"kFuncOfflineRange":        1007,
	"kFuncReplaceRange":        1008,
}

func (x FunctionID) String() string {
	return proto.EnumName(FunctionID_name, int32(x))
}
func (FunctionID) EnumDescriptor() ([]byte, []int) { return fileDescriptorFuncpb, []int{0} }

func init() {
	proto.RegisterEnum("funcpb.FunctionID", FunctionID_name, FunctionID_value)
}

func init() { proto.RegisterFile("funcpb.proto", fileDescriptorFuncpb) }

var fileDescriptorFuncpb = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xdd, 0x6e, 0xd3, 0x40,
	0x10, 0x85, 0x6b, 0x7e, 0x6a, 0xb1, 0x4d, 0xd3, 0xe9, 0x36, 0x20, 0x40, 0x22, 0x0f, 0xc0, 0x05,
	0x48, 0x14, 0x5e, 0xa0, 0x84, 0x96, 0x28, 0x11, 0x44, 0x09, 0x88, 0xeb, 0xcd, 0xe6, 0xd8, 0x0d,
	0xb1, 0x76, 0xa3, 0xed, 0x3a, 0xe5, 0x51, 0x78, 0xa4, 0x42, 0xb9, 0xe0, 0x11, 0x50, 0xb8, 0xe2,
	0x9f, 0x47, 0xa8, 0x66, 0xed, 0x3a, 0x6e, 0xee, 0x3c, 0x9f, 0xcf, 0x39, 0x9a, 0xd9, 0x19, 0xd1,
	0x48, 0x72, 0xa3, 0xe7, 0xe3, 0x47, 0x73, 0x67, 0xbd, 0x95, 0x9b, 0x45, 0x75, 0xbf, 0x95, 0xda,
	0xd4, 0x06, 0xf4, 0x98, 0xbf, 0x8a, 0xbf, 0x0f, 0xcf, 0x6f, 0x0a, 0x71, 0x98, 0x1b, 0xed, 0xa7,
	0xd6, 0x74, 0x3b, 0x52, 0x8a, 0xe6, 0x8c, 0xcb, 0x97, 0x50, 0xce, 0x8f, 0xa1, 0x3c, 0x6d, 0xc8,
	0x1d, 0xb1, 0x15, 0xd8, 0x50, 0x9d, 0x1e, 0xc1, 0x53, 0x54, 0x07, 0x83, 0xdc, 0xd3, 0xb5, 0xca,
	0x35, 0x54, 0xa7, 0x1d, 0x64, 0xf0, 0xa0, 0xeb, 0x72, 0x4f, 0xec, 0x5c, 0xb2, 0x17, 0x1f, 0xa0,
	0x73, 0x0f, 0xba, 0x51, 0x39, 0x47, 0xc8, 0xa0, 0x3d, 0x89, 0x0a, 0x74, 0xcd, 0x09, 0x9c, 0xa7,
	0xad, 0x0a, 0x94, 0x39, 0x8d, 0x0a, 0xbc, 0x9d, 0x4f, 0x94, 0x07, 0x6d, 0x4b, 0x12, 0x8d, 0x5e,
	0x08, 0xc6, 0x3c, 0x53, 0x1a, 0xd4, 0x94, 0xbb, 0x62, 0x3b, 0x48, 0xde, 0x29, 0xaf, 0x8f, 0xb9,
	0xc5, 0x27, 0x2c, 0x0a, 0x68, 0x90, 0x3b, 0x30, 0xd9, 0xbf, 0x2a, 0xe2, 0xb6, 0x9f, 0x5e, 0x45,
	0x1d, 0x64, 0xf4, 0x4c, 0x36, 0x85, 0x08, 0xa8, 0xb7, 0x18, 0xc1, 0xd3, 0xa4, 0x56, 0x73, 0x0a,
	0xaa, 0xa9, 0x7a, 0x8b, 0x03, 0x36, 0xb1, 0x28, 0x59, 0x87, 0xac, 0x4c, 0x6b, 0x4e, 0x4e, 0x3e,
	0x5e, 0x17, 0x31, 0x9c, 0xd6, 0xe0, 0x50, 0x99, 0x14, 0x0c, 0xdf, 0x57, 0x13, 0xf7, 0x16, 0x23,
	0xad, 0x0c, 0xcd, 0x64, 0x53, 0xdc, 0x0a, 0xa0, 0x6f, 0xf5, 0x8c, 0xce, 0x22, 0xd9, 0x2a, 0x5d,
	0x5c, 0x97, 0xcf, 0xf2, 0x29, 0x92, 0x74, 0xf9, 0x50, 0x26, 0x63, 0xdd, 0xe7, 0x48, 0xde, 0x16,
	0x54, 0x23, 0x87, 0xd6, 0x69, 0xd0, 0x79, 0x24, 0xf7, 0xca, 0x6d, 0xb1, 0x3d, 0x8c, 0x4e, 0x5f,
	0x56, 0xda, 0xe7, 0x0e, 0xca, 0x23, 0x74, 0x43, 0x3f, 0xe2, 0x0a, 0x17, 0xeb, 0x28, 0xf0, 0xcf,
	0x58, 0x3e, 0x10, 0x77, 0xcb, 0xe5, 0x9a, 0x14, 0x6f, 0x9c, 0x32, 0x27, 0x09, 0x5c, 0x1f, 0x6a,
	0x02, 0x47, 0xbf, 0x56, 0xae, 0xa2, 0xb9, 0xc2, 0xf5, 0x7b, 0x85, 0x8f, 0xe0, 0x07, 0x80, 0xeb,
	0x9a, 0xc4, 0xd2, 0x9f, 0x58, 0xde, 0x13, 0xad, 0xf2, 0x28, 0xfc, 0x2b, 0x3b, 0x41, 0xdf, 0xa6,
	0x7d, 0x2c, 0x90, 0xd1, 0xdf, 0x58, 0xde, 0x11, 0xbb, 0xe1, 0xd7, 0xeb, 0x24, 0xc9, 0xa6, 0xa6,
	0x4c, 0xfa, 0xb7, 0xe2, 0xe5, 0x0d, 0x14, 0xfc, 0x7f, 0x7c, 0x40, 0x67, 0xcb, 0x76, 0xf4, 0x75,
	0xd9, 0x8e, 0xbe, 0x2d, 0xdb, 0xd1, 0xc7, 0xef, 0xed, 0x8d, 0xf1, 0x66, 0x38, 0xf3, 0xfd, 0x8b,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xe9, 0xbe, 0xd1, 0x14, 0x03, 0x00, 0x00,
}
