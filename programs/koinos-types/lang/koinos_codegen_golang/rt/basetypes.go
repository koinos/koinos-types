package koinos

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"math/big"
	"sync"
	"unicode/utf8"

	"github.com/btcsuite/btcutil/base58"
)

const bigIntNumericLiteralMin int64 = -9007199254740991 // -1 << 53
const bigIntNumericLiteralMax int64 = 9007199254740991  // 1 << 53 - 1

// Serializeable type
type Serializeable interface {
	Serialize(vb *VariableBlob) *VariableBlob
}

// --------------------------------
//  String
// --------------------------------

// String type
type String string

// NewString type
func NewString() *String {
	o := String("")
	return &o
}

// Serialize String
func (n *String) Serialize(vb *VariableBlob) *VariableBlob {
	nvb := VariableBlob(make([]byte, len(*n)))
	copy(nvb, *n)

	return nvb.Serialize(vb)
}

// DeserializeString function
func DeserializeString(vb *VariableBlob) (uint64, *String, error) {
	bytes, vbPtr, err := DeserializeVariableBlob(vb)
	s := String("")
	if err != nil {
		return 0, &s, err
	}

	if !utf8.Valid(*vbPtr) {
		return 0, &s, errors.New("String is not UTF-8 encoded")
	}
	s = String(*vbPtr)

	return bytes, &s, nil
}

// --------------------------------
//  Boolean
// --------------------------------

// Boolean type
type Boolean bool

// NewBoolean factory
func NewBoolean() *Boolean {
	o := Boolean(false)
	return &o
}

// Serialize Boolean
func (n *Boolean) Serialize(vb *VariableBlob) *VariableBlob {
	var b byte
	if *n {
		b = 1
	}
	x := append(*vb, b)
	return &x
}

// DeserializeBoolean function
func DeserializeBoolean(vb *VariableBlob) (uint64, *Boolean, error) {
	var b Boolean

	if len(*vb) < 1 {
		return 0, &b, errors.New("Unexpected EOF")
	}

	if (*vb)[0] == 1 {
		b = true
	} else if (*vb)[0] != 0 {
		return 0, &b, errors.New("Boolean must be 0 or 1")
	}

	return 1, &b, nil
}

// --------------------------------
//  Int8
// --------------------------------

// Int8 type
type Int8 int8

// NewInt8 function
func NewInt8() *Int8 {
	o := Int8(0)
	return &o
}

// Serialize Int8
func (n *Int8) Serialize(vb *VariableBlob) *VariableBlob {
	ov := append(*vb, byte(*n))
	return &ov
}

// DeserializeInt8 function
func DeserializeInt8(vb *VariableBlob) (uint64, *Int8, error) {
	var i Int8

	if len(*vb) < 1 {
		return 0, &i, errors.New("Unexpected EOF")
	}

	i = Int8((*vb)[0])

	return 1, &i, nil
}

// --------------------------------
//  UInt8
// --------------------------------

// UInt8 type
type UInt8 uint8

// NewUInt8 factory
func NewUInt8() *UInt8 {
	o := UInt8(0)
	return &o
}

// Serialize UInt8
func (n *UInt8) Serialize(vb *VariableBlob) *VariableBlob {
	ov := append(*vb, byte(*n))
	return &ov
}

// DeserializeUInt8 function
func DeserializeUInt8(vb *VariableBlob) (uint64, *UInt8, error) {
	var i UInt8

	if len(*vb) < 1 {
		return 0, &i, errors.New("Unexpected EOF")
	}

	i = UInt8((*vb)[0])

	return 1, &i, nil
}

// --------------------------------
//  Int16
// --------------------------------

// Int16 type
type Int16 int16

// NewInt16 factory
func NewInt16() *Int16 {
	o := Int16(0)
	return &o
}

// Serialize Int16
func (n *Int16) Serialize(vb *VariableBlob) *VariableBlob {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(*n))
	ov := append(*vb, b...)
	return &ov
}

// DeserializeInt16 function
func DeserializeInt16(vb *VariableBlob) (uint64, *Int16, error) {
	var i Int16

	if len(*vb) < 2 {
		return 0, &i, errors.New("Unexpected EOF")
	}

	i = Int16(binary.BigEndian.Uint16(*vb))

	return 2, &i, nil
}

// --------------------------------
//  UInt16
// --------------------------------

// UInt16 type
type UInt16 uint16

