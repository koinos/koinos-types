package main

import (
	"bytes"
	"encoding/json"
	"koinos"
	"testing"
)

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
