import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_UINT32 = 0xFFFFFFFF;
export class UInt32 extends KoinosNumber {
  constructor(n: number = 0) {
    super(n, "UInt32", MAX_UINT32);
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeUint32(this.num);
  }
}

export default UInt32;