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
} from "./";

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
    expect(vb.deserialize(Multihash).equals(m)).toBe(true);
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