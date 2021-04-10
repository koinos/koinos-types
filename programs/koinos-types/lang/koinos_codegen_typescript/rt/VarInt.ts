import { KoinosBigInt } from "./KoinosBigInt";
import { KoinosNumber, NumberLike } from "./KoinosNumber";
import { VariableBlob } from "./VariableBlob";

export class VarInt {
  // TODO: consider using bigint instead of Number to construct the varint
  public num: number;

  constructor(number: NumberLike = 0) {
    let n: number;
    if (number instanceof KoinosBigInt) n = Number(number.num);
    else if (number instanceof KoinosNumber) n = number.num;
    else n = Number(number);
    this.num = n;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob();
    vb.buffer.writeVarint64(Number(this.num));
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): VarInt {
    if (vb.buffer.limit === 0) throw new Error("Unexpected EOF");
    const value = vb.buffer.readVarint64().toNumber();
    return new VarInt(value);
  }

  toNumber(): number {
    return this.num;
  }

  toString(radix?: number): string {
    return this.num.toString(radix);
  }
}

export default VarInt;
