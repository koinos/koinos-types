package main

import (
    "bytes"
    "testing"
    . "koinos"
)

func TestBasic(t *testing.T) {
   integer := Int64(-256)
   expected := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00}
   result := NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   _, integer2 := DeserializeInt64(result)
   if *integer2 != Int64(-256) {
      t.Errorf("*integer2 != Int64(-256)")
   }

   result = NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
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
      t.Errorf("*result != expected")
   }

   _, from_bin := DeserializeUInt128(result)
   if to_bin.Value.Cmp(&from_bin.Value) != 0 {
      t.Errorf("to_bin != from_bin")
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

   _, from_bin := DeserializeUInt160(result)
   if to_bin.Value.Cmp(&from_bin.Value) != 0 {
      t.Errorf("to_bin != from_bin")
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

   _, from_bin := DeserializeUInt256(result)
   if to_bin.Value.Cmp(&from_bin.Value) != 0 {
      t.Errorf("to_bin != from_bin")
   }
}


// TODO: This is utterly broken
//       Test against the C serialization
/*func TestMultihashVector(t *testing.T) {
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

   _, mhv := DeserializeVectorMultihash(vblob)
   for i := 0; i < 2; i++ {
      if multihash_vector[i].Id != (*mhv)[i].Id {
         t.Errorf("Id not equal")
      }

      if !bytes.Equal(multihash_vector[i].Digest, (*mhv)[i].Digest) {
         t.Errorf("Digest not equal")
      }
   }
}*/

func TestFixedBlob(t *testing.T) {
   expected := [20]byte{ 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
                         0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, }
   fixed_blob := FixedBlob20(expected)
   vblob := NewVariableBlob()
   vblob = fixed_blob.Serialize(vblob)

   if !bytes.Equal(*vblob, expected[:]) {
      t.Errorf("*vblob != expected")
   }
}

