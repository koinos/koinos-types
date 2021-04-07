import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_INT16 = 0x7FFF;
export const MIN_INT16 = -0x8000;
export class Int16 extends KoinosNumber {
  constructor(n: number = 0) {
    super(n, "Int16", MAX_INT16, MIN_INT16);
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeInt16(this.num);
  }
}

export default Int16;