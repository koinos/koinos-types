import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_INT16 = 0x7fff;
export const MIN_INT16 = -0x8000;
export class Int16 extends KoinosNumber {
  constructor(n = 0) {
    super(n, "Int16", MAX_INT16, MIN_INT16);
  }

  serialize(vb: VariableBlob): VariableBlob {
    vb.buffer.writeInt16(this.num);
    return vb;
  }

  static deserialize(vb: VariableBlob): Int16 {
    if (vb.buffer.limit < 2) throw new Error("Unexpected EOF");
    const value = vb.buffer.readInt16();
    return new Int16(value);
  }
}

export default Int16;