// NewUInt16 factory
func NewUInt16() *UInt16 {
	o := UInt16(0)
	return &o
}

// Serialize UInt16
func (n *UInt16) Serialize(vb *VariableBlob) *VariableBlob {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(*n))
	ov := append(*vb, b...)
	return &ov
}

// DeserializeUInt16 function
func DeserializeUInt16(vb *VariableBlob) (uint64, *UInt16, error) {
	var i UInt16

	if len(*vb) < 2 {
		return 0, &i, errors.New("Unexpected EOF")
	}

	i = UInt16(binary.BigEndian.Uint16(*vb))

	return 2, &i, nil
}

// --------------------------------
//  Int32
// --------------------------------

// Int32 type
type Int32 int32

// NewInt32 factory
func NewInt32() *Int32 {
	o := Int32(0)
	return &o
}

// Serialize Int32
func (n *Int32) Serialize(vb *VariableBlob) *VariableBlob {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(*n))
	ov := append(*vb, b...)
	return &ov
}

// DeserializeInt32 function
func DeserializeInt32(vb *VariableBlob) (uint64, *Int32, error) {
	var i Int32

	if len(*vb) < 4 {
		return 0, &i, errors.New("Unexpected EOF")
	}

	i = Int32(binary.BigEndian.Uint32(*vb))

	return 4, &i, nil
}

// --------------------------------
//  UInt32
// --------------------------------

// UInt32 type
type UInt32 uint32

// NewUInt32 factory
func NewUInt32() *UInt32 {
	o := UInt32(0)
	return &o
}

// Serialize UInt32
func (n *UInt32) Serialize(vb *VariableBlob) *VariableBlob {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(*n))
	ov := append(*vb, b...)
	return &ov
}

// DeserializeUInt32 function
func DeserializeUInt32(vb *VariableBlob) (uint64, *UInt32, error) {
	var i UInt32

	if len(*vb) < 4 {
		return 0, &i, errors.New("Unexpected EOF")
	}

	i = UInt32(binary.BigEndian.Uint32(*vb))

	return 4, &i, nil
}

// --------------------------------
//  Int64
// --------------------------------

// Int64 type
type Int64 int64

// NewInt64 factory
func NewInt64() *Int64 {
	o := Int64(0)
	return &o
}

// Serialize Int64
func (n *Int64) Serialize(vb *VariableBlob) *VariableBlob {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(*n))
	ov := append(*vb, b...)
	return &ov
}

// DeserializeInt64 function
func DeserializeInt64(vb *VariableBlob) (uint64, *Int64, error) {
	var i Int64

	if len(*vb) < 8 {
		return 0, &i, errors.New("Unexpected EOF")
	}

	i = Int64(binary.BigEndian.Uint64(*vb))

	return 8, &i, nil
}

// --------------------------------
//  UInt64
// --------------------------------

// UInt64 type
type UInt64 uint64

// NewUInt64 factory
func NewUInt64() *UInt64 {
	o := UInt64(0)
	return &o
}

// Serialize UInt64
func (n *UInt64) Serialize(vb *VariableBlob) *VariableBlob {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(*n))
	ov := append(*vb, b...)
	return &ov
}

// DeserializeUInt64 function
func DeserializeUInt64(vb *VariableBlob) (uint64, *UInt64, error) {
	var i UInt64

	if len(*vb) < 8 {
		return 0, &i, errors.New("Unexpected EOF")
	}

	i = UInt64(binary.BigEndian.Uint64(*vb))

	return 8, &i, nil
}

// ----------------------------------------
//  Int128
// ----------------------------------------

// Int128 type
type Int128 struct {
	Value big.Int
}

var (
	int128Max  *Int128
	int128Min  *Int128
	int128Once sync.Once
)

func initInt128Bounds() {
	int128Once.Do(func() {
		int128Max = NewInt128()
		nv, _ := int128Max.Value.SetString("170141183460469231731687303715884105727", 10)
		int128Max.Value = *nv
		int128Min = NewInt128()
		nv, _ = int128Min.Value.SetString("-170141183460469231731687303715884105728", 10)
		int128Min.Value = *nv
	})
}

// Int128Max upper bound
func Int128Max() Int128 {
	initInt128Bounds()
	return *int128Max
}

// Int128Min lower bound
func Int128Min() Int128 {
	initInt128Bounds()
	return *int128Min
}

