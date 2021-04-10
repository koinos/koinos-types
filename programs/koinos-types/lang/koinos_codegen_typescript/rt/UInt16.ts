import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";
import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_UINT16 = 0xffff;
export class UInt16 extends KoinosNumber {
  constructor(
    number: bigint | string | KoinosNumber | KoinosBigInt | number = 0
  ) {
    super(number, "UInt16", MAX_UINT16);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(2);
    vb.buffer.writeUint16(this.num);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): UInt16 {
    if (vb.buffer.limit < 2) throw new Error("Unexpected EOF");
    const value = vb.buffer.readUint16();
    return new UInt16(value);
  }
}

export default UInt16;
