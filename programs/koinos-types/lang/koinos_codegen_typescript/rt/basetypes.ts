import * as ByteBuffer from 'bytebuffer'
import * as Long from 'long'

export class KString {
  
  public _string: string;

  constructor(str: string = "") {
    this._string = str;
  }
  
  serialize(): VariableBlob {
    return new VariableBlob(ByteBuffer.fromUTF8(this._string));
  }
}

export class KBoolean {
  
  public _bool: boolean;

  constructor(bool: boolean = false) {
    this._bool = bool;
  }

  serialize(): VariableBlob {
    const buffer = new ByteBuffer(1);
    buffer.writeByte(this._bool ? 1: 0)
    return new VariableBlob(buffer);
  }
}

export class Int8 {

  public _number: number;

  constructor(n: number = 0) {
    this._number = n;
  }

  serialize(): VariableBlob {
    return new VariableBlob(new ByteBuffer(1).writeByte(this._number));
  }
}

export class UInt8 extends Int8 {}

export class Int16 {

  public _number: number;

  constructor(n: number = 0) {
    this._number = n;
  }

  serialize(): VariableBlob {
    return new VariableBlob(new ByteBuffer(2).writeInt16(this._number));
  }
}

export class UInt16 {

  public _number: number;

  constructor(n: number = 0) {
    this._number = n;
  }

  serialize(): VariableBlob {
    return new VariableBlob(new ByteBuffer(2).writeUint16(this._number));
  }
}

export class Int32 {

  public _number: number;

  constructor(n: number = 0) {
    this._number = n;
  }

  serialize(): VariableBlob {
    return new VariableBlob(new ByteBuffer(4).writeInt32(this._number));
  }
}

export class UInt32 {

  public _number: number;

  constructor(n: number = 0) {
    this._number = n;
  }

  serialize(): VariableBlob {
    return new VariableBlob(new ByteBuffer(4).writeUint32(this._number));
  }
}

export class Int64 {

  public _number: Long.Long;

  constructor(n?: number | Long.Long) {
    if(n instanceof Long)
      this._number = n;
    else
      this._number = new Long(n);
  }

  serialize(): VariableBlob {
    return new VariableBlob(new ByteBuffer(8).writeInt64(this._number));
  }
}

export class UInt64 {

  public _number: Long.Long;

  constructor(n?: number | Long.Long) {
    if(n instanceof Long)
      this._number = n;
    else
      this._number = new Long(n);
  }

  serialize(): VariableBlob {
    return new VariableBlob(new ByteBuffer(8).writeUint64(this._number));
  }
}

export class VariableBlob {

  public buffer: ByteBuffer;

  constructor(b: ByteBuffer | number = 0) {
    if(b instanceof ByteBuffer)
      this.buffer = b;
    else
      this.buffer = ByteBuffer.allocate(b);
  }

  serialize(): VariableBlob {
    if(this.buffer.offset !== 0) this.buffer.flip();
    const ser = new VariableBlob();
    ser.buffer
      .writeVarint64(this.buffer.limit)
      .append(this.buffer);
    return ser;
  }

  deserializeVariableBlob(): VariableBlob {
    if(this.buffer.offset !== 0) this.buffer.flip();
    const size = this.buffer.readVarint64().toNumber();
    if(size < 0)
      throw new Error("Could not deserialize variable blob");

    const { limit, offset } = this.buffer;
    if(limit < offset + size)
      throw new Error("Unexpected EOF");
    const vb = new VariableBlob(size);
    this.buffer.copyTo(vb.buffer, 0, offset, offset + size);
    return vb;
  }

  deserializeString(): KString {
    if(this.buffer.offset !== 0) this.buffer.flip();
    return new KString(this.buffer.toUTF8())
  }

  deserializeBoolean(): KBoolean {
    if(this.buffer.offset !== 0) this.buffer.flip();
    if(this.buffer.limit === 0)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readByte(0);
    if( value !== 0 && value !== 1)
      throw new Error("Boolean must be 0 or 1");
    return new KBoolean(!!value);
  }

  deserializeInt8(): Int8 {
    if(this.buffer.offset !== 0) this.buffer.flip();
    if(this.buffer.limit === 0)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readByte(0);
    return new Int8(value);
  }

  deserializeUInt8(): UInt8 {
    return this.deserializeInt8() as UInt8;
  }

  deserializeInt16(): Int16 {
    if(this.buffer.offset !== 0) this.buffer.flip();
    if(this.buffer.limit < 2)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readInt16();
    return new Int16(value);
  }

  deserializeUInt16(): UInt16 {
    if(this.buffer.offset !== 0) this.buffer.flip();
    if(this.buffer.limit < 2)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readUint16();
    return new UInt16(value);
  }

  deserializeInt32(): Int32 {
    if(this.buffer.offset !== 0) this.buffer.flip();
    if(this.buffer.limit < 4)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readInt32();
    return new Int32(value);
  }

  deserializeUInt32(): UInt32 {
    if(this.buffer.offset !== 0) this.buffer.flip();
    if(this.buffer.limit < 4)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readUint32();
    return new UInt32(value);
  }

  deserializeInt64(): Int64 {
    if(this.buffer.offset !== 0) this.buffer.flip();
    if(this.buffer.limit < 8)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readInt64();
    return new Int64(value);
  }

  deserializeUInt64(): UInt64 {
    if(this.buffer.offset !== 0) this.buffer.flip();
    if(this.buffer.limit < 4)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readUint64();
    return new UInt64(value);
  }

  toHex() {
    if(this.buffer.offset !== 0) this.buffer.flip();
    return this.buffer.toHex();
  }
}
