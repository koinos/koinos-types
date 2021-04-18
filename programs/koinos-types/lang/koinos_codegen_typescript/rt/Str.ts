import * as ByteBuffer from "bytebuffer";
import { VariableBlob } from "./VariableBlob";

export type StringLike = string | Str;

export class Str {
  public str: string;

  constructor(str: StringLike = "") {
    this.str = str instanceof Str ? str.str : str;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    const buffer = ByteBuffer.fromUTF8(this.str) as ByteBuffer;
    vb.buffer.writeVarint64(buffer.limit).append(buffer);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): Str {
    const subvb = vb.deserialize(VariableBlob);
    return new Str(subvb.buffer.toUTF8());
  }

  calcSerializedSize(): number {
    const size = new TextEncoder().encode(this.str).length;
    const header = Math.ceil(Math.log2(size + 1) / 7);
    return header + size;
  }

  toString(): string {
    return this.str;
  }

  toJSON(): string {
    return this.str;
  }
}

export default Str;
