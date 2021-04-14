import { VariableBlob } from "./VariableBlob";
import { NumberLike } from "./Num";
import { BigNum } from "./BigNum";

export const MAX_UINT160 = BigInt("0x" + "F".repeat(40));
export class Uint160 extends BigNum {
  constructor(number: NumberLike = 0) {
    super(number, 160, MAX_UINT160);
  }

  static deserialize(vb: VariableBlob): Uint160 {
    const num = new Uint160().deserializeBigInt(vb);
    return new Uint160(num);
  }
}

export default Uint160;