// NewInt128 factory
func NewInt128() *Int128 {
	result := Int128{}
	result.Value = *big.NewInt(0)
	return &result
}

// NewInt128FromString string factory
func NewInt128FromString(value string) (*Int128, error) {
	var result Int128 = Int128{}
	nv, ok := result.Value.SetString(value, 10)
	if ok == false {
		return nil, errors.New("Could not parse Int128")
	}

	max := Int128Max()
	min := Int128Min()
	if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
		return nil, errors.New("Int128 is out of bounds")
	}
	result.Value = *nv
	return &result, nil
}

// Serialize Int128
func (n *Int128) Serialize(vb *VariableBlob) *VariableBlob {
	s := SerializeBigInt(&n.Value, 16, true)
	ov := append(*vb, *s...)
	return &ov
}

// DeserializeInt128 function
func DeserializeInt128(vb *VariableBlob) (uint64, *Int128, error) {
	biPtr, err := DeserializeBigInt(vb, 16, true)
	i := Int128{}

	if err != nil {
		return 0, &i, err
	}

	i.Value = *biPtr

	return 16, &i, nil
}

// MarshalJSON Int128
func (n Int128) MarshalJSON() ([]byte, error) {
	if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
		return json.Marshal(&i)
	}

	s := n.Value.String()
	return json.Marshal(s)
}

// UnmarshalJSON Int128
func (n *Int128) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)

	if err != nil {
		var i int64
		if err = json.Unmarshal(b, &i); err != nil {
			return err
		}
		n.Value = *big.NewInt(i)
	} else {
		nv, err := NewInt128FromString(s)
		if err != nil {
			return err
		}
		*n = *nv
	}

	return nil
}

// ----------------------------------------
//  UInt128
// ----------------------------------------

// UInt128 type
type UInt128 struct {
	Value big.Int
}

var (
	uint128Max  *UInt128
	uint128Min  *UInt128
	uint128Once sync.Once
)

func initUInt128Bounds() {
	uint128Once.Do(func() {
		uint128Max = NewUInt128()
		nv, _ := uint128Max.Value.SetString("340282366920938463463374607431768211455", 10)
		uint128Max.Value = *nv
		uint128Min = NewUInt128()
	})
}

// UInt128Max upper bound
func UInt128Max() UInt128 {
	initUInt128Bounds()
	return *uint128Max
}

// UInt128Min lower bound
func UInt128Min() UInt128 {
	initUInt128Bounds()
	return *uint128Min
}

// NewUInt128 factory
func NewUInt128() *UInt128 {
	result := UInt128{}
	result.Value = *big.NewInt(0)
	return &result
}

// NewUInt128FromString factory
func NewUInt128FromString(value string) (*UInt128, error) {
	var result UInt128 = UInt128{}
	nv, ok := result.Value.SetString(value, 10)
	if ok == false {
		return nil, errors.New("Could not parse UInt128")
	}
	max := UInt128Max()
	min := UInt128Min()
	if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
		return nil, errors.New("UInt128 is out of bounds")
	}
	result.Value = *nv
	return &result, nil
}

// Serialize UInt128
func (n *UInt128) Serialize(vb *VariableBlob) *VariableBlob {
	x := SerializeBigInt(&n.Value, 16, false)
	ov := append(*vb, *x...)
	return &ov
}

// DeserializeUInt128 function
func DeserializeUInt128(vb *VariableBlob) (uint64, *UInt128, error) {
	biPtr, err := DeserializeBigInt(vb, 16, false)
	i := UInt128{}

	if err != nil {
		return 0, &i, err
	}

	i.Value = *biPtr

	return 16, &i, nil
}

// MarshalJSON UInt128
func (n UInt128) MarshalJSON() ([]byte, error) {
	if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
		return json.Marshal(&i)
	}

	s := n.Value.String()
	return json.Marshal(s)
}

// UnmarshalJSON UInt128
func (n *UInt128) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)

	if err != nil {
		var i int64
		if err = json.Unmarshal(b, &i); err != nil {
			return err
		}
		if i < 0 {
			return errors.New("UInt128 is out of bounds")
		}
		n.Value = *big.NewInt(i)
	} else {
		nv, err := NewUInt128FromString(s)
		if err != nil {
			return err
		}
		*n = *nv
	}

	return nil
}

// ----------------------------------------
//  Int160
// ----------------------------------------

