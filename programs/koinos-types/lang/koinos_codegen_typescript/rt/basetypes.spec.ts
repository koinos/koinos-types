import * as Koinos from "./basetypes";
import * as Long from 'long';

describe("Koinos Types - Typescript", () => {
  it("Serialize and desearialize", () => {
    expect.assertions(2);
    const vb = new Koinos.VariableBlob();
    const int32 = new Koinos.Int32(12345);
    int32.serialize(vb);
    const str = new Koinos.KString("test");
    str.serialize(vb);
    vb.buffer.flip();    
    expect(vb.deserializeInt32().toNumber()).toBe(12345);
    expect(vb.deserializeString().toString()).toBe("test");
  });
});