import multibase from "multibase";
import { VariableBlob, VariableBlobLike } from "./VariableBlob";
import { VarInt } from "./VarInt";
import { NumberLike } from "./Num";

export type MultihashLike =
  | Multihash
  | VariableBlobLike
  | {
      id: NumberLike;
      digest: VariableBlobLike;
    };

export class Multihash {
  public id: VarInt;

  public digest: VariableBlob;

  constructor(
    value: MultihashLike = {
      id: 0,
      digest: "z",
    }
  ) {
    const valueObj = value as {
      id: NumberLike;
      digest: VariableBlobLike;
    };

    if (value instanceof Multihash) {
      this.id = value.id;
      this.digest = value.digest;
    } else if (
      typeof valueObj.id !== "undefined" &&
      typeof valueObj.digest !== "undefined"
    ) {
      this.id = new VarInt(valueObj.id);
      this.digest = new VariableBlob(valueObj.digest);
    } else {
      // Although VariableBlob and Multihash use different
      // multibase encoding we can use the constructor of
      // VariableBlob because it uses a general function to decode
      // into bytes
      const multihash = Multihash.deserialize(
        new VariableBlob(value as VariableBlobLike)
      );
      this.id = multihash.id;
      this.digest = multihash.digest;
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

  toJSON(): string {
    return this.serialize().toJSON("z");
  }
}

export default Multihash;
