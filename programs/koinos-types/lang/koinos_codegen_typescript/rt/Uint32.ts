import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_UINT32 = 0xffffffff;
export class Uint32 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Uint32", MAX_UINT32);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.writeUint32(this.num);
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): Uint32 {
    const value = vb.readUint32();
    return new Uint32(value);
  }

  calcSerializedSize(): number {
    return 4;
  }
}

export default Uint32;
