import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_UINT8 = 0xff;
export class UInt8 extends KoinosNumber {
  constructor(n = 0) {
    super(n, "UInt8", MAX_UINT8);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(1);
    vb.buffer.writeByte(this.num);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): UInt8 {
    if (vb.buffer.limit === 0) throw new Error("Unexpected EOF");
    const value = vb.buffer.readByte();
    return new UInt8(value);
  }
}

export default UInt8;
