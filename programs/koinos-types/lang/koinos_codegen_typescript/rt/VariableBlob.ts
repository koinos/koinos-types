import * as ByteBuffer from "bytebuffer";
import KoinosBoolean from "./KoinosBoolean";
import KoinosString from "./KoinosString";

export interface KoinosClass {
  serialize(vb: VariableBlob): VariableBlob;
}

export interface KoinosClassBuilder<T extends KoinosClass> {
  new (): T;
  deserialize(vb: VariableBlob): T;
}

export class VariableBlob {
  public buffer: ByteBuffer;

  constructor(b: ByteBuffer | number = 0) {
    if (b instanceof ByteBuffer) this.buffer = b;
    else this.buffer = ByteBuffer.allocate(b) as ByteBuffer;
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

  serializeVariableBlob(): VariableBlob {
    const vb = new VariableBlob();
    vb.buffer.writeVarint64(this.buffer.limit).append(this.buffer);
    return vb;
  }

  serialize(k?: KoinosClass): VariableBlob {
    if (!k || k instanceof VariableBlob) return this.serializeVariableBlob()
    return k.serialize(this);
  }

  deserializeVariableBlob(): VariableBlob {
    const size = this.buffer.readVarint64().toNumber();
    if (size < 0) throw new Error("Could not deserialize variable blob");

    const { limit, offset } = this.buffer;
    if (limit < offset + size) throw new Error("Unexpected EOF");
    const subvb = new VariableBlob(size);
    this.buffer.copyTo(subvb.buffer, 0, offset, offset + size);
    this.buffer.offset += size;
    return subvb;
  }

  deserialize<T extends KoinosClass>(ClassT: KoinosClassBuilder<T>):T {
    return ClassT.deserialize(this);
  }

  toHex(): string {
    if (this.buffer.offset !== 0) this.flip();
    return this.buffer.toHex();
  }

  flip(): void {
    this.buffer.flip();
  }
}

export default VariableBlob;
