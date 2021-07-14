import { VariableBlob } from "./VariableBlob";
import { NumberLike } from "./Num";
import { BigNum } from "./BigNum";

export const MAX_UINT128 = BigInt("0x" + "F".repeat(32));
export class Uint128 extends BigNum {
  constructor(number: NumberLike = 0) {
    super(number, 128, MAX_UINT128);
  }

  static deserialize(vb: VariableBlob): Uint128 {
    const num = new Uint128().deserializeBigInt(vb);
    return new Uint128(num);
  }
}

export default Uint128;
