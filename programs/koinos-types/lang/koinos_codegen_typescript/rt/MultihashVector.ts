import { VariableBlob, VariableBlobLike } from "./VariableBlob";
import { VarInt } from "./VarInt";
import { NumberLike } from "./Num";

export interface JsonMultihashVector {
  id: NumberLike;
  digests: VariableBlobLike[];
}

export class MultihashVector {
  public id: VarInt;

  // TODO: consider using FixedBlob since all digests have the same size
  // and there is no header on each digest serialization
  public digests: VariableBlob[];

  constructor(
    json: JsonMultihashVector = {
      id: undefined,
      digests: [],
    }
  ) {
    this.id = new VarInt(json.id);
    this.digests = json.digests.map((digest) => new VariableBlob(digest));
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    // size of a single digest
    const sizeDigest = this.digests.length > 0 ? this.digests[0].length() : 0;

    vb.serialize(this.id);
    vb.serialize(new VarInt(sizeDigest));
    vb.serialize(new VarInt(this.digests.length));
    this.digests.forEach((digest) => {
      if (digest.length() !== sizeDigest)
        throw new Error("Multihash vector size mismatch");
      vb.buffer.append(digest.buffer);
    });
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): MultihashVector {
    const id = vb.deserialize(VarInt);
    const sizeDigest = vb.deserialize(VarInt).num;
    const len = vb.deserialize(VarInt).num;
    if (sizeDigest < 0)
      throw new Error("Could not deserialize multihash vector");
    const digests = new Array(len).fill(null).map(() => {
      const { limit, offset } = vb.buffer;
      if (limit < offset + sizeDigest) throw new Error("Unexpected EOF");
      const digest = new VariableBlob(sizeDigest);
      vb.buffer.copyTo(digest.buffer, 0, offset, offset + sizeDigest);
      vb.buffer.offset += sizeDigest;
      return digest;
    });
    return new MultihashVector({ id, digests });
  }

  calcSerializedSize(): number {
    const sizeDigest = this.digests.length > 0 ? this.digests[0].length() : 0;
    return (
      this.id.calcSerializedSize() +
      new VarInt(sizeDigest).calcSerializedSize() +
      new VarInt(this.digests.length).calcSerializedSize() +
      this.digests.length * sizeDigest
    );
  }

  toJSON(): JsonMultihashVector {
    return {
      id: this.id.toJSON(),
      digests: this.digests.map((digest) => digest.toJSON()),
    };
  }
}

export default MultihashVector;
