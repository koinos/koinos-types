import * as ByteBuffer from 'bytebuffer'

export class KString {
  
  public _string: string;

  constructor(str: string = "") {
    this._string = str;
  }
  
  serialize(vb: VariableBlob): VariableBlob {
    const buffer = ByteBuffer.fromUTF8(this._string);
    return vb.buffer
      .writeVarint64(buffer.limit)
      .append(buffer);
  }

  toString(): string {
    return this._string;
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

  toBoolean(): boolean {
    return this._bool;
  }

  toString(): string {
    return this._bool.toString();
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

  toNumber(): number {
    return this._number;
  }

  toString(radix?: number): string {
    return this._number.toString(radix);
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

  toNumber(): number {
    return this._number;
  }

  toString(radix?: number): string {
    return this._number.toString(radix);
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

  toNumber(): number {
    return this._number;
  }

  toString(radix?: number): string {
    return this._number.toString(radix);
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

  toNumber(): number {
    return this._number;
  }

  toString(radix?: number): string {
    return this._number.toString(radix);
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

  toNumber(): number {
    return this._number;
  }

  toString(radix?: number): string {
    return this._number.toString(radix);
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

  toNumber(): number {
    return this._number;
  }

  toString(radix?: number): string {
    return this._number.toString(radix);
  }
}

export class KBigInt {

  public _number: bigint;

  private bytes: number;

  constructor(number: bigint | string | number, bits: number, max: bigint, min: bigint = BigInt(0)) {
    const n = BigInt(number);
    const unsigned = min === BigInt(0);
    this.bytes = bits / 8;
    this._number = n;
    if(n < min || n > max )
      throw new Error(`${unsigned ? "U" : ""}Int${bits} is out of bounds`);
  }

  serialize(vb: VariableBlob): VariableBlob {
    let numString: string;
    if(this._number >= BigInt(0)) {
      numString = this._number.toString(16);
      numString = "0".repeat(2*this.bytes - numString.length) + numString;
    } else {
      numString = (BigInt("0x1" + "0".repeat(2*this.bytes)) + this._number).toString(16);
    }
    for(let i=0; i<2*this.bytes; i+=8) {
      const int32 = Number("0x" + numString.substring(i, i+8));
      vb.buffer.writeUint32(int32);
    }
    return vb;
  }

  toBigInt(): bigint {
    return this._number;
  }

  toString(radix?: number): string {
    return this._number.toString(radix);
  }
}

export const MAX_INT64  = BigInt("0x7" + "F".repeat(15));
export const MAX_INT128 = BigInt("0x7" + "F".repeat(31));
export const MAX_INT160 = BigInt("0x7" + "F".repeat(39));
export const MAX_INT256 = BigInt("0x7" + "F".repeat(63));

export const MIN_INT64  = -BigInt("0x8" + "0".repeat(15));
export const MIN_INT128 = -BigInt("0x8" + "0".repeat(31));
export const MIN_INT160 = -BigInt("0x8" + "0".repeat(39));
export const MIN_INT256 = -BigInt("0x8" + "0".repeat(63));

export const MAX_UINT64  = BigInt("0x" + "F".repeat(16));
export const MAX_UINT128 = BigInt("0x" + "F".repeat(32));
export const MAX_UINT160 = BigInt("0x" + "F".repeat(40));
export const MAX_UINT256 = BigInt("0x" + "F".repeat(64));

export class Int64 extends KBigInt {
  constructor(number: bigint | string | number) {
    super(number, 64, MAX_INT64, MIN_INT64);
  }
}

export class UInt64 extends KBigInt {
  constructor(number: bigint | string | number) {
    super(number, 64, MAX_UINT64);
  }
}

export class Int128 extends KBigInt {
  constructor(number: bigint | string | number) {
    super(number, 128, MAX_INT128, MIN_INT128);
  }
}

export class UInt128 extends KBigInt {
  constructor(number: bigint | string | number) {
    super(number, 128, MAX_UINT128);
  }
}

export class Int160 extends KBigInt {
  constructor(number: bigint | string | number) {
    super(number, 160, MAX_INT160, MIN_INT160);
  }
}

export class UInt160 extends KBigInt {
  constructor(number: bigint | string | number) {
    super(number, 160, MAX_UINT160);
  }
}

export class Int256 extends KBigInt {
  constructor(number: bigint | string | number) {
    super(number, 256, MAX_INT256, MIN_INT256);
  }
}

export class UInt256 extends KBigInt {
  constructor(number: bigint | string | number) {
    super(number, 256, MAX_UINT256);
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
    this.buffer.offset += size;
    return vb;
  }

  deserializeString(): KString {
    const vb = this.deserializeVariableBlob();
    return new KString(vb.buffer.toUTF8())
  }

  deserializeBoolean(): KBoolean {
    if(this.buffer.limit === 0)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readByte();
    if( value !== 0 && value !== 1)
      throw new Error("Boolean must be 0 or 1");
    return new KBoolean(!!value);
  }

  deserializeInt8(): Int8 {
    if(this.buffer.limit === 0)
      throw new Error("Unexpected EOF");
    const value = this.buffer.readByte();
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

  deserializeBigInt(bits: number, unsigned: boolean): bigint {
    const bytes = bits / 8;
    if(this.buffer.limit < bytes)
      throw new Error("Unexpected EOF");
    let numString = "0x";
    for(let i=0; i<bytes/4; i+=1) {
      const uint32Str = this.buffer.readUint32().toString(16);
      numString += "0".repeat(8 - uint32Str.length) + uint32Str;
    }
    if(!unsigned && Number(numString.substring(0,3)) >= 8) {
      // signed number and starts with bit 1 (negative number)
      return BigInt(numString) - BigInt("0x1" + "0".repeat(2*bytes));
    }
    return BigInt(numString);
  }

  deserializeInt64(): Int64 {
    return new Int64(this.deserializeBigInt(64, false));
  }
  
  deserializeUInt64(): UInt64 {
    return new UInt64(this.deserializeBigInt(64, true));
  }

  deserializeInt128(): Int128 {
    return new Int128(this.deserializeBigInt(128, false));
  }
  
  deserializeUInt128(): UInt128 {
    return new UInt128(this.deserializeBigInt(128, true));
  }

  deserializeInt160(): Int160 {
    return new Int160(this.deserializeBigInt(160, false));
  }
  
  deserializeUInt160(): UInt160 {
    return new UInt160(this.deserializeBigInt(160, true));
  }

  deserializeInt256(): Int256 {
    return new Int256(this.deserializeBigInt(256, false));
  }
  
  deserializeUInt256(): UInt256 {
    return new UInt256(this.deserializeBigInt(256, true));
  }  

  toHex() {
    if(this.buffer.offset !== 0) this.buffer.flip();
    return this.buffer.toHex();
  }
}
