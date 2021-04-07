import { VariableBlob } from "./VariableBlob";

export class KoinosBigInt {
  public num: bigint;

  private bytes: number;

  private unsigned: boolean;

  constructor(
    number: bigint | string | number,
    bits: number,
    max: bigint,
    min = BigInt(0)
  ) {
    const n = BigInt(number);
    this.unsigned = min === BigInt(0);
    this.bytes = bits / 8;
    this.num = n;
    if (n < min || n > max)
      throw new Error(`${this.unsigned ? "U" : ""}Int${bits} is out of bounds`);
  }

  serialize(vb: VariableBlob): VariableBlob {
    let numString: string;
    if (this.num >= BigInt(0)) {
      numString = this.num.toString(16);
      numString = "0".repeat(2 * this.bytes - numString.length) + numString;
    } else {
      numString = (
        BigInt("0x1" + "0".repeat(2 * this.bytes)) + this.num
      ).toString(16);
    }
    for (let i = 0; i < 2 * this.bytes; i += 8) {
      const int32 = Number("0x" + numString.substring(i, i + 8));
      vb.buffer.writeUint32(int32);
    }
    return vb;
  }

  deserializeBigInt(vb: VariableBlob): bigint {
    if (vb.buffer.limit < this.bytes) throw new Error("Unexpected EOF");
    let numString = "0x";
    for (let i = 0; i < this.bytes / 4; i += 1) {
      const uint32Str = vb.buffer.readUint32().toString(16);
      numString += "0".repeat(8 - uint32Str.length) + uint32Str;
    }
    if (!this.unsigned && Number(numString.substring(0, 3)) >= 8) {
      // signed number and starts with bit 1 (negative number)
      return BigInt(numString) - BigInt("0x1" + "0".repeat(2 * this.bytes));
    }
    return BigInt(numString);
  }

  toBigInt(): bigint {
    return this.num;
  }

  toString(radix?: number): string {
    return this.num.toString(radix);
  }
}

export default KoinosBigInt;
