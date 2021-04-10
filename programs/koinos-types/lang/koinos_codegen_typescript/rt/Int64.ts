import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";
import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_INT64 = BigInt("0x7" + "F".repeat(15));
export const MIN_INT64 = -BigInt("0x8" + "0".repeat(15));
export class Int64 extends KoinosBigInt {
  constructor(
    number: bigint | string | KoinosNumber | KoinosBigInt | number = 0
  ) {
    super(number, 64, MAX_INT64, MIN_INT64);
  }

  static deserialize(vb: VariableBlob): Int64 {
    const num = new Int64().deserializeBigInt(vb);
    return new Int64(num);
  }
}

export default Int64;
