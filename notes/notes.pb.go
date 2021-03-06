// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notes.proto

package notes

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Note struct {
	OwnerId              string   `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Directory            string   `protobuf:"bytes,2,opt,name=directory,proto3" json:"directory,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Labels               []string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{0}
}

func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (m *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(m, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *Note) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *Note) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Note) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Note) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type FileNote struct {
	Labels               []string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileNote) Reset()         { *m = FileNote{} }
func (m *FileNote) String() string { return proto.CompactTextString(m) }
func (*FileNote) ProtoMessage()    {}
func (*FileNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{1}
}

func (m *FileNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileNote.Unmarshal(m, b)
}
func (m *FileNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileNote.Marshal(b, m, deterministic)
}
func (m *FileNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileNote.Merge(m, src)
}
func (m *FileNote) XXX_Size() int {
	return xxx_messageInfo_FileNote.Size(m)
}
func (m *FileNote) XXX_DiscardUnknown() {
	xxx_messageInfo_FileNote.DiscardUnknown(m)
}

var xxx_messageInfo_FileNote proto.InternalMessageInfo

func (m *FileNote) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *FileNote) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type DeleteRequest struct {
	Directory            string   `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{2}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateRequest struct {
	Directory            string   `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{3}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *CreateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateRequest struct {
	Directory            string   `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Labels               []string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{4}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *UpdateRequest) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type GetNoteRequest struct {
	Directory            string   `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNoteRequest) Reset()         { *m = GetNoteRequest{} }
func (m *GetNoteRequest) String() string { return proto.CompactTextString(m) }
func (*GetNoteRequest) ProtoMessage()    {}
func (*GetNoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{5}
}

func (m *GetNoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNoteRequest.Unmarshal(m, b)
}
func (m *GetNoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNoteRequest.Marshal(b, m, deterministic)
}
func (m *GetNoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNoteRequest.Merge(m, src)
}
func (m *GetNoteRequest) XXX_Size() int {
	return xxx_messageInfo_GetNoteRequest.Size(m)
}
func (m *GetNoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNoteRequest proto.InternalMessageInfo

func (m *GetNoteRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *GetNoteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDirectoryNotesRequest struct {
	Directory            string   `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDirectoryNotesRequest) Reset()         { *m = GetDirectoryNotesRequest{} }
func (m *GetDirectoryNotesRequest) String() string { return proto.CompactTextString(m) }
func (*GetDirectoryNotesRequest) ProtoMessage()    {}
func (*GetDirectoryNotesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{6}
}

func (m *GetDirectoryNotesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDirectoryNotesRequest.Unmarshal(m, b)
}
func (m *GetDirectoryNotesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDirectoryNotesRequest.Marshal(b, m, deterministic)
}
func (m *GetDirectoryNotesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDirectoryNotesRequest.Merge(m, src)
}
func (m *GetDirectoryNotesRequest) XXX_Size() int {
	return xxx_messageInfo_GetDirectoryNotesRequest.Size(m)
}
func (m *GetDirectoryNotesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDirectoryNotesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDirectoryNotesRequest proto.InternalMessageInfo

func (m *GetDirectoryNotesRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

type GetDirectoryNotesResponse struct {
	Notes                []string `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDirectoryNotesResponse) Reset()         { *m = GetDirectoryNotesResponse{} }
func (m *GetDirectoryNotesResponse) String() string { return proto.CompactTextString(m) }
func (*GetDirectoryNotesResponse) ProtoMessage()    {}
func (*GetDirectoryNotesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffefd935cd6c4a4a, []int{7}
}

func (m *GetDirectoryNotesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDirectoryNotesResponse.Unmarshal(m, b)
}
func (m *GetDirectoryNotesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDirectoryNotesResponse.Marshal(b, m, deterministic)
}
func (m *GetDirectoryNotesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDirectoryNotesResponse.Merge(m, src)
}
func (m *GetDirectoryNotesResponse) XXX_Size() int {
	return xxx_messageInfo_GetDirectoryNotesResponse.Size(m)
}
func (m *GetDirectoryNotesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDirectoryNotesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDirectoryNotesResponse proto.InternalMessageInfo

func (m *GetDirectoryNotesResponse) GetNotes() []string {
	if m != nil {
		return m.Notes
	}
	return nil
}

func init() {
	proto.RegisterType((*Note)(nil), "notes.Note")
	proto.RegisterType((*FileNote)(nil), "notes.FileNote")
	proto.RegisterType((*DeleteRequest)(nil), "notes.DeleteRequest")
	proto.RegisterType((*CreateRequest)(nil), "notes.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "notes.UpdateRequest")
	proto.RegisterType((*GetNoteRequest)(nil), "notes.GetNoteRequest")
	proto.RegisterType((*GetDirectoryNotesRequest)(nil), "notes.GetDirectoryNotesRequest")
	proto.RegisterType((*GetDirectoryNotesResponse)(nil), "notes.GetDirectoryNotesResponse")
}

func init() { proto.RegisterFile("notes.proto", fileDescriptor_ffefd935cd6c4a4a) }

var fileDescriptor_ffefd935cd6c4a4a = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x26, 0x69, 0xba, 0xdd, 0x9e, 0x65, 0x17, 0x1c, 0xd6, 0x25, 0x1b, 0x8b, 0x96, 0x41, 0xa4,
	0xac, 0x90, 0xb0, 0xeb, 0x8d, 0x88, 0x7a, 0x63, 0xb5, 0x78, 0xe3, 0x45, 0xc1, 0x4b, 0x91, 0xa4,
	0x39, 0x96, 0x91, 0x74, 0x26, 0x26, 0x53, 0xa4, 0xd4, 0x5e, 0xe8, 0x2b, 0xf8, 0x08, 0x3e, 0x92,
	0xaf, 0xe0, 0x83, 0x48, 0x66, 0xa6, 0x6d, 0x06, 0x1b, 0xe9, 0xf6, 0x2e, 0xe7, 0xef, 0xfb, 0xbe,
	0x39, 0xf9, 0x0e, 0x9c, 0x70, 0x21, 0xb1, 0x0c, 0xf3, 0x42, 0x48, 0x41, 0xda, 0x2a, 0x08, 0x7a,
	0x53, 0x21, 0xa6, 0x19, 0x46, 0x71, 0xce, 0xa2, 0x98, 0x73, 0x21, 0x63, 0xc9, 0x04, 0x37, 0x4d,
	0xc1, 0x3d, 0x53, 0x55, 0x51, 0x32, 0xff, 0x14, 0xe1, 0x2c, 0x97, 0x0b, 0x5d, 0xa4, 0xdf, 0x1d,
	0xf0, 0xde, 0x09, 0x89, 0xe4, 0x12, 0x8e, 0xc5, 0x57, 0x8e, 0xc5, 0x47, 0x96, 0xfa, 0x4e, 0xdf,
	0x19, 0x74, 0xc7, 0x1d, 0x15, 0xbf, 0x4d, 0x49, 0x0f, 0xba, 0x29, 0x2b, 0x70, 0x22, 0x45, 0xb1,
	0xf0, 0x5d, 0x55, 0xdb, 0x26, 0xc8, 0x19, 0xb8, 0x2c, 0xf5, 0x5b, 0x2a, 0xed, 0xb2, 0x94, 0xf8,
	0xd0, 0x99, 0x08, 0x2e, 0x91, 0x4b, 0xdf, 0xd3, 0x38, 0x26, 0x24, 0x17, 0x70, 0x94, 0xc5, 0x09,
	0x66, 0xa5, 0xdf, 0xee, 0xb7, 0x06, 0xdd, 0xb1, 0x89, 0xe8, 0x73, 0x38, 0x7e, 0xc3, 0x32, 0x54,
	0x32, 0xb6, 0x3d, 0x4e, 0xbd, 0xa7, 0x8e, 0xea, 0x5a, 0xa8, 0xf4, 0x05, 0x9c, 0x0e, 0x31, 0x43,
	0x89, 0x63, 0xfc, 0x32, 0xc7, 0x52, 0xda, 0x72, 0x9d, 0xdd, 0x72, 0xdd, 0xb5, 0xdc, 0x6a, 0xfc,
	0x55, 0x81, 0xf1, 0xa1, 0xe3, 0x02, 0x4e, 0xdf, 0xe7, 0xe9, 0xa1, 0xe3, 0xf5, 0x67, 0xb5, 0x9a,
	0x96, 0xe5, 0x59, 0xcb, 0x7a, 0x09, 0x67, 0x23, 0x94, 0xd5, 0xae, 0x0e, 0x13, 0xfc, 0x14, 0xfc,
	0x11, 0xca, 0xe1, 0xba, 0x5e, 0x01, 0x95, 0x7b, 0x21, 0xd1, 0x6b, 0xb8, 0xdc, 0x31, 0x59, 0xe6,
	0x82, 0x97, 0x48, 0xce, 0x41, 0x7b, 0xd1, 0xfc, 0x36, 0x1d, 0xdc, 0xfc, 0xf2, 0xa0, 0xad, 0xfa,
	0xc8, 0x37, 0xb8, 0xf3, 0xcf, 0x30, 0x79, 0x10, 0x6a, 0x33, 0x37, 0x09, 0x0a, 0xfa, 0xcd, 0x0d,
	0x9a, 0x97, 0x3e, 0xfc, 0xf1, 0xfb, 0xcf, 0x4f, 0xf7, 0x3e, 0xe9, 0x45, 0xaa, 0x33, 0xda, 0xc8,
	0x8d, 0x96, 0x9b, 0xcf, 0x15, 0xf9, 0x00, 0x1d, 0xb3, 0x34, 0x72, 0x77, 0x0b, 0x59, 0x5b, 0x62,
	0x70, 0x62, 0xd2, 0x55, 0x8e, 0x86, 0x0a, 0x74, 0x40, 0x1e, 0xfd, 0x0f, 0x54, 0xd5, 0xa2, 0x25,
	0x4b, 0x57, 0x64, 0x06, 0xa0, 0x3d, 0xa4, 0x18, 0xce, 0x0d, 0x94, 0x65, 0xab, 0xe0, 0x22, 0xd4,
	0x67, 0x18, 0xae, 0xcf, 0x30, 0x7c, 0x5d, 0x9d, 0x21, 0xbd, 0x56, 0x5c, 0x8f, 0xe9, 0x9e, 0x5c,
	0xcf, 0x9c, 0x2b, 0x92, 0x00, 0x68, 0xcf, 0x59, 0x74, 0x96, 0x0d, 0xed, 0xf7, 0x18, 0x8e, 0x9b,
	0x5b, 0x70, 0x7c, 0x06, 0xd0, 0x57, 0x65, 0x71, 0x58, 0x87, 0xd6, 0xf8, 0x24, 0xb3, 0xbe, 0xab,
	0x3d, 0xe9, 0x92, 0x23, 0x35, 0xff, 0xe4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0x82, 0x68,
	0x7e, 0xdb, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotesClient is the client API for Notes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotesClient interface {
	// Get list of notes in directory
	//
	// Returns a list of a user's notes for a specific directory.
	GetDirectoryNotes(ctx context.Context, in *GetDirectoryNotesRequest, opts ...grpc.CallOption) (*GetDirectoryNotesResponse, error)
	// Get a single note
	//
	// Returns a single note in a particular directory.
	GetNote(ctx context.Context, in *GetNoteRequest, opts ...grpc.CallOption) (*Note, error)
	// Create a note
	//
	// Creates a note in a particular directory with a specified ID.
	CreateNote(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Update a note's contents
	//
	// Updates a note and returns the updated entry.
	UpdateNote(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Note, error)
	// Delete a note
	//
	// Deletes a particular note by directory and ID.
	DeleteNote(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type notesClient struct {
	cc grpc.ClientConnInterface
}

func NewNotesClient(cc grpc.ClientConnInterface) NotesClient {
	return &notesClient{cc}
}

func (c *notesClient) GetDirectoryNotes(ctx context.Context, in *GetDirectoryNotesRequest, opts ...grpc.CallOption) (*GetDirectoryNotesResponse, error) {
	out := new(GetDirectoryNotesResponse)
	err := c.cc.Invoke(ctx, "/notes.Notes/GetDirectoryNotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) GetNote(ctx context.Context, in *GetNoteRequest, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/notes.Notes/GetNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) CreateNote(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/notes.Notes/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) UpdateNote(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/notes.Notes/UpdateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) DeleteNote(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/notes.Notes/DeleteNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotesServer is the server API for Notes service.
type NotesServer interface {
	// Get list of notes in directory
	//
	// Returns a list of a user's notes for a specific directory.
	GetDirectoryNotes(context.Context, *GetDirectoryNotesRequest) (*GetDirectoryNotesResponse, error)
	// Get a single note
	//
	// Returns a single note in a particular directory.
	GetNote(context.Context, *GetNoteRequest) (*Note, error)
	// Create a note
	//
	// Creates a note in a particular directory with a specified ID.
	CreateNote(context.Context, *CreateRequest) (*empty.Empty, error)
	// Update a note's contents
	//
	// Updates a note and returns the updated entry.
	UpdateNote(context.Context, *UpdateRequest) (*Note, error)
	// Delete a note
	//
	// Deletes a particular note by directory and ID.
	DeleteNote(context.Context, *DeleteRequest) (*empty.Empty, error)
}

// UnimplementedNotesServer can be embedded to have forward compatible implementations.
type UnimplementedNotesServer struct {
}

func (*UnimplementedNotesServer) GetDirectoryNotes(ctx context.Context, req *GetDirectoryNotesRequest) (*GetDirectoryNotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDirectoryNotes not implemented")
}
func (*UnimplementedNotesServer) GetNote(ctx context.Context, req *GetNoteRequest) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNote not implemented")
}
func (*UnimplementedNotesServer) CreateNote(ctx context.Context, req *CreateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (*UnimplementedNotesServer) UpdateNote(ctx context.Context, req *UpdateRequest) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNote not implemented")
}
func (*UnimplementedNotesServer) DeleteNote(ctx context.Context, req *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNote not implemented")
}

func RegisterNotesServer(s *grpc.Server, srv NotesServer) {
	s.RegisterService(&_Notes_serviceDesc, srv)
}

func _Notes_GetDirectoryNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDirectoryNotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).GetDirectoryNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/GetDirectoryNotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).GetDirectoryNotes(ctx, req.(*GetDirectoryNotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_GetNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).GetNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/GetNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).GetNote(ctx, req.(*GetNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).CreateNote(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_UpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).UpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/UpdateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).UpdateNote(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/DeleteNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).DeleteNote(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notes.Notes",
	HandlerType: (*NotesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDirectoryNotes",
			Handler:    _Notes_GetDirectoryNotes_Handler,
		},
		{
			MethodName: "GetNote",
			Handler:    _Notes_GetNote_Handler,
		},
		{
			MethodName: "CreateNote",
			Handler:    _Notes_CreateNote_Handler,
		},
		{
			MethodName: "UpdateNote",
			Handler:    _Notes_UpdateNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _Notes_DeleteNote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notes.proto",
}
