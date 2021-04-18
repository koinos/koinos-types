import * as ByteBuffer from "bytebuffer";
import * as bs58 from "bs58";
import { VariableBlob, VariableBlobLike } from "./VariableBlob";

export class FixedBlob {
  public buffer: ByteBuffer;

  public size: number;

  constructor(b?: VariableBlobLike, size = 0) {
    this.size = size;
    this.buffer = ByteBuffer.allocate(size) as ByteBuffer;
    if (b instanceof VariableBlob) {
      if (b.length() !== size)
        throw new Error(`Invalid blob size: ${b.length()}. Expected ${size}`);
      this.buffer = b.buffer;
    } else if (b instanceof ByteBuffer) {
      if (b.limit !== size)
        throw new Error(`Invalid blob size: ${b.limit}. Expected ${size}`);
      this.buffer = b;
    } else if (typeof b === "string") {
      if (b[0] !== "z") throw new Error(`Unknown encoding: ${b[0]}`);
      const buffer = new ByteBuffer();
      buffer.buffer = bs58.decode(b.slice(1));
      if (buffer.buffer.length !== size)
        throw new Error(
          `Invalid blob size: ${buffer.buffer.length}. Expected ${size}`
        );
      buffer.limit = size;
      this.buffer = buffer;
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
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    this.buffer.copyTo(vb.buffer, vb.buffer.offset, 0, this.size);
    vb.buffer.offset += this.size;
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob, size: number): FixedBlob {
    const { limit, offset } = vb.buffer;
    if (limit < offset + size) throw new Error("Unexpected EOF");
    const subfb = new FixedBlob(null, size);
    vb.buffer.copyTo(subfb.buffer, 0, offset, offset + size);
    vb.buffer.offset += size;
    return subfb;
  }

  calcSerializedSize(): number {
    return this.size;
  }

  toJSON(): string {
    return "z" + bs58.encode(this.buffer.buffer);
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
