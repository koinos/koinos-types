import * as bs58 from "bs58";
import { FixedBlob } from "./FixedBlob";
import { VarInt } from "./VarInt";
import { Vector } from "./Vector";

export type VariableBlobLike = string | VariableBlob | FixedBlob;

export interface KoinosClass {
  serialize(vb?: VariableBlob): VariableBlob;

  calcSerializedSize(): number;

  toJSON(): unknown;
}

export interface KoinosClassBuilder<T extends KoinosClass> {
  new (json?: any, blobSize?: number): T;
  deserialize(vb: VariableBlob, blobSize?: number): T;
}

export class VariableBlob {
  public buffer: Uint8Array;

  public offset: number;

  constructor(b: VariableBlobLike | number = 0) {
    if (b instanceof VariableBlob || b instanceof FixedBlob) {
      this.buffer = b.buffer;
    } else if (typeof b === "string") {
      if (b[0] !== "z") throw new Error(`Unknown encoding: ${b[0]}`);
      this.buffer = bs58.decode(b.slice(1));
    } else if (typeof b === "number") {
      this.buffer = new Uint8Array(b);
    } else {
      this.buffer = new Uint8Array(0);
    }
    this.offset = 0;
  }

  length(): number {
    return this.buffer.length;
  }

  equals(vb: VariableBlob): boolean {
    const size1 = this.length();
    const size2 = vb.length();

    if (size1 !== size2) return false;

    for (let i = 0; i < size1; i += 1)
      if (this.buffer[i] !== vb.buffer[i]) return false;

    return true;
  }

  serializeVariableBlob(blob?: VariableBlob): VariableBlob {
    let vbModified: VariableBlob;
    let vbInput: VariableBlob;
    if (blob) {
      vbModified = this;
      vbInput = blob;
    } else {
      vbModified = new VariableBlob(this.calcSerializedSize());
      vbInput = this;
    }
    vbModified.serialize(new VarInt(vbInput.length()));
    vbModified.buffer.set(vbInput.buffer, vbModified.offset);
    vbModified.offset += vbInput.length();
    if (!blob) vbModified.offset = 0;
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

    if (this.length() < this.offset + size) throw new Error("Unexpected EOF");
    const subvb = new VariableBlob(size);
    subvb.buffer.set(this.buffer.subarray(this.offset, this.offset + size));
    this.offset += size;
    return subvb;
  }

  deserialize<T extends KoinosClass>(
    ClassT: KoinosClassBuilder<T> | typeof VariableBlob,
    size?: number
  ): T {
    if (new ClassT() instanceof VariableBlob)
      return (this.deserializeVariableBlob() as unknown) as T;
    if (new ClassT() instanceof FixedBlob)
      return (ClassT as KoinosClassBuilder<T>).deserialize(this, size);
    return (ClassT as KoinosClassBuilder<T>).deserialize(this);
  }

  deserializeVector<T extends KoinosClass>(
    ClassT: KoinosClassBuilder<T>,
    blobSize?: number
  ): Vector<T> {
    return Vector.deserialize(ClassT, this, blobSize);
  }

  calcSerializedSize(): number {
    const header = Math.ceil(Math.log2(this.length() + 1) / 7);
    return header + this.length();
  }

  toJSON(): string {
    return "z" + bs58.encode(this.buffer);
  }
}

export default VariableBlob;
