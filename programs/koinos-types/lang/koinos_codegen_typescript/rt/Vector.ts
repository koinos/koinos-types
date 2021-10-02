import { VariableBlob, KoinosClass, KoinosClassBuilder } from "./VariableBlob";
import { VarInt, sizeVarInt } from "./VarInt";
import { FixedBlob } from "./FixedBlob";
import multibase from "multibase";

export class Vector<T extends KoinosClass> {
  public items: T[];

  constructor(
    ClassT: KoinosClassBuilder<T>,
    items: (T | unknown)[] = [],
    blobSize?: number
  ) {
    this.items = items.map((item) => {
      if (item instanceof ClassT) return item;
      if (new FixedBlob() instanceof ClassT) return new ClassT(item, blobSize);
      return new ClassT(item);
    });
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.serialize(new VarInt(this.items.length));
    this.items.forEach((item) => vb.serialize(item));
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize<K extends KoinosClass>(
    ClassT: KoinosClassBuilder<K>,
    vb: VariableBlob,
    blobSize?: number
  ): Vector<K> {
    const len = vb.deserialize(VarInt).toNumber();
    const items = new Array(len).fill(null).map(() => {
      if (new FixedBlob() instanceof ClassT)
        return vb.deserialize(ClassT, blobSize);
      return vb.deserialize(ClassT);
    });
    return new Vector<K>(ClassT, items);
  }

  calcSerializedSize(): number {
    const header = sizeVarInt(this.items.length);
    const sizes = this.items.map((item) => item.calcSerializedSize());
    const dataSize = sizes.reduce((t, v) => t + v, 0);
    return header + dataSize;
  }

  toJSON(nameOrCode?: multibase.BaseNameOrCode): unknown[] {
    return this.items.map((item) => item.toJSON(nameOrCode));
  }
}

export default Vector;
