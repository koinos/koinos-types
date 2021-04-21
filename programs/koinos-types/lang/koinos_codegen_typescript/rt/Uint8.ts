import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";

export const MAX_UINT8 = 0xff;
export class Uint8 extends Num {
  constructor(number: NumberLike = 0) {
    super(number, "Uint8", MAX_UINT8);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.writeUint8(this.num);
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): Uint8 {
    const value = vb.readUint8();
    return new Uint8(value);
  }

  calcSerializedSize(): number {
    return 1;
  }
}

export default Uint8;
