// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/movies/v1/movies.proto

package moviesv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Genre       string                 `protobuf:"bytes,3,opt,name=genre,proto3" json:"genre,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ReleaseDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,1000,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime  *timestamppb.Timestamp `protobuf:"bytes,1001,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	mi := &file_api_movies_v1_movies_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_api_movies_v1_movies_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_api_movies_v1_movies_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *Movie) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Movie) GetReleaseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReleaseDate
	}
	return nil
}

func (x *Movie) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Movie) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type CreateMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *CreateMovieRequest) Reset() {
	*x = CreateMovieRequest{}
	mi := &file_api_movies_v1_movies_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMovieRequest) ProtoMessage() {}

func (x *CreateMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movies_v1_movies_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMovieRequest.ProtoReflect.Descriptor instead.
func (*CreateMovieRequest) Descriptor() ([]byte, []int) {
	return file_api_movies_v1_movies_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMovieRequest) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type DeleteMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMovieRequest) Reset() {
	*x = DeleteMovieRequest{}
	mi := &file_api_movies_v1_movies_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMovieRequest) ProtoMessage() {}

func (x *DeleteMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movies_v1_movies_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMovieRequest.ProtoReflect.Descriptor instead.
func (*DeleteMovieRequest) Descriptor() ([]byte, []int) {
	return file_api_movies_v1_movies_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteMovieRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UndeleteMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UndeleteMovieRequest) Reset() {
	*x = UndeleteMovieRequest{}
	mi := &file_api_movies_v1_movies_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UndeleteMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UndeleteMovieRequest) ProtoMessage() {}

func (x *UndeleteMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movies_v1_movies_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UndeleteMovieRequest.ProtoReflect.Descriptor instead.
func (*UndeleteMovieRequest) Descriptor() ([]byte, []int) {
	return file_api_movies_v1_movies_proto_rawDescGZIP(), []int{3}
}

func (x *UndeleteMovieRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie      *Movie                 `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateMovieRequest) Reset() {
	*x = UpdateMovieRequest{}
	mi := &file_api_movies_v1_movies_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMovieRequest) ProtoMessage() {}

func (x *UpdateMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movies_v1_movies_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMovieRequest.ProtoReflect.Descriptor instead.
func (*UpdateMovieRequest) Descriptor() ([]byte, []int) {
	return file_api_movies_v1_movies_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMovieRequest) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

func (x *UpdateMovieRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type GetMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMovieRequest) Reset() {
	*x = GetMovieRequest{}
	mi := &file_api_movies_v1_movies_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieRequest) ProtoMessage() {}

func (x *GetMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movies_v1_movies_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieRequest.ProtoReflect.Descriptor instead.
func (*GetMovieRequest) Descriptor() ([]byte, []int) {
	return file_api_movies_v1_movies_proto_rawDescGZIP(), []int{5}
}

func (x *GetMovieRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListMoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListMoviesRequest) Reset() {
	*x = ListMoviesRequest{}
	mi := &file_api_movies_v1_movies_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMoviesRequest) ProtoMessage() {}

func (x *ListMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_movies_v1_movies_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMoviesRequest.ProtoReflect.Descriptor instead.
func (*ListMoviesRequest) Descriptor() ([]byte, []int) {
	return file_api_movies_v1_movies_proto_rawDescGZIP(), []int{6}
}

func (x *ListMoviesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListMoviesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListMoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies        []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListMoviesResponse) Reset() {
	*x = ListMoviesResponse{}
	mi := &file_api_movies_v1_movies_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMoviesResponse) ProtoMessage() {}

func (x *ListMoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_movies_v1_movies_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMoviesResponse.ProtoReflect.Descriptor instead.
func (*ListMoviesResponse) Descriptor() ([]byte, []int) {
	return file_api_movies_v1_movies_proto_rawDescGZIP(), []int{7}
}

func (x *ListMoviesResponse) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

func (x *ListMoviesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_api_movies_v1_movies_proto protoreflect.FileDescriptor

var file_api_movies_v1_movies_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd,
	0x02, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3d, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3c,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xe8, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xe9, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x32, 0xea, 0x41, 0x2f, 0x0a,
	0x0e, 0x79, 0x61, 0x6e, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x2f, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x0e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x7d, 0x2a,
	0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x32, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x45,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x43, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1d, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x10, 0x0a,
	0x0e, 0x79, 0x61, 0x6e, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x2f, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0xba,
	0x48, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x14, 0x55, 0x6e,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1d,
	0xe0, 0x41, 0x02, 0xfa, 0x41, 0x10, 0x0a, 0x0e, 0x79, 0x61, 0x6e, 0x74, 0x6f, 0x2e, 0x61, 0x72,
	0x2f, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x88, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x09, 0xe0,
	0x41, 0x02, 0xba, 0x48, 0x03, 0xc0, 0x01, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x40, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1d, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x10, 0x0a, 0x0e, 0x79, 0x61, 0x6e, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x2f, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6a, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a,
	0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xd6, 0x03, 0x0a, 0x13, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x21, 0xda, 0x41, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x0a, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x64, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x22, 0x1c, 0xda, 0x41, 0x02, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6f,
	0x0a, 0x0d, 0x55, 0x6e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x6e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x23, 0xda, 0x41, 0x02, 0x69,
	0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x3a, 0x75, 0x6e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x7d, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x21,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x35, 0xda, 0x41, 0x0e, 0x69, 0x64, 0x2c, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x3a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x32, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x69, 0x64, 0x7d, 0x32, 0xdf,
	0x01, 0x0a, 0x13, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x1f, 0xda, 0x41, 0x05, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x65, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x42, 0xf7, 0x02, 0x92, 0x41, 0xdb, 0x01, 0x12, 0xa7, 0x01, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x20, 0x41, 0x50, 0x49, 0x12, 0x76, 0x41, 0x20, 0x77, 0x65, 0x62, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x20, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x2c, 0x20, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x69, 0x6e,
	0x67, 0x20, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x74,
	0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64,
	0x20, 0x69, 0x6e, 0x20, 0x72, 0x65, 0x61, 0x6c, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x20, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x67, 0x72, 0x61, 0x64, 0x65, 0x20,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x22, 0x0e, 0x0a,
	0x05, 0x59, 0x61, 0x6e, 0x74, 0x6f, 0x1a, 0x05, 0x79, 0x61, 0x6e, 0x74, 0x6f, 0x2a, 0x0d, 0x0a,
	0x0b, 0x4d, 0x49, 0x54, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x32, 0x03, 0x30, 0x2e,
	0x38, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x08, 0x0a, 0x06, 0x52, 0x45, 0x41, 0x44,
	0x4d, 0x45, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x4d, 0x58, 0xaa, 0x02, 0x0d, 0x41,
	0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x41,
	0x70, 0x69, 0x5c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x41,
	0x70, 0x69, 0x5c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_movies_v1_movies_proto_rawDescOnce sync.Once
	file_api_movies_v1_movies_proto_rawDescData = file_api_movies_v1_movies_proto_rawDesc
)

func file_api_movies_v1_movies_proto_rawDescGZIP() []byte {
	file_api_movies_v1_movies_proto_rawDescOnce.Do(func() {
		file_api_movies_v1_movies_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_movies_v1_movies_proto_rawDescData)
	})
	return file_api_movies_v1_movies_proto_rawDescData
}

var file_api_movies_v1_movies_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_movies_v1_movies_proto_goTypes = []any{
	(*Movie)(nil),                 // 0: api.movies.v1.Movie
	(*CreateMovieRequest)(nil),    // 1: api.movies.v1.CreateMovieRequest
	(*DeleteMovieRequest)(nil),    // 2: api.movies.v1.DeleteMovieRequest
	(*UndeleteMovieRequest)(nil),  // 3: api.movies.v1.UndeleteMovieRequest
	(*UpdateMovieRequest)(nil),    // 4: api.movies.v1.UpdateMovieRequest
	(*GetMovieRequest)(nil),       // 5: api.movies.v1.GetMovieRequest
	(*ListMoviesRequest)(nil),     // 6: api.movies.v1.ListMoviesRequest
	(*ListMoviesResponse)(nil),    // 7: api.movies.v1.ListMoviesResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 9: google.protobuf.FieldMask
}
var file_api_movies_v1_movies_proto_depIdxs = []int32{
	8,  // 0: api.movies.v1.Movie.release_date:type_name -> google.protobuf.Timestamp
	8,  // 1: api.movies.v1.Movie.create_time:type_name -> google.protobuf.Timestamp
	8,  // 2: api.movies.v1.Movie.update_time:type_name -> google.protobuf.Timestamp
	0,  // 3: api.movies.v1.CreateMovieRequest.movie:type_name -> api.movies.v1.Movie
	0,  // 4: api.movies.v1.UpdateMovieRequest.movie:type_name -> api.movies.v1.Movie
	9,  // 5: api.movies.v1.UpdateMovieRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 6: api.movies.v1.ListMoviesResponse.movies:type_name -> api.movies.v1.Movie
	1,  // 7: api.movies.v1.MoviesWriterService.CreateMovie:input_type -> api.movies.v1.CreateMovieRequest
	2,  // 8: api.movies.v1.MoviesWriterService.DeleteMovie:input_type -> api.movies.v1.DeleteMovieRequest
	3,  // 9: api.movies.v1.MoviesWriterService.UndeleteMovie:input_type -> api.movies.v1.UndeleteMovieRequest
	4,  // 10: api.movies.v1.MoviesWriterService.UpdateMovie:input_type -> api.movies.v1.UpdateMovieRequest
	5,  // 11: api.movies.v1.MoviesReaderService.GetMovie:input_type -> api.movies.v1.GetMovieRequest
	6,  // 12: api.movies.v1.MoviesReaderService.ListMovies:input_type -> api.movies.v1.ListMoviesRequest
	0,  // 13: api.movies.v1.MoviesWriterService.CreateMovie:output_type -> api.movies.v1.Movie
	0,  // 14: api.movies.v1.MoviesWriterService.DeleteMovie:output_type -> api.movies.v1.Movie
	0,  // 15: api.movies.v1.MoviesWriterService.UndeleteMovie:output_type -> api.movies.v1.Movie
	0,  // 16: api.movies.v1.MoviesWriterService.UpdateMovie:output_type -> api.movies.v1.Movie
	0,  // 17: api.movies.v1.MoviesReaderService.GetMovie:output_type -> api.movies.v1.Movie
	7,  // 18: api.movies.v1.MoviesReaderService.ListMovies:output_type -> api.movies.v1.ListMoviesResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_movies_v1_movies_proto_init() }
func file_api_movies_v1_movies_proto_init() {
	if File_api_movies_v1_movies_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_movies_v1_movies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_movies_v1_movies_proto_goTypes,
		DependencyIndexes: file_api_movies_v1_movies_proto_depIdxs,
		MessageInfos:      file_api_movies_v1_movies_proto_msgTypes,
	}.Build()
	File_api_movies_v1_movies_proto = out.File
	file_api_movies_v1_movies_proto_rawDesc = nil
	file_api_movies_v1_movies_proto_goTypes = nil
	file_api_movies_v1_movies_proto_depIdxs = nil
}
