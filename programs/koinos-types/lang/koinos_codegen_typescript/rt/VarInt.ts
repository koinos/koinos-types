import { BigNum, isInt64 } from "./BigNum";
import { Num, NumberLike } from "./Num";
import { VariableBlob } from "./VariableBlob";

const ZERO = BigInt(0);
const ONE = BigInt(1);
const SEVEN = BigInt(7);
const MSB_ONE = BigInt(128);
const MAX_SAFE_INTEGER = BigInt(Number.MAX_SAFE_INTEGER);

export function sizeVarInt(n: number): number {
  return n === 0 ? 1 : Math.ceil(Math.log2(n + 1) / 7);
}

export class VarInt {
  public num: bigint;

  constructor(number: NumberLike = 0) {
    let n: bigint;
    if (number instanceof BigNum) n = number.num;
    else if (number instanceof Num) n = BigInt(number.num);
    else n = BigInt(number);
    this.num = n;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    if (this.num === ZERO) {
      vb.writeUint8(0);
      if (!blob) vb.resetCursor();
      return vb;
    }

    let n = this.num;
    let shift7 = this.num;
    while (n > ZERO) {
      shift7 = n >> SEVEN;
      const group7 = n - (shift7 << SEVEN);
      const byte = shift7 > ZERO ? MSB_ONE + group7 : group7;
      vb.writeUint8(Number(byte));
      n = shift7;
    }
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): VarInt {
    let i = ZERO;
    let n = ZERO;
    let endVarInt = false;
    while (!endVarInt) {
      const byte = BigInt(vb.readUint8());
      endVarInt = byte < MSB_ONE;
      const mod = endVarInt ? byte : byte - MSB_ONE;
      n += mod << (SEVEN * i);
      i += ONE;
    }
    return new VarInt(n);
  }

  calcSerializedSize(): number {
    if (this.num + ONE <= MAX_SAFE_INTEGER) return sizeVarInt(Number(this.num));
    return Math.ceil(this.num.toString(2).length / 7);
  }

  toNumber(): number {
    return Number(this.num);
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

export default VarInt;
