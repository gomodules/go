package jsoniter

import (
	"encoding/base64"
	"reflect"
	"unsafe"
	"github.com/v2pro/plz/reflect2"
)

func createEncoderOfNative(cfg *frozenConfig, prefix string, typ reflect.Type) ValEncoder {
	if typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.Uint8 {
		sliceDecoder := decoderOfSlice(cfg, prefix, typ)
		return &base64Codec{sliceDecoder: sliceDecoder}
	}
	typeName := typ.String()
	kind := typ.Kind()
	switch kind {
	case reflect.String:
		if typeName != "string" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*string)(nil)).Elem())
		}
		return &stringCodec{}
	case reflect.Int:
		if typeName != "int" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*int)(nil)).Elem())
		}
		return &intCodec{}
	case reflect.Int8:
		if typeName != "int8" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*int8)(nil)).Elem())
		}
		return &int8Codec{}
	case reflect.Int16:
		if typeName != "int16" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*int16)(nil)).Elem())
		}
		return &int16Codec{}
	case reflect.Int32:
		if typeName != "int32" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*int32)(nil)).Elem())
		}
		return &int32Codec{}
	case reflect.Int64:
		if typeName != "int64" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*int64)(nil)).Elem())
		}
		return &int64Codec{}
	case reflect.Uint:
		if typeName != "uint" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*uint)(nil)).Elem())
		}
		return &uintCodec{}
	case reflect.Uint8:
		if typeName != "uint8" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*uint8)(nil)).Elem())
		}
		return &uint8Codec{}
	case reflect.Uint16:
		if typeName != "uint16" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*uint16)(nil)).Elem())
		}
		return &uint16Codec{}
	case reflect.Uint32:
		if typeName != "uint32" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*uint32)(nil)).Elem())
		}
		return &uint32Codec{}
	case reflect.Uintptr:
		if typeName != "uintptr" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*uintptr)(nil)).Elem())
		}
		return &uintptrCodec{}
	case reflect.Uint64:
		if typeName != "uint64" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*uint64)(nil)).Elem())
		}
		return &uint64Codec{}
	case reflect.Float32:
		if typeName != "float32" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*float32)(nil)).Elem())
		}
		return &float32Codec{}
	case reflect.Float64:
		if typeName != "float64" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*float64)(nil)).Elem())
		}
		return &float64Codec{}
	case reflect.Bool:
		if typeName != "bool" {
			return encoderOfType(cfg, prefix, reflect.TypeOf((*bool)(nil)).Elem())
		}
		return &boolCodec{}
	}
	return nil
}

func createDecoderOfNative(cfg *frozenConfig, prefix string, typ reflect.Type) ValDecoder {
	if typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.Uint8 {
		sliceDecoder := decoderOfSlice(cfg, prefix, typ)
		return &base64Codec{sliceDecoder: sliceDecoder}
	}
	typeName := typ.String()
	switch typ.Kind() {
	case reflect.String:
		if typeName != "string" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*string)(nil)).Elem())
		}
		return &stringCodec{}
	case reflect.Int:
		if typeName != "int" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*int)(nil)).Elem())
		}
		return &intCodec{}
	case reflect.Int8:
		if typeName != "int8" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*int8)(nil)).Elem())
		}
		return &int8Codec{}
	case reflect.Int16:
		if typeName != "int16" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*int16)(nil)).Elem())
		}
		return &int16Codec{}
	case reflect.Int32:
		if typeName != "int32" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*int32)(nil)).Elem())
		}
		return &int32Codec{}
	case reflect.Int64:
		if typeName != "int64" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*int64)(nil)).Elem())
		}
		return &int64Codec{}
	case reflect.Uint:
		if typeName != "uint" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*uint)(nil)).Elem())
		}
		return &uintCodec{}
	case reflect.Uint8:
		if typeName != "uint8" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*uint8)(nil)).Elem())
		}
		return &uint8Codec{}
	case reflect.Uint16:
		if typeName != "uint16" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*uint16)(nil)).Elem())
		}
		return &uint16Codec{}
	case reflect.Uint32:
		if typeName != "uint32" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*uint32)(nil)).Elem())
		}
		return &uint32Codec{}
	case reflect.Uintptr:
		if typeName != "uintptr" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*uintptr)(nil)).Elem())
		}
		return &uintptrCodec{}
	case reflect.Uint64:
		if typeName != "uint64" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*uint64)(nil)).Elem())
		}
		return &uint64Codec{}
	case reflect.Float32:
		if typeName != "float32" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*float32)(nil)).Elem())
		}
		return &float32Codec{}
	case reflect.Float64:
		if typeName != "float64" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*float64)(nil)).Elem())
		}
		return &float64Codec{}
	case reflect.Bool:
		if typeName != "bool" {
			return decoderOfType(cfg, prefix, reflect.TypeOf((*bool)(nil)).Elem())
		}
		return &boolCodec{}
	}
	return nil
}

