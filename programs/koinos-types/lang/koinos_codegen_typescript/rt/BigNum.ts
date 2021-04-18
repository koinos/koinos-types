import { VariableBlob } from "./VariableBlob";
import { Num, NumberLike } from "./Num";
import { VarInt } from "./VarInt";

const MAX_INT64 = BigInt("0x7" + "F".repeat(15));
const MIN_INT64 = -BigInt("0x8" + "0".repeat(15));
function isInt64(n: bigint): boolean {
  return n >= MIN_INT64 && n <= MAX_INT64;
}

export class BigNum {
  public num: bigint;

  private bytes: number;

  private unsigned: boolean;

  constructor(number: NumberLike, bits: number, max: bigint, min = BigInt(0)) {
    let n: bigint;
    if (number instanceof BigNum) n = number.num;
    else if (number instanceof Num) n = BigInt(number.num);
    else if (number instanceof VarInt) n = BigInt(number.num);
    else n = BigInt(number);
    this.unsigned = min === BigInt(0);
    this.bytes = bits / 8;
    this.num = n;
    if (n < min || n > max)
      throw new Error(`${this.unsigned ? "U" : ""}Int${bits} is out of bounds`);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
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
      const uint32 = Number("0x" + numString.substring(i, i + 8));
      new DataView(vb.buffer.buffer).setUint32(vb.offset, uint32);
      vb.offset += 4;
    }
    if (!blob) vb.offset = 0;
    return vb;
  }

  deserializeBigInt(vb: VariableBlob): bigint {
    if (vb.length() < vb.offset + this.bytes) throw new Error("Unexpected EOF");
    let numString = "0x";
    for (let i = 0; i < this.bytes / 4; i += 1) {
      const uint32Str = new DataView(vb.buffer.buffer).getUint32(vb.offset).toString(16);
      vb.offset += 4;
      numString += "0".repeat(8 - uint32Str.length) + uint32Str;
    }
    if (!this.unsigned && Number(numString.substring(0, 3)) >= 8) {
      // signed number and starts with bit 1 (negative number)
      return BigInt(numString) - BigInt("0x1" + "0".repeat(2 * this.bytes));
    }
    return BigInt(numString);
  }

  calcSerializedSize(): number {
    return this.bytes;
  }

  toBigInt(): bigint {
    return this.num;
  }

  toJSON(): bigint | string {
    return isInt64(this.num) ? this.num : this.num.toString();
  }

  toString(radix?: number): string {
    return this.num.toString(radix);
  }
}

export default BigNum;
