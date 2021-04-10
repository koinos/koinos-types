import { VariableBlob } from "./VariableBlob";
import { KoinosNumber } from "./KoinosNumber";
import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_UINT160 = BigInt("0x" + "F".repeat(40));
export class UInt160 extends KoinosBigInt {
  constructor(
    number: bigint | string | KoinosNumber | KoinosBigInt | number = 0
  ) {
    super(number, 160, MAX_UINT160);
  }

  static deserialize(vb: VariableBlob): UInt160 {
    const num = new UInt160().deserializeBigInt(vb);
    return new UInt160(num);
  }
}

export default UInt160;