type stringCodec struct {
}

func (codec *stringCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*((*string)(ptr)) = iter.ReadString()
}

func (codec *stringCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	str := *((*string)(ptr))
	stream.WriteString(str)
}

func (codec *stringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*string)(ptr)) == ""
}

type intCodec struct {
}

func (codec *intCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*int)(ptr)) = iter.ReadInt()
	}
}

func (codec *intCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteInt(*((*int)(ptr)))
}

func (codec *intCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*int)(ptr)) == 0
}

type uintptrCodec struct {
}

func (codec *uintptrCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*uintptr)(ptr)) = uintptr(iter.ReadUint64())
	}
}

func (codec *uintptrCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteUint64(uint64(*((*uintptr)(ptr))))
}

func (codec *uintptrCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*uintptr)(ptr)) == 0
}

type int8Codec struct {
}

func (codec *int8Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*int8)(ptr)) = iter.ReadInt8()
	}
}

func (codec *int8Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteInt8(*((*int8)(ptr)))
}

func (codec *int8Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*int8)(ptr)) == 0
}

type int16Codec struct {
}

func (codec *int16Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*int16)(ptr)) = iter.ReadInt16()
	}
}

func (codec *int16Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteInt16(*((*int16)(ptr)))
}

func (codec *int16Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*int16)(ptr)) == 0
}

type int32Codec struct {
}

func (codec *int32Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*int32)(ptr)) = iter.ReadInt32()
	}
}

func (codec *int32Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteInt32(*((*int32)(ptr)))
}

func (codec *int32Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*int32)(ptr)) == 0
}

type int64Codec struct {
}

func (codec *int64Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*int64)(ptr)) = iter.ReadInt64()
	}
}

func (codec *int64Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteInt64(*((*int64)(ptr)))
}

func (codec *int64Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*int64)(ptr)) == 0
}

type uintCodec struct {
}

func (codec *uintCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*uint)(ptr)) = iter.ReadUint()
		return
	}
}

func (codec *uintCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteUint(*((*uint)(ptr)))
}

func (codec *uintCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*uint)(ptr)) == 0
}

type uint8Codec struct {
}

func (codec *uint8Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*uint8)(ptr)) = iter.ReadUint8()
	}
}

func (codec *uint8Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteUint8(*((*uint8)(ptr)))
}

func (codec *uint8Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*uint8)(ptr)) == 0
}

type uint16Codec struct {
}

func (codec *uint16Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*uint16)(ptr)) = iter.ReadUint16()
	}
}

func (codec *uint16Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteUint16(*((*uint16)(ptr)))
}

func (codec *uint16Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*uint16)(ptr)) == 0
}

type uint32Codec struct {
}

func (codec *uint32Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*uint32)(ptr)) = iter.ReadUint32()
	}
}

func (codec *uint32Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteUint32(*((*uint32)(ptr)))
}

func (codec *uint32Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*uint32)(ptr)) == 0
}

type uint64Codec struct {
}

func (codec *uint64Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*uint64)(ptr)) = iter.ReadUint64()
	}
}

func (codec *uint64Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteUint64(*((*uint64)(ptr)))
}

func (codec *uint64Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*uint64)(ptr)) == 0
}

type float32Codec struct {
}

func (codec *float32Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*float32)(ptr)) = iter.ReadFloat32()
	}
}

func (codec *float32Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteFloat32(*((*float32)(ptr)))
}

func (codec *float32Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*float32)(ptr)) == 0
}

type float64Codec struct {
}

func (codec *float64Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*float64)(ptr)) = iter.ReadFloat64()
	}
}

func (codec *float64Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteFloat64(*((*float64)(ptr)))
}

func (codec *float64Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*float64)(ptr)) == 0
}

type boolCodec struct {
}

func (codec *boolCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if !iter.ReadNil() {
		*((*bool)(ptr)) = iter.ReadBool()
	}
}

func (codec *boolCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteBool(*((*bool)(ptr)))
}

func (codec *boolCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return !(*((*bool)(ptr)))
}

type emptyInterfaceCodec struct {
}

