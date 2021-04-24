import { VariableBlob, KoinosClass, KoinosClassBuilder } from "./VariableBlob";
import VarInt from "./VarInt";

export class Variant<
  A extends KoinosClass,
  B extends KoinosClass,
  C extends KoinosClass,
  D extends KoinosClass,
  E extends KoinosClass,
  F extends KoinosClass,
  G extends KoinosClass,
  H extends KoinosClass,
  I extends KoinosClass,
  J extends KoinosClass
> {
  public value: A | B | C | D | E | F | G | H | I | J;

  public Classes: KoinosClassBuilder<A | B | C | D | E | F | G | H | I | J>[];

  constructor(
    value: A | B | C | D | E | F | G | H | I | J,
    Classes: KoinosClassBuilder<A | B | C | D | E | F | G | H | I | J>[]
  ) {
    this.value = value;
    this.Classes = Classes;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    for (let i = 0; i < this.Classes.length; i += 1)
      if (this.value instanceof this.Classes[i]) {
        vb.serialize(new VarInt(i));
        break;
      }
    vb.serialize(this.value);
    if (!blob) vb.resetCursor();
    return vb;
  }

  deserializeVariant(vb: VariableBlob): Variant<A, B, C, D, E, F, G, H, I, J> {
    const i = vb.deserialize(VarInt).toNumber();
    if (i >= this.Classes.length) throw new Error("Unknown variant tag");
    const value = vb.deserialize(this.Classes[i]);
    return new Variant(value, this.Classes);
  }

  calcSerializedSize(): number {
    return 1 + this.value.calcSerializedSize();
  }

  toJSON() {
    return this.value.toJSON();
  }
}
