import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_INT8 = 0x7f;
export const MIN_INT8 = -0x80;
export class Int8 extends KoinosNumber {
  constructor(n = 0) {
    super(n, "Int8", MAX_INT8, MIN_INT8);
  }

  serialize(vb: VariableBlob): VariableBlob {
    vb.buffer.writeByte(this.num);
    return vb;
  }

  static deserialize(vb: VariableBlob): Int8 {
    if (vb.buffer.limit === 0) throw new Error("Unexpected EOF");
    const value = vb.buffer.readByte();
    return new Int8(value);
  }
}

export default Int8;
