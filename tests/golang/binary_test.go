package main

import (
   "koinos"
   "testing"
   "bytes"
)

func TestBoolean(t *testing.T) {
   value := koinos.Boolean(true)
   expected := []byte{0x01}
   result := koinos.NewVariableBlob()
   result = value.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, value2, err := koinos.DeserializeBoolean(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *value2 != true {
      t.Errorf("*integer2 != true (%t != true)", *value2)
   }
   if size != 1 {
      t.Errorf("size != 1 (%d != 1)", size)
   }

   result = &koinos.VariableBlob{0x00}
   size, value2, err = koinos.DeserializeBoolean(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *value2 != false {
      t.Errorf("*integer2 != true (%t != false)", *value2)
   }
   if size != 1 {
      t.Errorf("size != 1 (%d != 1)", size)
   }

   vb := koinos.VariableBlob{0x02}
   _, _, err = koinos.DeserializeInt64(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }

   vb = koinos.VariableBlob{}
   _, _, err = koinos.DeserializeInt64(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt8(t *testing.T) {
   integer := koinos.Int8(-128)
   expected := []byte{0x80}
   result := koinos.NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := koinos.DeserializeInt8(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != koinos.Int8(-128) {
      t.Errorf("*integer2 != Int8(-128) (%d != %d)", *integer2, koinos.Int8(-128))
   }
   if size != 1 {
      t.Errorf("size != 1 (%d != 1)", size)
   }

   result = koinos.NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := koinos.VariableBlob{}
   _, _, err = koinos.DeserializeInt8(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt8(t *testing.T) {
   integer := koinos.UInt8(255)
   expected := []byte{0xFF}
   result := koinos.NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := koinos.DeserializeUInt8(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != koinos.UInt8(255) {
      t.Errorf("*integer2 != UInt8(255) (%d != %d)", *integer2, koinos.UInt8(255))
   }
   if size != 1 {
      t.Errorf("size != 1 (%d != 1)", size)
   }

   result = koinos.NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := koinos.VariableBlob{}
   _, _, err = koinos.DeserializeUInt8(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt16(t *testing.T) {
   integer := koinos.Int16(-32768)
   expected := []byte{0x80, 0x00}
   result := koinos.NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := koinos.DeserializeInt16(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != koinos.Int16(-32768) {
      t.Errorf("*integer2 != Int16(-32768) (%d != %d)", *integer2, koinos.Int16(-32768))
   }
   if size != 2 {
      t.Errorf("size != 2 (%d != 2)", size)
   }

   result = koinos.NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := koinos.VariableBlob{0x08}
   _, _, err = koinos.DeserializeInt16(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt16(t *testing.T) {
   integer := koinos.UInt16(65535)
   expected := []byte{0xFF, 0xFF}
   result := koinos.NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := koinos.DeserializeUInt16(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != koinos.UInt16(65535) {
      t.Errorf("*integer2 != UInt32(65535) (%d != %d)", *integer2, koinos.UInt16(65535))
   }
   if size != 2 {
      t.Errorf("size != 2 (%d != 2)", size)
   }

   result = koinos.NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := koinos.VariableBlob{0xFF}
   _, _, err = koinos.DeserializeUInt16(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt32(t *testing.T) {
   integer := koinos.Int32(-2147483648)
   expected := []byte{0x80, 0x00, 0x00, 0x00}
   result := koinos.NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := koinos.DeserializeInt32(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != koinos.Int32(-2147483648) {
      t.Errorf("*integer2 != Int32(-2147483648) (%d != %d)", *integer2, koinos.Int32(-2147483648))
   }
   if size != 4 {
      t.Errorf("size != 4 (%d != 4)", size)
   }

   result = koinos.NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := koinos.VariableBlob{0x08, 0x00, 0x00}
   _, _, err = koinos.DeserializeInt32(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt32(t *testing.T) {
   integer := koinos.UInt32(4294967295)
   expected := []byte{0xFF, 0xFF, 0xFF, 0xFF}
   result := koinos.NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := koinos.DeserializeUInt32(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != koinos.UInt32(4294967295) {
      t.Errorf("*integer2 != UInt32(4294967295) (%d != %d)", *integer2, koinos.UInt32(4294967295))
   }
   if size != 4 {
      t.Errorf("size != 4 (%d != 4)", size)
   }

   result = koinos.NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := koinos.VariableBlob{0xFF, 0xFF, 0xFF}
   _, _, err = koinos.DeserializeUInt32(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt64(t *testing.T) {
   integer := koinos.Int64(-256)
   expected := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00}
   result := koinos.NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := koinos.DeserializeInt64(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != koinos.Int64(-256) {
      t.Errorf("*integer2 != Int64(-256) (%d != %d)", *integer2, koinos.Int64(-256))
   }
   if size != 8 {
      t.Errorf("size != 8 (%d != 8)", size)
   }

   result = koinos.NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := koinos.VariableBlob{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
   _, _, err = koinos.DeserializeInt64(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt64(t *testing.T) {
   integer := koinos.UInt64(18446744073709551615)
   expected := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
   result := koinos.NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
   size, integer2, err := koinos.DeserializeUInt64(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *integer2 != koinos.UInt64(18446744073709551615) {
      t.Errorf("*integer2 != UInt64(18446744073709551615) (%d != %d)", *integer2, koinos.UInt64(18446744073709551615))
   }
   if size != 8 {
      t.Errorf("size != 8 (%d != 8)", size)
   }

   result = koinos.NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := koinos.VariableBlob{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
   _, _, err = koinos.DeserializeUInt64(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt128(t *testing.T) {
   integer := koinos.NewInt128FromString("-170141183460469231731687303715884105728")
   expected := []byte{
      0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
   }
   result := koinos.NewVariableBlob()
   result = integer.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d)", result)
   }
   size, integer2, err := koinos.DeserializeInt128(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }

   if koinos.NewInt128FromString("-170141183460469231731687303715884105728").Value.Cmp(&integer2.Value) != 0 {
      t.Errorf("*integer2 != Int128(-170141183460469231731687303715884105728) (%s != %s)", (*integer2).Value.String(), koinos.NewInt128FromString("-170141183460469231731687303715884105728").Value.String())
   }
   if size != 16 {
      t.Errorf("size != 16 (%d != 16)", size)
   }

   result = koinos.NewVariableBlob()
   result = integer2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := koinos.VariableBlob{
      0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = koinos.DeserializeInt128(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt128(t *testing.T) {
   toBin := koinos.NewUInt128FromString("36893488147419103231")
   result := koinos.NewVariableBlob()
   result = toBin.Serialize(result)

   expected := []byte{
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   _, fromBin, err := koinos.DeserializeUInt128(result)
   if err != nil {
      t.Errorf("err != nil")
   }
   if toBin.Value.Cmp(&fromBin.Value) != 0 {
      t.Errorf("toBin != fromBin")
   }

   vb := koinos.VariableBlob{
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = koinos.DeserializeUInt128(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt160(t *testing.T) {
   toBin := koinos.NewInt160FromString("-730750818665451459101842416358141509827966271488")
   result := koinos.NewVariableBlob()
   result = toBin.Serialize(result)

   expected := []byte{
      0x80, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
   }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }

   _, fromBin, err := koinos.DeserializeInt160(result)
   if err != nil {
      t.Errorf("err != nil")
   }
   if toBin.Value.Cmp(&fromBin.Value) != 0 {
      t.Errorf("toBin != fromBin")
   }

   vb := koinos.VariableBlob{
      0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = koinos.DeserializeInt160(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt160(t *testing.T) {
   toBin := koinos.NewUInt160FromString("680564733841876926926749214863536422911")
   result := koinos.NewVariableBlob()
   result = toBin.Serialize(result)

   expected := []byte{
      0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }

   _, fromBin, err := koinos.DeserializeUInt160(result)
   if err != nil {
      t.Errorf("err != nil")
   }
   if toBin.Value.Cmp(&fromBin.Value) != 0 {
      t.Errorf("toBin != fromBin")
   }

   vb := koinos.VariableBlob{
      0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = koinos.DeserializeUInt160(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestInt256(t *testing.T) {
   toBin := koinos.NewInt256FromString("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
   result := koinos.NewVariableBlob()
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

   _, fromBin, err := koinos.DeserializeInt256(result)
   if err != nil {
      t.Errorf("err != nil")
   }
   if toBin.Value.Cmp(&fromBin.Value) != 0 {
      t.Errorf("toBin != fromBin (%s != %s)", toBin.Value.String(), fromBin.Value.String())
   }

   vb := koinos.VariableBlob{
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = koinos.DeserializeInt256(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestUInt256(t *testing.T) {
   toBin := koinos.NewUInt256FromString("680564733841876926926749214863536422911")
   result := koinos.NewVariableBlob()
   result = toBin.Serialize(result)

   expected := []byte{
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }

   _, fromBin, err := koinos.DeserializeUInt256(result)
   if err != nil {
      t.Errorf("err != nil")
   }
   if toBin.Value.Cmp(&fromBin.Value) != 0 {
      t.Errorf("toBin != fromBin")
   }

   vb := koinos.VariableBlob{
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
      0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
   }
   _, _, err = koinos.DeserializeUInt256(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestMultihash(t *testing.T) {
   m := koinos.Multihash{Id: 1, Digest: koinos.VariableBlob{ 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A }}
   result := koinos.NewVariableBlob()
   result = m.Serialize(result)

   expected := []byte{ 0x01, 0x06, 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A }

   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected")
   }
}

func TestMultihashVector(t *testing.T) {
   variableBlob := koinos.NewVariableBlob()
   *variableBlob = append(*variableBlob, 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A)
   variableBlob2 := koinos.NewVariableBlob()
   *variableBlob2 = append(*variableBlob2, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06)
   var multihashVector koinos.MultihashVector
   multihashVector.Id = 1
   multihashVector.Digests = append(multihashVector.Digests, *variableBlob)
   multihashVector.Digests = append(multihashVector.Digests, *variableBlob2)

   result := koinos.NewVariableBlob()
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

   failBytes := []byte{
      0x01,                               // hash_id
      0x06,                               // hash length
      0x02,                               // num hashes
      0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A, // digest_a
      0x01, 0x02, 0x03, 0x04, 0x05,       // digest_b
   }
   failBlob := koinos.NewVariableBlob()
   *failBlob = append(*failBlob, failBytes...)

   _, _, err := koinos.DeserializeMultihashVector(failBlob)
   if err == nil {
      t.Errorf("Expected multihash vector size mismatch")
   }

   defer func() {
      if r := recover(); r == nil {
         t.Errorf("Expected panic on mismatching multihash vector size")
      } else {
         if r != "Multihash vector size mismatch" {
            t.Errorf("Expected panic on mismatching multihash vector size, rather than: %s", r)
         }
      }
   }()
   vblobFail := koinos.NewVariableBlob()
   *vblobFail = append(*vblobFail, 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A)
   vblobFail2 := koinos.NewVariableBlob()
   *vblobFail2 = append(*vblobFail2, 0x01, 0x02, 0x03, 0x04, 0x05)
   var mhvFail koinos.MultihashVector
   mhvFail.Id = 1
   mhvFail.Digests = append(mhvFail.Digests, *vblobFail)
   mhvFail.Digests = append(mhvFail.Digests, *vblobFail2)

   failSerialize := koinos.NewVariableBlob()
   mhvFail.Serialize(failSerialize)
}

func TestFixedBlob(t *testing.T) {
   expected := [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
      0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
   fixedBlob := koinos.FixedBlob20(expected)
   vblob := koinos.NewVariableBlob()
   vblob = fixedBlob.Serialize(vblob)

   if !bytes.Equal(*vblob, expected[:]) {
      t.Errorf("*vblob != expected")
   }

   vb := koinos.VariableBlob{
      0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
      0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
   }
   _, _, err := koinos.DeserializeFixedBlob20(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }
}

func TestString(t *testing.T) {
   msg := koinos.String("Hello World!")
   expected := []byte{0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20,
      0x57, 0x6f, 0x72, 0x6c, 0x64, 0x21}
   result := koinos.NewVariableBlob()
   result = msg.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }
   size, msg2, err := koinos.DeserializeString(result)
   if err != nil {
      t.Errorf("err != nil (%s)", err)
   }
   if *msg2 != msg {
      t.Errorf("*msg2 != msg (%s != %s)", *msg2, msg)
   }
   if size != 13 {
      t.Errorf("size != 13 (%d != 13)", size)
   }

   result = koinos.NewVariableBlob()
   result = msg2.Serialize(result)
   if !bytes.Equal(*result, expected) {
      t.Errorf("*result != expected (%d != %d)", *result, expected)
   }

   vb := koinos.VariableBlob{0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20,
      0x57, 0x6f, 0x72, 0x6c, 0x64}
   _, _, err = koinos.DeserializeString(&vb)
   if err == nil {
      t.Errorf("err == nil")
   }

   vb = koinos.VariableBlob{0x0c, 0xc7, 0xc0, 0x6c, 0x6c, 0x6f, 0x20,
      0x57, 0x6f, 0x72, 0x6c, 0x64, 0x21}
   _, _, err = koinos.DeserializeString(&vb)
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

   variant := koinos.SystemCallTarget{Value: koinos.UInt64(0)}
   vb := koinos.NewVariableBlob()
   _ = variant.Serialize(vb)
}

func TestVariableBlob(t *testing.T) {
   variableBlob := &koinos.VariableBlob{0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A}
   result := koinos.NewVariableBlob()
   result = variableBlob.Serialize(result)

   expected := []byte{0x06, 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A}

   if !bytes.Equal(*result, expected) {
      t.Errorf("result != expected")
   }
}
