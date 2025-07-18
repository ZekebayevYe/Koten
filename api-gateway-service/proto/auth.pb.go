// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/auth.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	FullName      string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	House         string                 `protobuf:"bytes,4,opt,name=house,proto3" json:"house,omitempty"`
	Street        string                 `protobuf:"bytes,5,opt,name=street,proto3" json:"street,omitempty"`
	Apartment     string                 `protobuf:"bytes,6,opt,name=apartment,proto3" json:"apartment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_proto_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *RegisterRequest) GetHouse() string {
	if x != nil {
		return x.House
	}
	return ""
}

func (x *RegisterRequest) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *RegisterRequest) GetApartment() string {
	if x != nil {
		return x.Apartment
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_proto_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetMyProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMyProfileRequest) Reset() {
	*x = GetMyProfileRequest{}
	mi := &file_proto_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMyProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyProfileRequest) ProtoMessage() {}

func (x *GetMyProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyProfileRequest.ProtoReflect.Descriptor instead.
func (*GetMyProfileRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *GetMyProfileRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	FullName      string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	House         string                 `protobuf:"bytes,3,opt,name=house,proto3" json:"house,omitempty"`
	Street        string                 `protobuf:"bytes,4,opt,name=street,proto3" json:"street,omitempty"`
	Apartment     string                 `protobuf:"bytes,5,opt,name=apartment,proto3" json:"apartment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProfileRequest) Reset() {
	*x = UpdateProfileRequest{}
	mi := &file_proto_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileRequest) ProtoMessage() {}

func (x *UpdateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProfileRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateProfileRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateProfileRequest) GetHouse() string {
	if x != nil {
		return x.House
	}
	return ""
}

func (x *UpdateProfileRequest) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *UpdateProfileRequest) GetApartment() string {
	if x != nil {
		return x.Apartment
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	mi := &file_proto_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *AuthResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserProfile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	FullName      string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Role          string                 `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	House         string                 `protobuf:"bytes,4,opt,name=house,proto3" json:"house,omitempty"`
	Street        string                 `protobuf:"bytes,5,opt,name=street,proto3" json:"street,omitempty"`
	Apartment     string                 `protobuf:"bytes,6,opt,name=apartment,proto3" json:"apartment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	mi := &file_proto_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *UserProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProfile) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UserProfile) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserProfile) GetHouse() string {
	if x != nil {
		return x.House
	}
	return ""
}

func (x *UserProfile) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *UserProfile) GetApartment() string {
	if x != nil {
		return x.Apartment
	}
	return ""
}

var File_proto_auth_proto protoreflect.FileDescriptor

const file_proto_auth_proto_rawDesc = "" +
	"\n" +
	"\x10proto/auth.proto\x12\x04auth\"\xac\x01\n" +
	"\x0fRegisterRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x1b\n" +
	"\tfull_name\x18\x03 \x01(\tR\bfullName\x12\x14\n" +
	"\x05house\x18\x04 \x01(\tR\x05house\x12\x16\n" +
	"\x06street\x18\x05 \x01(\tR\x06street\x12\x1c\n" +
	"\tapartment\x18\x06 \x01(\tR\tapartment\"@\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"+\n" +
	"\x13GetMyProfileRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"\x95\x01\n" +
	"\x14UpdateProfileRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\x1b\n" +
	"\tfull_name\x18\x02 \x01(\tR\bfullName\x12\x14\n" +
	"\x05house\x18\x03 \x01(\tR\x05house\x12\x16\n" +
	"\x06street\x18\x04 \x01(\tR\x06street\x12\x1c\n" +
	"\tapartment\x18\x05 \x01(\tR\tapartment\"$\n" +
	"\fAuthResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"\xa0\x01\n" +
	"\vUserProfile\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1b\n" +
	"\tfull_name\x18\x02 \x01(\tR\bfullName\x12\x12\n" +
	"\x04role\x18\x03 \x01(\tR\x04role\x12\x14\n" +
	"\x05house\x18\x04 \x01(\tR\x05house\x12\x16\n" +
	"\x06street\x18\x05 \x01(\tR\x06street\x12\x1c\n" +
	"\tapartment\x18\x06 \x01(\tR\tapartment2\xfd\x01\n" +
	"\vAuthService\x129\n" +
	"\fRegisterUser\x12\x15.auth.RegisterRequest\x1a\x12.auth.AuthResponse\x123\n" +
	"\tLoginUser\x12\x12.auth.LoginRequest\x1a\x12.auth.AuthResponse\x12<\n" +
	"\fGetMyProfile\x12\x19.auth.GetMyProfileRequest\x1a\x11.auth.UserProfile\x12@\n" +
	"\x0fUpdateMyProfile\x12\x1a.auth.UpdateProfileRequest\x1a\x11.auth.UserProfileB#Z!auth-service/internal/proto;protob\x06proto3"

var (
	file_proto_auth_proto_rawDescOnce sync.Once
	file_proto_auth_proto_rawDescData []byte
)

func file_proto_auth_proto_rawDescGZIP() []byte {
	file_proto_auth_proto_rawDescOnce.Do(func() {
		file_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_auth_proto_rawDesc), len(file_proto_auth_proto_rawDesc)))
	})
	return file_proto_auth_proto_rawDescData
}

var file_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_auth_proto_goTypes = []any{
	(*RegisterRequest)(nil),      // 0: auth.RegisterRequest
	(*LoginRequest)(nil),         // 1: auth.LoginRequest
	(*GetMyProfileRequest)(nil),  // 2: auth.GetMyProfileRequest
	(*UpdateProfileRequest)(nil), // 3: auth.UpdateProfileRequest
	(*AuthResponse)(nil),         // 4: auth.AuthResponse
	(*UserProfile)(nil),          // 5: auth.UserProfile
}
var file_proto_auth_proto_depIdxs = []int32{
	0, // 0: auth.AuthService.RegisterUser:input_type -> auth.RegisterRequest
	1, // 1: auth.AuthService.LoginUser:input_type -> auth.LoginRequest
	2, // 2: auth.AuthService.GetMyProfile:input_type -> auth.GetMyProfileRequest
	3, // 3: auth.AuthService.UpdateMyProfile:input_type -> auth.UpdateProfileRequest
	4, // 4: auth.AuthService.RegisterUser:output_type -> auth.AuthResponse
	4, // 5: auth.AuthService.LoginUser:output_type -> auth.AuthResponse
	5, // 6: auth.AuthService.GetMyProfile:output_type -> auth.UserProfile
	5, // 7: auth.AuthService.UpdateMyProfile:output_type -> auth.UserProfile
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auth_proto_init() }
func file_proto_auth_proto_init() {
	if File_proto_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_auth_proto_rawDesc), len(file_proto_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_proto_goTypes,
		DependencyIndexes: file_proto_auth_proto_depIdxs,
		MessageInfos:      file_proto_auth_proto_msgTypes,
	}.Build()
	File_proto_auth_proto = out.File
	file_proto_auth_proto_goTypes = nil
	file_proto_auth_proto_depIdxs = nil
}