// Int160 type
type Int160 struct {
	Value big.Int
}

var (
	int160Max  *Int160
	int160Min  *Int160
	int160Once sync.Once
)

func initInt160Bounds() {
	int160Once.Do(func() {
		int160Max = NewInt160()
		nv, _ := int160Max.Value.SetString("730750818665451459101842416358141509827966271487", 10)
		int160Max.Value = *nv
		int160Min = NewInt160()
		nv, _ = int160Min.Value.SetString("-730750818665451459101842416358141509827966271488", 10)
		int160Min.Value = *nv
	})
}

// Int160Max upper bound
func Int160Max() Int160 {
	initInt160Bounds()
	return *int160Max
}

// Int160Min lower bound
func Int160Min() Int160 {
	initInt160Bounds()
	return *int160Min
}

// NewInt160 factory
func NewInt160() *Int160 {
	result := Int160{}
	result.Value = *big.NewInt(0)
	return &result
}

// NewInt160FromString factory
func NewInt160FromString(value string) (*Int160, error) {
	var result Int160 = Int160{}
	nv, ok := result.Value.SetString(value, 10)
	if ok == false {
		return nil, errors.New("Could not parse Int160")
	}
	max := Int160Max()
	min := Int160Min()
	if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
		return nil, errors.New("Int160 is out of bounds")
	}
	result.Value = *nv
	return &result, nil
}

// Serialize Int160
func (n *Int160) Serialize(vb *VariableBlob) *VariableBlob {
	x := SerializeBigInt(&n.Value, 20, true)
	ov := append(*vb, *x...)
	return &ov
}

// DeserializeInt160 function
func DeserializeInt160(vb *VariableBlob) (uint64, *Int160, error) {
	biPtr, err := DeserializeBigInt(vb, 20, true)
	i := Int160{}

	if err != nil {
		return 0, &i, err
	}

	i.Value = *biPtr

	return 20, &i, nil
}

// MarshalJSON Int160
func (n *Int160) MarshalJSON() ([]byte, error) {
	if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
		return json.Marshal(i)
	}

	s := n.Value.String()
	return json.Marshal(s)
}

// UnmarshalJSON Int160
func (n *Int160) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)

	if err != nil {
		var i int64
		if err = json.Unmarshal(b, &i); err != nil {
			return err
		}
		n.Value = *big.NewInt(i)
	} else {
		nv, err := NewInt160FromString(s)
		if err != nil {
			return err
		}
		*n = *nv
	}

	return nil
}

// ----------------------------------------
//  UInt160
// ----------------------------------------

// UInt160 type
type UInt160 struct {
	Value big.Int
}

var (
	uint160Max  *UInt160
	uint160Min  *UInt160
	uint160Once sync.Once
)

func initUInt160Bounds() {
	uint160Once.Do(func() {
		uint160Max = NewUInt160()
		nv, _ := uint160Max.Value.SetString("1461501637330902918203684832716283019655932542975", 10)
		uint160Max.Value = *nv
		uint160Min = NewUInt160()
	})
}

// UInt160Max upper bound
func UInt160Max() UInt160 {
	initUInt160Bounds()
	return *uint160Max
}

// UInt160Min lower bound
func UInt160Min() UInt160 {
	initUInt160Bounds()
	return *uint160Min
}

// NewUInt160 factory
func NewUInt160() *UInt160 {
	result := UInt160{}
	result.Value = *big.NewInt(0)
	return &result
}

// NewUInt160FromString factory
func NewUInt160FromString(value string) (*UInt160, error) {
	var result UInt160 = UInt160{}
	nv, ok := result.Value.SetString(value, 10)
	if ok == false {
		return nil, errors.New("Could not parse UInt160")
	}
	max := UInt160Max()
	min := UInt160Min()
	if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
		return nil, errors.New("UInt160 is out of bounds")
	}
	result.Value = *nv
	return &result, nil
}

// Serialize UInt160
func (n *UInt160) Serialize(vb *VariableBlob) *VariableBlob {
	x := SerializeBigInt(&n.Value, 20, false)
	ov := append(*vb, *x...)
	return &ov
}

// DeserializeUInt160 function
func DeserializeUInt160(vb *VariableBlob) (uint64, *UInt160, error) {
	biPtr, err := DeserializeBigInt(vb, 20, false)
	i := UInt160{}

	if err != nil {
		return 0, &i, err
	}

	i.Value = *biPtr

	return 20, &i, nil
}

