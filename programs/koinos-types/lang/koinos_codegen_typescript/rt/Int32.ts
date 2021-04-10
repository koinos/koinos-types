import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";
import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_INT32 = 0x7fffffff;
export const MIN_INT32 = -0x80000000;
export class Int32 extends KoinosNumber {
  constructor(
    number: bigint | string | KoinosNumber | KoinosBigInt | number = 0
  ) {
    super(number, "Int32", MAX_INT32, MIN_INT32);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(4);
    vb.buffer.writeInt32(this.num);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): Int32 {
    if (vb.buffer.limit < 4) throw new Error("Unexpected EOF");
    const value = vb.buffer.readInt32();
    return new Int32(value);
  }
}

export default Int32;
