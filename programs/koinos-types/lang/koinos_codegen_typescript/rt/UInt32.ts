import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";

export const MAX_UINT32 = 0xffffffff;
export class UInt32 extends KoinosNumber {
  constructor(n = 0) {
    super(n, "UInt32", MAX_UINT32);
  }

  serialize(vb: VariableBlob): VariableBlob {
    vb.buffer.writeUint32(this.num);
    return vb;
  }
}

export default UInt32;
