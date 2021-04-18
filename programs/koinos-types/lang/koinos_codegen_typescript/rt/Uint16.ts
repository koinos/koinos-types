import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_UINT16 = 0xffff;
export class Uint16 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Uint16", MAX_UINT16);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    new DataView(vb.buffer.buffer).setUint16(vb.offset, this.num);
    vb.offset += 2;
    if (!blob) vb.offset = 0;
    return vb;
  }

  static deserialize(vb: VariableBlob): Uint16 {
    if (vb.length() < vb.offset + 2) throw new Error("Unexpected EOF");
    const value = new DataView(vb.buffer.buffer).getUint16(vb.offset);
    vb.offset += 2;
    return new Uint16(value);
  }

  calcSerializedSize(): number {
    return 2;
  }
}

export default Uint16;
