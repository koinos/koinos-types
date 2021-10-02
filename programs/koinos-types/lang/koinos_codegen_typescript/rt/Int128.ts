import { VariableBlob } from "./VariableBlob";
import { NumberLike } from "./Num";
import { BigNum } from "./BigNum";

export const MAX_INT128 = BigInt("0x7" + "F".repeat(31));
export const MIN_INT128 = -BigInt("0x8" + "0".repeat(31));
export class Int128 extends BigNum {
  constructor(number: NumberLike = 0) {
    super(number, 128, MAX_INT128, MIN_INT128);
  }

  static deserialize(vb: VariableBlob): Int128 {
    const num = new Int128().deserializeBigInt(vb);
    return new Int128(num);
  }
}

export default Int128;
