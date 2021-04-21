import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_UINT16 = 0xffff;
export class Uint16 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Uint16", MAX_UINT16);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.writeUint16(this.num);
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): Uint16 {
    const value = vb.readUint16();
    return new Uint16(value);
  }

  calcSerializedSize(): number {
    return 2;
  }
}

export default Uint16;
