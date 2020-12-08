package koinos_test

import (
	"bytes"
	"encoding/json"
	"github.com/koinos/koinos-types-golang"
	"testing"
)

func TestBoolean(t *testing.T) {
	b := koinos.NewBoolean()
	if *b != false {
		t.Errorf("Boolean default value incorrect. Expected: false")
	}

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
	_, _, err = koinos.DeserializeBoolean(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}

	vb = koinos.VariableBlob{}
	_, _, err = koinos.DeserializeBoolean(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
}

func TestInt8(t *testing.T) {
	i := koinos.NewInt8()
	if *i != 0 {
		t.Errorf("Int8 default value incorrect. Expected: 0")
	}

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
	i := koinos.NewUInt8()
	if *i != 0 {
		t.Errorf("UInt8 default value incorrect. Expected: 0")
	}

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
	i := koinos.NewInt16()
	if *i != 0 {
		t.Errorf("Int16 default value incorrect. Expected: 0")
	}

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
	i := koinos.NewUInt16()
	if *i != 0 {
		t.Errorf("UInt6 default value incorrect. Expected: 0")
	}

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
	i := koinos.NewInt32()
	if *i != 0 {
		t.Errorf("Int32 default value incorrect. Expected: 0")
	}

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
	i := koinos.NewUInt32()
	if *i != 0 {
		t.Errorf("UInt32 default value incorrect. Expected: 0")
	}

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
	i := koinos.NewInt64()
	if *i != 0 {
		t.Errorf("Int64 default value incorrect. Expected: 0")
	}

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
	i := koinos.NewUInt64()
	if *i != 0 {
		t.Errorf("UInt64 default value incorrect. Expected: 0")
	}

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

func TestTimestampType(t *testing.T) {
	timestamp := koinos.NewTimestampType()
	defaultTimestamp := koinos.NewTimestampType()
	*defaultTimestamp = 0
	if *timestamp != *defaultTimestamp {
		t.Errorf("Unexpected default initialization of TimestampType")
	}
	*timestamp = 18446744073709551615
	expected := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	result := koinos.NewVariableBlob()
	result = timestamp.Serialize(result)
	if !bytes.Equal(*result, expected) {
		t.Errorf("*result != expected")
	}
	size, integer2, err := koinos.DeserializeTimestampType(result)
	if err != nil {
		t.Errorf("err != nil (%s)", err)
	}
	if *integer2 != 18446744073709551615 {
		t.Errorf("*integer2 != 18446744073709551615 (%d != %d)", *integer2, uint64(18446744073709551615))
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
	_, _, err = koinos.DeserializeTimestampType(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
}

func TestBlockHeightType(t *testing.T) {
	blockHeight := koinos.NewBlockHeightType()
	defaultBlockHeight := koinos.NewBlockHeightType()
	*defaultBlockHeight = 0
	if *blockHeight != *defaultBlockHeight {
		t.Errorf("Unexpected default initialization of BlockHeightType")
	}
	*blockHeight = 18446744073709551615
	expected := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	result := koinos.NewVariableBlob()
	result = blockHeight.Serialize(result)
	if !bytes.Equal(*result, expected) {
		t.Errorf("*result != expected")
	}
	size, integer2, err := koinos.DeserializeBlockHeightType(result)
	if err != nil {
		t.Errorf("err != nil (%s)", err)
	}
	if *integer2 != 18446744073709551615 {
		t.Errorf("*integer2 != 18446744073709551615 (%d != %d)", *integer2, uint64(18446744073709551615))
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
	_, _, err = koinos.DeserializeBlockHeightType(&vb)
	if err == nil {
		t.Errorf("err == nil")
	}
}

func TestInt128(t *testing.T) {
	integer, _ := koinos.NewInt128FromString("-170141183460469231731687303715884105728")
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

	expectedValue, _ := koinos.NewInt128FromString("-170141183460469231731687303715884105728")

	if expectedValue.Value.Cmp(&integer2.Value) != 0 {
		t.Errorf("*integer2 != Int128(-170141183460469231731687303715884105728) (%s != %s)", (*integer2).Value.String(), expectedValue.Value.String())
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
	toBin, _ := koinos.NewUInt128FromString("36893488147419103231")
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
	toBin, _ := koinos.NewInt160FromString("-730750818665451459101842416358141509827966271488")
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
	toBin, _ := koinos.NewUInt160FromString("680564733841876926926749214863536422911")
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
	toBin, _ := koinos.NewInt256FromString("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
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
	toBin, _ := koinos.NewUInt256FromString("680564733841876926926749214863536422911")
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
	m := koinos.NewMultihash()
	m.ID = 1
	m.Digest = koinos.VariableBlob{0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A}
	result := koinos.NewVariableBlob()
	result = m.Serialize(result)

	expected := []byte{0x01, 0x06, 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A}

	if !bytes.Equal(*result, expected) {
		t.Errorf("*result != expected")
	}

	_, resultMh, e := koinos.DeserializeMultihash(result)
	if e != nil {
		t.Errorf("Failed to deserialize multihash vector")
	}
	if resultMh.ID != m.ID {
		t.Errorf("Multihash vector Id mismatch")
	}
	if !m.Equals(resultMh) {
		t.Errorf("Multihash mismatch")
	}

	_, _, err := koinos.DeserializeMultihash(&koinos.VariableBlob{0x01, 0x06, 0x04, 0x08, 0x0F, 0x10, 0x17})
	if err.Error() != "Unexpected EOF" {
		t.Errorf("Expected an EOF error")
	}

	mA := koinos.NewMultihash()
	mA.ID = 1
	mA.Digest = koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	mB := koinos.NewMultihash()
	mB.ID = 2
	mB.Digest = koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	mC := koinos.NewMultihash()
	mC.ID = 2
	mC.Digest = koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	mD := koinos.NewMultihash()
	mD.ID = 1
	mD.Digest = koinos.VariableBlob{0x00, 0x00, 0x00, 0x00, 0x00}

	if !mD.LessThan(mA) {
		t.Errorf("Multihash LessThan has unexpected results")
	}
	if !mA.LessThan(mB) {
		t.Errorf("Multihash LessThan has unexpected results")
	}
	if mB.LessThan(mA) {
		t.Errorf("Multihash LessThan has unexpected results")
	}
	if !mB.GreaterThan(mA) {
		t.Errorf("Multihash GreaterThan has unexpected results")
	}
	if !mB.Equals(mC) {
		t.Errorf("Multihash Equals has unexpected results")
	}

	_, _, err = koinos.DeserializeMultihash(&koinos.VariableBlob{})
	if err.Error() != "Could not deserialize multihash id" {
		t.Errorf("Expected a failure to deserialize multihash Id")
	}
}

func TestMultihashVector(t *testing.T) {
	variableBlob := koinos.NewVariableBlob()
	*variableBlob = append(*variableBlob, 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A)
	variableBlob2 := koinos.NewVariableBlob()
	*variableBlob2 = append(*variableBlob2, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06)
	multihashVector := koinos.NewMultihashVector()
	multihashVector.ID = 1
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

	_, resultMh, e := koinos.DeserializeMultihashVector(result)
	if e != nil {
		t.Errorf("Failed to deserialize multihash vector")
	}
	if resultMh.ID != multihashVector.ID {
		t.Errorf("Multihash vector Id mismatch")
	}
	for i, s := range multihashVector.Digests {
		if !bytes.Equal(s, resultMh.Digests[i]) {
			t.Errorf("Multihash vector digest mismatch")
		}
	}

	failBytes := []byte{
		0x01,                               // hash_id
		0x06,                               // hash length
		0x02,                               // num hashes
		0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A, // digest_a
		0x01, 0x02, 0x03, 0x04, 0x05, // digest_b
	}
	failBlob := koinos.NewVariableBlob()
	*failBlob = append(*failBlob, failBytes...)

	_, _, err := koinos.DeserializeMultihashVector(failBlob)
	if err.Error() != "Unexpected EOF" {
		t.Errorf("Expected an EOF error")
	}

	_, _, err = koinos.DeserializeMultihashVector(&koinos.VariableBlob{})
	if err.Error() != "Could not deserialize multihash vector id" {
		t.Errorf("Expected an failure deserializing multihash vector Id")
	}

	_, _, err = koinos.DeserializeMultihashVector(&koinos.VariableBlob{0x01})
	if err.Error() != "Could not deserialize multihash vector hash size" {
		t.Errorf("Expected an failure deserializing multihash hash size")
	}

	_, _, err = koinos.DeserializeMultihashVector(&koinos.VariableBlob{0x01, 0x02})
	if err.Error() != "Could not deserialize multihash vector size" {
		t.Errorf("Expected an failure deserializing multihash vector size")
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
	mhvFail.ID = 1
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
	s := koinos.NewString()
	if *s != "" {
		t.Errorf("String default value incorrect. Expected: \"\"")
	}

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

	variableBlob = &koinos.VariableBlob{0x80}
	bytes, result, err := koinos.DeserializeVariableBlob(variableBlob)
	if err == nil {
		t.Errorf("err == nil")
	}
	if bytes != 0 {
		t.Errorf("bytes != 0 (%d != 0)", bytes)
	}
}

func TestBooleanJson(t *testing.T) {
	value := koinos.Boolean(true)
	bytes, err := json.Marshal(value)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	var result koinos.Boolean
	if string(bytes) != "true" {
		t.Errorf("Unexpected JSON output")
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if result != value || result != koinos.Boolean(true) {
		t.Errorf("The resulting values are unequal")
	}
}

func TestInt8Json(t *testing.T) {
	value := koinos.Int8(-128)
	bytes, err := json.Marshal(value)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != "-128" {
		t.Errorf("Unexpected JSON output")
	}
	var result koinos.Int8
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if result != value || result != koinos.Int8(-128) {
		t.Errorf("The resulting values are unequal")
	}
}

func TestUInt8Json(t *testing.T) {
	value := koinos.UInt8(255)
	bytes, err := json.Marshal(value)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != "255" {
		t.Errorf("Unexpected JSON output")
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
	if string(bytes) != "65535" {
		t.Errorf("Unexpected JSON output")
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
	if string(bytes) != "-32768" {
		t.Errorf("Unexpected JSON output")
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
	if string(bytes) != "-2147483648" {
		t.Errorf("Unexpected JSON output")
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
	if string(bytes) != "4294967295" {
		t.Errorf("Unexpected JSON output")
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
	if string(bytes) != "-9223372036854775808" {
		t.Errorf("Unexpected JSON output")
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
	if string(bytes) != "18446744073709551615" {
		t.Errorf("Unexpected JSON output")
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
	value, _ := koinos.NewInt128FromString("-170141183460469231731687303715884105728")
	bytes, err := json.Marshal(value)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != "\"-170141183460469231731687303715884105728\"" {
		t.Errorf("Unexpected JSON output")
	}
	var result koinos.Int128
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	value, _ = koinos.NewInt128FromString("10")
	bytes = []byte("10")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	bytes = []byte("\"foobar\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	// JSON Bounds
	bytes = []byte("9007199254740991")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected := "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740991\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("9007199254740992")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("err == nil")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740992\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("-9007199254740991")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "-9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"-9007199254740991\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "-9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("-9007199254740992")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("err == nil")
	}
	expected = "\"-9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"-9007199254740992\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "\"-9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	// Int64 Bounds
	bytes = []byte("9223372036854775807")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"9223372036854775807\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("9223372036854775808")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"9223372036854775808\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("-9223372036854775808")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"-9223372036854775808\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("-9223372036854775809")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"-9223372036854775809\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	// Int128 Bounds
	bytes = []byte("\"170141183460469231731687303715884105727\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"170141183460469231731687303715884105728\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"-170141183460469231731687303715884105728\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"-170141183460469231731687303715884105729\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}
}

func TestUInt128Json(t *testing.T) {
	value, _ := koinos.NewUInt128FromString("340282366920938463463374607431768211455")
	bytes, err := json.Marshal(value)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != "\"340282366920938463463374607431768211455\"" {
		t.Errorf("Unexpected JSON output")
	}
	var result koinos.UInt128
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	value, _ = koinos.NewUInt128FromString("10")
	bytes = []byte("10")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	bytes = []byte("\"foobar\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	// JSON Bounds
	bytes = []byte("9007199254740991")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected := "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740991\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("9007199254740992")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("err == nil")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740992\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	// UInt64 Bounds
	bytes = []byte("9223372036854775807")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"9223372036854775807\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("9223372036854775808")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"9223372036854775808\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("0")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"0\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("-1")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"-1\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	// Int128 Bounds
	bytes = []byte("\"340282366920938463463374607431768211455\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"340282366920938463463374607431768211456\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}
}

func TestInt160Json(t *testing.T) {
	value, _ := koinos.NewInt160FromString("-730750818665451459101842416358141509827966271488")
	bytes, err := json.Marshal(value)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != "\"-730750818665451459101842416358141509827966271488\"" {
		t.Errorf("Unexpected JSON output")
	}
	var result koinos.Int160
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	value, _ = koinos.NewInt160FromString("10")
	bytes = []byte("10")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	bytes = []byte("\"foobar\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	// JSON Bounds
	bytes = []byte("9007199254740991")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected := "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740991\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("9007199254740992")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("err == nil")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740992\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("-9007199254740991")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "-9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"-9007199254740991\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "-9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("-9007199254740992")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("err == nil")
	}
	expected = "\"-9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"-9007199254740992\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "\"-9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	// Int64 Bounds
	bytes = []byte("9223372036854775807")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"9223372036854775807\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("9223372036854775808")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"9223372036854775808\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("-9223372036854775808")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"-9223372036854775808\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("-9223372036854775809")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"-9223372036854775809\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	// Int128 Bounds
	bytes = []byte("\"730750818665451459101842416358141509827966271487\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"730750818665451459101842416358141509827966271488\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"-730750818665451459101842416358141509827966271488\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"-730750818665451459101842416358141509827966271489\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}
}

func TestUInt160Json(t *testing.T) {
	value, _ := koinos.NewUInt160FromString("1461501637330902918203684832716283019655932542975")
	bytes, err := json.Marshal(value)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != "\"1461501637330902918203684832716283019655932542975\"" {
		t.Errorf("Unexpected JSON output")
	}
	var result koinos.UInt160
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	value, _ = koinos.NewUInt160FromString("10")
	bytes = []byte("10")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	bytes = []byte("\"foobar\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	// JSON Bounds
	bytes = []byte("9007199254740991")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected := "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740991\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("9007199254740992")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("err == nil")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740992\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	// UInt64 Bounds
	bytes = []byte("9223372036854775807")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"9223372036854775807\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("9223372036854775808")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"9223372036854775808\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("0")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"0\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("-1")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"-1\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	// Int128 Bounds
	bytes = []byte("\"1461501637330902918203684832716283019655932542975\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"1461501637330902918203684832716283019655932542976\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}
}

func TestInt256Json(t *testing.T) {
	value, _ := koinos.NewInt256FromString("-57896044618658097711785492504343953926634992332820282019728792003956564819968")
	bytes, err := json.Marshal(value)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != "\"-57896044618658097711785492504343953926634992332820282019728792003956564819968\"" {
		t.Errorf("Unexpected JSON output")
	}
	var result koinos.Int256
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	value, _ = koinos.NewInt256FromString("10")
	bytes = []byte("10")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	bytes = []byte("\"foobar\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	// JSON Bounds
	bytes = []byte("9007199254740991")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected := "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740991\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("9007199254740992")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("err == nil")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740992\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("-9007199254740991")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "-9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"-9007199254740991\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "-9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("-9007199254740992")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("err == nil")
	}
	expected = "\"-9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"-9007199254740992\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "\"-9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	// Int64 Bounds
	bytes = []byte("9223372036854775807")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"9223372036854775807\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("9223372036854775808")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"9223372036854775808\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("-9223372036854775808")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"-9223372036854775808\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("-9223372036854775809")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"-9223372036854775809\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	// Int128 Bounds
	bytes = []byte("\"57896044618658097711785492504343953926634992332820282019728792003956564819967\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"57896044618658097711785492504343953926634992332820282019728792003956564819968\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"-57896044618658097711785492504343953926634992332820282019728792003956564819968\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"-57896044618658097711785492504343953926634992332820282019728792003956564819969\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}
}

func TestUInt256Json(t *testing.T) {
	value, _ := koinos.NewUInt256FromString("115792089237316195423570985008687907853269984665640564039457584007913129639935")
	bytes, err := json.Marshal(value)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != "\"115792089237316195423570985008687907853269984665640564039457584007913129639935\"" {
		t.Errorf("Unexpected JSON output")
	}
	var result koinos.UInt256
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	value, _ = koinos.NewUInt256FromString("10")
	bytes = []byte("10")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if (*value).Value.Cmp(&result.Value) != 0 {
		t.Errorf("The resulting values are unequal (%s != %s)", result.Value.String(), value.Value.String())
	}

	bytes = []byte("\"foobar\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	// JSON Bounds
	bytes = []byte("9007199254740991")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected := "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740991\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "9007199254740991"
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("9007199254740992")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("err == nil")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	bytes = []byte("\"9007199254740992\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected = "\"9007199254740992\""
	bytes, err = json.Marshal(&result)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(bytes) != expected {
		t.Errorf("Incorrect JSON serialization. Expected: %s, Was: %s", expected, string(bytes))
	}

	// UInt64 Bounds
	bytes = []byte("9223372036854775807")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"9223372036854775807\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("9223372036854775808")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"9223372036854775808\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("0")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"0\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("-1")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	bytes = []byte("\"-1\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	// Int128 Bounds
	bytes = []byte("\"115792089237316195423570985008687907853269984665640564039457584007913129639935\"")
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	bytes = []byte("\"115792089237316195423570985008687907853269984665640564039457584007913129639936\"")
	err = json.Unmarshal(bytes, &result)
	if err == nil {
		t.Errorf("err == nil")
	}
}

func TestMultihashJson(t *testing.T) {
	value := koinos.Multihash{ID: 1, Digest: koinos.VariableBlob{0x01, 0x02, 0x03, 0x04}}
	b, err := json.Marshal(&value)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(b) != "{\"hash\":1,\"digest\":\"z2VfUX\"}" {
		t.Errorf("Unexpected JSON output")
	}
	var result koinos.Multihash
	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	expected := []byte{0x01, 0x02, 0x03, 0x04}
	if !value.Equals(&result) || result.ID != 1 || !bytes.Equal(result.Digest, expected) {
		t.Errorf("The resulting values are unequal")
	}
}

func TestMultihashVectorJson(t *testing.T) {
	variableBlob := koinos.NewVariableBlob()
	*variableBlob = append(*variableBlob, 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A)
	variableBlob2 := koinos.NewVariableBlob()
	*variableBlob2 = append(*variableBlob2, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06)
	var value koinos.MultihashVector
	value.ID = 1
	value.Digests = append(value.Digests, *variableBlob)
	value.Digests = append(value.Digests, *variableBlob2)
	b, err := json.Marshal(&value)
	if err != nil {
		t.Errorf("An error occurred while encoding to JSON")
	}
	if string(b) != "{\"hash\":1,\"digests\":[\"z31SRtpx1\",\"zW7LcTy7\"]}" {
		t.Errorf("Unexpected JSON output")
	}
	var result koinos.MultihashVector
	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}

	if result.ID != value.ID || !bytes.Equal(result.Digests[0], value.Digests[0]) || !bytes.Equal(result.Digests[1], value.Digests[1]) {
		t.Errorf("The resulting values are unequal")
	}

	var mhv koinos.MultihashVector
	by := []byte(`{"hash":"1","digests":["z31SRtpx1","zW7yj"]}`)
	err = json.Unmarshal(by, &mhv)
	if err.Error() != "json: cannot unmarshal string into Go struct field .hash of type uint64" {
		t.Errorf("Expected failure parsing hash id")
	}

	by = []byte(`{"hash":1,"digests":["31SRtpx1"]}`)
	err = json.Unmarshal(by, &mhv)
	if err.Error() != "Unknown encoding: 3" {
		t.Errorf("Expected unknown encoding")
	}

	by = []byte(`{"hash":1,"digests":["z31SRtpx1","zW7yj"]}`)

	err = json.Unmarshal(by, &mhv)
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
	mhvFail.ID = 1
	mhvFail.Digests = append(mhvFail.Digests, *vblobFail)
	mhvFail.Digests = append(mhvFail.Digests, *vblobFail2)
	_, e := json.Marshal(&mhvFail)
	if e == nil {
		t.Errorf("Unexpected success with mismatching multihash vector sizes")
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
	if string(b) != "\"zpEbmSWqJdBuR5GAyhdJxAcGKAf\"" {
		t.Errorf("Unexpected JSON output")
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
	if string(bytes) != "\"alice bob charlie\"" {
		t.Errorf("Unexpected JSON output")
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
	if string(b) != "\"zW7LcTy7\"" {
		t.Errorf("Unexpected JSON output")
	}
	var result koinos.VariableBlob
	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if !bytes.Equal(value, result) {
		t.Errorf("The resulting values are unequal (%x != %x)", result, value)
	}

	raw := []byte("10")
	err = json.Unmarshal(raw, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	raw = []byte("\"fo0bar\"")
	err = json.Unmarshal(raw, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	raw = []byte("\"zfo0bar\"")
	err = json.Unmarshal(raw, &result)
	if err == nil {
		t.Errorf("err == nil")
	}

	raw = []byte("\"z\"")
	err = json.Unmarshal(raw, &result)
	expected := make([]byte, 0)
	if err != nil {
		t.Errorf("An error occurred while decoding from JSON")
	}
	if !bytes.Equal(result, expected) {
		t.Errorf("The resulting values are unequal (%x != %x)", result, expected)
	}
}
