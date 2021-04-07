import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_INT32 = 0x7fffffff;
export const MIN_INT32 = -0x80000000;
export class Int32 extends KoinosNumber {
  constructor(n = 0) {
    super(n, "Int32", MAX_INT32, MIN_INT32);
  }

  serialize(vb: VariableBlob): VariableBlob {
    vb.buffer.writeInt32(this.num);
    return vb;
  }
}

export default Int32;
