import * as ByteBuffer from "bytebuffer";
import { VariableBlob } from "./VariableBlob";

export class KoinosString {
  public str: string;

  constructor(str = "") {
    this.str = str;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob();
    const buffer = ByteBuffer.fromUTF8(this.str) as ByteBuffer;
    vb.buffer.writeVarint64(buffer.limit).append(buffer);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): KoinosString {
    const subvb = vb.deserialize(VariableBlob);
    return new KoinosString(subvb.buffer.toUTF8());
  }

  toString(): string {
    return this.str;
  }
}

export default KoinosString;
