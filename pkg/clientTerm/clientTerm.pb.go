// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: clientTerm.proto

package clientTerm

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

type AddNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	NoteText string `protobuf:"bytes,2,opt,name=noteText,proto3" json:"noteText,omitempty"`
}

func (x *AddNoteRequest) Reset() {
	*x = AddNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientTerm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoteRequest) ProtoMessage() {}

func (x *AddNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clientTerm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNoteRequest.ProtoReflect.Descriptor instead.
func (*AddNoteRequest) Descriptor() ([]byte, []int) {
	return file_clientTerm_proto_rawDescGZIP(), []int{0}
}

func (x *AddNoteRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddNoteRequest) GetNoteText() string {
	if x != nil {
		return x.NoteText
	}
	return ""
}

type NoteIDReplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	NoteID   string `protobuf:"bytes,2,opt,name=noteID,proto3" json:"noteID,omitempty"`
}

func (x *NoteIDReplay) Reset() {
	*x = NoteIDReplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientTerm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteIDReplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteIDReplay) ProtoMessage() {}

func (x *NoteIDReplay) ProtoReflect() protoreflect.Message {
	mi := &file_clientTerm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteIDReplay.ProtoReflect.Descriptor instead.
func (*NoteIDReplay) Descriptor() ([]byte, []int) {
	return file_clientTerm_proto_rawDescGZIP(), []int{1}
}

func (x *NoteIDReplay) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NoteIDReplay) GetNoteID() string {
	if x != nil {
		return x.NoteID
	}
	return ""
}

type GetNotesByUsername struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetNotesByUsername) Reset() {
	*x = GetNotesByUsername{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientTerm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotesByUsername) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotesByUsername) ProtoMessage() {}

func (x *GetNotesByUsername) ProtoReflect() protoreflect.Message {
	mi := &file_clientTerm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotesByUsername.ProtoReflect.Descriptor instead.
func (*GetNotesByUsername) Descriptor() ([]byte, []int) {
	return file_clientTerm_proto_rawDescGZIP(), []int{2}
}

func (x *GetNotesByUsername) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetNoteByID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	NoteID   string `protobuf:"bytes,2,opt,name=noteID,proto3" json:"noteID,omitempty"`
}

func (x *GetNoteByID) Reset() {
	*x = GetNoteByID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientTerm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteByID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteByID) ProtoMessage() {}

func (x *GetNoteByID) ProtoReflect() protoreflect.Message {
	mi := &file_clientTerm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteByID.ProtoReflect.Descriptor instead.
func (*GetNoteByID) Descriptor() ([]byte, []int) {
	return file_clientTerm_proto_rawDescGZIP(), []int{3}
}

func (x *GetNoteByID) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetNoteByID) GetNoteID() string {
	if x != nil {
		return x.NoteID
	}
	return ""
}

type RequestedNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteID   string `protobuf:"bytes,1,opt,name=noteID,proto3" json:"noteID,omitempty"`
	NoteText string `protobuf:"bytes,2,opt,name=noteText,proto3" json:"noteText,omitempty"`
}

func (x *RequestedNote) Reset() {
	*x = RequestedNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientTerm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestedNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestedNote) ProtoMessage() {}

func (x *RequestedNote) ProtoReflect() protoreflect.Message {
	mi := &file_clientTerm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestedNote.ProtoReflect.Descriptor instead.
func (*RequestedNote) Descriptor() ([]byte, []int) {
	return file_clientTerm_proto_rawDescGZIP(), []int{4}
}

func (x *RequestedNote) GetNoteID() string {
	if x != nil {
		return x.NoteID
	}
	return ""
}

func (x *RequestedNote) GetNoteText() string {
	if x != nil {
		return x.NoteText
	}
	return ""
}

type RequestedNotes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteID   []string `protobuf:"bytes,1,rep,name=noteID,proto3" json:"noteID,omitempty"`
	NoteText []string `protobuf:"bytes,2,rep,name=noteText,proto3" json:"noteText,omitempty"`
}