func (codec *emptyInterfaceCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	existing := *((*interface{})(ptr))

	// Checking for both typed and untyped nil pointers.
	if existing != nil &&
		reflect.TypeOf(existing).Kind() == reflect.Ptr &&
		!reflect.ValueOf(existing).IsNil() {

		var ptrToExisting interface{}
		for {
			elem := reflect.ValueOf(existing).Elem()
			if elem.Kind() != reflect.Ptr || elem.IsNil() {
				break
			}
			ptrToExisting = existing
			existing = elem.Interface()
		}

		if iter.ReadNil() {
			if ptrToExisting != nil {
				nilPtr := reflect.Zero(reflect.TypeOf(ptrToExisting).Elem())
				reflect.ValueOf(ptrToExisting).Elem().Set(nilPtr)
			} else {
				*((*interface{})(ptr)) = nil
			}
		} else {
			iter.ReadVal(existing)
		}

		return
	}

	if iter.ReadNil() {
		*((*interface{})(ptr)) = nil
	} else {
		*((*interface{})(ptr)) = iter.Read()
	}
}

func (codec *emptyInterfaceCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	obj := *((*interface{})(ptr))
	stream.WriteVal(obj)
}

func (codec *emptyInterfaceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	emptyInterface := (*emptyInterface)(ptr)
	return emptyInterface.typ == nil
}

type nonEmptyInterfaceCodec struct {
}

func (codec *nonEmptyInterfaceCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if iter.WhatIsNext() == NilValue {
		iter.skipFourBytes('n', 'u', 'l', 'l')
		*((*interface{})(ptr)) = nil
		return
	}
	nonEmptyInterface := (*nonEmptyInterface)(ptr)
	if nonEmptyInterface.itab == nil {
		iter.ReportError("read non-empty interface", "do not know which concrete type to decode to")
		return
	}
	var i interface{}
	e := (*emptyInterface)(unsafe.Pointer(&i))
	e.typ = nonEmptyInterface.itab.typ
	e.word = nonEmptyInterface.word
	iter.ReadVal(&i)
	if e.word == nil {
		nonEmptyInterface.itab = nil
	}
	nonEmptyInterface.word = e.word
}

func (codec *nonEmptyInterfaceCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	nonEmptyInterface := (*nonEmptyInterface)(ptr)
	var i interface{}
	if nonEmptyInterface.itab != nil {
		e := (*emptyInterface)(unsafe.Pointer(&i))
		e.typ = nonEmptyInterface.itab.typ
		e.word = nonEmptyInterface.word
	}
	stream.WriteVal(i)
}

func (codec *nonEmptyInterfaceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	nonEmptyInterface := (*nonEmptyInterface)(ptr)
	return nonEmptyInterface.word == nil
}

type dynamicEncoder struct {
	valType reflect2.Type
}

func (encoder *dynamicEncoder) Encode(ptr unsafe.Pointer, stream *Stream) {
	obj := encoder.valType.UnsafeIndirect(ptr)
	stream.WriteVal(obj)
}

func (encoder *dynamicEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	return encoder.valType.UnsafeIndirect(ptr) == nil
}

type base64Codec struct {
	sliceType *reflect2.UnsafeSliceType
	sliceDecoder ValDecoder
}

func (codec *base64Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if iter.ReadNil() {
		codec.sliceType.UnsafeSetNil(ptr)
		return
	}
	switch iter.WhatIsNext() {
	case StringValue:
		encoding := base64.StdEncoding
		src := iter.SkipAndReturnBytes()
		src = src[1: len(src)-1]
		decodedLen := encoding.DecodedLen(len(src))
		dst := make([]byte, decodedLen)
		len, err := encoding.Decode(dst, src)
		if err != nil {
			iter.ReportError("decode base64", err.Error())
		} else {
			dst = dst[:len]
			codec.sliceType.UnsafeSet(ptr, unsafe.Pointer(&dst))
		}
	case ArrayValue:
		codec.sliceDecoder.Decode(ptr, iter)
	default:
		iter.ReportError("base64Codec", "invalid input")
	}
}

func (codec *base64Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	src := *((*[]byte)(ptr))
	if len(src) == 0 {
		stream.WriteNil()
		return
	}
	encoding := base64.StdEncoding
	stream.writeByte('"')
	size := encoding.EncodedLen(len(src))
	buf := make([]byte, size)
	encoding.Encode(buf, src)
	stream.buf = append(stream.buf, buf...)
	stream.writeByte('"')
}

func (codec *base64Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(*((*[]byte)(ptr))) == 0
}