// MarshalJSON UInt160
func (n UInt160) MarshalJSON() ([]byte, error) {
	if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
		return json.Marshal(&i)
	}

	s := n.Value.String()
	return json.Marshal(s)
}

// UnmarshalJSON UInt160
func (n *UInt160) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)

	if err != nil {
		var i int64
		if err = json.Unmarshal(b, &i); err != nil {
			return err
		}
		if i < 0 {
			return errors.New("UInt160 is out of bounds")
		}
		n.Value = *big.NewInt(i)
	} else {
		nv, err := NewUInt160FromString(s)
		if err != nil {
			return err
		}
		*n = *nv
	}

	return nil
}

// ----------------------------------------
//  Int256
// ----------------------------------------

// Int256 type
type Int256 struct {
	Value big.Int
}

var (
	int256Max  *Int256
	int256Min  *Int256
	int256Once sync.Once
)

func initInt256Bounds() {
	int256Once.Do(func() {
		int256Max = NewInt256()
		nv, _ := int256Max.Value.SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		int256Max.Value = *nv
		int256Min = NewInt256()
		nv, _ = int256Min.Value.SetString("-57896044618658097711785492504343953926634992332820282019728792003956564819968", 10)
		int256Min.Value = *nv
	})
}

// Int256Max upper bound
func Int256Max() Int256 {
	initInt256Bounds()
	return *int256Max
}

// Int256Min lower bound
func Int256Min() Int256 {
	initInt256Bounds()
	return *int256Min
}

// NewInt256 factory
func NewInt256() *Int256 {
	result := Int256{}
	result.Value = *big.NewInt(0)
	return &result
}

// NewInt256FromString factory
func NewInt256FromString(value string) (*Int256, error) {
	var result Int256 = Int256{}
	nv, ok := result.Value.SetString(value, 10)
	if ok == false {
		return nil, errors.New("Could not parse Int256")
	}
	max := Int256Max()
	min := Int256Min()
	if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
		return nil, errors.New("Int256 is out of bounds")
	}
	result.Value = *nv
	return &result, nil
}

// Serialize Int256
func (n *Int256) Serialize(vb *VariableBlob) *VariableBlob {
	x := SerializeBigInt(&n.Value, 32, true)
	ov := append(*vb, *x...)
	return &ov
}

// DeserializeInt256 function
func DeserializeInt256(vb *VariableBlob) (uint64, *Int256, error) {
	biPtr, err := DeserializeBigInt(vb, 32, true)
	i := Int256{}

	if err != nil {
		return 0, &i, err
	}

	i.Value = *biPtr

	return 32, &i, nil
}

// MarshalJSON Int256
func (n Int256) MarshalJSON() ([]byte, error) {
	if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
		return json.Marshal(&i)
	}

	s := n.Value.String()
	return json.Marshal(s)
}

// UnmarshalJSON Int256
func (n *Int256) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)

	if err != nil {
		var i int64
		if err = json.Unmarshal(b, &i); err != nil {
			return err
		}
		n.Value = *big.NewInt(i)
	} else {
		nv, err := NewInt256FromString(s)
		if err != nil {
			return err
		}
		*n = *nv
	}

	return nil
}

// ----------------------------------------
//  UInt256
// ----------------------------------------

// UInt256 type
type UInt256 struct {
	Value big.Int
}

var (
	uint256Max  *UInt256
	uint256Min  *UInt256
	uint256Once sync.Once
)

func initUInt256Bounds() {
	uint256Once.Do(func() {
		uint256Max = NewUInt256()
		nv, _ := uint256Max.Value.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
		uint256Max.Value = *nv
		uint256Min = NewUInt256()
	})
}

// UInt256Max upper bound
func UInt256Max() UInt256 {
	initUInt256Bounds()
	return *uint256Max
}

// UInt256Min lower bound
func UInt256Min() UInt256 {
	initUInt256Bounds()
	return *uint256Min
}

// NewUInt256 factory
func NewUInt256() *UInt256 {
	result := UInt256{}
	result.Value = *big.NewInt(0)
	return &result
}

