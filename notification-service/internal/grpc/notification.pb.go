// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: internal/grpc/notification.proto

package grpc

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

type EmailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Street        string                 `protobuf:"bytes,2,opt,name=street,proto3" json:"street,omitempty"`
	House         string                 `protobuf:"bytes,3,opt,name=house,proto3" json:"house,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmailRequest) Reset() {
	*x = EmailRequest{}
	mi := &file_internal_grpc_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequest) ProtoMessage() {}

func (x *EmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequest.ProtoReflect.Descriptor instead.
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_notification_proto_rawDescGZIP(), []int{0}
}

func (x *EmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailRequest) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *EmailRequest) GetHouse() string {
	if x != nil {
		return x.House
	}
	return ""
}

type Notification struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	SendAt        int64                  `protobuf:"varint,4,opt,name=send_at,json=sendAt,proto3" json:"send_at,omitempty"`
	Street        string                 `protobuf:"bytes,5,opt,name=street,proto3" json:"street,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_internal_grpc_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_internal_grpc_notification_proto_rawDescGZIP(), []int{1}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Notification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Notification) GetSendAt() int64 {
	if x != nil {
		return x.SendAt
	}
	return 0
}

func (x *Notification) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_internal_grpc_notification_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_notification_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_internal_grpc_notification_proto_rawDescGZIP(), []int{2}
}

type NotificationList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Notification        `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationList) Reset() {
	*x = NotificationList{}
	mi := &file_internal_grpc_notification_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationList) ProtoMessage() {}

func (x *NotificationList) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_notification_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationList.ProtoReflect.Descriptor instead.
func (*NotificationList) Descriptor() ([]byte, []int) {
	return file_internal_grpc_notification_proto_rawDescGZIP(), []int{3}
}

func (x *NotificationList) GetItems() []*Notification {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_internal_grpc_notification_proto protoreflect.FileDescriptor

const file_internal_grpc_notification_proto_rawDesc = "" +
	"\n" +
	" internal/grpc/notification.proto\x12\fnotification\"R\n" +
	"\fEmailRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x16\n" +
	"\x06street\x18\x02 \x01(\tR\x06street\x12\x14\n" +
	"\x05house\x18\x03 \x01(\tR\x05house\"\x7f\n" +
	"\fNotification\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\x12\x17\n" +
	"\asend_at\x18\x04 \x01(\x03R\x06sendAt\x12\x16\n" +
	"\x06street\x18\x05 \x01(\tR\x06street\"\a\n" +
	"\x05Empty\"D\n" +
	"\x10NotificationList\x120\n" +
	"\x05items\x18\x01 \x03(\v2\x1a.notification.NotificationR\x05items2\x9d\x02\n" +
	"\x13NotificationService\x12<\n" +
	"\tSubscribe\x12\x1a.notification.EmailRequest\x1a\x13.notification.Empty\x12>\n" +
	"\vUnsubscribe\x12\x1a.notification.EmailRequest\x1a\x13.notification.Empty\x12E\n" +
	"\x12CreateNotification\x12\x1a.notification.Notification\x1a\x13.notification.Empty\x12A\n" +
	"\n" +
	"GetHistory\x12\x13.notification.Empty\x1a\x1e.notification.NotificationListB@Z>github.com/ZekebayevYe/notification-service/internal/grpc;grpcb\x06proto3"

var (
	file_internal_grpc_notification_proto_rawDescOnce sync.Once
	file_internal_grpc_notification_proto_rawDescData []byte
)

func file_internal_grpc_notification_proto_rawDescGZIP() []byte {
	file_internal_grpc_notification_proto_rawDescOnce.Do(func() {
		file_internal_grpc_notification_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_grpc_notification_proto_rawDesc), len(file_internal_grpc_notification_proto_rawDesc)))
	})
	return file_internal_grpc_notification_proto_rawDescData
}

var file_internal_grpc_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_grpc_notification_proto_goTypes = []any{
	(*EmailRequest)(nil),     // 0: notification.EmailRequest
	(*Notification)(nil),     // 1: notification.Notification
	(*Empty)(nil),            // 2: notification.Empty
	(*NotificationList)(nil), // 3: notification.NotificationList
}
var file_internal_grpc_notification_proto_depIdxs = []int32{
	1, // 0: notification.NotificationList.items:type_name -> notification.Notification
	0, // 1: notification.NotificationService.Subscribe:input_type -> notification.EmailRequest
	0, // 2: notification.NotificationService.Unsubscribe:input_type -> notification.EmailRequest
	1, // 3: notification.NotificationService.CreateNotification:input_type -> notification.Notification
	2, // 4: notification.NotificationService.GetHistory:input_type -> notification.Empty
	2, // 5: notification.NotificationService.Subscribe:output_type -> notification.Empty
	2, // 6: notification.NotificationService.Unsubscribe:output_type -> notification.Empty
	2, // 7: notification.NotificationService.CreateNotification:output_type -> notification.Empty
	3, // 8: notification.NotificationService.GetHistory:output_type -> notification.NotificationList
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_grpc_notification_proto_init() }
func file_internal_grpc_notification_proto_init() {
	if File_internal_grpc_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_grpc_notification_proto_rawDesc), len(file_internal_grpc_notification_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpc_notification_proto_goTypes,
		DependencyIndexes: file_internal_grpc_notification_proto_depIdxs,
		MessageInfos:      file_internal_grpc_notification_proto_msgTypes,
	}.Build()
	File_internal_grpc_notification_proto = out.File
	file_internal_grpc_notification_proto_goTypes = nil
	file_internal_grpc_notification_proto_depIdxs = nil
}
