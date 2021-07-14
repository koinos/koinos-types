import { VariableBlob } from "./VariableBlob";
import { NumberLike } from "./Num";
import { BigNum } from "./BigNum";

export const MAX_INT160 = BigInt("0x7" + "F".repeat(39));
export const MIN_INT160 = -BigInt("0x8" + "0".repeat(39));
export class Int160 extends BigNum {
  constructor(number: NumberLike = 0) {
    super(number, 160, MAX_INT160, MIN_INT160);
  }

  static deserialize(vb: VariableBlob): Int160 {
    const num = new Int160().deserializeBigInt(vb);
    return new Int160(num);
  }
}

export default Int160;
