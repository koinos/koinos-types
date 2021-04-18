import * as bs58 from "bs58";
import { VariableBlob, VariableBlobLike } from "./VariableBlob";

export class FixedBlob {
  public buffer: Uint8Array;

  public size: number;

  constructor(b?: VariableBlobLike, size = 0) {
    this.size = size;
    if (b instanceof VariableBlob || b instanceof FixedBlob) {
      if (b.length() !== size)
        throw new Error(`Invalid blob size: ${b.length()}. Expected ${size}`);
      this.buffer = b.buffer;
    } else if (typeof b === "string") {
      if (b[0] !== "z") throw new Error(`Unknown encoding: ${b[0]}`);
      const buffer = bs58.decode(b.slice(1));
      if (buffer.length !== size)
        throw new Error(
          `Invalid blob size: ${buffer.length}. Expected ${size}`
        );
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
      if (this.buffer[i] !== vb.buffer[i]) return false;

    return true;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.buffer.set(this.buffer, vb.offset);
    vb.offset += this.length();
    if (!blob) vb.offset = 0;
    return vb;
  }

  static deserialize(vb: VariableBlob, size: number): FixedBlob {
    if (vb.length() < vb.offset + size) throw new Error("Unexpected EOF");
    const subfb = new FixedBlob(null, size);
    subfb.buffer.set(vb.buffer.subarray(vb.offset, vb.offset + size));
    vb.offset += size;
    return subfb;
  }

  calcSerializedSize(): number {
    return this.size;
  }

  toJSON(): string {
    return "z" + bs58.encode(this.buffer);
  }
}

export default FixedBlob;
