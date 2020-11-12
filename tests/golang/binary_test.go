package main

import (
   "bytes"
   . "koinos"
   "testing"
)

func TestBoolean(t *testing.T) {
   value := Boolean(true)
   expected := []byte{0x01}
   result := NewVariableBlob()
   result = value.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, value2, err := DeserializeBoolean(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *value2 != true {
      t.Errorf("*integer2 != true (%t != true)", *value2)
   }
   if size != 1 {
      t.Errorf("size != 1 (%d != 1)", size)
   }

   result = &VariableBlob{0x00}
   size, value2, err = DeserializeBoolean(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *value2 != false {
      t.Errorf("*integer2 != true (%t != false)", *value2)
   }
   if size != 1 {
      t.Errorf("size != 1 (%d != 1)", size)
   }

   vb := VariableBlob{0x02}
   _, _, err = DeserializeInt64(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }

   vb = VariableBlob{}
   _, _, err = DeserializeInt64(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt8(t *testing.T) {
   integer := Int8(-128)
   expected := []byte{0x80}
   result := NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := DeserializeInt8(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != Int8(-128) {
      t.Errorf("*integer2 != Int8(-128) (%d != %d)", *integer2, Int8(-128))
   }
   if size != 1 {
      t.Errorf("size != 1 (%d != 1)", size)
   }

   result = NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := VariableBlob{}
   _, _, err = DeserializeInt8(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt8(t *testing.T) {
   integer := UInt8(255)
   expected := []byte{0xFF}
   result := NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := DeserializeUInt8(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != UInt8(255) {
      t.Errorf("*integer2 != UInt8(255) (%d != %d)", *integer2, UInt8(255))
   }
   if size != 1 {
      t.Errorf("size != 1 (%d != 1)", size)
   }

   result = NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := VariableBlob{}
   _, _, err = DeserializeUInt8(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt16(t *testing.T) {
   integer := Int16(-32768)
   expected := []byte{0x80, 0x00}
   result := NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := DeserializeInt16(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != Int16(-32768) {
      t.Errorf("*integer2 != Int16(-32768) (%d != %d)", *integer2, Int16(-32768))
   }
   if size != 2 {
      t.Errorf("size != 2 (%d != 2)", size)
   }

   result = NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := VariableBlob{0x08}
   _, _, err = DeserializeInt16(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt16(t *testing.T) {
   integer := UInt16(65535)
   expected := []byte{0xFF, 0xFF}
   result := NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := DeserializeUInt16(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != UInt16(65535) {
      t.Errorf("*integer2 != UInt32(65535) (%d != %d)", *integer2, UInt16(65535))
   }
   if size != 2 {
      t.Errorf("size != 2 (%d != 2)", size)
   }

   result = NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := VariableBlob{0xFF}
   _, _, err = DeserializeUInt16(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt32(t *testing.T) {
   integer := Int32(-2147483648)
   expected := []byte{0x80, 0x00, 0x00, 0x00}
   result := NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := DeserializeInt32(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != Int32(-2147483648) {
      t.Errorf("*integer2 != Int32(-2147483648) (%d != %d)", *integer2, Int32(-2147483648))
   }
   if size != 4 {
      t.Errorf("size != 4 (%d != 4)", size)
   }

   result = NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := VariableBlob{0x08, 0x00, 0x00}
   _, _, err = DeserializeInt32(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt32(t *testing.T) {
   integer := UInt32(4294967295)
   expected := []byte{0xFF, 0xFF, 0xFF, 0xFF}
   result := NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := DeserializeUInt32(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != UInt32(4294967295) {
      t.Errorf("*integer2 != UInt32(4294967295) (%d != %d)", *integer2, UInt32(4294967295))
   }
   if size != 4 {
      t.Errorf("size != 4 (%d != 4)", size)
   }

   result = NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := VariableBlob{0xFF, 0xFF, 0xFF}
   _, _, err = DeserializeUInt32(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt64(t *testing.T) {
   integer := Int64(-256)
   expected := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00}
   result := NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := DeserializeInt64(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != Int64(-256) {
      t.Errorf("*integer2 != Int64(-256) (%d != %d)", *integer2, Int64(-256))
   }
   if size != 8 {
      t.Errorf("size != 8 (%d != 8)", size)
   }

   result = NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := VariableBlob{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
   _, _, err = DeserializeInt64(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt64(t *testing.T) {
   integer := UInt64(18446744073709551615)
   expected := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
   result := NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := DeserializeUInt64(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != UInt64(18446744073709551615) {
      t.Errorf("*integer2 != UInt64(18446744073709551615) (%d != %d)", *integer2, UInt64(18446744073709551615))
   }
   if size != 8 {
      t.Errorf("size != 8 (%d != 8)", size)
   }

   result = NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := VariableBlob{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
   _, _, err = DeserializeUInt64(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt128(t *testing.T) {
   integer := NewInt128("-170141183460469231731687303715884105728")
   expected := []byte{
      0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
   }
   result := NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d)", result)
   }
   size, integer2, err := DeserializeInt128(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }

   if NewInt128("-170141183460469231731687303715884105728").Value.Cmp(&integer2.Value) != 0 {
      t.Errorf("*integer2 != Int128(-170141183460469231731687303715884105728) (%s != %s)", (*integer2).Value.String(), NewInt128("-170141183460469231731687303715884105728").Value.String())
   }
   if size != 16 {
      t.Errorf("size != 16 (%d != 16)", size)
   }

   result = NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := VariableBlob{
      0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = DeserializeInt128(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt128(t *testing.T) {
   to_bin := NewUInt128("36893488147419103231")
   result := NewVariableBlob()
   result = to_bin.Serialize(result)

   expected := []byte{
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   _, from_bin, err := DeserializeUInt128(result)
   if err != nil {
      t.Errorf("err != nil")
   }
   if to_bin.Value.Cmp(&from_bin.Value) != 0 {
      t.Errorf("to_bin != from_bin")
   }

   vb := VariableBlob{
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = DeserializeUInt128(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt160(t *testing.T) {
   toBin := NewInt160("-730750818665451459101842416358141509827966271488")
   result := NewVariableBlob()
   result = toBin.Serialize(result)

   expected := []byte{
      0x80, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
   }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }

   _, fromBin, err := DeserializeInt160(result)
   if err != nil {
      t.Errorf("err != nil")
   }
   if toBin.Value.Cmp(&fromBin.Value) != 0 {
      t.Errorf("to_bin != from_bin")
   }

   vb := VariableBlob{
      0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = DeserializeInt160(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt160(t *testing.T) {
   toBin := NewUInt160("680564733841876926926749214863536422911")
   result := NewVariableBlob()
   result = toBin.Serialize(result)

   expected := []byte{
      0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }

   _, fromBin, err := DeserializeUInt160(result)
   if err != nil {
      t.Errorf("err != nil")
   }
   if toBin.Value.Cmp(&fromBin.Value) != 0 {
      t.Errorf("toBin != fromBin")
   }

   vb := VariableBlob{
      0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = DeserializeUInt160(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt256(t *testing.T) {
   toBin := NewInt256("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
   result := NewVariableBlob()
   result = toBin.Serialize(result)

   expected := []byte{
      0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
   }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }

   _, fromBin, err := DeserializeInt256(result)
   if err != nil {
      t.Errorf("err != nil")
   }
   if toBin.Value.Cmp(&fromBin.Value) != 0 {
      t.Errorf("toBin != fromBin (%s != %s)", toBin.Value.String(), fromBin.Value.String())
   }

   vb := VariableBlob{
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = DeserializeInt256(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt256(t *testing.T) {
   to_bin := NewUInt256("680564733841876926926749214863536422911")
   result := NewVariableBlob()
   result = to_bin.Serialize(result)

   expected := []byte{
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }

   _, from_bin, err := DeserializeUInt256(result)
   if err != nil {
      t.Errorf("err != nil")
   }
   if to_bin.Value.Cmp(&from_bin.Value) != 0 {
      t.Errorf("to_bin != from_bin")
   }

   vb := VariableBlob{
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = DeserializeUInt256(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestMultihash(t *testing.T) {
   m := Multihash{Id: 1, Digest: VariableBlob{ 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A }}
   result := NewVariableBlob()
   result = m.Serialize(result)

   expected := []byte{ 0x01, 0x06, 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
}

func TestMultihashVector(t *testing.T) {
   variable_blob := NewVariableBlob()
   *variable_blob = append(*variable_blob, 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A)
   variable_blob2 := NewVariableBlob()
   *variable_blob2 = append(*variable_blob2, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06)
   var multihashVector MultihashVector
   multihashVector.Id = 1
   multihashVector.Digests = append(multihashVector.Digests, *variable_blob)
   multihashVector.Digests = append(multihashVector.Digests, *variable_blob2)

   result := NewVariableBlob()
   result = multihashVector.Serialize(result)

   expected := []byte{
      0x01,                               // hash_id
      0x06,                               // hash length
      0x02,                               // num hashes
      0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A, // digest_a
      0x01, 0x02, 0x03, 0x04, 0x05, 0x06, // digest_b
   }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
}

func TestFixedBlob(t *testing.T) {
   expected := [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
      0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
   fixed_blob := FixedBlob20(expected)
   vblob := NewVariableBlob()
   vblob = fixed_blob.Serialize(vblob)

   if !bytes.Equal(*vblob, expected[:]) {
      t.Errorf("*vblob != expected")
   }

   vb := VariableBlob{
      0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
      0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
   }
   _, _, err := DeserializeFixedBlob20(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestString(t *testing.T) {
   msg := String("Hello World!")
   expected := []byte{0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20,
      0x57, 0x6f, 0x72, 0x6c, 0x64, 0x21}
   result := NewVariableBlob()
   result = msg.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }
   size, msg2, err := DeserializeString(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *msg2 != msg {
      t.Errorf("*msg2 != msg (%s != %s)", *msg2, msg)
   }
   if size != 13 {
      t.Errorf("size != 13 (%d != 13)", size)
   }

   result = NewVariableBlob()
   result = msg2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := VariableBlob{0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20,
      0x57, 0x6f, 0x72, 0x6c, 0x64}
   _, _, err = DeserializeString(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }

   vb = VariableBlob{0x0c, 0xc7, 0xc0, 0x6c, 0x6c, 0x6f, 0x20,
      0x57, 0x6f, 0x72, 0x6c, 0x64, 0x21}
   _, _, err = DeserializeString(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestVariant(t *testing.T) {
   defer func() {
      if r := recover(); r == nil {
         t.Errorf("Serializing an incorrect variant tag did not panic")
      }
   }()

   variant := SystemCallTarget{Value: UInt64(0)}
   vb := NewVariableBlob()
   _ = variant.Serialize(vb)
}

func TestVariableBlob(t *testing.T) {
   variableBlob := &VariableBlob{0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A}
   result := NewVariableBlob()
   result = variableBlob.Serialize(result)

   expected := []byte{0x06, 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A}

   if !bytes.Equal(*result, expected) {
      t.Errorf("result != expected")
   }
}
