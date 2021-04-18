import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_UINT8 = 0xff;
export class Uint8 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Uint8", MAX_UINT8);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    new DataView(vb.buffer.buffer).setUint8(vb.offset, this.num);
    vb.offset += 1;
    if (!blob) vb.offset = 0;
    return vb;
  }

  static deserialize(vb: VariableBlob): Uint8 {
    if (vb.length() < vb.offset + 1) throw new Error("Unexpected EOF");
    const value = new DataView(vb.buffer.buffer).getUint8(vb.offset);
    vb.offset += 1;
    return new Uint8(value);
  }

  calcSerializedSize(): number {
    return 1;
  }
}

export default Uint8;
