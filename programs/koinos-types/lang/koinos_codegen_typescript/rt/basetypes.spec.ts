import * as Koinos from "./basetypes";

describe("Koinos Types - Typescript", () => {
  it("Serialize and desearialize", () => {
    expect.assertions(8);
    const vb = new Koinos.VariableBlob();
    new Koinos.KBoolean(true).serialize(vb);
    new Koinos.KString("test").serialize(vb);
    new Koinos.Int8(100).serialize(vb);
    new Koinos.UInt8(100).serialize(vb);
    new Koinos.Int16(1000).serialize(vb);
    new Koinos.UInt16(1000).serialize(vb);
    new Koinos.Int32(10000).serialize(vb);
    new Koinos.UInt32(10000).serialize(vb);
    //new Koinos.Int64("100000").serialize(vb);
    //new Koinos.UInt64("100000").serialize(vb);
    vb.buffer.flip();
    expect(vb.deserializeBoolean().toBoolean()).toBe(true);
    expect(vb.deserializeString().toString()).toBe("test");
    expect(vb.deserializeInt8().toNumber()).toBe(100);
    expect(vb.deserializeUInt8().toNumber()).toBe(100);
    expect(vb.deserializeInt16().toNumber()).toBe(1000);
    expect(vb.deserializeUInt16().toNumber()).toBe(1000);
    expect(vb.deserializeInt32().toNumber()).toBe(10000);
    expect(vb.deserializeUInt32().toNumber()).toBe(10000);
    //expect(vb.deserializeInt64().toString()).toBe("100000");
    //expect(vb.deserializeUInt64().toString()).toBe("100000");
  });
});