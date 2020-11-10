package main

import (
    "bytes"
    "testing"
    . "koinos"
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
      t.Errorf("size != 1 (%d != 1)", size )
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
      t.Errorf("size != 1 (%d != 1)", size )
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

func TestBasic(t *testing.T) {
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
      t.Errorf("size != 8 (%d != 8)", size )
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

func TestUInt160(t *testing.T) {
   to_bin := NewUInt160("680564733841876926926749214863536422911")
   result := NewVariableBlob()
   result = to_bin.Serialize(result)

   expected := []byte{
      0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }

   _, from_bin, err := DeserializeUInt160(result)
   if err != nil {
      t.Errorf("err != nil")
   }
   if to_bin.Value.Cmp(&from_bin.Value) != 0 {
      t.Errorf("to_bin != from_bin")
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

func TestMultihashVector(t *testing.T) {
   multihash_vector := VectorMultihash(make([]Multihash, 0))
   variable_blob := NewVariableBlob()
   *variable_blob = append(*variable_blob, "alice"...)
   variable_blob2 := NewVariableBlob()
   *variable_blob2 = append(*variable_blob2, "bob"...)
   multihash_vector = append(multihash_vector, Multihash{Id: 1, Digest: *variable_blob}, Multihash{Id: 2, Digest: *variable_blob2})

   vblob := NewVariableBlob()
   vblob = multihash_vector.Serialize(vblob)

   expected := []byte{
      0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x05, 0x61, 0x6c, 0x69, 0x63, 0x65,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x03, 0x62, 0x6f, 0x62, }
   if !bytes.Equal(expected, *vblob) {
      t.Errorf("expected != *vblob")
   }

   _, mhv, err := DeserializeVectorMultihash(vblob)
   if err != nil {
      t.Errorf("err != nil")
   }
   for i := 0; i < 2; i++ {
      if multihash_vector[i].Id != (*mhv)[i].Id {
         t.Errorf("Id not equal")
      }

      if !bytes.Equal(multihash_vector[i].Digest, (*mhv)[i].Digest) {
         t.Errorf("Digest not equal")
      }
   }

   vb := VariableBlob{
      0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x05, 0x61, 0x6c, 0x69, 0x63, 0x65,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x03, 0x62, 0x6f,
   }
   _, _, err = DeserializeVectorMultihash(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestFixedBlob(t *testing.T) {
   expected := [20]byte{ 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
                         0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, }
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
                      0x57, 0x6f, 0x72, 0x6c, 0x64, 0x21, }
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
      t.Errorf("size != 13 (%d != 13)", size )
   }

   result = NewVariableBlob()
   result = msg2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := VariableBlob{0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20,
                      0x57, 0x6f, 0x72, 0x6c, 0x64, }
   _, _, err = DeserializeString(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }

   vb = VariableBlob{0x0c, 0xc7, 0xc0, 0x6c, 0x6c, 0x6f, 0x20,
                     0x57, 0x6f, 0x72, 0x6c, 0x64, 0x21, }
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

   variant := SystemCallTarget{Value:UInt64(0)}
   vb := NewVariableBlob()
   _ = variant.Serialize(vb)
}
