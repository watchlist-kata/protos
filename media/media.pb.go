// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: media.proto

package media

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

// Общая модель для медиа (для нашей базы данных)
type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                           // ID в вашей базе данных
	TmdbId        int64  `protobuf:"varint,2,opt,name=tmdb_id,json=tmdbId,proto3" json:"tmdb_id,omitempty"`                     // ID в TMDB
	Type          string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`                                        // Тип: movie / tv
	Title         string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`                                      // Название
	OriginalTitle string `protobuf:"bytes,5,opt,name=original_title,json=originalTitle,proto3" json:"original_title,omitempty"` // Оригинальное название (если есть)
	Description   string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`                          // Описание
	ReleaseDate   string `protobuf:"bytes,7,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`       // Дата релиза
	Poster        string `protobuf:"bytes,8,opt,name=poster,proto3" json:"poster,omitempty"`                                    // URL постера
}

func (x *Media) Reset() {
	*x = Media{}
	mi := &file_media_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Media) GetTmdbId() int64 {
	if x != nil {
		return x.TmdbId
	}
	return 0
}

func (x *Media) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Media) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Media) GetOriginalTitle() string {
	if x != nil {
		return x.OriginalTitle
	}
	return ""
}

func (x *Media) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Media) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *Media) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

// Модель для результата поиска TMDB
type TMDBMedia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adult            bool    `protobuf:"varint,1,opt,name=adult,proto3" json:"adult,omitempty"`                                              // Для взрослых
	BackdropPath     string  `protobuf:"bytes,2,opt,name=backdrop_path,json=backdropPath,proto3" json:"backdrop_path,omitempty"`             // Путь к фоновому изображению
	Id               int64   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`                                                    // ID в TMDB
	Title            string  `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`                                               // Название (для фильмов)
	OriginalLanguage string  `protobuf:"bytes,5,opt,name=original_language,json=originalLanguage,proto3" json:"original_language,omitempty"` // Оригинальный язык
	OriginalTitle    string  `protobuf:"bytes,6,opt,name=original_title,json=originalTitle,proto3" json:"original_title,omitempty"`          // Оригинальное название (для фильмов)
	Overview         string  `protobuf:"bytes,7,opt,name=overview,proto3" json:"overview,omitempty"`                                         // Описание
	PosterPath       string  `protobuf:"bytes,8,opt,name=poster_path,json=posterPath,proto3" json:"poster_path,omitempty"`                   // Путь к постеру
	MediaType        string  `protobuf:"bytes,9,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`                      // Тип медиа (movie, tv, person)
	GenreIds         []int32 `protobuf:"varint,10,rep,packed,name=genre_ids,json=genreIds,proto3" json:"genre_ids,omitempty"`                // Список ID жанров
	Popularity       float32 `protobuf:"fixed32,11,opt,name=popularity,proto3" json:"popularity,omitempty"`                                  // Популярность
	ReleaseDate      string  `protobuf:"bytes,12,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`               // Дата релиза
	Video            bool    `protobuf:"varint,13,opt,name=video,proto3" json:"video,omitempty"`                                             // Видео
	VoteAverage      float32 `protobuf:"fixed32,14,opt,name=vote_average,json=voteAverage,proto3" json:"vote_average,omitempty"`             // Средний рейтинг
	VoteCount        int32   `protobuf:"varint,15,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`                    // Количество голосов в рейтинге
	Name             string  `protobuf:"bytes,16,opt,name=name,proto3" json:"name,omitempty"`                                                // Название / имя (для сериалов и людей)
	OriginalName     string  `protobuf:"bytes,17,opt,name=original_name,json=originalName,proto3" json:"original_name,omitempty"`            // Оригинальное название (для сериалов)
	FirstAirDate     string  `protobuf:"bytes,18,opt,name=first_air_date,json=firstAirDate,proto3" json:"first_air_date,omitempty"`          // Дата первого эфира (для сериалов)
}

func (x *TMDBMedia) Reset() {
	*x = TMDBMedia{}
	mi := &file_media_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TMDBMedia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TMDBMedia) ProtoMessage() {}

