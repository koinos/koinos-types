export class KoinosNumber {
  public num: number;

  constructor(n = 0, name: string, max: number, min = 0) {
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
