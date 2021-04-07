import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_INT8 = 0x7F;
export const MIN_INT8 = -0x80;
export class Int8 extends KoinosNumber {
  constructor(n: number = 0) {
    super(n, "Int8", MAX_INT8, MIN_INT8);
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeByte(this.num);
  }
}

export default Int8;