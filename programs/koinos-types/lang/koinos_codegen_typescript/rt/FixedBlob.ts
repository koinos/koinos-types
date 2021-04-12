import * as ByteBuffer from "bytebuffer";
import { VariableBlob, VariableBlobLike, remove0xPrefix } from "./VariableBlob";

export class FixedBlob {
  public buffer: ByteBuffer;

  public size: number;

  constructor(size: number, b?: VariableBlobLike) {
    this.size = size;
    this.buffer = ByteBuffer.allocate(size) as ByteBuffer;
    if (b instanceof VariableBlob) {
      this.buffer = b.buffer;
    } else if (b instanceof ByteBuffer) {
      this.buffer = b;
    } else if (typeof b === "string") {
      this.buffer = ByteBuffer.fromHex(remove0xPrefix(b)) as ByteBuffer;
    }
  }

  length(): number {
    return this.size;
  }

  equals(vb: FixedBlob | VariableBlob): boolean {
    const size1 = this.length();
    const size2 = vb.length();

    if (size1 !== size2) return false;

    for (let i = 0; i < size1; i += 1)
      if (this.buffer.buffer[i] !== vb.buffer.buffer[i]) return false;

    return true;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob();
    this.buffer.copyTo(vb.buffer, vb.buffer.offset, 0, this.size);
    vb.buffer.offset += this.size;
    if (!blob) vb.flip();
    return vb;
  }

  toHex(): string {
    if (this.buffer.offset !== 0) this.flip();
    return this.buffer.toHex();
  }

  flip(): FixedBlob {
    this.buffer.flip();
    return this;
  }
}

export default FixedBlob;
