import { VariableBlob } from "./VariableBlob";
import { NumberLike } from "./Num";
import { BigNum } from "./BigNum";

export const MAX_UINT64 = BigInt("0x" + "F".repeat(16));
export class Uint64 extends BigNum {
  constructor(number: NumberLike = 0) {
    super(number, 64, MAX_UINT64);
  }

  static deserialize(vb: VariableBlob): Uint64 {
    const num = new Uint64().deserializeBigInt(vb);
    return new Uint64(num);
  }
}

export default Uint64;
