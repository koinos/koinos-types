import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_UINT16 = 0xFFFF;
export class UInt16 extends KoinosNumber {
  constructor(n: number = 0) {
    super(n, "UInt16", MAX_UINT16);
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeUint16(this.num);
  }
}

export default UInt16;