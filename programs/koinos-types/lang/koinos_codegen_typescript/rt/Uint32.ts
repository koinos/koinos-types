import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_UINT32 = 0xffffffff;
export class Uint32 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Uint32", MAX_UINT32);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(4);
    vb.buffer.writeUint32(this.num);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): Uint32 {
    if (vb.buffer.limit < 4) throw new Error("Unexpected EOF");
    const value = vb.buffer.readUint32();
    return new Uint32(value);
  }
}

export default Uint32;
