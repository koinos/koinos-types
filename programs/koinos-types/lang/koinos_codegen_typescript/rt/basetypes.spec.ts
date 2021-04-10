import {
  VariableBlob,
  KoinosBoolean,
  KoinosString,
  Int8,
  Int16,
  Int32,
  Int64,
  Int128,
  Int160,
  Int256,
  UInt8,
  UInt16,
  UInt32,
  UInt64,
  UInt128,
  UInt160,
  UInt256,
  TimestampType,
  BlockHeightType,
  Multihash,
  Opaque,
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
    expect.assertions(24);
    const vb1 = new VariableBlob("ce123456af");
    const multihash = new Multihash({
      id: 123,
      digest: "0xff1234567890",
    });

    const vb = new VariableBlob()
      .serialize(vb1)
      .serialize(new KoinosBoolean(true))
      .serialize(new KoinosString("test"))
      .serialize(new Int8(100))
      .serialize(new UInt8(100))
      .serialize(new Int16(1000))
      .serialize(new UInt16(1000))
      .serialize(new Int32(10000))
      .serialize(new UInt32(10000))
      .serialize(new Int64("100000"))
      .serialize(new Int64("-100000"))
      .serialize(new UInt64("100000"))
      .serialize(new Int128("100000"))
      .serialize(new Int128("-100000"))
      .serialize(new UInt128("100000"))
      .serialize(new Int160("100000"))
      .serialize(new Int160("-100000"))
      .serialize(new UInt160("100000"))
      .serialize(new Int256("100000"))
      .serialize(new Int256("-100000"))
      .serialize(new UInt256("100000"))
      .serialize(new TimestampType("1234567890"))
      .serialize(new BlockHeightType(123456))
      .serialize(multihash)
      .flip();

    expect(vb.deserialize(VariableBlob).equals(vb1)).toBe(true);
    expect(vb.deserialize(KoinosBoolean).toBoolean()).toBe(true);
    expect(vb.deserialize(KoinosString).toString()).toBe("test");
    expect(vb.deserialize(Int8).toNumber()).toBe(100);
    expect(vb.deserialize(UInt8).toNumber()).toBe(100);
    expect(vb.deserialize(Int16).toNumber()).toBe(1000);
    expect(vb.deserialize(UInt16).toNumber()).toBe(1000);
    expect(vb.deserialize(Int32).toNumber()).toBe(10000);
    expect(vb.deserialize(UInt32).toNumber()).toBe(10000);
    expect(vb.deserialize(Int64).toString()).toBe("100000");
    expect(vb.deserialize(Int64).toString()).toBe("-100000");
    expect(vb.deserialize(UInt64).toString()).toBe("100000");
    expect(vb.deserialize(Int128).toString()).toBe("100000");
    expect(vb.deserialize(Int128).toString()).toBe("-100000");
    expect(vb.deserialize(UInt128).toString()).toBe("100000");
    expect(vb.deserialize(Int160).toString()).toBe("100000");
    expect(vb.deserialize(Int160).toString()).toBe("-100000");
    expect(vb.deserialize(UInt160).toString()).toBe("100000");
    expect(vb.deserialize(Int256).toString()).toBe("100000");
    expect(vb.deserialize(Int256).toString()).toBe("-100000");
    expect(vb.deserialize(UInt256).toString()).toBe("100000");
    expect(vb.deserialize(TimestampType).toString()).toBe("1234567890");
    expect(vb.deserialize(BlockHeightType).toString()).toBe("123456");
    expect(vb.deserialize(Multihash).equals(multihash)).toBe(true);
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