func (x *TMDBMedia) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TMDBMedia.ProtoReflect.Descriptor instead.
func (*TMDBMedia) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{1}
}

func (x *TMDBMedia) GetAdult() bool {
	if x != nil {
		return x.Adult
	}
	return false
}

func (x *TMDBMedia) GetBackdropPath() string {
	if x != nil {
		return x.BackdropPath
	}
	return ""
}

func (x *TMDBMedia) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TMDBMedia) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TMDBMedia) GetOriginalLanguage() string {
	if x != nil {
		return x.OriginalLanguage
	}
	return ""
}

func (x *TMDBMedia) GetOriginalTitle() string {
	if x != nil {
		return x.OriginalTitle
	}
	return ""
}

func (x *TMDBMedia) GetOverview() string {
	if x != nil {
		return x.Overview
	}
	return ""
}

func (x *TMDBMedia) GetPosterPath() string {
	if x != nil {
		return x.PosterPath
	}
	return ""
}

func (x *TMDBMedia) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *TMDBMedia) GetGenreIds() []int32 {
	if x != nil {
		return x.GenreIds
	}
	return nil
}

func (x *TMDBMedia) GetPopularity() float32 {
	if x != nil {
		return x.Popularity
	}
	return 0
}

func (x *TMDBMedia) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *TMDBMedia) GetVideo() bool {
	if x != nil {
		return x.Video
	}
	return false
}

func (x *TMDBMedia) GetVoteAverage() float32 {
	if x != nil {
		return x.VoteAverage
	}
	return 0
}

func (x *TMDBMedia) GetVoteCount() int32 {
	if x != nil {
		return x.VoteCount
	}
	return 0
}

func (x *TMDBMedia) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TMDBMedia) GetOriginalName() string {
	if x != nil {
		return x.OriginalName
	}
	return ""
}

func (x *TMDBMedia) GetFirstAirDate() string {
	if x != nil {
		return x.FirstAirDate
	}
	return ""
}

// Запросы и ответы
type GetMediaByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMediaByIDRequest) Reset() {
	*x = GetMediaByIDRequest{}
	mi := &file_media_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMediaByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMediaByIDRequest) ProtoMessage() {}

func (x *GetMediaByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMediaByIDRequest.ProtoReflect.Descriptor instead.
func (*GetMediaByIDRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{2}
}

func (x *GetMediaByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMediasByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetMediasByNameRequest) Reset() {
	*x = GetMediasByNameRequest{}
	mi := &file_media_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMediasByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMediasByNameRequest) ProtoMessage() {}

func (x *GetMediasByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMediasByNameRequest.ProtoReflect.Descriptor instead.
func (*GetMediasByNameRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{3}
}

func (x *GetMediasByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SaveMediaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Media *Media `protobuf:"bytes,1,opt,name=media,proto3" json:"media,omitempty"`
}

func (x *SaveMediaRequest) Reset() {
	*x = SaveMediaRequest{}
	mi := &file_media_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveMediaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveMediaRequest) ProtoMessage() {}

func (x *SaveMediaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveMediaRequest.ProtoReflect.Descriptor instead.
func (*SaveMediaRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{4}
}

func (x *SaveMediaRequest) GetMedia() *Media {
	if x != nil {
		return x.Media
	}
	return nil
}

type MediaList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Medias []*Media `protobuf:"bytes,1,rep,name=medias,proto3" json:"medias,omitempty"`
}

func (x *MediaList) Reset() {
	*x = MediaList{}
	mi := &file_media_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaList) ProtoMessage() {}

func (x *MediaList) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaList.ProtoReflect.Descriptor instead.
func (*MediaList) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{5}
}

func (x *MediaList) GetMedias() []*Media {
	if x != nil {
		return x.Medias
	}
	return nil
}

type SearchTMDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SearchTMDBRequest) Reset() {
	*x = SearchTMDBRequest{}
	mi := &file_media_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchTMDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTMDBRequest) ProtoMessage() {}

