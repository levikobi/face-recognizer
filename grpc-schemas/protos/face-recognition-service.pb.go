// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/face-recognition-service.proto

package face_recognition_service

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

type AddPersonResponse_Status int32

const (
	AddPersonResponse_SUCCESS AddPersonResponse_Status = 0
	AddPersonResponse_FAILURE AddPersonResponse_Status = 1
)

// Enum value maps for AddPersonResponse_Status.
var (
	AddPersonResponse_Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
	}
	AddPersonResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
	}
)

func (x AddPersonResponse_Status) Enum() *AddPersonResponse_Status {
	p := new(AddPersonResponse_Status)
	*p = x
	return p
}

func (x AddPersonResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddPersonResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_face_recognition_service_proto_enumTypes[0].Descriptor()
}

func (AddPersonResponse_Status) Type() protoreflect.EnumType {
	return &file_protos_face_recognition_service_proto_enumTypes[0]
}

func (x AddPersonResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddPersonResponse_Status.Descriptor instead.
func (AddPersonResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_protos_face_recognition_service_proto_rawDescGZIP(), []int{1, 0}
}

type AddPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *AddPersonRequest) Reset() {
	*x = AddPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_face_recognition_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPersonRequest) ProtoMessage() {}

func (x *AddPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_face_recognition_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPersonRequest.ProtoReflect.Descriptor instead.
func (*AddPersonRequest) Descriptor() ([]byte, []int) {
	return file_protos_face_recognition_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddPersonRequest) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type AddPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddPersonResponse) Reset() {
	*x = AddPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_face_recognition_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPersonResponse) ProtoMessage() {}

func (x *AddPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_face_recognition_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPersonResponse.ProtoReflect.Descriptor instead.
func (*AddPersonResponse) Descriptor() ([]byte, []int) {
	return file_protos_face_recognition_service_proto_rawDescGZIP(), []int{1}
}

type FindPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features []float32 `protobuf:"fixed32,1,rep,packed,name=features,proto3" json:"features,omitempty"`
}

func (x *FindPersonRequest) Reset() {
	*x = FindPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_face_recognition_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPersonRequest) ProtoMessage() {}

func (x *FindPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_face_recognition_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPersonRequest.ProtoReflect.Descriptor instead.
func (*FindPersonRequest) Descriptor() ([]byte, []int) {
	return file_protos_face_recognition_service_proto_rawDescGZIP(), []int{2}
}

func (x *FindPersonRequest) GetFeatures() []float32 {
	if x != nil {
		return x.Features
	}
	return nil
}

type FindPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
}

func (x *FindPersonResponse) Reset() {
	*x = FindPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_face_recognition_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPersonResponse) ProtoMessage() {}

func (x *FindPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_face_recognition_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPersonResponse.ProtoReflect.Descriptor instead.
func (*FindPersonResponse) Descriptor() ([]byte, []int) {
	return file_protos_face_recognition_service_proto_rawDescGZIP(), []int{3}
}

func (x *FindPersonResponse) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features []float32 `protobuf:"fixed32,1,rep,packed,name=features,proto3" json:"features,omitempty"`
	Name     string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_face_recognition_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_protos_face_recognition_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_protos_face_recognition_service_proto_rawDescGZIP(), []int{4}
}

func (x *Person) GetFeatures() []float32 {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_protos_face_recognition_service_proto protoreflect.FileDescriptor

var file_protos_face_recognition_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x72, 0x65,
	0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22,
	0x3a, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x22, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55,
	0x52, 0x45, 0x10, 0x01, 0x22, 0x2f, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x08, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x38, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x02, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32,
	0x9c, 0x01, 0x0a, 0x0f, 0x46, 0x61, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x52,
	0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x76,
	0x69, 0x6b, 0x6f, 0x62, 0x69, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x72, 0x65, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x72,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_face_recognition_service_proto_rawDescOnce sync.Once
	file_protos_face_recognition_service_proto_rawDescData = file_protos_face_recognition_service_proto_rawDesc
)

func file_protos_face_recognition_service_proto_rawDescGZIP() []byte {
	file_protos_face_recognition_service_proto_rawDescOnce.Do(func() {
		file_protos_face_recognition_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_face_recognition_service_proto_rawDescData)
	})
	return file_protos_face_recognition_service_proto_rawDescData
}

var file_protos_face_recognition_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_face_recognition_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_face_recognition_service_proto_goTypes = []interface{}{
	(AddPersonResponse_Status)(0), // 0: protos.AddPersonResponse.Status
	(*AddPersonRequest)(nil),      // 1: protos.AddPersonRequest
	(*AddPersonResponse)(nil),     // 2: protos.AddPersonResponse
	(*FindPersonRequest)(nil),     // 3: protos.FindPersonRequest
	(*FindPersonResponse)(nil),    // 4: protos.FindPersonResponse
	(*Person)(nil),                // 5: protos.Person
}
var file_protos_face_recognition_service_proto_depIdxs = []int32{
	5, // 0: protos.AddPersonRequest.person:type_name -> protos.Person
	5, // 1: protos.FindPersonResponse.persons:type_name -> protos.Person
	1, // 2: protos.FaceRecognition.AddPerson:input_type -> protos.AddPersonRequest
	3, // 3: protos.FaceRecognition.FindPerson:input_type -> protos.FindPersonRequest
	2, // 4: protos.FaceRecognition.AddPerson:output_type -> protos.AddPersonResponse
	4, // 5: protos.FaceRecognition.FindPerson:output_type -> protos.FindPersonResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_face_recognition_service_proto_init() }
func file_protos_face_recognition_service_proto_init() {
	if File_protos_face_recognition_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_face_recognition_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPersonRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_face_recognition_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPersonResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_face_recognition_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPersonRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_face_recognition_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPersonResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_face_recognition_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_face_recognition_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_face_recognition_service_proto_goTypes,
		DependencyIndexes: file_protos_face_recognition_service_proto_depIdxs,
		EnumInfos:         file_protos_face_recognition_service_proto_enumTypes,
		MessageInfos:      file_protos_face_recognition_service_proto_msgTypes,
	}.Build()
	File_protos_face_recognition_service_proto = out.File
	file_protos_face_recognition_service_proto_rawDesc = nil
	file_protos_face_recognition_service_proto_goTypes = nil
	file_protos_face_recognition_service_proto_depIdxs = nil
}
