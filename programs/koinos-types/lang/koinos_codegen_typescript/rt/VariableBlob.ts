import * as bs58 from "bs58";
import { FixedBlob } from "./FixedBlob";
import { Opaque } from "./Opaque";
import { Variant } from "./Variant";
import { VarInt } from "./VarInt";
import { Vector } from "./Vector";

export type VariableBlobLike = string | Uint8Array | VariableBlob | FixedBlob;

export interface KoinosClass {
  serialize(vb?: VariableBlob): VariableBlob;

  calcSerializedSize(): number;

  toJSON(): unknown;
}

export interface KoinosClassBuilder<T extends KoinosClass> {
  new (json?: unknown, blobSize?: number): T;
  deserialize(vb: VariableBlob, blobSize?: number): T;
}

export class VariableBlob {
  public buffer: Uint8Array;

  public offset: number;

  constructor(b: VariableBlobLike | number = 0) {
    if (b instanceof VariableBlob || b instanceof FixedBlob) {
      this.buffer = b.buffer;
    } else if (b instanceof Uint8Array) {
      this.buffer = b;
    } else if (typeof b === "string") {
      if (b[0] !== "z") throw new Error(`Unknown encoding: ${b[0]}`);
      this.buffer = new Uint8Array(bs58.decode(b.slice(1)));
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

  equals(vb: FixedBlob | VariableBlob): boolean {
    if (this.length() !== vb.length()) return false;

    for (let i = 0; i < this.length(); i += 1)
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
    vbModified.write(vbInput.buffer);
    if (!blob) vbModified.resetCursor();
    return vbModified;
  }

  serialize<T extends KoinosClass>(k?: T): VariableBlob {
    if (!k || k instanceof VariableBlob)
      return this.serializeVariableBlob((k as unknown) as VariableBlob);
    return k.serialize(this);
  }

  deserializeVariableBlob(): VariableBlob {
    const size = this.deserialize(VarInt).toNumber();
    const subBuffer = this.read(size);
    return new VariableBlob(subBuffer);
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

  deserializeOpaque<T extends KoinosClass>(
    ClassT: KoinosClassBuilder<T>
  ): Opaque<T> {
    return Opaque.deserialize(ClassT, this);
  }

  deserializeVariant<
    A extends KoinosClass,
    B extends KoinosClass,
    C extends KoinosClass,
    D extends KoinosClass,
    E extends KoinosClass,
    F extends KoinosClass,
    G extends KoinosClass,
    H extends KoinosClass,
    I extends KoinosClass,
    J extends KoinosClass
  >(
    variant: Variant<A, B, C, D, E, F, G, H, I, J>
  ): Variant<A, B, C, D, E, F, G, H, I, J> {
    return variant.deserializeVariant(this);
  }

  calcSerializedSize(): number {
    const header = Math.ceil(Math.log2(this.length() + 1) / 7);
    return header + this.length();
  }

  toJSON(): string {
    return "z" + bs58.encode(this.buffer);
  }

  resetCursor(): VariableBlob {
    this.offset = 0;
    return this;
  }

  checkRemaining(size: number, canResize = false): void {
    if (this.offset + size > this.buffer.length) {
      if (!canResize) throw new Error("Unexpected EOF");
      const newSize = this.offset + size;
      const buffer = new Uint8Array(newSize);
      buffer.set(this.buffer);
      this.buffer = buffer;
    }
  }

  write(bytes: Uint8Array): void {
    this.checkRemaining(bytes.length, true);
    this.buffer.set(bytes, this.offset);
    this.offset += bytes.length;
  }

  writeInt8(n: number): void {
    this.checkRemaining(1, true);
    new DataView(this.buffer.buffer).setInt8(this.offset, n);
    this.offset += 1;
  }

  writeInt16(n: number): void {
    this.checkRemaining(2, true);
    new DataView(this.buffer.buffer).setInt16(this.offset, n);
    this.offset += 2;
  }

  writeInt32(n: number): void {
    this.checkRemaining(4, true);
    new DataView(this.buffer.buffer).setInt32(this.offset, n);
    this.offset += 4;
  }

  writeUint8(n: number): void {
    this.checkRemaining(1, true);
    // new DataView(this.buffer.buffer).setUint8(this.offset, n);
    this.buffer[this.offset] = n;
    this.offset += 1;
  }

  writeUint16(n: number): void {
    this.checkRemaining(2, true);
    new DataView(this.buffer.buffer).setUint16(this.offset, n);
    this.offset += 2;
  }

  writeUint32(n: number): void {
    this.checkRemaining(4, true);
    new DataView(this.buffer.buffer).setUint32(this.offset, n);
    this.offset += 4;
  }

  read(size: number): Uint8Array {
    this.checkRemaining(size);
    const subBuffer = new Uint8Array(size);
    subBuffer.set(this.buffer.subarray(this.offset, this.offset + size));
    this.offset += size;
    return subBuffer;
  }

  readInt8(): number {
    this.checkRemaining(1);
    const n = new DataView(this.buffer.buffer).getInt8(this.offset);
    this.offset += 1;
    return n;
  }

  readInt16(): number {
    this.checkRemaining(2);
    const n = new DataView(this.buffer.buffer).getInt16(this.offset);
    this.offset += 2;
    return n;
  }

  readInt32(): number {
    this.checkRemaining(4);
    const n = new DataView(this.buffer.buffer).getInt32(this.offset);
    this.offset += 4;
    return n;
  }

  readUint8(): number {
    this.checkRemaining(1);
    // const n = new DataView(this.buffer.buffer).getUint8(this.offset);
    const n = this.buffer[this.offset];
    this.offset += 1;
    return n;
  }

  readUint16(): number {
    this.checkRemaining(2);
    const n = new DataView(this.buffer.buffer).getUint16(this.offset);
    this.offset += 2;
    return n;
  }

  readUint32(): number {
    this.checkRemaining(4);
    const n = new DataView(this.buffer.buffer).getUint32(this.offset);
    this.offset += 4;
    return n;
  }
}

export default VariableBlob;
