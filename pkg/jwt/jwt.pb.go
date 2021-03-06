// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/gbs-jwt/pkg/jwt/jwt.proto

package jwt

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Secret struct {
	Name   string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Claims map[string]string `protobuf:"bytes,2,rep,name=claims,proto3" json:"claims,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Secret) Reset()      { *m = Secret{} }
func (*Secret) ProtoMessage() {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb1916878479de91, []int{0}
}
func (m *Secret) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return m.Size()
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Secret) GetClaims() map[string]string {
	if m != nil {
		return m.Claims
	}
	return nil
}

type SecretList struct {
	Secrets []Secret `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets"`
}

func (m *SecretList) Reset()      { *m = SecretList{} }
func (*SecretList) ProtoMessage() {}
func (*SecretList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb1916878479de91, []int{1}
}
func (m *SecretList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SecretList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SecretList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SecretList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretList.Merge(m, src)
}
func (m *SecretList) XXX_Size() int {
	return m.Size()
}
func (m *SecretList) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretList.DiscardUnknown(m)
}

var xxx_messageInfo_SecretList proto.InternalMessageInfo

func (m *SecretList) GetSecrets() []Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func init() {
	proto.RegisterType((*Secret)(nil), "jwt.Secret")
	golang_proto.RegisterType((*Secret)(nil), "jwt.Secret")
	proto.RegisterMapType((map[string]string)(nil), "jwt.Secret.ClaimsEntry")
	golang_proto.RegisterMapType((map[string]string)(nil), "jwt.Secret.ClaimsEntry")
	proto.RegisterType((*SecretList)(nil), "jwt.SecretList")
	golang_proto.RegisterType((*SecretList)(nil), "jwt.SecretList")
}

func init() {
	proto.RegisterFile("github.com/elojah/gbs-jwt/pkg/jwt/jwt.proto", fileDescriptor_fb1916878479de91)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/gbs-jwt/pkg/jwt/jwt.proto", fileDescriptor_fb1916878479de91)
}

var fileDescriptor_fb1916878479de91 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xfa, 0x40,
	0x18, 0xc6, 0xef, 0x05, 0xfe, 0xfc, 0xe3, 0xb1, 0x98, 0x8b, 0x89, 0x0d, 0xc3, 0x2b, 0x61, 0x22,
	0x21, 0xf4, 0x12, 0x5d, 0xc4, 0x11, 0x75, 0x73, 0xc2, 0x4f, 0xd0, 0x92, 0xb3, 0x50, 0x5a, 0x8e,
	0xb4, 0x57, 0x1b, 0x36, 0x47, 0x47, 0x3f, 0x86, 0x1f, 0xc1, 0x91, 0x91, 0xb1, 0x23, 0x93, 0xb1,
	0xd7, 0xc5, 0x91, 0xd1, 0xd1, 0xf4, 0x8a, 0x49, 0xa7, 0x7b, 0x7e, 0xef, 0xfb, 0x3e, 0xcf, 0x93,
	0x1c, 0x1d, 0x7a, 0x0b, 0x35, 0x4f, 0x5c, 0x7b, 0x26, 0x43, 0x2e, 0x02, 0xe9, 0x3b, 0x73, 0xee,
	0xb9, 0xf1, 0xc8, 0x4f, 0x15, 0x5f, 0x2f, 0x3d, 0x5e, 0xbe, 0x7e, 0xaa, 0xec, 0x75, 0x24, 0x95,
	0x64, 0x4d, 0x3f, 0x55, 0xdd, 0x51, 0xcd, 0xe1, 0x49, 0x4f, 0x72, 0xb3, 0x73, 0x93, 0x27, 0x43,
	0x06, 0x8c, 0xaa, 0x3c, 0xfd, 0x57, 0xa0, 0xed, 0x47, 0x31, 0x8b, 0x84, 0x62, 0x8c, 0xb6, 0x56,
	0x4e, 0x28, 0x2c, 0xe8, 0xc1, 0xe0, 0x64, 0x6a, 0x34, 0xe3, 0xb4, 0x3d, 0x0b, 0x9c, 0x45, 0x18,
	0x5b, 0x8d, 0x5e, 0x73, 0xd0, 0xb9, 0x3c, 0xb7, 0xcb, 0xba, 0xca, 0x60, 0xdf, 0x9a, 0xcd, 0xfd,
	0x4a, 0x45, 0x9b, 0xe9, 0xf1, 0xac, 0x3b, 0xa6, 0x9d, 0xda, 0x98, 0x9d, 0xd2, 0xe6, 0x52, 0x6c,
	0x8e, 0x91, 0xa5, 0x64, 0x67, 0xf4, 0xdf, 0xb3, 0x13, 0x24, 0xc2, 0x6a, 0x98, 0x59, 0x05, 0x37,
	0x8d, 0x6b, 0xe8, 0x8f, 0x29, 0xad, 0x82, 0x1f, 0x16, 0xb1, 0x62, 0x43, 0xfa, 0x3f, 0x36, 0x14,
	0x5b, 0x60, 0xaa, 0x3b, 0xb5, 0xea, 0x49, 0x6b, 0xf7, 0x79, 0x41, 0xa6, 0x7f, 0x17, 0x93, 0xbb,
	0x2c, 0x47, 0xb2, 0xcf, 0x91, 0x1c, 0x72, 0x84, 0x9f, 0x1c, 0xe1, 0x45, 0x23, 0xbc, 0x6b, 0x84,
	0x0f, 0x8d, 0xb0, 0xd5, 0x08, 0x3b, 0x8d, 0x90, 0x69, 0x84, 0x2f, 0x8d, 0xf0, 0xad, 0x91, 0x1c,
	0x34, 0xc2, 0x5b, 0x81, 0x64, 0x5b, 0x20, 0x64, 0x05, 0x92, 0x7d, 0x81, 0xc4, 0x6d, 0x9b, 0x2f,
	0xb9, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x49, 0x43, 0x98, 0xe5, 0x75, 0x01, 0x00, 0x00,
}

