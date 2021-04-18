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
    new DataView(vb.buffer.buffer).setInt8(vb.offset, this.num);
    vb.offset += 1;
    if (!blob) vb.offset = 0;
    return vb;
  }

  static deserialize(vb: VariableBlob): Int8 {
    if (vb.length() < vb.offset + 1) throw new Error("Unexpected EOF");
    const value = new DataView(vb.buffer.buffer).getInt8(vb.offset);
    vb.offset += 1;
    return new Int8(value);
  }

  calcSerializedSize(): number {
    return 1;
  }
}

export default Int8;
