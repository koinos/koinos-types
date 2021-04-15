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
  Opaque,
  Vector,
} from ".";

describe("Koinos Types - Typescript", () => {
  it("should compare VariableBlobs", () => {
    expect.assertions(3);
    const num1 = new Int16(1000);
    const num2 = new Int64(1000);
    const vb1 = new VariableBlob(9).serialize(num1).flip();
    const vb2 = new VariableBlob(5).serialize(num1).flip();
    const vb3 = new VariableBlob(5).serialize(num2).flip();
    const vb4 = new VariableBlob(vb3);
    expect(vb1.equals(vb2)).toBe(true);
    expect(vb1.equals(vb3)).toBe(false);
    expect(vb3.equals(vb4)).toBe(true);
  });

  it("Serialize and desearialize", () => {
    expect.assertions(27);
    const vb1 = new VariableBlob("z26UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGS");
    const multihash = new Multihash({
      id: 123,
      digest: "z36UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGS",
    });

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
      .serialize(new Vector(Int8, [2, 4, 6]))
      .serialize(new FixedBlob(7, "z36UjcYNBG9"))
      .flip();

    expect(vb.deserialize(VariableBlob).equals(vb1)).toBe(true);
    expect(vb.deserialize(Bool).toBoolean()).toBe(true);
    expect(vb.deserialize(Str).toString()).toBe("test");
    expect(vb.deserialize(Int8).toNumber()).toBe(100);
    expect(vb.deserialize(Uint8).toNumber()).toBe(100);
    expect(vb.deserialize(Int16).toNumber()).toBe(1000);
    expect(vb.deserialize(Uint16).toNumber()).toBe(1000);
    expect(vb.deserialize(Int32).toNumber()).toBe(10000);
    expect(vb.deserialize(Uint32).toNumber()).toBe(10000);
    expect(vb.deserialize(Int64).toString()).toBe("100000");
    expect(vb.deserialize(Int64).toString()).toBe("-100000");
    expect(vb.deserialize(Uint64).toString()).toBe("100000");
    expect(vb.deserialize(Int128).toString()).toBe("100000");
    expect(vb.deserialize(Int128).toString()).toBe("-100000");
    expect(vb.deserialize(Uint128).toString()).toBe("100000");
    expect(vb.deserialize(Int160).toString()).toBe("100000");
    expect(vb.deserialize(Int160).toString()).toBe("-100000");
    expect(vb.deserialize(Uint160).toString()).toBe("100000");
    expect(vb.deserialize(Int256).toString()).toBe("100000");
    expect(vb.deserialize(Int256).toString()).toBe("-100000");
    expect(vb.deserialize(Uint256).toString()).toBe("100000");
    expect(vb.deserialize(TimestampType).toString()).toBe("1234567890");
    expect(vb.deserialize(BlockHeightType).toString()).toBe("123456");
    expect(vb.deserialize(Multihash).equals(multihash)).toBe(true);
    expect(vb.deserializeVector(Int8).toJSON()).toStrictEqual([2, 4, 6]);
    expect(vb.deserializeFixedBlob(7).toJSON()).toBe("z36UjcYNBG9");

    expect(JSONbig.stringify(multihash.toJSON())).toBe(
      '{"id":123,"digest":"z36UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGS"}'
    );
  });

  it("should create an opaque class", () => {
    expect.assertions(10);
    const num = new Int32(123456);
    const opaque = new Opaque(Int32, num);
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
});
