import * as JSONbig from "json-bigint";
import {
  VariableBlob,
  FixedBlob,
  Bool,
  Str,
  Int8,
  Int16,
  Int32,
  Int64,
  Int128,
  Int160,
  Int256,
  Uint8,
  Uint16,
  Uint32,
  Uint64,
  Uint128,
  Uint160,
  Uint256,
  TimestampType,
  BlockHeightType,
  Multihash,
  MultihashVector,
  Opaque,
  Optional,
  VarInt,
  Vector,
  MAX_INT64,
  MIN_INT64,
} from ".";

describe("Koinos Types - Typescript", () => {
  it("should compare VariableBlobs", () => {
    expect.assertions(3);
    const num1 = new Int16(1000);
    const num2 = new Int64(1000);
    const vb1 = new VariableBlob(8).serialize(num1);
    const vb2 = new VariableBlob(8).serialize(num1);
    const vb3 = new VariableBlob(8).serialize(num2);
    const vb4 = new VariableBlob(vb3);
    expect(vb1.equals(vb2)).toBe(true);
    expect(vb1.equals(vb3)).toBe(false);
    expect(vb3.equals(vb4)).toBe(true);
  });

  it("Serialize and desearialize", () => {
    expect.assertions(31);
    const vb1 = new VariableBlob("Mkw96mR+Hh71IWwJoT/2lJXBDl5Q=");
    const multihash = new Multihash("zHdxh1kBNvtADThNaTBNvDqyMwzJYi8x1Do7Vg");
    const jsonMultihashVector = {
      id: BigInt(1234),
      digests: [
        "z16UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGA",
        "z26UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGB",
        "z36UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGC",
        "z46UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGD",
      ],
    };
    const multihashVector = new MultihashVector(jsonMultihashVector);

    const vb = new VariableBlob()
      .serialize(vb1)
      .serialize(new Bool(true))
      .serialize(new Str("test"))
      .serialize(new Int8(100))
      .serialize(new Uint8(100))
      .serialize(new Int16(1000))
      .serialize(new Uint16(1000))
      .serialize(new Int32(10000))
      .serialize(new Uint32(10000))
      .serialize(new Int64("100000"))
      .serialize(new Int64("-100000"))
      .serialize(new Uint64("100000"))
      .serialize(new Int128("100000"))
      .serialize(new Int128("-100000"))
      .serialize(new Uint128("100000"))
      .serialize(new Int160("100000"))
      .serialize(new Int160("-100000"))
      .serialize(new Uint160("100000"))
      .serialize(new Int256("100000"))
      .serialize(new Int256("-100000"))
      .serialize(new Uint256("100000"))
      .serialize(new TimestampType("1234567890"))
      .serialize(new BlockHeightType(123456))
      .serialize(multihash)
      .serialize(multihashVector)
      .serialize(new Opaque(Str, "test opaque"))
      .serialize(new Optional(Str))
      .serialize(new Optional(Str, "optional with value"))
      .serialize(new Vector(Int8, [2, 4, 6]))
      .serialize(new FixedBlob("MN0TkNkaw7g==", 7))
      .resetCursor();

    expect(vb.deserialize(VariableBlob).equals(vb1)).toBe(true);
    expect(vb.deserialize(Bool).toJSON()).toBe(true);
    expect(vb.deserialize(Str).toJSON()).toBe("test");
    expect(vb.deserialize(Int8).toJSON()).toBe(100);
    expect(vb.deserialize(Uint8).toJSON()).toBe(100);
    expect(vb.deserialize(Int16).toJSON()).toBe(1000);
    expect(vb.deserialize(Uint16).toJSON()).toBe(1000);
    expect(vb.deserialize(Int32).toJSON()).toBe(10000);
    expect(vb.deserialize(Uint32).toJSON()).toBe(10000);
    expect(vb.deserialize(Int64).toJSON()).toBe(BigInt(100000));
    expect(vb.deserialize(Int64).toJSON()).toBe(BigInt(-100000));
    expect(vb.deserialize(Uint64).toJSON()).toBe(BigInt(100000));
    expect(vb.deserialize(Int128).toJSON()).toBe(BigInt(100000));
    expect(vb.deserialize(Int128).toJSON()).toBe(BigInt(-100000));
    expect(vb.deserialize(Uint128).toJSON()).toBe(BigInt(100000));
    expect(vb.deserialize(Int160).toJSON()).toBe(BigInt(100000));
    expect(vb.deserialize(Int160).toJSON()).toBe(BigInt(-100000));
    expect(vb.deserialize(Uint160).toJSON()).toBe(BigInt(100000));
    expect(vb.deserialize(Int256).toJSON()).toBe(BigInt(100000));
    expect(vb.deserialize(Int256).toJSON()).toBe(BigInt(-100000));
    expect(vb.deserialize(Uint256).toJSON()).toBe(BigInt(100000));
    expect(vb.deserialize(TimestampType).toJSON()).toBe(BigInt(1234567890));
    expect(vb.deserialize(BlockHeightType).toJSON()).toBe(BigInt(123456));
    expect(vb.deserialize(Multihash).equals(multihash)).toBe(true);
    expect(vb.deserialize(MultihashVector).toJSON()).toStrictEqual(
      jsonMultihashVector
    );
    expect(vb.deserializeOpaque(Str).toJSON()).toBe("test opaque");
    expect(vb.deserializeOptional(Str).toJSON()).not.toBeDefined();
    expect(vb.deserializeOptional(Str).toJSON()).toBe("optional with value");
    expect(vb.deserializeVector(Int8).toJSON()).toStrictEqual([2, 4, 6]);
    expect(vb.deserialize(FixedBlob, 7).toJSON()).toBe("MN0TkNkaw7g==");

    expect(multihash.toJSON()).toBe("zHdxh1kBNvtADThNaTBNvDqyMwzJYi8x1Do7Vg");
  });

  it("should transform integers to json as number or string depending on value", () => {
    expect.assertions(4);
    // serialized as numbers
    expect(JSONbig.stringify(new Int256(MAX_INT64).toJSON())).toBe(
      "9223372036854775807"
    );
    expect(JSONbig.stringify(new Int256(MIN_INT64).toJSON())).toBe(
      "-9223372036854775808"
    );

    // serialized as strings
    expect(JSONbig.stringify(new Int256(MAX_INT64 + BigInt(1)).toJSON())).toBe(
      '"9223372036854775808"'
    );
    expect(JSONbig.stringify(new Int256(MIN_INT64 - BigInt(1)).toJSON())).toBe(
      '"-9223372036854775809"'
    );
  });

  it("should create an opaque class", () => {
    expect.assertions(10);
    const opaque = new Opaque(Int32, 123456);
    expect(opaque.isUnboxed()).toBe(true);
    expect(opaque.isMutable()).toBe(true);
    opaque.box();
    expect(opaque.isUnboxed()).toBe(false);
    expect(opaque.isMutable()).toBe(false);
    opaque.unbox();
    expect(opaque.isUnboxed()).toBe(true);
    expect(opaque.isMutable()).toBe(false);
    opaque.makeMutable();
    expect(opaque.isUnboxed()).toBe(true);
    expect(opaque.isMutable()).toBe(true);
    opaque.makeImmutable();
    expect(opaque.isUnboxed()).toBe(true);
    expect(opaque.isMutable()).toBe(false);
  });

  it("should calculate the size required for serialization", () => {
    expect.assertions(25);

    const vb = new VariableBlob("zABCDEF1234567");
    expect(vb.calcSerializedSize()).toBe(vb.serialize().length());

    const fb = new FixedBlob("z36UjcYNBG9", 7);
    expect(fb.calcSerializedSize()).toBe(fb.serialize().length());

    const bool = new Bool(true);
    expect(bool.calcSerializedSize()).toBe(bool.serialize().length());

    const str = new Str("test!$^#🙂🍻");
    expect(str.calcSerializedSize()).toBe(str.serialize().length());

    const int8 = new Int8(123);
    expect(int8.calcSerializedSize()).toBe(int8.serialize().length());

    const int16 = new Int16(12345);
    expect(int16.calcSerializedSize()).toBe(int16.serialize().length());

    const int32 = new Int32(123456);
    expect(int32.calcSerializedSize()).toBe(int32.serialize().length());

    const int64 = new Int64(123456);
    expect(int64.calcSerializedSize()).toBe(int64.serialize().length());

    const int128 = new Int128(123456);
    expect(int128.calcSerializedSize()).toBe(int128.serialize().length());

    const int160 = new Int160(123456);
    expect(int160.calcSerializedSize()).toBe(int160.serialize().length());

    const int256 = new Int256(123456);
    expect(int256.calcSerializedSize()).toBe(int256.serialize().length());

    const uint8 = new Uint8(123);
    expect(uint8.calcSerializedSize()).toBe(int8.serialize().length());

    const uint16 = new Uint16(12345);
    expect(uint16.calcSerializedSize()).toBe(uint16.serialize().length());

    const uint32 = new Uint32(123456);
    expect(uint32.calcSerializedSize()).toBe(uint32.serialize().length());

    const uint64 = new Uint64(123456);
    expect(uint64.calcSerializedSize()).toBe(uint64.serialize().length());

    const uint128 = new Uint128(123456);
    expect(uint128.calcSerializedSize()).toBe(uint128.serialize().length());

    const uint160 = new Uint160(123456);
    expect(uint160.calcSerializedSize()).toBe(uint160.serialize().length());

    const uint256 = new Uint256(123456);
    expect(uint256.calcSerializedSize()).toBe(uint256.serialize().length());

    const varint = new VarInt(1234567);
    expect(varint.calcSerializedSize()).toBe(varint.serialize().length());

    const multihash = new Multihash({
      id: 123,
      digest: "z36UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGS",
    });
    expect(multihash.calcSerializedSize()).toBe(multihash.serialize().length());

    const multihashVector = new MultihashVector({
      id: 1234,
      digests: [
        "z16UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGA",
        "z26UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGB",
        "z36UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGC",
        "z46UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGD",
      ],
    });
    expect(multihashVector.calcSerializedSize()).toBe(
      multihashVector.serialize().length()
    );

    const vector = new Vector(Str, ["alice", "bob", "carl"]);
    expect(vector.calcSerializedSize()).toBe(vector.serialize().length());

    const opaque = new Opaque(Str, "test!$^#🙂🍻");
    expect(opaque.calcSerializedSize()).toBe(opaque.serialize().length());

    const optional = new Optional(Uint8, 0);
    expect(optional.calcSerializedSize()).toBe(optional.serialize().length());

    const optionalEmpty = new Optional(Uint8);
    expect(optionalEmpty.calcSerializedSize()).toBe(
      optionalEmpty.serialize().length()
    );
  });
});
