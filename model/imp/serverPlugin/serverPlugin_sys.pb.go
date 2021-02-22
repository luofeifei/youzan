// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serverPlugin_sys.proto

package serverPlugin

import (
	context "context"
	fmt "fmt"
	tars "github.com/TarsCloud/TarsGo/tars"
	model "github.com/TarsCloud/TarsGo/tars/model"
	requestf "github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	current "github.com/TarsCloud/TarsGo/tars/util/current"
	tools "github.com/TarsCloud/TarsGo/tars/util/tools"
	proto "github.com/golang/protobuf/proto"
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

type ReqApiPlugin struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Coid                 int64    `protobuf:"varint,2,opt,name=coid,proto3" json:"coid,omitempty"`
	FeatureId            int64    `protobuf:"varint,3,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	PluginList           []byte   `protobuf:"bytes,4,opt,name=plugin_list,json=pluginList,proto3" json:"plugin_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqApiPlugin) Reset()         { *m = ReqApiPlugin{} }
func (m *ReqApiPlugin) String() string { return proto.CompactTextString(m) }
func (*ReqApiPlugin) ProtoMessage()    {}
func (*ReqApiPlugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_486a34343ec1ca4c, []int{0}
}

func (m *ReqApiPlugin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqApiPlugin.Unmarshal(m, b)
}
func (m *ReqApiPlugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqApiPlugin.Marshal(b, m, deterministic)
}
func (m *ReqApiPlugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqApiPlugin.Merge(m, src)
}
func (m *ReqApiPlugin) XXX_Size() int {
	return xxx_messageInfo_ReqApiPlugin.Size(m)
}
func (m *ReqApiPlugin) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqApiPlugin.DiscardUnknown(m)
}

var xxx_messageInfo_ReqApiPlugin proto.InternalMessageInfo

func (m *ReqApiPlugin) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ReqApiPlugin) GetCoid() int64 {
	if m != nil {
		return m.Coid
	}
	return 0
}

func (m *ReqApiPlugin) GetFeatureId() int64 {
	if m != nil {
		return m.FeatureId
	}
	return 0
}

func (m *ReqApiPlugin) GetPluginList() []byte {
	if m != nil {
		return m.PluginList
	}
	return nil
}

// 默认输出
type Result struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_486a34343ec1ca4c, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Result) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Result) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqApiPlugin)(nil), "serverPlugin.ReqApiPlugin")
	proto.RegisterType((*Result)(nil), "serverPlugin.Result")
}

func init() { proto.RegisterFile("serverPlugin_sys.proto", fileDescriptor_486a34343ec1ca4c) }

var fileDescriptor_486a34343ec1ca4c = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x8d, 0xa9, 0x85, 0x8e, 0x39, 0x94, 0x41, 0x24, 0x14, 0xc4, 0x92, 0x53, 0x4f, 0x39,
	0xe8, 0x2f, 0x50, 0x44, 0x11, 0x14, 0x64, 0xff, 0x40, 0x49, 0xdd, 0xb1, 0x2c, 0x44, 0x37, 0xee,
	0xcc, 0x0a, 0xf9, 0xf7, 0xb2, 0xb3, 0x41, 0x82, 0xb7, 0x6f, 0xdf, 0x2c, 0xf3, 0xde, 0x1b, 0xb8,
	0x64, 0x0a, 0x3f, 0x14, 0xde, 0xfa, 0x78, 0x74, 0x5f, 0x7b, 0x1e, 0xb9, 0x1d, 0x82, 0x17, 0x8f,
	0xd5, 0x5c, 0x6f, 0x04, 0x2a, 0x43, 0xdf, 0x77, 0x83, 0xcb, 0x6f, 0x5c, 0x43, 0x19, 0x9d, 0xad,
	0x8b, 0x6d, 0xb1, 0x2b, 0x4d, 0x42, 0x44, 0x58, 0xbc, 0x7b, 0x67, 0xeb, 0x53, 0x95, 0x94, 0xf1,
	0x0a, 0xe0, 0x83, 0x3a, 0x89, 0x81, 0xf6, 0xce, 0xd6, 0xa5, 0x4e, 0x56, 0x93, 0xf2, 0x6c, 0xf1,
	0x1a, 0xce, 0x87, 0x6c, 0xdb, 0x3b, 0x96, 0x7a, 0xb1, 0x2d, 0x76, 0x95, 0x81, 0x2c, 0xbd, 0x38,
	0x96, 0xe6, 0x1e, 0x96, 0x86, 0x38, 0xf6, 0x92, 0xb7, 0x5b, 0x52, 0xc3, 0x33, 0xa3, 0x9c, 0x32,
	0x7c, 0xf2, 0x51, 0x0d, 0x57, 0x26, 0x61, 0xfa, 0x75, 0xf0, 0x76, 0x9c, 0x36, 0x29, 0xdf, 0xbc,
	0x42, 0xc9, 0x23, 0xe3, 0x23, 0xac, 0x9f, 0x48, 0xfe, 0x0a, 0x3c, 0x74, 0xd2, 0xe1, 0xa6, 0x9d,
	0x77, 0x6c, 0xe7, 0x05, 0x37, 0x17, 0xff, 0x67, 0x29, 0x46, 0x73, 0x72, 0x58, 0xea, 0x75, 0x6e,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x58, 0xc5, 0xf9, 0x37, 0x01, 0x00, 0x00,
}

