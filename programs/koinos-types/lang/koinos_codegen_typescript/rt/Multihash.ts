import { VariableBlob, VariableBlobLike } from "./VariableBlob";
import { VarInt } from "./VarInt";
import { NumberLike } from "./Num";

export type MultihashLike = Multihash | {
  id: NumberLike;
  digest: VariableBlobLike;
}

export class Multihash {
  public id: VarInt;

  public digest: VariableBlob;

  constructor(
    value: MultihashLike = {
      id: undefined,
      digest: undefined,
    }
  ) {
    if (value instanceof Multihash) {
      this.id = value.id;
      this.digest = value.digest;
    } else {
      this.id = new VarInt(value.id);
      this.digest = new VariableBlob(value.digest);
    }
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
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.serialize(this.id);
    vb.serialize(this.digest);
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): Multihash {
    const id = vb.deserialize(VarInt);
    const digest = vb.deserialize(VariableBlob);
    return new Multihash({ id, digest });
  }

  calcSerializedSize(): number {
    return this.id.calcSerializedSize() + this.digest.calcSerializedSize();
  }

  toJSON(): MultihashLike {
    return {
      id: this.id.toJSON(),
      digest: this.digest.toJSON(),
    };
  }
}

export default Multihash;
