import { KoinosBigInt } from "./KoinosBigInt";
import { VariableBlob } from "./VariableBlob";

export const MAX_UINT160 = BigInt("0x" + "F".repeat(40));
export class UInt160 extends KoinosBigInt {
  constructor(number: bigint | string | number = 0) {
    super(number, 160, MAX_UINT160);
  }

  static deserialize(vb: VariableBlob): UInt160 {
    const num = new UInt160().deserializeBigInt(vb);
    return new UInt160(num);
  }
}

export default UInt160;
