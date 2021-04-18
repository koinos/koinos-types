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
    new DataView(vb.buffer.buffer).setInt16(vb.offset, this.num);
    vb.offset += 2;
    if (!blob) vb.offset = 0;
    return vb;
  }

  static deserialize(vb: VariableBlob): Int16 {
    if (vb.length() < vb.offset + 2) throw new Error("Unexpected EOF");
    const value = new DataView(vb.buffer.buffer).getInt16(vb.offset);
    vb.offset += 2;
    return new Int16(value);
  }

  calcSerializedSize(): number {
    return 2;
  }
}

export default Int16;
