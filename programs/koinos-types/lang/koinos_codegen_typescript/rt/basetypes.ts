import * as ByteBuffer from 'bytebuffer'
import * as Long from 'long'

export class KString {
  
  public _string: string;

  constructor(str: string = "") {
    this._string = str;
  }
  
  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.append(ByteBuffer.fromUTF8(this._string));
  }
}

export class KBoolean {
  
  public _bool: boolean;

  constructor(bool: boolean = false) {
    this._bool = bool;
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeByte(this._bool ? 1: 0);
  }
}

export const MAX_INT8 = 0x7F;
export const MIN_INT8 = -0x80;
export class Int8 {

  public _number: number;

  constructor(n: number = 0) {
    if(n < MIN_INT8 || n > MAX_INT8)
      throw new Error("Int8 is out of bounds");
    this._number = n;
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeByte(this._number);
  }
}

export const MAX_UINT8 = 0xFF;
export class UInt8 {

  public _number: number;

  constructor(n: number = 0) {
    if(n < 0 || n > MAX_UINT8)
      throw new Error("UInt8 is out of bounds");
    this._number = n;
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeByte(this._number);
  }
}

export const MAX_INT16 = 0x7FFF;
export const MIN_INT16 = -0x8000;
export class Int16 {

  public _number: number;

  constructor(n: number = 0) {
    if(n < MIN_INT16 || n > MAX_INT16)
      throw new Error("Int16 is out of bounds");
    this._number = n;
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeInt16(this._number);
  }
}

export const MAX_UINT16 = 0xFFFF;
export class UInt16 {
  
  public _number: number;

  constructor(n: number = 0) {
    if(n < 0 || n > MAX_UINT16)
      throw new Error("UInt16 is out of bounds");  
    this._number = n;
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeUint16(this._number);
  }
}

export const MAX_INT32 = 0x7FFFFFFF;
export const MIN_INT32 = -0x80000000;
export class Int32 {

  public _number: number;

  constructor(n: number = 0) {
    if(n < MIN_INT32 || n > MAX_INT32)
      throw new Error("Int32 is out of bounds");
    this._number = n;
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeInt32(this._number);
  }
}

export const MAX_UINT32 = 0xFFFFFFFF;
export class UInt32 {

  public _number: number;

  constructor(n: number = 0) {
    if(n < 0 || n > MAX_UINT32)
      throw new Error("UInt32 is out of bounds");
    this._number = n;
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeUint32(this._number);
  }
}

export const MAX_INT64 = Long.MAX_VALUE;
export const MIN_INT64 = Long.MIN_VALUE;
export class Int64 {

  public _number: Long.Long;

  constructor(n: number | string | Long.Long) {
    if (n instanceof Long)
      this._number = n;
    else if(typeof n === "string")
      this._number = Long.fromString(n);
    else {
      const number = n as unknown as number;
      if (number < MIN_INT32 || number > MAX_INT32)
        throw new Error("Int64 low number out of bounds");
      this._number = Long.fromNumber(number);
    }
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeInt64(this._number);
  }
}

export const MAX_UINT64 = Long.MAX_UNSIGNED_VALUE;
export class UInt64 {

  public _number: Long.Long;

  constructor(n: number | string | Long.Long) {
    if (n instanceof Long)
      this._number = n;
    else if(typeof n === "string")
      this._number = Long.fromString(n);
    else {
      const number = n as unknown as number;
      if (number < 0 || number > MAX_UINT32)
        throw new Error("UInt64 low number out of bounds");
      this._number = Long.fromNumber(number);
    }
  }

  serialize(vb: VariableBlob): VariableBlob {
    return vb.buffer.writeUint64(this._number);
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

  serialize(vb: VariableBlob): VariableBlob {
    vb.buffer
      .writeVarint64(this.buffer.limit)
      .append(this.buffer);
    return vb;
  }

  deserializeVariableBlob(): VariableBlob {
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
    return new KString(this.buffer.toUTF8())
  }

  deserializeBoolean(): KBoolean {
    if(this.buffer.limit === 0)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readByte(0);
    if( value !== 0 && value !== 1)
      throw new Error("Boolean must be 0 or 1");
    return new KBoolean(!!value);
  }

  deserializeInt8(): Int8 {
    if(this.buffer.limit === 0)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readByte(0);
    return new Int8(value);
  }

  deserializeUInt8(): UInt8 {
    return this.deserializeInt8() as UInt8;
  }

  deserializeInt16(): Int16 {
    if(this.buffer.limit < 2)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readInt16();
    return new Int16(value);
  }

  deserializeUInt16(): UInt16 {
    if(this.buffer.limit < 2)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readUint16();
    return new UInt16(value);
  }

  deserializeInt32(): Int32 {
    if(this.buffer.limit < 4)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readInt32();
    return new Int32(value);
  }

  deserializeUInt32(): UInt32 {
    if(this.buffer.limit < 4)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readUint32();
    return new UInt32(value);
  }

  deserializeInt64(): Int64 {
    if(this.buffer.limit < 8)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readInt64();
    return new Int64(value);
  }

  deserializeUInt64(): UInt64 {
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