// This following code was generated by tarsrpc
// Gernerated from serverPlugin_sys.proto
type Sys struct {
	s model.Servant
}

//SetServant is required by the servant interface.
func (obj *Sys) SetServant(s model.Servant) {
	obj.s = s
}

//AddServant is required by the servant interface
func (obj *Sys) AddServant(imp impSys, objStr string) {
	tars.AddServant(obj, imp, objStr)
}

////AddServant adds servant  for the service with context
func (obj *Sys) AddServantWithContext(imp impSysWithContext, objStr string) {
	tars.AddServantWithContext(obj, imp, objStr)
}

//TarsSetTimeout is required by the servant interface. t is the timeout in ms.
func (obj *Sys) TarsSetTimeout(t int) {
	obj.s.TarsSetTimeout(t)
}

//TarsSetProtocol is required by the servant interface. t is the protocol.
func (obj *Sys) TarsSetProtocol(p model.Protocol) {
	obj.s.TarsSetProtocol(p)
}

type impSys interface {
	GetApiPluginData(input ReqApiPlugin) (output Result, err error)
}

type impSysWithContext interface {
	GetApiPluginData(ctx context.Context, input ReqApiPlugin) (output Result, err error)
}

//Dispatch is used to call the user implement of the defined method.
func (obj *Sys) Dispatch(ctx context.Context, val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	input := tools.Int8ToByte(req.SBuffer)
	var output []byte
	funcName := req.SFuncName
	switch funcName {

	case "GetApiPluginData":
		inputDefine := ReqApiPlugin{}
		if err = proto.Unmarshal(input, &inputDefine); err != nil {
			return err
		}
		var res Result
		if withContext == false {
			imp := val.(impSys)
			res, err = imp.GetApiPluginData(inputDefine)
			if err != nil {
				return err
			}
		} else {
			imp := val.(impSysWithContext)
			res, err = imp.GetApiPluginData(ctx, inputDefine)
			if err != nil {
				return err
			}
		}
		output, err = proto.Marshal(&res)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(output),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}
	return nil
}

// GetApiPluginData is client rpc method as defined
func (obj *Sys) GetApiPluginData(input ReqApiPlugin, _opt ...map[string]string) (output Result, err error) {
	ctx := context.Background()
	return obj.GetApiPluginDataWithContext(ctx, input, _opt...)
}

// GetApiPluginDataWithContext is client rpc method as defined
func (obj *Sys) GetApiPluginDataWithContext(ctx context.Context, input ReqApiPlugin, _opt ...map[string]string) (output Result, err error) {
	var inputMarshal []byte
	inputMarshal, err = proto.Marshal(&input)
	if err != nil {
		return output, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}

	resp := new(requestf.ResponsePacket)

	err = obj.s.Tars_invoke(ctx, 0, "GetApiPluginData", inputMarshal, _status, _context, resp)
	if err != nil {
		return output, err
	}
	if err = proto.Unmarshal(tools.Int8ToByte(resp.SBuffer), &output); err != nil {
		return output, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range resp.Status {
			_status[k] = v
		}
	}

	return output, nil
}