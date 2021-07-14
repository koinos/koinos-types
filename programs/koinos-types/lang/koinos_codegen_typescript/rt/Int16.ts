import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_INT16 = 0x7fff;
export const MIN_INT16 = -0x8000;
export class Int16 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Int16", MAX_INT16, MIN_INT16);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.writeInt16(this.num);
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): Int16 {
    const value = vb.readInt16();
    return new Int16(value);
  }

  calcSerializedSize(): number {
    return 2;
  }
}

export default Int16;
