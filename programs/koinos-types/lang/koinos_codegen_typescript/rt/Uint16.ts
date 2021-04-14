import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_UINT16 = 0xffff;
export class Uint16 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Uint16", MAX_UINT16);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(2);
    vb.buffer.writeUint16(this.num);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): Uint16 {
    if (vb.buffer.limit < 2) throw new Error("Unexpected EOF");
    const value = vb.buffer.readUint16();
    return new Uint16(value);
  }
}

export default Uint16;
