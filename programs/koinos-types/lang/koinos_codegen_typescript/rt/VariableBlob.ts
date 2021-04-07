import * as ByteBuffer from "bytebuffer";

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

  serialize(vb: VariableBlob): VariableBlob {
    vb.buffer.writeVarint64(this.buffer.limit).append(this.buffer);
    return vb;
  }

  static deserialize(vb: VariableBlob): VariableBlob {
    const size = vb.buffer.readVarint64().toNumber();
    if (size < 0) throw new Error("Could not deserialize variable blob");

    const { limit, offset } = vb.buffer;
    if (limit < offset + size) throw new Error("Unexpected EOF");
    const subvb = new VariableBlob(size);
    vb.buffer.copyTo(subvb.buffer, 0, offset, offset + size);
    vb.buffer.offset += size;
    return subvb;
  }

  toHex(): string {
    if (this.buffer.offset !== 0) this.buffer.flip();
    return this.buffer.toHex();
  }

  flip(): void {
    this.buffer.flip();
  }
}

export default VariableBlob;
