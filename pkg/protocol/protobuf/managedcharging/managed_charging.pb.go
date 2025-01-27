// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: managed_charging.proto

package managedcharging

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The reasons why the site controller may recommend no charge.
// The site controller will only return the highest priority reason to the vehicle.
type ChargeOnSolarNoChargeReason int32

const (
	ChargeOnSolarNoChargeReason_CHARGE_ON_SOLAR_NO_CHARGE_REASON_INVALID ChargeOnSolarNoChargeReason = 0
	// The Powerwall is being prioritized over the vehicle to charge.
	ChargeOnSolarNoChargeReason_CHARGE_ON_SOLAR_NO_CHARGE_REASON_POWERWALL_CHARGE_PRIORITY ChargeOnSolarNoChargeReason = 1
	// There is not enough solar for the vehicle to charge effectively.
	ChargeOnSolarNoChargeReason_CHARGE_ON_SOLAR_NO_CHARGE_REASON_INSUFFICIENT_SOLAR ChargeOnSolarNoChargeReason = 2
	// The site controller is prioritizing export to the grid. This can
	// happen when the site controller is in autonomous mode and it is
	// most economical to export excess solar to the grid, or during a
	// virtual power plant event.
	ChargeOnSolarNoChargeReason_CHARGE_ON_SOLAR_NO_CHARGE_REASON_GRID_EXPORT_PRIORITY ChargeOnSolarNoChargeReason = 3
	// Another vehicle is charging on solar at this location and has priority.
	ChargeOnSolarNoChargeReason_CHARGE_ON_SOLAR_NO_CHARGE_REASON_ALTERNATE_VEHICLE_CHARGE_PRIORITY ChargeOnSolarNoChargeReason = 4
)

// Enum value maps for ChargeOnSolarNoChargeReason.
var (
	ChargeOnSolarNoChargeReason_name = map[int32]string{
		0: "CHARGE_ON_SOLAR_NO_CHARGE_REASON_INVALID",
		1: "CHARGE_ON_SOLAR_NO_CHARGE_REASON_POWERWALL_CHARGE_PRIORITY",
		2: "CHARGE_ON_SOLAR_NO_CHARGE_REASON_INSUFFICIENT_SOLAR",
		3: "CHARGE_ON_SOLAR_NO_CHARGE_REASON_GRID_EXPORT_PRIORITY",
		4: "CHARGE_ON_SOLAR_NO_CHARGE_REASON_ALTERNATE_VEHICLE_CHARGE_PRIORITY",
	}
	ChargeOnSolarNoChargeReason_value = map[string]int32{
		"CHARGE_ON_SOLAR_NO_CHARGE_REASON_INVALID":                           0,
		"CHARGE_ON_SOLAR_NO_CHARGE_REASON_POWERWALL_CHARGE_PRIORITY":         1,
		"CHARGE_ON_SOLAR_NO_CHARGE_REASON_INSUFFICIENT_SOLAR":                2,
		"CHARGE_ON_SOLAR_NO_CHARGE_REASON_GRID_EXPORT_PRIORITY":              3,
		"CHARGE_ON_SOLAR_NO_CHARGE_REASON_ALTERNATE_VEHICLE_CHARGE_PRIORITY": 4,
	}
)

func (x ChargeOnSolarNoChargeReason) Enum() *ChargeOnSolarNoChargeReason {
	p := new(ChargeOnSolarNoChargeReason)
	*p = x
	return p
}

func (x ChargeOnSolarNoChargeReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChargeOnSolarNoChargeReason) Descriptor() protoreflect.EnumDescriptor {
	return file_managed_charging_proto_enumTypes[0].Descriptor()
}

func (ChargeOnSolarNoChargeReason) Type() protoreflect.EnumType {
	return &file_managed_charging_proto_enumTypes[0]
}

func (x ChargeOnSolarNoChargeReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChargeOnSolarNoChargeReason.Descriptor instead.
func (ChargeOnSolarNoChargeReason) EnumDescriptor() ([]byte, []int) {
	return file_managed_charging_proto_rawDescGZIP(), []int{0}
}

var File_managed_charging_proto protoreflect.FileDescriptor

var file_managed_charging_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x64, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x2a, 0xc7, 0x02, 0x0a, 0x1b, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x4f, 0x6e, 0x53, 0x6f, 0x6c, 0x61, 0x72, 0x4e, 0x6f, 0x43, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x28, 0x43, 0x48, 0x41,
	0x52, 0x47, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x53, 0x4f, 0x4c, 0x41, 0x52, 0x5f, 0x4e, 0x4f, 0x5f,
	0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x3e, 0x0a, 0x3a, 0x43, 0x48, 0x41, 0x52, 0x47,
	0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x53, 0x4f, 0x4c, 0x41, 0x52, 0x5f, 0x4e, 0x4f, 0x5f, 0x43, 0x48,
	0x41, 0x52, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x57, 0x45,
	0x52, 0x57, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x50, 0x52, 0x49,
	0x4f, 0x52, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12, 0x37, 0x0a, 0x33, 0x43, 0x48, 0x41, 0x52, 0x47,
	0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x53, 0x4f, 0x4c, 0x41, 0x52, 0x5f, 0x4e, 0x4f, 0x5f, 0x43, 0x48,
	0x41, 0x52, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x53, 0x55,
	0x46, 0x46, 0x49, 0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x4f, 0x4c, 0x41, 0x52, 0x10, 0x02,
	0x12, 0x39, 0x0a, 0x35, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x53, 0x4f,
	0x4c, 0x41, 0x52, 0x5f, 0x4e, 0x4f, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x49, 0x44, 0x5f, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54,
	0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x10, 0x03, 0x12, 0x46, 0x0a, 0x42, 0x43,
	0x48, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x53, 0x4f, 0x4c, 0x41, 0x52, 0x5f, 0x4e,
	0x4f, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x41, 0x4c, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x54, 0x45, 0x5f, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c,
	0x45, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54,
	0x59, 0x10, 0x04, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x65, 0x73, 0x6c, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x63, 0x68, 0x61, 0x72, 0x67,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managed_charging_proto_rawDescOnce sync.Once
	file_managed_charging_proto_rawDescData = file_managed_charging_proto_rawDesc
)

func file_managed_charging_proto_rawDescGZIP() []byte {
	file_managed_charging_proto_rawDescOnce.Do(func() {
		file_managed_charging_proto_rawDescData = protoimpl.X.CompressGZIP(file_managed_charging_proto_rawDescData)
	})
	return file_managed_charging_proto_rawDescData
}

var file_managed_charging_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_managed_charging_proto_goTypes = []interface{}{
	(ChargeOnSolarNoChargeReason)(0), // 0: ManagedCharging.ChargeOnSolarNoChargeReason
}
var file_managed_charging_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_managed_charging_proto_init() }
func file_managed_charging_proto_init() {
	if File_managed_charging_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managed_charging_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_managed_charging_proto_goTypes,
		DependencyIndexes: file_managed_charging_proto_depIdxs,
		EnumInfos:         file_managed_charging_proto_enumTypes,
	}.Build()
	File_managed_charging_proto = out.File
	file_managed_charging_proto_rawDesc = nil
	file_managed_charging_proto_goTypes = nil
	file_managed_charging_proto_depIdxs = nil
}
