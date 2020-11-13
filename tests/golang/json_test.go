package main

import (
   "koinos"
   "testing"
   "bytes"
   "encoding/json"
)

func TestBooleanJson(t *testing.T) {
   value := koinos.Boolean(true)
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.Boolean
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if result != value || result != koinos.Boolean(true) {
      t.Errorf("The resulting values are unequal")
   }
}

func TestInt8Json(t *testing.T) {
   value := koinos.Int8(-127)
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.Int8
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if result != value || result != koinos.Int8(-127) {
      t.Errorf("The resulting values are unequal")
   }
}

func TestUInt8Json(t *testing.T) {
   value := koinos.UInt8(255)
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.UInt8
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if result != value || result != koinos.UInt8(255) {
      t.Errorf("The resulting values are unequal")
   }
}

func TestUInt16Json(t *testing.T) {
   value := koinos.UInt16(65535)
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.UInt16
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if result != value || result != koinos.UInt16(65535) {
      t.Errorf("The resulting values are unequal")
   }
}

func TestInt16Json(t *testing.T) {
   value := koinos.Int16(-32768)
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.Int16
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if result != value || result != koinos.Int16(-32768) {
      t.Errorf("The resulting values are unequal")
   }
}

func TestInt32Json(t *testing.T) {
   value := koinos.Int32(-2147483648)
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.Int32
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if result != value || result != koinos.Int32(-2147483648) {
      t.Errorf("The resulting values are unequal")
   }
}

func TestUInt32Json(t *testing.T) {
   value := koinos.UInt32(4294967295)
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.UInt32
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if result != value || result != koinos.UInt32(4294967295) {
      t.Errorf("The resulting values are unequal")
   }
}

func TestInt64Json(t *testing.T) {
   value := koinos.Int64(-9223372036854775808)
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.Int64
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if result != value || result != koinos.Int64(-9223372036854775808) {
      t.Errorf("The resulting values are unequal")
   }
}

func TestUInt64Json(t *testing.T) {
   value := koinos.UInt64(18446744073709551615)
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.UInt64
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if result != value || result != koinos.UInt64(18446744073709551615) {
      t.Errorf("The resulting values are unequal")
   }
}

func TestInt128Json(t *testing.T) {
   value := koinos.NewInt128("-170141183460469231731687303715884105728")
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.Int128
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if (*value).Value.Cmp(&result.Value) != 0 || koinos.NewInt128("-170141183460469231731687303715884105728").Value.Cmp(&result.Value) != 0 {
      t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
   }
}

func TestUInt128Json(t *testing.T) {
   value := koinos.NewUInt128("340282366920938463463374607431768211455")
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.UInt128
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if (*value).Value.Cmp(&result.Value) != 0 || koinos.NewUInt128("340282366920938463463374607431768211455").Value.Cmp(&result.Value) != 0 {
      t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
   }
}

func TestInt160Json(t *testing.T) {
   value := koinos.NewInt160("-730750818665451459101842416358141509827966271488")
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.Int160
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if (*value).Value.Cmp(&result.Value) != 0 || koinos.NewInt160("-730750818665451459101842416358141509827966271488").Value.Cmp(&result.Value) != 0 {
      t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
   }
}

func TestUInt160Json(t *testing.T) {
   value := koinos.NewUInt160("1461501637330902918203684832716283019655932542975")
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.UInt160
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if (*value).Value.Cmp(&result.Value) != 0 || koinos.NewUInt160("1461501637330902918203684832716283019655932542975").Value.Cmp(&result.Value) != 0 {
      t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
   }
}

func TestInt256Json(t *testing.T) {
   value := koinos.NewInt256("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.Int256
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if (*value).Value.Cmp(&result.Value) != 0 || koinos.NewInt256("-57896044618658097711785492504343953926634992332820282019728792003956564819968").Value.Cmp(&result.Value) != 0 {
      t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
   }
}

func TestUInt256Json(t *testing.T) {
   value := koinos.NewUInt256("115792089237316195423570985008687907853269984665640564039457584007913129639935")
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.UInt256
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if (*value).Value.Cmp(&result.Value) != 0 || koinos.NewUInt256("115792089237316195423570985008687907853269984665640564039457584007913129639935").Value.Cmp(&result.Value) != 0 {
      t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
   }
}

func TestMultihashJson(t *testing.T) {
   value := koinos.Multihash{Id: 1, Digest: koinos.VariableBlob{0x01, 0x02, 0x03, 0x04}}
   b, err := json.Marshal(&value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.Multihash
   err = json.Unmarshal(b, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   expected := []byte{0x01, 0x02, 0x03, 0x04}
   if !value.Equals(&result) || result.Id != 1 || !bytes.Equal(result.Digest, expected) {
      t.Errorf("The resulting values are unequal")
   }
}

func TestMultihashVectorJson(t *testing.T) {
   variableBlob := koinos.NewVariableBlob()
   *variableBlob = append(*variableBlob, 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A)
   variableBlob2 := koinos.NewVariableBlob()
   *variableBlob2 = append(*variableBlob2, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06)
   var value koinos.MultihashVector
   value.Id = 1
   value.Digests = append(value.Digests, *variableBlob)
   value.Digests = append(value.Digests, *variableBlob2)
   b, err := json.Marshal(&value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.MultihashVector
   err = json.Unmarshal(b, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }

   if result.Id != value.Id || !bytes.Equal(result.Digests[0], value.Digests[0]) || !bytes.Equal(result.Digests[1], value.Digests[1]) {
      t.Errorf("The resulting values are unequal")
   }
}

func TestFixedBlobJson(t *testing.T) {
   expected := [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10,
      0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20}
   value := koinos.FixedBlob20(expected)
   b, err := json.Marshal(&value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.FixedBlob20
   err = json.Unmarshal(b, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if !bytes.Equal(result[:], expected[:]) {
      t.Errorf("The resulting values are unequal (% x != % x)", result, value)
   }
}

func TestStringJson(t *testing.T) {
   value := koinos.String("alice bob charlie")
   bytes, err := json.Marshal(value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.String
   err = json.Unmarshal(bytes, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if value != result {
      t.Errorf("The resulting values are unequal (%s != %s)", result, value)
   }
}

func TestVariableBlobJson(t *testing.T) {
   value := koinos.VariableBlob{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}
   b, err := json.Marshal(&value)
   if err != nil {
      t.Errorf("An error occurred while encoding to JSON")
   }
   var result koinos.VariableBlob
   err = json.Unmarshal(b, &result)
   if err != nil {
      t.Errorf("An error occurred while decoding from JSON")
   }
   if !bytes.Equal(value, result) {
      t.Errorf("The resulting values are unequal (%x != %x)", result, value)
   }
}
