import { VariableBlob, KoinosClass, KoinosClassBuilder } from "./VariableBlob";
import { VarInt } from "./VarInt";

export class Vector<T extends KoinosClass> {
  public items: T[];

  constructor(items: T[] = []) {
    this.items = items;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob();
    vb.serialize(new VarInt(this.items.length));
    this.items.forEach((item) => vb.serialize(item));
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize<K extends KoinosClass>(
    ClassT: KoinosClassBuilder<K>,
    vb: VariableBlob
  ): Vector<K> {
    const len = vb.deserialize(VarInt).num;
    const items = new Array(len).fill(null).map(() => vb.deserialize(ClassT));
    return new Vector<K>(items);
  }
}

export default Vector;
