// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/auth/v2alpha/external_auth.proto

package v2alpha

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_rpc "github.com/gogo/googleapis/google/rpc"
import _ "github.com/lyft/protoc-gen-validate/validate"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CheckRequest struct {
	// The request attributes.
	Attributes *AttributeContext `protobuf:"bytes,1,opt,name=attributes" json:"attributes,omitempty"`
}

func (m *CheckRequest) Reset()                    { *m = CheckRequest{} }
func (m *CheckRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()               {}
func (*CheckRequest) Descriptor() ([]byte, []int) { return fileDescriptorExternalAuth, []int{0} }

func (m *CheckRequest) GetAttributes() *AttributeContext {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type CheckResponse struct {
	// Status `OK` allows the request. Any other status indicates the request should be denied.
	Status *google_rpc.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *CheckResponse) Reset()                    { *m = CheckResponse{} }
func (m *CheckResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()               {}
func (*CheckResponse) Descriptor() ([]byte, []int) { return fileDescriptorExternalAuth, []int{1} }

func (m *CheckResponse) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// An optional message that contains HTTP response attributes. This message is
// used when the authorization service needs to send custom responses to the
// downstream client or, to modify/add request headers being dispatched to the upstream.
type CheckResponse_HttpResponse struct {
	// Http status code.
	StatusCode uint32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// Http entity headers.
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Http entity body.
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *CheckResponse_HttpResponse) Reset()         { *m = CheckResponse_HttpResponse{} }
func (m *CheckResponse_HttpResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse_HttpResponse) ProtoMessage()    {}
func (*CheckResponse_HttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorExternalAuth, []int{1, 0}
}

func (m *CheckResponse_HttpResponse) GetStatusCode() uint32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *CheckResponse_HttpResponse) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *CheckResponse_HttpResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "envoy.service.auth.v2alpha.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "envoy.service.auth.v2alpha.CheckResponse")
	proto.RegisterType((*CheckResponse_HttpResponse)(nil), "envoy.service.auth.v2alpha.CheckResponse.HttpResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authorization service

type AuthorizationClient interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := grpc.Invoke(ctx, "/envoy.service.auth.v2alpha.Authorization/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authorization service

type AuthorizationServer interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.auth.v2alpha.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.auth.v2alpha.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/auth/v2alpha/external_auth.proto",
}

func (m *CheckRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Attributes != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExternalAuth(dAtA, i, uint64(m.Attributes.Size()))
		n1, err := m.Attributes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *CheckResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExternalAuth(dAtA, i, uint64(m.Status.Size()))
		n2, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *CheckResponse_HttpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckResponse_HttpResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StatusCode != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintExternalAuth(dAtA, i, uint64(m.StatusCode))
	}
	if len(m.Headers) > 0 {
		for k, _ := range m.Headers {
			dAtA[i] = 0x12
			i++
			v := m.Headers[k]
			mapSize := 1 + len(k) + sovExternalAuth(uint64(len(k))) + 1 + len(v) + sovExternalAuth(uint64(len(v)))
			i = encodeVarintExternalAuth(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintExternalAuth(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintExternalAuth(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintExternalAuth(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	return i, nil
}

func encodeVarintExternalAuth(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CheckRequest) Size() (n int) {
	var l int
	_ = l
	if m.Attributes != nil {
		l = m.Attributes.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	return n
}

func (m *CheckResponse) Size() (n int) {
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	return n
}

func (m *CheckResponse_HttpResponse) Size() (n int) {
	var l int
	_ = l
	if m.StatusCode != 0 {
		n += 1 + sovExternalAuth(uint64(m.StatusCode))
	}
	if len(m.Headers) > 0 {
		for k, v := range m.Headers {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovExternalAuth(uint64(len(k))) + 1 + len(v) + sovExternalAuth(uint64(len(v)))
			n += mapEntrySize + 1 + sovExternalAuth(uint64(mapEntrySize))
		}
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	return n
}

func sovExternalAuth(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExternalAuth(x uint64) (n int) {
	return sovExternalAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CheckRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CheckRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attributes == nil {
				m.Attributes = &AttributeContext{}
			}
			if err := m.Attributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CheckResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CheckResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &google_rpc.Status{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CheckResponse_HttpResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HttpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusCode", wireType)
			}
			m.StatusCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatusCode |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Headers == nil {
				m.Headers = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExternalAuth
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExternalAuth
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthExternalAuth
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExternalAuth
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthExternalAuth
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipExternalAuth(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthExternalAuth
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Headers[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExternalAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExternalAuth
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthExternalAuth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExternalAuth
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipExternalAuth(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthExternalAuth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExternalAuth   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/service/auth/v2alpha/external_auth.proto", fileDescriptorExternalAuth)
}

var fileDescriptorExternalAuth = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x99, 0xdd, 0x6e, 0x97, 0xbe, 0xcd, 0x42, 0x19, 0x04, 0x43, 0x0e, 0xcb, 0xd2, 0x53,
	0x2c, 0x3a, 0x81, 0x78, 0x91, 0xde, 0xda, 0x20, 0xf4, 0xe0, 0x29, 0xde, 0xa4, 0xb2, 0xcc, 0x26,
	0x8f, 0x26, 0x34, 0x66, 0xe2, 0xcc, 0x4b, 0x68, 0xfc, 0x04, 0x7e, 0x16, 0x3f, 0x82, 0x27, 0x8f,
	0x1e, 0x3d, 0x7a, 0x94, 0xbd, 0x09, 0x7e, 0x08, 0xe9, 0x4c, 0x52, 0xd6, 0x83, 0xa5, 0xb7, 0x97,
	0xf9, 0xff, 0xe7, 0x97, 0xdf, 0x23, 0x01, 0x81, 0x75, 0xa7, 0xfa, 0xc8, 0xa0, 0xee, 0xca, 0x0c,
	0x23, 0xd9, 0x52, 0x11, 0x75, 0xb1, 0xac, 0x9a, 0x42, 0x46, 0x78, 0x4b, 0xa8, 0x6b, 0x59, 0x6d,
	0xee, 0x4e, 0x45, 0xa3, 0x15, 0x29, 0x1e, 0xd8, 0xbe, 0x18, 0xfa, 0xc2, 0x26, 0x43, 0x3f, 0x88,
	0x1f, 0x60, 0x49, 0x22, 0x5d, 0x6e, 0x5b, 0xc2, 0x4d, 0xa6, 0x6a, 0xc2, 0x5b, 0x72, 0xbc, 0xe0,
	0xe9, 0xb5, 0x52, 0xd7, 0x15, 0x46, 0xba, 0xc9, 0x22, 0x43, 0x92, 0x5a, 0x33, 0x06, 0x9d, 0xac,
	0xca, 0x5c, 0x12, 0x46, 0xe3, 0xe0, 0x82, 0x93, 0x2b, 0xf0, 0x92, 0x02, 0xb3, 0x9b, 0x14, 0x3f,
	0xb6, 0x68, 0x88, 0xbf, 0x01, 0xb8, 0x87, 0x1b, 0x9f, 0xad, 0x59, 0xb8, 0x88, 0x9f, 0x8b, 0xff,
	0x6b, 0x8a, 0xf3, 0xb1, 0x9d, 0x38, 0x93, 0x74, 0xef, 0xfe, 0xc9, 0x97, 0x09, 0x2c, 0x07, 0xbc,
	0x69, 0x54, 0x6d, 0x90, 0x9f, 0xc2, 0xa1, 0x13, 0x1b, 0xd8, 0x5c, 0x38, 0x65, 0xa1, 0x9b, 0x4c,
	0xbc, 0xb5, 0x49, 0x3a, 0x34, 0x82, 0x3f, 0x0c, 0xbc, 0x4b, 0xa2, 0xe6, 0xfe, 0xf2, 0x0b, 0x58,
	0xb8, 0x68, 0x93, 0xa9, 0x1c, 0x2d, 0x61, 0x79, 0xe1, 0x7d, 0xfd, 0xfd, 0x6d, 0x3a, 0x3f, 0x9d,
	0x1d, 0xff, 0x3c, 0x08, 0xf3, 0x14, 0x5c, 0x21, 0x51, 0x39, 0xf2, 0xf7, 0x30, 0x2f, 0x50, 0xe6,
	0xa8, 0x8d, 0x3f, 0x59, 0x4f, 0xc3, 0x45, 0x9c, 0x3c, 0xb4, 0xc8, 0x3f, 0x9e, 0x62, 0xff, 0xbd,
	0xe2, 0xd2, 0x51, 0x5e, 0xd7, 0xa4, 0xfb, 0x74, 0x64, 0x72, 0x0e, 0x07, 0x5b, 0x95, 0xf7, 0xfe,
	0x74, 0xcd, 0xc2, 0xa3, 0xd4, 0xce, 0xc1, 0x19, 0x78, 0xfb, 0x65, 0x7e, 0x0c, 0xd3, 0x1b, 0xec,
	0xad, 0xe9, 0x51, 0x7a, 0x37, 0xf2, 0x27, 0x30, 0xeb, 0x64, 0xd5, 0xa2, 0x3f, 0xb1, 0x67, 0xee,
	0xe1, 0x6c, 0xf2, 0x8a, 0xc5, 0x1f, 0x60, 0x79, 0xde, 0x52, 0xa1, 0x74, 0xf9, 0x49, 0x52, 0xa9,
	0x6a, 0x7e, 0x05, 0x33, 0x2b, 0xc5, 0xc3, 0x47, 0x78, 0xdb, 0xcf, 0x17, 0x3c, 0x7b, 0xf4, 0x86,
	0x17, 0xfe, 0xf7, 0xdd, 0x8a, 0xfd, 0xd8, 0xad, 0xd8, 0xaf, 0xdd, 0x8a, 0xbd, 0x9b, 0x0f, 0xa5,
	0xcf, 0x8c, 0x6d, 0x0f, 0xed, 0xaf, 0xf1, 0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0xbd,
	0x8f, 0x11, 0xce, 0x02, 0x00, 0x00,
}
