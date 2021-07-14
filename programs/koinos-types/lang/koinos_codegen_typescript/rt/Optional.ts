import { VariableBlob, KoinosClass, KoinosClassBuilder } from "./VariableBlob";
import { Bool } from "./Bool";

export class Optional<T extends KoinosClass> {
  public value: T;

  constructor(ClassT: KoinosClassBuilder<T>, value?: T | unknown) {
    if (value instanceof ClassT) {
      this.value = value;
    } else if (value === null || typeof value === "undefined") {
      this.value = undefined;
    } else {
      this.value = new ClassT(value);
    }
  }

  hasValue(): boolean {
    return this.value !== null && typeof this.value !== "undefined";
  }

  reset(): void {
    this.value = undefined;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    if (this.hasValue()) {
      vb.writeUint8(1);
      vb.serialize(this.value);
    } else {
      vb.writeUint8(0);
    }
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize<K extends KoinosClass>(
    ClassT: KoinosClassBuilder<K>,
    vb: VariableBlob
  ): Optional<K> {
    const hasValue = vb.deserialize(Bool).bool;
    if (!hasValue) return new Optional<K>(ClassT);
    const value = vb.deserialize(ClassT);
    return new Optional<K>(ClassT, value);
  }

  calcSerializedSize(): number {
    return this.hasValue() ? 1 + this.value.calcSerializedSize() : 1;
  }

  toJSON(): unknown {
    if (!this.hasValue()) return undefined;
    return this.value.toJSON();
  }
}

export default Optional;
