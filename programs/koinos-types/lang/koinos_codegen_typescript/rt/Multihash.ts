import { VariableBlob, VariableBlobLike } from "./VariableBlob";
import { Uint64 } from "./Uint64";
import { NumberLike } from "./Num";
import { VarInt } from "./VarInt";

export interface JsonMultihash {
  id: NumberLike;
  digest: VariableBlobLike;
}

export class Multihash {
  public id: Uint64;

  public digest: VariableBlob;

  constructor(
    json: JsonMultihash = {
      id: 0,
      digest: "",
    }
  ) {
    this.id = new Uint64(json.id);
    this.digest = new VariableBlob(json.digest);
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

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob();
    vb.serialize(new VarInt(this.id));
    vb.serialize(this.digest);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): Multihash {
    const id = vb.deserialize(VarInt);
    const digest = vb.deserialize(VariableBlob);
    return new Multihash({ id, digest });
  }
}

export default Multihash;
