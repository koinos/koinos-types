import * as ByteBuffer from "bytebuffer";
import { VariableBlob } from "./VariableBlob";

export class KoinosString {
  public str: string;

  constructor(str = "") {
    this.str = str;
  }

  serialize(vb: VariableBlob): VariableBlob {
    const buffer = ByteBuffer.fromUTF8(this.str) as ByteBuffer;
    vb.buffer.writeVarint64(buffer.limit).append(buffer);
    return vb;
  }

  toString(): string {
    return this.str;
  }
}

export default KoinosString;