func (this *Secret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret)
	if !ok {
		that2, ok := that.(Secret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if len(this.Claims) != len(that1.Claims) {
		return false
	}
	for i := range this.Claims {
		if this.Claims[i] != that1.Claims[i] {
			return false
		}
	}
	return true
}
func (this *SecretList) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SecretList)
	if !ok {
		that2, ok := that.(SecretList)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Secrets) != len(that1.Secrets) {
		return false
	}
	for i := range this.Secrets {
		if !this.Secrets[i].Equal(&that1.Secrets[i]) {
			return false
		}
	}
	return true
}
func (this *Secret) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&jwt.Secret{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	keysForClaims := make([]string, 0, len(this.Claims))
	for k, _ := range this.Claims {
		keysForClaims = append(keysForClaims, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForClaims)
	mapStringForClaims := "map[string]string{"
	for _, k := range keysForClaims {
		mapStringForClaims += fmt.Sprintf("%#v: %#v,", k, this.Claims[k])
	}
	mapStringForClaims += "}"
	if this.Claims != nil {
		s = append(s, "Claims: "+mapStringForClaims+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SecretList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&jwt.SecretList{")
	if this.Secrets != nil {
		vs := make([]Secret, len(this.Secrets))
		for i := range vs {
			vs[i] = this.Secrets[i]
		}
		s = append(s, "Secrets: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringJwt(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Secret) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Secret) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Secret) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Claims) > 0 {
		for k := range m.Claims {
			v := m.Claims[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintJwt(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintJwt(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintJwt(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SecretList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SecretList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SecretList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Secrets) > 0 {
		for iNdEx := len(m.Secrets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Secrets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintJwt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintJwt(dAtA []byte, offset int, v uint64) int {
	offset -= sovJwt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedSecret(r randyJwt, easy bool) *Secret {
	this := &Secret{}
	this.Name = string(randStringJwt(r))
	if r.Intn(5) != 0 {
		v1 := r.Intn(10)
		this.Claims = make(map[string]string)
		for i := 0; i < v1; i++ {
			this.Claims[randStringJwt(r)] = randStringJwt(r)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSecretList(r randyJwt, easy bool) *SecretList {
	this := &SecretList{}
	if r.Intn(5) != 0 {
		v2 := r.Intn(5)
		this.Secrets = make([]Secret, v2)
		for i := 0; i < v2; i++ {
			v3 := NewPopulatedSecret(r, easy)
			this.Secrets[i] = *v3
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyJwt interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneJwt(r randyJwt) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringJwt(r randyJwt) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneJwt(r)
	}
	return string(tmps)
}
func randUnrecognizedJwt(r randyJwt, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldJwt(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldJwt(dAtA []byte, r randyJwt, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateJwt(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateJwt(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateJwt(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateJwt(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateJwt(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateJwt(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateJwt(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Secret) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if len(m.Claims) > 0 {
		for k, v := range m.Claims {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovJwt(uint64(len(k))) + 1 + len(v) + sovJwt(uint64(len(v)))
			n += mapEntrySize + 1 + sovJwt(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *SecretList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Secrets) > 0 {
		for _, e := range m.Secrets {
			l = e.Size()
			n += 1 + l + sovJwt(uint64(l))
		}
	}
	return n
}

func sovJwt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJwt(x uint64) (n int) {
	return sovJwt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Secret) String() string {
	if this == nil {
		return "nil"
	}
	keysForClaims := make([]string, 0, len(this.Claims))
	for k, _ := range this.Claims {
		keysForClaims = append(keysForClaims, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForClaims)
	mapStringForClaims := "map[string]string{"
	for _, k := range keysForClaims {
		mapStringForClaims += fmt.Sprintf("%v: %v,", k, this.Claims[k])
	}
	mapStringForClaims += "}"
	s := strings.Join([]string{`&Secret{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Claims:` + mapStringForClaims + `,`,
		`}`,
	}, "")
	return s
}
func (this *SecretList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForSecrets := "[]Secret{"
	for _, f := range this.Secrets {
		repeatedStringForSecrets += strings.Replace(strings.Replace(f.String(), "Secret", "Secret", 1), `&`, ``, 1) + ","
	}
	repeatedStringForSecrets += "}"
	s := strings.Join([]string{`&SecretList{`,
		`Secrets:` + repeatedStringForSecrets + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringJwt(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Secret) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJwt
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Secret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Secret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claims", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claims == nil {
				m.Claims = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowJwt
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowJwt
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthJwt
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthJwt
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowJwt
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthJwt
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthJwt
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipJwt(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthJwt
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Claims[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJwt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJwt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthJwt
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
func (m *SecretList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJwt
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SecretList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecretList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secrets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secrets = append(m.Secrets, Secret{})
			if err := m.Secrets[len(m.Secrets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJwt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJwt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthJwt
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
func skipJwt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJwt
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
					return 0, ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJwt
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
			if length < 0 {
				return 0, ErrInvalidLengthJwt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupJwt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthJwt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthJwt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJwt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupJwt = fmt.Errorf("proto: unexpected end of group")
)
