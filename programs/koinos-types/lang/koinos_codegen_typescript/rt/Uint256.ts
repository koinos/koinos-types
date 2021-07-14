import { VariableBlob } from "./VariableBlob";
import { NumberLike } from "./Num";
import { BigNum } from "./BigNum";

export const MAX_UINT256 = BigInt("0x" + "F".repeat(64));
export class Uint256 extends BigNum {
  constructor(number: NumberLike = 0) {
    super(number, 256, MAX_UINT256);
  }

  static deserialize(vb: VariableBlob): Uint256 {
    const num = new Uint256().deserializeBigInt(vb);
    return new Uint256(num);
  }
}

export default Uint256;
