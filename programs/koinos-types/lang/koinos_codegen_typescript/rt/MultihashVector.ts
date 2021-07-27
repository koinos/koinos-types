import { VariableBlob, VariableBlobLike } from "./VariableBlob";
import { VarInt } from "./VarInt";
import { NumberLike } from "./Num";
import { FixedBlob } from "./FixedBlob";
import { Vector } from "./Vector";

export interface JsonMultihashVector {
  id: NumberLike;
  digests: VariableBlobLike[];
}

export class MultihashVector {
  public id: VarInt;

  public digests: Vector<FixedBlob>;

  constructor(
    json: JsonMultihashVector = {
      id: undefined,
      digests: [],
    }
  ) {
    this.id = new VarInt(json.id);
    const sizeDigest =
      json.digests.length > 0 ? new VariableBlob(json.digests[0]).length() : 0;
    this.digests = new Vector(FixedBlob, json.digests, sizeDigest);
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    // size of a single digest
    const sizeDigest =
      this.digests.items.length > 0 ? this.digests.items[0].size : 0;

    vb.serialize(this.id);
    vb.serialize(new VarInt(sizeDigest));
    vb.serialize(this.digests);
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): MultihashVector {
    const multihashVector = new MultihashVector();
    multihashVector.id = vb.deserialize(VarInt);
    const sizeDigest = vb.deserialize(VarInt).toNumber();
    multihashVector.digests = vb.deserializeVector(FixedBlob, sizeDigest);
    return multihashVector;
  }

  calcSerializedSize(): number {
    const sizeDigest =
      this.digests.items.length > 0 ? this.digests.items[0].size : 0;
    return (
      this.id.calcSerializedSize() +
      new VarInt(sizeDigest).calcSerializedSize() +
      this.digests.calcSerializedSize()
    );
  }

  toJSON(): JsonMultihashVector {
    return {
      id: this.id.toJSON(),
      digests: this.digests.toJSON("z") as string[],
    };
  }
}

export default MultihashVector;