// NewUInt256FromString factory
func NewUInt256FromString(value string) (*UInt256, error) {
	var result UInt256 = UInt256{}
	nv, ok := result.Value.SetString(value, 10)
	if ok == false {
		return nil, errors.New("Could not parse UInt256")
	}
	max := UInt256Max()
	min := UInt256Min()
	if nv.Cmp(&max.Value) == 1 || nv.Cmp(&min.Value) == -1 {
		return nil, errors.New("UInt256 is out of bounds")
	}
	result.Value = *nv
	return &result, nil
}

// Serialize UInt256
func (n *UInt256) Serialize(vb *VariableBlob) *VariableBlob {
	x := SerializeBigInt(&n.Value, 32, false)
	ov := append(*vb, *x...)
	return &ov
}

// DeserializeUInt256 function
func DeserializeUInt256(vb *VariableBlob) (uint64, *UInt256, error) {
	biPtr, err := DeserializeBigInt(vb, 32, false)
	i := UInt256{}

	if err != nil {
		return 0, &i, err
	}

	i.Value = *biPtr

	return 32, &i, nil
}

// MarshalJSON UInt256
func (n UInt256) MarshalJSON() ([]byte, error) {
	if i := n.Value.Int64(); n.Value.IsInt64() && i <= bigIntNumericLiteralMax && i >= bigIntNumericLiteralMin {
		return json.Marshal(&i)
	}

	s := n.Value.String()
	return json.Marshal(s)
}

// UnmarshalJSON UInt256
func (n *UInt256) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)

	if err != nil {
		var i int64
		if err = json.Unmarshal(b, &i); err != nil {
			return err
		}
		if i < 0 {
			return errors.New("UInt256 is out of bounds")
		}
		n.Value = *big.NewInt(i)
	} else {
		nv, err := NewUInt256FromString(s)
		if err != nil {
			return err
		}
		*n = *nv
	}

	return nil
}

// --------------------------------
//  VariableBlob
// --------------------------------

// VariableBlob type
type VariableBlob []byte

// NewVariableBlob factory
// TODO: Make this variadic for size and size_hint
func NewVariableBlob() *VariableBlob {
	vb := VariableBlob(make([]byte, 0))
	return &vb
}

// Serialize VariableBlob
func (n *VariableBlob) Serialize(vb *VariableBlob) *VariableBlob {
	header := make([]byte, binary.MaxVarintLen64)
	bytes := binary.PutUvarint(header, uint64(len(*n)))
	ovb := append(*vb, header[:bytes]...)
	ovb = append(ovb, *n...)
	return &ovb
}

// DeserializeVariableBlob function
func DeserializeVariableBlob(vb *VariableBlob) (uint64, *VariableBlob, error) {
	size, bytes := binary.Uvarint(*vb)
	var result VariableBlob = VariableBlob(make([]byte, 0, size))
	if bytes <= 0 {
		return 0, &result, errors.New("Could not deserialize variable blob size")
	}

	if len(*vb) < bytes+int(size) {
		return 0, &result, errors.New("Unexpected EOF")
	}

	ovb := append(result, (*vb)[bytes:uint64(bytes)+size]...)
	return uint64(uint64(bytes) + size), &ovb, nil
}

// MarshalJSON VariableBlob
func (n VariableBlob) MarshalJSON() ([]byte, error) {
	s := EncodeBytes(n)
	return json.Marshal(s)
}

// UnmarshalJSON VariabeBlob
func (n *VariableBlob) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	db, err := DecodeBytes(s)
	if err != nil {
		return err
	}
	if len(db) == 0 && len(s[1:]) > 0 {
		return errors.New("Unable to decode base58")
	}

	*n = db
	return nil
}

// --------------------------------
//  TimestampType
// --------------------------------

// TimestampType type
type TimestampType UInt64

// NewTimestampType factory
func NewTimestampType() *TimestampType {
	o := TimestampType(0)
	return &o
}

// Serialize TimestampType
func (n *TimestampType) Serialize(vb *VariableBlob) *VariableBlob {
	un := UInt64(*n)
	return un.Serialize(vb)
}

// DeserializeTimestampType function
func DeserializeTimestampType(vb *VariableBlob) (uint64, *TimestampType, error) {
	i, x, err := DeserializeUInt64(vb)
	ox := TimestampType(*x)
	return i, &ox, err
}

// --------------------------------
//  BlockHeightType
// --------------------------------

// BlockHeightType type
type BlockHeightType UInt64

// NewBlockHeightType factory
func NewBlockHeightType() *BlockHeightType {
	o := BlockHeightType(0)
	return &o
}

