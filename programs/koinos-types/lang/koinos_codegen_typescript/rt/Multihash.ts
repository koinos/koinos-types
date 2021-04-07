import { VariableBlob } from "./VariableBlob";
import { UInt64 } from "./UInt64";

export class Multihash {
  public id: UInt64;

  public digest: VariableBlob;

  constructor() {
    this.id = new UInt64(0);
    this.digest = new VariableBlob();
  }

  equals(m: Multihash): boolean {
    return this.id.num === m.id.num && this.digest.equals(m.digest);
  }

  lessThan(m: Multihash): boolean {
    if (this.id.num !== m.id.num) return this.id.num < m.id.num;
    return this.digest.length() < m.digest.length();
  }

  greaterThan(m: Multihash): boolean {
    return !this.equals(m) && !this.lessThan(m);
  }

  serialize(vb: VariableBlob): VariableBlob {
    // TODO: Use Long instead of Number to construct the varint
    vb.buffer.writeVarint64(Number(this.id.num));
    this.digest.serialize(vb);
    return vb;
  }

  static deserialize(vb: VariableBlob): Multihash {
    const id = vb.buffer.readVarint64().toNumber();
    const digest = VariableBlob.deserialize(vb);
    const multihash = new Multihash();
    multihash.id = new UInt64(id);
    multihash.digest = digest;
    return multihash;
  }
}

export default Multihash;
