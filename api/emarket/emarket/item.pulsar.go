// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package emarket

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Item             protoreflect.MessageDescriptor
	fd_Item_id          protoreflect.FieldDescriptor
	fd_Item_name        protoreflect.FieldDescriptor
	fd_Item_productType protoreflect.FieldDescriptor
	fd_Item_amount      protoreflect.FieldDescriptor
	fd_Item_price       protoreflect.FieldDescriptor
	fd_Item_discounted  protoreflect.FieldDescriptor
	fd_Item_creator     protoreflect.FieldDescriptor
	fd_Item_tags        protoreflect.FieldDescriptor
)

func init() {
	file_emarket_emarket_item_proto_init()
	md_Item = File_emarket_emarket_item_proto.Messages().ByName("Item")
	fd_Item_id = md_Item.Fields().ByName("id")
	fd_Item_name = md_Item.Fields().ByName("name")
	fd_Item_productType = md_Item.Fields().ByName("productType")
	fd_Item_amount = md_Item.Fields().ByName("amount")
	fd_Item_price = md_Item.Fields().ByName("price")
	fd_Item_discounted = md_Item.Fields().ByName("discounted")
	fd_Item_creator = md_Item.Fields().ByName("creator")
	fd_Item_tags = md_Item.Fields().ByName("tags")
}

var _ protoreflect.Message = (*fastReflection_Item)(nil)

type fastReflection_Item Item

func (x *Item) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Item)(x)
}

