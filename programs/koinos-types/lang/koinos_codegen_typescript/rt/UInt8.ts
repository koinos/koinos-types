import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_UINT8 = 0xFF;
export class UInt8 extends KoinosNumber {
  constructor(n: number = 0) {
    super(n, "UInt8", MAX_UINT8);
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeByte(this.num);
  }
}

export default UInt8;