func (x *RequestedNotes) Reset() {
	*x = RequestedNotes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientTerm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestedNotes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestedNotes) ProtoMessage() {}

func (x *RequestedNotes) ProtoReflect() protoreflect.Message {
	mi := &file_clientTerm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestedNotes.ProtoReflect.Descriptor instead.
func (*RequestedNotes) Descriptor() ([]byte, []int) {
	return file_clientTerm_proto_rawDescGZIP(), []int{5}
}

func (x *RequestedNotes) GetNoteID() []string {
	if x != nil {
		return x.NoteID
	}
	return nil
}

func (x *RequestedNotes) GetNoteText() []string {
	if x != nil {
		return x.NoteText
	}
	return nil
}

var File_clientTerm_proto protoreflect.FileDescriptor

var file_clientTerm_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x48,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x22, 0x42, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x65,
	0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x44, 0x22, 0x30, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x41,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74,
	0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49,
	0x44, 0x22, 0x43, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x6f,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f,
	0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f,
	0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x22, 0x44, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x32, 0xe1, 0x01, 0x0a,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x42, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x65,
	0x72, 0x6d, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12,
	0x4e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x3f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x44, 0x12, 0x17, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x19, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x65,
	0x72, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x65,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x6c, 0x70, 0x68, 0x6f, 0x6e, 0x6e, 0x73, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x54, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x41, 0x70, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x65, 0x72, 0x6d, 0x3b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clientTerm_proto_rawDescOnce sync.Once
	file_clientTerm_proto_rawDescData = file_clientTerm_proto_rawDesc
)

func file_clientTerm_proto_rawDescGZIP() []byte {
	file_clientTerm_proto_rawDescOnce.Do(func() {
		file_clientTerm_proto_rawDescData = protoimpl.X.CompressGZIP(file_clientTerm_proto_rawDescData)
	})
	return file_clientTerm_proto_rawDescData
}

var file_clientTerm_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_clientTerm_proto_goTypes = []interface{}{
	(*AddNoteRequest)(nil),     // 0: clientTerm.AddNoteRequest
	(*NoteIDReplay)(nil),       // 1: clientTerm.NoteIDReplay
	(*GetNotesByUsername)(nil), // 2: clientTerm.GetNotesByUsername
	(*GetNoteByID)(nil),        // 3: clientTerm.GetNoteByID
	(*RequestedNote)(nil),      // 4: clientTerm.RequestedNote
	(*RequestedNotes)(nil),     // 5: clientTerm.RequestedNotes
}
var file_clientTerm_proto_depIdxs = []int32{
	0, // 0: clientTerm.clientTerm.CreateNote:input_type -> clientTerm.AddNoteRequest
	2, // 1: clientTerm.clientTerm.GetNotesUsername:input_type -> clientTerm.GetNotesByUsername
	3, // 2: clientTerm.clientTerm.GetNoteID:input_type -> clientTerm.GetNoteByID
	1, // 3: clientTerm.clientTerm.CreateNote:output_type -> clientTerm.NoteIDReplay
	5, // 4: clientTerm.clientTerm.GetNotesUsername:output_type -> clientTerm.RequestedNotes
	4, // 5: clientTerm.clientTerm.GetNoteID:output_type -> clientTerm.RequestedNote
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_clientTerm_proto_init() }
func file_clientTerm_proto_init() {
	if File_clientTerm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clientTerm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNoteRequest); i {
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
		file_clientTerm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteIDReplay); i {
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
		file_clientTerm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotesByUsername); i {
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
		file_clientTerm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoteByID); i {
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
		file_clientTerm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestedNote); i {
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
		file_clientTerm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestedNotes); i {
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
			RawDescriptor: file_clientTerm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clientTerm_proto_goTypes,
		DependencyIndexes: file_clientTerm_proto_depIdxs,
		MessageInfos:      file_clientTerm_proto_msgTypes,
	}.Build()
	File_clientTerm_proto = out.File
	file_clientTerm_proto_rawDesc = nil
	file_clientTerm_proto_goTypes = nil
	file_clientTerm_proto_depIdxs = nil
}