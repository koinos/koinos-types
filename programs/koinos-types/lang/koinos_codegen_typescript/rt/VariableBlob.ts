import * as ByteBuffer from "bytebuffer";
import * as bs58 from "bs58";
import { FixedBlob } from "./FixedBlob";
import { VarInt } from "./VarInt";
import { Vector } from "./Vector";

export type VariableBlobLike = string | ByteBuffer | VariableBlob;

export interface KoinosClass {
  serialize(vb?: VariableBlob): VariableBlob;

  toJSON(): unknown;
}

export interface KoinosClassBuilder<T extends KoinosClass> {
  new (json?: any): T;
  deserialize(vb: VariableBlob): T;
}

export function remove0xPrefix(str: string): string {
  return str.startsWith("0x") ? str.slice(2) : str;
}

export class VariableBlob {
  public buffer: ByteBuffer;

  constructor(b: VariableBlobLike | number = 0) {
    if (b instanceof VariableBlob) {
      this.buffer = b.buffer;
    } else if (b instanceof ByteBuffer) {
      this.buffer = b;
    } else if (typeof b === "string") {
      if (b[0] !== "z") throw new Error(`Unknown encoding: ${b[0]}`);
      this.buffer = new ByteBuffer();
      this.buffer.buffer = bs58.decode(b.slice(1));
      this.buffer.limit = this.buffer.buffer.length;
    } else {
      this.buffer = ByteBuffer.allocate(b) as ByteBuffer;
    }
  }

  length(): number {
    return Math.max(this.buffer.offset, this.buffer.limit);
  }

  equals(vb: VariableBlob): boolean {
    const size1 = this.length();
    const size2 = vb.length();

    if (size1 !== size2) return false;

    for (let i = 0; i < size1; i += 1)
      if (this.buffer.buffer[i] !== vb.buffer.buffer[i]) return false;

    return true;
  }

  serializeVariableBlob(blob?: VariableBlob): VariableBlob {
    let vbModified: VariableBlob;
    let vbInput: VariableBlob;
    if (blob) {
      vbModified = this;
      vbInput = blob;
    } else {
      vbModified = new VariableBlob();
      vbInput = this;
    }
    vbModified.serialize(new VarInt(vbInput.buffer.limit));
    vbModified.buffer.append(vbInput.buffer);
    if (!blob) vbModified.flip();
    return vbModified;
  }

  serialize<T extends KoinosClass>(k?: T): VariableBlob {
    if (!k || k instanceof VariableBlob)
      return this.serializeVariableBlob((k as unknown) as VariableBlob);
    return k.serialize(this);
  }

  deserializeVariableBlob(): VariableBlob {
    const size = this.deserialize(VarInt).toNumber();
    if (size < 0) throw new Error("Could not deserialize variable blob");

    const { limit, offset } = this.buffer;
    if (limit < offset + size) throw new Error("Unexpected EOF");
    const subvb = new VariableBlob(size);
    this.buffer.copyTo(subvb.buffer, 0, offset, offset + size);
    this.buffer.offset += size;
    return subvb;
  }

  deserialize<T extends KoinosClass>(
    ClassT: KoinosClassBuilder<T> | typeof VariableBlob
  ): T {
    if (new ClassT() instanceof VariableBlob)
      return (this.deserializeVariableBlob() as unknown) as T;
    return (ClassT as KoinosClassBuilder<T>).deserialize(this);
  }

  deserializeVector<T extends KoinosClass>(
    ClassT: KoinosClassBuilder<T>
  ): Vector<T> {
    return Vector.deserialize(ClassT, this);
  }

  deserializeFixedBlob(size: number): FixedBlob {
    const { limit, offset } = this.buffer;
    if (limit < offset + size) throw new Error("Unexpected EOF");
    const subfb = new FixedBlob(size);
    this.buffer.copyTo(subfb.buffer, 0, offset, offset + size);
    this.buffer.offset += size;
    return subfb;
  }

  toJSON(): string {
    return "z" + bs58.encode(this.buffer.buffer);
  }

  toHex(): string {
    if (this.buffer.offset !== 0) this.flip();
    return this.buffer.toHex();
  }

  flip(): VariableBlob {
    this.buffer.flip();
    return this;
  }
}

export default VariableBlob;
