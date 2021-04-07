import * as ByteBuffer from "bytebuffer";
import { KoinosBoolean } from "./KoinosBoolean";
import { KoinosString } from "./KoinosString";
import { Int8 } from "./Int8";
import { Int16 } from "./Int16";
import { Int32 } from "./Int32";
import { Int64 } from "./Int64";
import { Int128 } from "./Int128";
import { Int160 } from "./Int160";
import { Int256 } from "./Int256";
import { UInt8 } from "./UInt8";
import { UInt16 } from "./UInt16";
import { UInt32 } from "./UInt32";
import { UInt64 } from "./UInt64";
import { UInt128 } from "./UInt128";
import { UInt160 } from "./UInt160";
import { UInt256 } from "./UInt256";
import { TimestampType } from "./TimestampType";
import { BlockHeightType } from "./BlockHeightType";
import { Multihash } from "./Multihash";

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

  deserialize<T>(TCtor: new () => T): T {
    const t = new TCtor();
    if (t instanceof VariableBlob)
      return (this.deserializeVariableBlob() as unknown) as T;
    if (t instanceof KoinosBoolean)
      return (this.deserializeBoolean() as unknown) as T;
    if (t instanceof KoinosString)
      return (this.deserializeString() as unknown) as T;
    if (t instanceof Int8) return (this.deserializeInt8() as unknown) as T;
    if (t instanceof UInt8) return (this.deserializeUInt8() as unknown) as T;
    if (t instanceof Int16) return (this.deserializeInt16() as unknown) as T;
    if (t instanceof UInt16) return (this.deserializeUInt16() as unknown) as T;
    if (t instanceof Int32) return (this.deserializeInt32() as unknown) as T;
    if (t instanceof UInt32) return (this.deserializeUInt32() as unknown) as T;
    if (t instanceof Int64) return (this.deserializeInt64() as unknown) as T;
    if (t instanceof UInt64) return (this.deserializeUInt64() as unknown) as T;
    if (t instanceof Int128) return (this.deserializeInt128() as unknown) as T;
    if (t instanceof UInt128)
      return (this.deserializeUInt128() as unknown) as T;
    if (t instanceof Int160) return (this.deserializeInt160() as unknown) as T;
    if (t instanceof UInt160)
      return (this.deserializeUInt160() as unknown) as T;
    if (t instanceof Int256) return (this.deserializeInt256() as unknown) as T;
    if (t instanceof UInt256)
      return (this.deserializeUInt256() as unknown) as T;
    if (t instanceof TimestampType)
      return (this.deserializeTimestampType() as unknown) as T;
    if (t instanceof BlockHeightType)
      return (this.deserializeBlockHeightType() as unknown) as T;
    if (t instanceof Multihash)
      return (this.deserializeMultihash() as unknown) as T;
    throw new Error("Invalid class");
  }

  deserializeVariableBlob(): VariableBlob {
    const size = this.buffer.readVarint64().toNumber();
    if (size < 0) throw new Error("Could not deserialize variable blob");

    const { limit, offset } = this.buffer;
    if (limit < offset + size) throw new Error("Unexpected EOF");
    const vb = new VariableBlob(size);
    this.buffer.copyTo(vb.buffer, 0, offset, offset + size);
    this.buffer.offset += size;
    return vb;
  }

  deserializeString(): KoinosString {
    const vb = this.deserializeVariableBlob();
    return new KoinosString(vb.buffer.toUTF8());
  }

  deserializeBoolean(): KoinosBoolean {
    if (this.buffer.limit === 0) throw new Error("Unexpected EOF");
    const value = this.buffer.readByte();
    if (value !== 0 && value !== 1) throw new Error("Boolean must be 0 or 1");
    return new KoinosBoolean(!!value);
  }

  deserializeInt8(): Int8 {
    if (this.buffer.limit === 0) throw new Error("Unexpected EOF");
    const value = this.buffer.readByte();
    return new Int8(value);
  }

  deserializeUInt8(): UInt8 {
    return this.deserializeInt8() as UInt8;
  }

  deserializeInt16(): Int16 {
    if (this.buffer.limit < 2) throw new Error("Unexpected EOF");
    const value = this.buffer.readInt16();
    return new Int16(value);
  }

  deserializeUInt16(): UInt16 {
    if (this.buffer.limit < 2) throw new Error("Unexpected EOF");
    const value = this.buffer.readUint16();
    return new UInt16(value);
  }

  deserializeInt32(): Int32 {
    if (this.buffer.limit < 4) throw new Error("Unexpected EOF");
    const value = this.buffer.readInt32();
    return new Int32(value);
  }

  deserializeUInt32(): UInt32 {
    if (this.buffer.limit < 4) throw new Error("Unexpected EOF");
    const value = this.buffer.readUint32();
    return new UInt32(value);
  }

  deserializeBigInt(bits: number, unsigned: boolean): bigint {
    const bytes = bits / 8;
    if (this.buffer.limit < bytes) throw new Error("Unexpected EOF");
    let numString = "0x";
    for (let i = 0; i < bytes / 4; i += 1) {
      const uint32Str = this.buffer.readUint32().toString(16);
      numString += "0".repeat(8 - uint32Str.length) + uint32Str;
    }
    if (!unsigned && Number(numString.substring(0, 3)) >= 8) {
      // signed number and starts with bit 1 (negative number)
      return BigInt(numString) - BigInt("0x1" + "0".repeat(2 * bytes));
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

  deserializeTimestampType(): TimestampType {
    return new TimestampType(this.deserializeBigInt(64, true));
  }

  deserializeBlockHeightType(): BlockHeightType {
    return new BlockHeightType(this.deserializeBigInt(64, true));
  }

  deserializeMultihash(): Multihash {
    const id = this.buffer.readVarint64().toNumber();
    const digest = this.deserializeVariableBlob();
    const multihash = new Multihash();
    multihash.id = new UInt64(id);
    multihash.digest = digest;
    return multihash;
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
