import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_INT8 = 0x7f;
export const MIN_INT8 = -0x80;
export class Int8 extends KoinosNumber {
  constructor(n = 0) {
    super(n, "Int8", MAX_INT8, MIN_INT8);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(1);
    vb.buffer.writeByte(this.num);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): Int8 {
    if (vb.buffer.limit === 0) throw new Error("Unexpected EOF");
    const value = vb.buffer.readByte();
    return new Int8(value);
  }
}

export default Int8;