// Serialize BlockHeightType
func (n *BlockHeightType) Serialize(vb *VariableBlob) *VariableBlob {
	un := UInt64(*n)
	return un.Serialize(vb)
}

// DeserializeBlockHeightType function
func DeserializeBlockHeightType(vb *VariableBlob) (uint64, *BlockHeightType, error) {
	i, x, err := DeserializeUInt64(vb)
	ox := BlockHeightType(*x)
	return i, &ox, err
}

// --------------------------------
//  Multihash
// --------------------------------

// Multihash type
type Multihash struct {
	ID     UInt64       `json:"hash"`
	Digest VariableBlob `json:"digest"`
}

// NewMultihash factory
func NewMultihash() *Multihash {
	o := Multihash{}
	o.ID = UInt64(0)
	o.Digest = *NewVariableBlob()
	return &o
}

// Equals Multihash comparison
func (m0 *Multihash) Equals(m1 *Multihash) bool {
	return (m0.ID == m1.ID) && bytes.Equal(m0.Digest, m1.Digest)
}

// LessThan Multihash comparison
func (m0 *Multihash) LessThan(m1 *Multihash) bool {
	if m0.ID < m1.ID {
		return true
	}
	if m0.ID > m1.ID {
		return false
	}
	return (len(m0.Digest) - len(m1.Digest)) < 0
}

// GreaterThan Multihash comparison
func (m0 *Multihash) GreaterThan(m1 *Multihash) bool {
	return !m0.Equals(m1) && !m0.LessThan(m1)
}

// Serialize Multihash
func (m0 *Multihash) Serialize(vb *VariableBlob) *VariableBlob {
	vb = EncodeVarint(vb, uint64(m0.ID))
	return m0.Digest.Serialize(vb)
}

// DeserializeMultihash function
func DeserializeMultihash(vb *VariableBlob) (uint64, *Multihash, error) {
	omh := Multihash{}
	id, isize := binary.Uvarint(*vb)
	if isize <= 0 {
		return 0, &omh, errors.New("Could not deserialize multihash id")
	}
	rvb := (*vb)[isize:]
	dsize, d, err := DeserializeVariableBlob(&rvb)
	if err != nil {
		return 0, &omh, err
	}
	omh.ID = UInt64(id)
	omh.Digest = *d
	return uint64(isize) + dsize, &omh, nil
}

// --------------------------------
//  Utility Functions
// --------------------------------

// EncodeBytes utility function
func EncodeBytes(b []byte) string {
	return "z" + base58.Encode(b)
}

// DecodeBytes utility function
func DecodeBytes(s string) ([]byte, error) {
	if len(s) <= 1 {
		return make([]byte, 0), nil
	}

	switch s[0] {
	case 'z':
		return base58.Decode(s[1:]), nil
	default:
		return nil, errors.New("Unknown encoding: " + string(s[0]))
	}
}

// SerializeBigInt helper function
func SerializeBigInt(num *big.Int, byteSize int, signed bool) *VariableBlob {
	v := VariableBlob(make([]byte, byteSize))

	if signed && num.Sign() == -1 {
		x := big.NewInt(1)
		x = x.Add(x, num)
		v = x.FillBytes(v)
		for i := 0; i < byteSize; i++ {
			v[i] = ^v[i]
		}
		return &v
	}

	v = num.FillBytes(v)
	return &v
}

// DeserializeBigInt helper function
func DeserializeBigInt(vb *VariableBlob, byteSize int, signed bool) (*big.Int, error) {
	num := new(big.Int)

	if len(*vb) < byteSize {
		return num, errors.New("Unexpected EOF")
	}

	if signed && (0x80&(*vb)[0]) == 0x80 {
		v := VariableBlob(make([]byte, byteSize))
		for i := 0; i < byteSize; i++ {
			v[i] = ^((*vb)[i])
		}
		neg := big.NewInt(-1)
		return num.SetBytes(v).Mul(neg, num).Add(neg, num), nil
	}

	return num.SetBytes((*vb)[:byteSize]), nil
}

// EncodeVarint utility function
func EncodeVarint(vb *VariableBlob, value uint64) *VariableBlob {
	header := make([]byte, binary.MaxVarintLen64)
	bytes := binary.PutUvarint(header, value)
	*vb = append(*vb, header[:bytes]...)
	return vb
}