func (x *SearchTMDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTMDBRequest.ProtoReflect.Descriptor instead.
func (*SearchTMDBRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{6}
}

func (x *SearchTMDBRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SearchTMDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*TMDBMedia `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"` // Список результатов поиска TMDB
}

func (x *SearchTMDBResponse) Reset() {
	*x = SearchTMDBResponse{}
	mi := &file_media_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchTMDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTMDBResponse) ProtoMessage() {}

func (x *SearchTMDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTMDBResponse.ProtoReflect.Descriptor instead.
func (*SearchTMDBResponse) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{7}
}

func (x *SearchTMDBResponse) GetResults() []*TMDBMedia {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_media_proto protoreflect.FileDescriptor

var file_media_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x22, 0xde, 0x01, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x6d, 0x64, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x74, 0x6d, 0x64, 0x62, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0xb3, 0x04, 0x0a, 0x09, 0x54, 0x4d, 0x44, 0x42, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x63,
	0x6b, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x74, 0x68, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x49, 0x64,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f,
	0x74, 0x65, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x76, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x61,
	0x69, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x41, 0x69, 0x72, 0x44, 0x61, 0x74, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x36, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x22, 0x31, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x4d, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x4d,
	0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2e, 0x54, 0x4d, 0x44, 0x42, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0xb9, 0x02, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x12, 0x42, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x34, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12,
	0x41, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x4d, 0x44, 0x42, 0x12, 0x18, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x4d, 0x44, 0x42,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x4d, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x6b, 0x61, 0x74, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_media_proto_rawDescOnce sync.Once
	file_media_proto_rawDescData = file_media_proto_rawDesc
)

func file_media_proto_rawDescGZIP() []byte {
	file_media_proto_rawDescOnce.Do(func() {
		file_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_media_proto_rawDescData)
	})
	return file_media_proto_rawDescData
}

var file_media_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_media_proto_goTypes = []any{
	(*Media)(nil),                  // 0: media.Media
	(*TMDBMedia)(nil),              // 1: media.TMDBMedia
	(*GetMediaByIDRequest)(nil),    // 2: media.GetMediaByIDRequest
	(*GetMediasByNameRequest)(nil), // 3: media.GetMediasByNameRequest
	(*SaveMediaRequest)(nil),       // 4: media.SaveMediaRequest
	(*MediaList)(nil),              // 5: media.MediaList
	(*SearchTMDBRequest)(nil),      // 6: media.SearchTMDBRequest
	(*SearchTMDBResponse)(nil),     // 7: media.SearchTMDBResponse
}
var file_media_proto_depIdxs = []int32{
	0, // 0: media.SaveMediaRequest.media:type_name -> media.Media
	0, // 1: media.MediaList.medias:type_name -> media.Media
	1, // 2: media.SearchTMDBResponse.results:type_name -> media.TMDBMedia
	2, // 3: media.MediaService.GetMediaByID:input_type -> media.GetMediaByIDRequest
	3, // 4: media.MediaService.GetMediasByName:input_type -> media.GetMediasByNameRequest
	4, // 5: media.MediaService.SaveMedia:input_type -> media.SaveMediaRequest
	4, // 6: media.MediaService.UpdateMedia:input_type -> media.SaveMediaRequest
	6, // 7: media.MediaService.SearchTMDB:input_type -> media.SearchTMDBRequest
	0, // 8: media.MediaService.GetMediaByID:output_type -> media.Media
	5, // 9: media.MediaService.GetMediasByName:output_type -> media.MediaList
	0, // 10: media.MediaService.SaveMedia:output_type -> media.Media
	0, // 11: media.MediaService.UpdateMedia:output_type -> media.Media
	7, // 12: media.MediaService.SearchTMDB:output_type -> media.SearchTMDBResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_media_proto_init() }
func file_media_proto_init() {
	if File_media_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_media_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_media_proto_goTypes,
		DependencyIndexes: file_media_proto_depIdxs,
		MessageInfos:      file_media_proto_msgTypes,
	}.Build()
	File_media_proto = out.File
	file_media_proto_rawDesc = nil
	file_media_proto_goTypes = nil
	file_media_proto_depIdxs = nil
}
