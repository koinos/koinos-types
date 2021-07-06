import multibase from "multibase";
import { VariableBlob, VariableBlobLike } from "./VariableBlob";

export class FixedBlob {
  public buffer: Uint8Array;

  public size: number;

  constructor(b: VariableBlobLike | number = 0, size = 0) {
    this.size = size;
    let buffer: Uint8Array;
    if (b instanceof VariableBlob || b instanceof FixedBlob) {
      buffer = b.buffer;
    } else if (b instanceof Uint8Array) {
      buffer = b;
    } else if (typeof b === "string") {
      buffer = multibase.decode(b);
    } else if (typeof b === "number") {
      buffer = new Uint8Array(b);
    } else {
      buffer = new Uint8Array(size);
    }

    if (buffer.length !== size)
      throw new Error(`Invalid blob size: ${buffer.length}. Expected ${size}`);
    this.buffer = buffer;
  }

  length(): number {
    return this.size;
  }

  equals(vb: FixedBlob | VariableBlob): boolean {
    if (this.length() !== vb.length()) return false;

    for (let i = 0; i < this.length(); i += 1)
      if (this.buffer[i] !== vb.buffer[i]) return false;

    return true;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.write(this.buffer);
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob, size: number): FixedBlob {
    const subBuffer = vb.read(size);
    return new FixedBlob(subBuffer, size);
  }

  calcSerializedSize(): number {
    return this.size;
  }

  toJSON(nameOrCode: multibase.BaseNameOrCode = "M"): string {
    return new TextDecoder().decode(multibase.encode(nameOrCode, this.buffer));
  }
}

export default FixedBlob;
