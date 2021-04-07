import { KoinosBigInt } from "./KoinosBigInt";
import { VariableBlob } from "./VariableBlob";

export const MAX_UINT128 = BigInt("0x" + "F".repeat(32));
export class UInt128 extends KoinosBigInt {
  constructor(number: bigint | string | number = 0) {
    super(number, 128, MAX_UINT128);
  }

  static deserialize(vb: VariableBlob): UInt128 {
    const num = new UInt128().deserializeBigInt(vb);
    return new UInt128(num);
  }
}

export default UInt128;
