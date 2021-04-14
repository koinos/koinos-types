import { BigNum } from "./BigNum";
import { VarInt } from "./VarInt";

export type NumberLike = number | bigint | string | Num | BigNum | VarInt;

export class Num {
  public num: number;

  constructor(number: NumberLike = 0, name: string, max: number, min = 0) {
    let n: number;
    if (number instanceof BigNum) n = Number(number.num);
    else if (number instanceof Num) n = number.num;
    else if (number instanceof VarInt) n = number.num;
    else n = Number(number);
    if (n < min || n > max) throw new Error(`${name} is out of bounds`);
    this.num = n;
  }

  toNumber(): number {
    return this.num;
  }

  toString(radix?: number): string {
    return this.num.toString(radix);
  }
}

export default Num;
