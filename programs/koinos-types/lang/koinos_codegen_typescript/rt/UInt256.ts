import { VariableBlob } from "./VariableBlob";
import { NumberLike } from "./KoinosNumber";
import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_UINT256 = BigInt("0x" + "F".repeat(64));
export class UInt256 extends KoinosBigInt {
  constructor(number: NumberLike = 0) {
    super(number, 256, MAX_UINT256);
  }

  static deserialize(vb: VariableBlob): UInt256 {
    const num = new UInt256().deserializeBigInt(vb);
    return new UInt256(num);
  }
}

export default UInt256;
