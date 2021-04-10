import { KoinosBigInt } from "./KoinosBigInt";

export class KoinosNumber {
  public num: number;

  constructor(
    number: bigint | string | KoinosNumber | KoinosBigInt | number = 0,
    name: string,
    max: number,
    min = 0
  ) {
    let n: number;
    if (number instanceof KoinosBigInt) n = Number(number.num);
    else if (number instanceof KoinosNumber) n = number.num;
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

export default KoinosNumber;