func (x *Item) slowProtoReflect() protoreflect.Message {
	mi := &file_emarket_emarket_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Item_messageType fastReflection_Item_messageType
var _ protoreflect.MessageType = fastReflection_Item_messageType{}

type fastReflection_Item_messageType struct{}

func (x fastReflection_Item_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Item)(nil)
}
func (x fastReflection_Item_messageType) New() protoreflect.Message {
	return new(fastReflection_Item)
}
func (x fastReflection_Item_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Item
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Item) Descriptor() protoreflect.MessageDescriptor {
	return md_Item
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Item) Type() protoreflect.MessageType {
	return _fastReflection_Item_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Item) New() protoreflect.Message {
	return new(fastReflection_Item)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Item) Interface() protoreflect.ProtoMessage {
	return (*Item)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Item) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Item_id, value) {
			return
		}
	}
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_Item_name, value) {
			return
		}
	}
	if x.ProductType != "" {
		value := protoreflect.ValueOfString(x.ProductType)
		if !f(fd_Item_productType, value) {
			return
		}
	}
	if x.Amount != int32(0) {
		value := protoreflect.ValueOfInt32(x.Amount)
		if !f(fd_Item_amount, value) {
			return
		}
	}
	if x.Price != int32(0) {
		value := protoreflect.ValueOfInt32(x.Price)
		if !f(fd_Item_price, value) {
			return
		}
	}
	if x.Discounted != false {
		value := protoreflect.ValueOfBool(x.Discounted)
		if !f(fd_Item_discounted, value) {
			return
		}
	}
	if x.Creator != "" {
		value := protoreflect.ValueOfString(x.Creator)
		if !f(fd_Item_creator, value) {
			return
		}
	}
	if x.Tags != "" {
		value := protoreflect.ValueOfString(x.Tags)
		if !f(fd_Item_tags, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Item) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "emarket.emarket.Item.id":
		return x.Id != uint64(0)
	case "emarket.emarket.Item.name":
		return x.Name != ""
	case "emarket.emarket.Item.productType":
		return x.ProductType != ""
	case "emarket.emarket.Item.amount":
		return x.Amount != int32(0)
	case "emarket.emarket.Item.price":
		return x.Price != int32(0)
	case "emarket.emarket.Item.discounted":
		return x.Discounted != false
	case "emarket.emarket.Item.creator":
		return x.Creator != ""
	case "emarket.emarket.Item.tags":
		return x.Tags != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emarket.emarket.Item"))
		}
		panic(fmt.Errorf("message emarket.emarket.Item does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Item) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "emarket.emarket.Item.id":
		x.Id = uint64(0)
	case "emarket.emarket.Item.name":
		x.Name = ""
	case "emarket.emarket.Item.productType":
		x.ProductType = ""
	case "emarket.emarket.Item.amount":
		x.Amount = int32(0)
	case "emarket.emarket.Item.price":
		x.Price = int32(0)
	case "emarket.emarket.Item.discounted":
		x.Discounted = false
	case "emarket.emarket.Item.creator":
		x.Creator = ""
	case "emarket.emarket.Item.tags":
		x.Tags = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emarket.emarket.Item"))
		}
		panic(fmt.Errorf("message emarket.emarket.Item does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Item) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "emarket.emarket.Item.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "emarket.emarket.Item.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "emarket.emarket.Item.productType":
		value := x.ProductType
		return protoreflect.ValueOfString(value)
	case "emarket.emarket.Item.amount":
		value := x.Amount
		return protoreflect.ValueOfInt32(value)
	case "emarket.emarket.Item.price":
		value := x.Price
		return protoreflect.ValueOfInt32(value)
	case "emarket.emarket.Item.discounted":
		value := x.Discounted
		return protoreflect.ValueOfBool(value)
	case "emarket.emarket.Item.creator":
		value := x.Creator
		return protoreflect.ValueOfString(value)
	case "emarket.emarket.Item.tags":
		value := x.Tags
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emarket.emarket.Item"))
		}
		panic(fmt.Errorf("message emarket.emarket.Item does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Item) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "emarket.emarket.Item.id":
		x.Id = value.Uint()
	case "emarket.emarket.Item.name":
		x.Name = value.Interface().(string)
	case "emarket.emarket.Item.productType":
		x.ProductType = value.Interface().(string)
	case "emarket.emarket.Item.amount":
		x.Amount = int32(value.Int())
	case "emarket.emarket.Item.price":
		x.Price = int32(value.Int())
	case "emarket.emarket.Item.discounted":
		x.Discounted = value.Bool()
	case "emarket.emarket.Item.creator":
		x.Creator = value.Interface().(string)
	case "emarket.emarket.Item.tags":
		x.Tags = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emarket.emarket.Item"))
		}
		panic(fmt.Errorf("message emarket.emarket.Item does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Item) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "emarket.emarket.Item.id":
		panic(fmt.Errorf("field id of message emarket.emarket.Item is not mutable"))
	case "emarket.emarket.Item.name":
		panic(fmt.Errorf("field name of message emarket.emarket.Item is not mutable"))
	case "emarket.emarket.Item.productType":
		panic(fmt.Errorf("field productType of message emarket.emarket.Item is not mutable"))
	case "emarket.emarket.Item.amount":
		panic(fmt.Errorf("field amount of message emarket.emarket.Item is not mutable"))
	case "emarket.emarket.Item.price":
		panic(fmt.Errorf("field price of message emarket.emarket.Item is not mutable"))
	case "emarket.emarket.Item.discounted":
		panic(fmt.Errorf("field discounted of message emarket.emarket.Item is not mutable"))
	case "emarket.emarket.Item.creator":
		panic(fmt.Errorf("field creator of message emarket.emarket.Item is not mutable"))
	case "emarket.emarket.Item.tags":
		panic(fmt.Errorf("field tags of message emarket.emarket.Item is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emarket.emarket.Item"))
		}
		panic(fmt.Errorf("message emarket.emarket.Item does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Item) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "emarket.emarket.Item.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "emarket.emarket.Item.name":
		return protoreflect.ValueOfString("")
	case "emarket.emarket.Item.productType":
		return protoreflect.ValueOfString("")
	case "emarket.emarket.Item.amount":
		return protoreflect.ValueOfInt32(int32(0))
	case "emarket.emarket.Item.price":
		return protoreflect.ValueOfInt32(int32(0))
	case "emarket.emarket.Item.discounted":
		return protoreflect.ValueOfBool(false)
	case "emarket.emarket.Item.creator":
		return protoreflect.ValueOfString("")
	case "emarket.emarket.Item.tags":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emarket.emarket.Item"))
		}
		panic(fmt.Errorf("message emarket.emarket.Item does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Item) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in emarket.emarket.Item", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Item) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Item) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Item) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Item) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Item)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ProductType)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Amount != 0 {
			n += 1 + runtime.Sov(uint64(x.Amount))
		}
		if x.Price != 0 {
			n += 1 + runtime.Sov(uint64(x.Price))
		}
		if x.Discounted {
			n += 2
		}
		l = len(x.Creator)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Tags)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Item)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Tags) > 0 {
			i -= len(x.Tags)
			copy(dAtA[i:], x.Tags)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Tags)))
			i--
			dAtA[i] = 0x42
		}
		if len(x.Creator) > 0 {
			i -= len(x.Creator)
			copy(dAtA[i:], x.Creator)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Creator)))
			i--
			dAtA[i] = 0x3a
		}
		if x.Discounted {
			i--
			if x.Discounted {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x30
		}
		if x.Price != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Price))
			i--
			dAtA[i] = 0x28
		}
		if x.Amount != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Amount))
			i--
			dAtA[i] = 0x20
		}
		if len(x.ProductType) > 0 {
			i -= len(x.ProductType)
			copy(dAtA[i:], x.ProductType)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProductType)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Item)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Item: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Item: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProductType", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ProductType = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				x.Amount = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Amount |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
				}
				x.Price = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Price |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Discounted", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Discounted = bool(v != 0)
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Creator = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Tags = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: emarket/emarket/item.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProductType string `protobuf:"bytes,3,opt,name=productType,proto3" json:"productType,omitempty"`
	Amount      int32  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Price       int32  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Discounted  bool   `protobuf:"varint,6,opt,name=discounted,proto3" json:"discounted,omitempty"`
	Creator     string `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	Tags        string `protobuf:"bytes,8,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emarket_emarket_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_emarket_emarket_item_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetProductType() string {
	if x != nil {
		return x.ProductType
	}
	return ""
}

