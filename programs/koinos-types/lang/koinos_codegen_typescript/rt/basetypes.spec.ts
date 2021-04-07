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
    expect.assertions(2);
    const vb1 = new VariableBlob(9);
    const vb2 = new VariableBlob(5);
    const vb3 = new VariableBlob(5);

    const num1 = new Int16(1000);
    const num2 = new Int64(1000);
    num1.serialize(vb1);
    num1.serialize(vb2);
    num2.serialize(vb3);
    vb1.flip();
    vb2.flip();
    vb3.flip();
    expect(vb1.equals(vb2)).toBe(true);
    expect(vb1.equals(vb3)).toBe(false);
  });

  it("Serialize and desearialize", () => {
    expect.assertions(24);
    const vb = new VariableBlob();

    const vb1 = new VariableBlob();
    new KoinosString("test variable blob").serialize(vb1);

    vb1.flip();
    vb1.serialize(vb);
    new KoinosBoolean(true).serialize(vb);
    new KoinosString("test").serialize(vb);
    new Int8(100).serialize(vb);
    new UInt8(100).serialize(vb);
    new Int16(1000).serialize(vb);
    new UInt16(1000).serialize(vb);
    new Int32(10000).serialize(vb);
    new UInt32(10000).serialize(vb);
    new Int64("100000").serialize(vb);
    new Int64("-100000").serialize(vb);
    new UInt64("100000").serialize(vb);
    new Int128("100000").serialize(vb);
    new Int128("-100000").serialize(vb);
    new UInt128("100000").serialize(vb);
    new Int160("100000").serialize(vb);
    new Int160("-100000").serialize(vb);
    new UInt160("100000").serialize(vb);
    new Int256("100000").serialize(vb);
    new Int256("-100000").serialize(vb);
    new UInt256("100000").serialize(vb);
    new TimestampType("1234567890").serialize(vb);
    new BlockHeightType(123456).serialize(vb);
    const m = new Multihash();
    m.id = new UInt64(123);
    m.digest = new VariableBlob();
    new KoinosString("digest").serialize(m.digest);
    m.digest.flip();
    m.serialize(vb);

    vb.flip();
    expect(VariableBlob.deserialize(vb).equals(vb1)).toBe(true);
    expect(KoinosBoolean.deserialize(vb).toBoolean()).toBe(true);
    expect(KoinosString.deserialize(vb).toString()).toBe("test");
    expect(Int8.deserialize(vb).toNumber()).toBe(100);
    expect(UInt8.deserialize(vb).toNumber()).toBe(100);
    expect(Int16.deserialize(vb).toNumber()).toBe(1000);
    expect(UInt16.deserialize(vb).toNumber()).toBe(1000);
    expect(Int32.deserialize(vb).toNumber()).toBe(10000);
    expect(UInt32.deserialize(vb).toNumber()).toBe(10000);
    expect(Int64.deserialize(vb).toString()).toBe("100000");
    expect(Int64.deserialize(vb).toString()).toBe("-100000");
    expect(UInt64.deserialize(vb).toString()).toBe("100000");
    expect(Int128.deserialize(vb).toString()).toBe("100000");
    expect(Int128.deserialize(vb).toString()).toBe("-100000");
    expect(UInt128.deserialize(vb).toString()).toBe("100000");
    expect(Int160.deserialize(vb).toString()).toBe("100000");
    expect(Int160.deserialize(vb).toString()).toBe("-100000");
    expect(UInt160.deserialize(vb).toString()).toBe("100000");
    expect(Int256.deserialize(vb).toString()).toBe("100000");
    expect(Int256.deserialize(vb).toString()).toBe("-100000");
    expect(UInt256.deserialize(vb).toString()).toBe("100000");
    expect(TimestampType.deserialize(vb).toString()).toBe("1234567890");
    expect(BlockHeightType.deserialize(vb).toString()).toBe("123456");
    expect(Multihash.deserialize(vb).equals(m)).toBe(true);
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
