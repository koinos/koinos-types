import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_UINT32 = 0xffffffff;
export class Uint32 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Uint32", MAX_UINT32);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    new DataView(vb.buffer.buffer).setUint32(vb.offset, this.num);
    vb.offset += 4;
    if (!blob) vb.offset = 0;
    return vb;
  }

  static deserialize(vb: VariableBlob): Uint32 {
    if (vb.length() < vb.offset + 4) throw new Error("Unexpected EOF");
    const value = new DataView(vb.buffer.buffer).getUint32(vb.offset);
    vb.offset += 4;
    return new Uint32(value);
  }

  calcSerializedSize(): number {
    return 4;
  }
}

export default Uint32;
