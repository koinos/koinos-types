import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_UINT16 = 0xffff;
export class UInt16 extends KoinosNumber {
  constructor(n = 0) {
    super(n, "UInt16", MAX_UINT16);
  }

  serialize(vb: VariableBlob): VariableBlob {
    vb.buffer.writeUint16(this.num);
    return vb;
  }

  static deserialize(vb: VariableBlob): UInt16 {
    if (vb.buffer.limit < 2) throw new Error("Unexpected EOF");
    const value = vb.buffer.readUint16();
    return new UInt16(value);
  }
}

export default UInt16;
