import * as Koinos from "./basetypes";

describe("Koinos Types - Typescript", () => {
  it("should compare VariableBlobs", () => {
    expect.assertions(2);
    const vb1 = new Koinos.VariableBlob(9);
    const vb2 = new Koinos.VariableBlob(5);
    const vb3 = new Koinos.VariableBlob(5);
    
    const num1 = new Koinos.Int16(1000);
    const num2 = new Koinos.Int64(1000);
    num1.serialize(vb1);
    num1.serialize(vb2);
    num2.serialize(vb3);
    vb1.buffer.flip();
    vb2.buffer.flip();
    vb3.buffer.flip();
    expect(vb1.equals(vb2)).toBe(true);
    expect(vb1.equals(vb3)).toBe(false);
  });
  it("Serialize and desearialize", () => {
    expect.assertions(24);
    const vb = new Koinos.VariableBlob();
    
    const vb1 = new Koinos.VariableBlob();
    new Koinos.KString("test variable blob").serialize(vb1);

    vb1.buffer.flip();
    vb1.serialize(vb);
    new Koinos.KBoolean(true).serialize(vb);
    new Koinos.KString("test").serialize(vb);
    new Koinos.Int8(100).serialize(vb);
    new Koinos.UInt8(100).serialize(vb);
    new Koinos.Int16(1000).serialize(vb);
    new Koinos.UInt16(1000).serialize(vb);
    new Koinos.Int32(10000).serialize(vb);
    new Koinos.UInt32(10000).serialize(vb);
    new Koinos.Int64("100000").serialize(vb);
    new Koinos.Int64("-100000").serialize(vb);
    new Koinos.UInt64("100000").serialize(vb);
    new Koinos.Int128("100000").serialize(vb);
    new Koinos.Int128("-100000").serialize(vb);
    new Koinos.UInt128("100000").serialize(vb);
    new Koinos.Int160("100000").serialize(vb);
    new Koinos.Int160("-100000").serialize(vb);
    new Koinos.UInt160("100000").serialize(vb);
    new Koinos.Int256("100000").serialize(vb);
    new Koinos.Int256("-100000").serialize(vb);
    new Koinos.UInt256("100000").serialize(vb);
    new Koinos.TimestampType("1234567890").serialize(vb);
    new Koinos.BlockHeightType(123456).serialize(vb);
    const m = new Koinos.Multihash();
    m.id = new Koinos.UInt64(123);
    m.digest = new Koinos.VariableBlob();
    new Koinos.KString("digest").serialize(m.digest);
    m.digest.buffer.flip();
    m.serialize(vb);

    vb.buffer.flip();
    expect(vb.deserialize(Koinos.VariableBlob).equals(vb1)).toBe(true);
    expect(vb.deserialize(Koinos.KBoolean).toBoolean()).toBe(true);
    expect(vb.deserializeString().toString()).toBe("test");
    expect(vb.deserializeInt8().toNumber()).toBe(100);
    expect(vb.deserializeUInt8().toNumber()).toBe(100);
    expect(vb.deserializeInt16().toNumber()).toBe(1000);
    expect(vb.deserializeUInt16().toNumber()).toBe(1000);
    expect(vb.deserializeInt32().toNumber()).toBe(10000);
    expect(vb.deserializeUInt32().toNumber()).toBe(10000);
    expect(vb.deserializeInt64().toString()).toBe("100000");
    expect(vb.deserializeInt64().toString()).toBe("-100000");
    expect(vb.deserializeUInt64().toString()).toBe("100000");
    expect(vb.deserializeInt128().toString()).toBe("100000");
    expect(vb.deserializeInt128().toString()).toBe("-100000");
    expect(vb.deserializeUInt128().toString()).toBe("100000");
    expect(vb.deserializeInt160().toString()).toBe("100000");
    expect(vb.deserializeInt160().toString()).toBe("-100000");
    expect(vb.deserializeUInt160().toString()).toBe("100000");
    expect(vb.deserializeInt256().toString()).toBe("100000");
    expect(vb.deserializeInt256().toString()).toBe("-100000");
    expect(vb.deserializeUInt256().toString()).toBe("100000");
    expect(vb.deserializeTimestampType().toString()).toBe("1234567890");
    expect(vb.deserializeBlockHeightType().toString()).toBe("123456");
    expect(vb.deserializeMultihash().equals(m)).toBe(true);
  });

  it("should create an opaque class", () => {
    expect.assertions(10);
    const num = new Koinos.Int32(123456);
    const opaque = new Koinos.Opaque(Koinos.Int32, num);
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