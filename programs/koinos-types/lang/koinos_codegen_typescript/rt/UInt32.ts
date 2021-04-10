import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";
import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_UINT32 = 0xffffffff;
export class UInt32 extends KoinosNumber {
  constructor(
    number: bigint | string | KoinosNumber | KoinosBigInt | number = 0
  ) {
    super(number, "UInt32", MAX_UINT32);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(4);
    vb.buffer.writeUint32(this.num);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): UInt32 {
    if (vb.buffer.limit < 4) throw new Error("Unexpected EOF");
    const value = vb.buffer.readUint32();
    return new UInt32(value);
  }
}

export default UInt32;
