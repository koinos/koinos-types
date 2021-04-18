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
    new DataView(vb.buffer.buffer).setInt32(vb.offset, this.num);
    vb.offset += 4;
    if (!blob) vb.offset = 0;
    return vb;
  }

  static deserialize(vb: VariableBlob): Int32 {
    if (vb.length() < vb.offset + 4) throw new Error("Unexpected EOF");
    const value = new DataView(vb.buffer.buffer).getInt16(vb.offset);
    vb.offset += 4;
    return new Int32(value);
  }

  calcSerializedSize(): number {
    return 4;
  }
}

export default Int32;
