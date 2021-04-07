import { KoinosBigInt } from "./KoinosBigInt";
import { VariableBlob } from "./VariableBlob";

export const MAX_UINT64 = BigInt("0x" + "F".repeat(16));
export class UInt64 extends KoinosBigInt {
  constructor(number: bigint | string | number = 0) {
    super(number, 64, MAX_UINT64);
  }

  static deserialize(vb: VariableBlob): UInt64 {
    const num = new UInt64().deserializeBigInt(vb);
    return new UInt64(num);
  }
}

export default UInt64;
