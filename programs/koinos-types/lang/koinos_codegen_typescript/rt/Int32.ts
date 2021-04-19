import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_INT32 = 0x7fffffff;
export const MIN_INT32 = -0x80000000;
export class Int32 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Int32", MAX_INT32, MIN_INT32);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.writeInt32(this.num);
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): Int32 {
    const value = vb.readInt32();
    return new Int32(value);
  }

  calcSerializedSize(): number {
    return 4;
  }
}

export default Int32;
