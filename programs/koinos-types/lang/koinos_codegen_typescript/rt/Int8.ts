import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_INT8 = 0x7f;
export const MIN_INT8 = -0x80;
export class Int8 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Int8", MAX_INT8, MIN_INT8);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.writeInt8(this.num);
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): Int8 {
    const value = vb.readInt8();
    return new Int8(value);
  }

  calcSerializedSize(): number {
    return 1;
  }
}

export default Int8;
