import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_UINT8 = 0xff;
export class Uint8 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Uint8", MAX_UINT8);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(1);
    vb.buffer.writeByte(this.num);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): Uint8 {
    if (vb.buffer.limit === 0) throw new Error("Unexpected EOF");
    const value = vb.buffer.readByte();
    return new Uint8(value);
  }
}

export default Uint8;