func (x *Item) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Item) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Item) GetDiscounted() bool {
	if x != nil {
		return x.Discounted
	}
	return false
}

func (x *Item) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Item) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

var File_emarket_emarket_item_proto protoreflect.FileDescriptor

var file_emarket_emarket_item_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x22, 0xc8, 0x01,
	0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x42, 0x9a, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x42, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1b, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2f, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0xa2, 0x02, 0x03, 0x45, 0x45, 0x58,
	0xaa, 0x02, 0x0f, 0x45, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x45, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0xca, 0x02, 0x0f, 0x45, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5c, 0x45, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0xe2, 0x02, 0x1b, 0x45, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5c, 0x45,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x10, 0x45, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x3a, 0x3a, 0x45, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_emarket_emarket_item_proto_rawDescOnce sync.Once
	file_emarket_emarket_item_proto_rawDescData = file_emarket_emarket_item_proto_rawDesc
)

func file_emarket_emarket_item_proto_rawDescGZIP() []byte {
	file_emarket_emarket_item_proto_rawDescOnce.Do(func() {
		file_emarket_emarket_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_emarket_emarket_item_proto_rawDescData)
	})
	return file_emarket_emarket_item_proto_rawDescData
}

var file_emarket_emarket_item_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_emarket_emarket_item_proto_goTypes = []interface{}{
	(*Item)(nil), // 0: emarket.emarket.Item
}
var file_emarket_emarket_item_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_emarket_emarket_item_proto_init() }
func file_emarket_emarket_item_proto_init() {
	if File_emarket_emarket_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_emarket_emarket_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_emarket_emarket_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_emarket_emarket_item_proto_goTypes,
		DependencyIndexes: file_emarket_emarket_item_proto_depIdxs,
		MessageInfos:      file_emarket_emarket_item_proto_msgTypes,
	}.Build()
	File_emarket_emarket_item_proto = out.File
	file_emarket_emarket_item_proto_rawDesc = nil
	file_emarket_emarket_item_proto_goTypes = nil
	file_emarket_emarket_item_proto_depIdxs = nil
}
