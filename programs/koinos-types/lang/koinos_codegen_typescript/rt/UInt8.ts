import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_UINT8 = 0xff;
export class UInt8 extends KoinosNumber {
  constructor(n = 0) {
    super(n, "UInt8", MAX_UINT8);
  }

  serialize(vb: VariableBlob): VariableBlob {
    vb.buffer.writeByte(this.num);
    return vb;
  }

  static deserialize(vb: VariableBlob): UInt8 {
    if (vb.buffer.limit === 0) throw new Error("Unexpected EOF");
    const value = vb.buffer.readByte();
    return new UInt8(value);
  }
}

export default UInt8;
