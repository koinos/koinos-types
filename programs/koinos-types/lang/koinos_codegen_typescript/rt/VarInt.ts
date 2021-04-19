import { BigNum } from "./BigNum";
import { Num, NumberLike } from "./Num";
import { VariableBlob } from "./VariableBlob";

export class VarInt {
  // TODO: consider using bigint instead of Number to construct the varint
  public num: number;

  constructor(number: NumberLike = 0) {
    let n: number;
    if (number instanceof BigNum) n = Number(number.num);
    else if (number instanceof Num) n = number.num;
    else n = Number(number);
    this.num = n;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    if (this.num === 0) {
      vb.writeUint8(0);
      return vb;
    }

    let n = this.num;
    let shift7 = this.num;
    while (n > 0) {
      shift7 = n >> 7;
      const group7 = n - (shift7 << 7);
      if (shift7 > 0) vb.writeUint8(group7 + 128);
      else vb.writeUint8(group7);
      n = shift7;
    }
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): VarInt {
    let i = 0;
    let n = 0;
    let endVarInt = false;
    while (!endVarInt) {
      const byte = vb.readUint8();
      endVarInt = byte < 128;
      const mod = endVarInt ? byte : byte - 128;
      n += mod << (7 * i);
      i += 1;
    }
    return new VarInt(n);
  }

  calcSerializedSize(): number {
    return Math.ceil(Math.log2(this.num + 1) / 7);
  }

  toNumber(): number {
    return this.num;
  }

  toJSON(): number {
    return this.num;
  }

  toString(radix?: number): string {
    return this.num.toString(radix);
  }
}

export default VarInt;
