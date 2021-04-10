import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";
import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_INT256 = BigInt("0x7" + "F".repeat(63));
export const MIN_INT256 = -BigInt("0x8" + "0".repeat(63));
export class Int256 extends KoinosBigInt {
  constructor(
    number: bigint | string | KoinosNumber | KoinosBigInt | number = 0
  ) {
    super(number, 256, MAX_INT256, MIN_INT256);
  }

  static deserialize(vb: VariableBlob): Int256 {
    const num = new Int256().deserializeBigInt(vb);
    return new Int256(num);
  }
}

export default Int256;
