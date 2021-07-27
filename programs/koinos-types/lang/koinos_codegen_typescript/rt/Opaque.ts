import { VariableBlob, KoinosClass, KoinosClassBuilder } from "./VariableBlob";

export class Opaque<T extends KoinosClass> {
  private native: T;

  private blob: VariableBlob;

  private ClassT: KoinosClassBuilder<T>;

  constructor(
    ClassT: KoinosClassBuilder<T>,
    data?: T | VariableBlob | unknown
  ) {
    this.ClassT = ClassT;
    if (data instanceof ClassT) {
      this.native = data;
    } else if (data instanceof VariableBlob) {
      this.blob = data;
    } else {
      this.native = new ClassT(data);
    }
  }

  unbox(): void {
    if (!this.native && this.blob)
      this.native = this.blob.deserialize(this.ClassT);
  }

  box(): void {
    if (this.native) {
      if (!this.blob) this.blob = this.native.serialize();
      this.native = null;
    }
  }

  isUnboxed(): boolean {
    return !!this.native;
  }

  makeImmutable(): void {
    if (this.native && !this.blob) this.blob = this.native.serialize();
  }

  makeMutable(): void {
    if (!this.native) this.unbox();
    if (this.native && this.blob) this.blob = null;
  }

  isMutable(): boolean {
    return !!this.native && !this.blob;
  }

  getBlob(): VariableBlob {
    if (this.native && !this.blob) this.blob = this.native.serialize();
    return this.blob;
  }

  getNative(): T {
    if (!this.native) throw new Error("Opaque type not unboxed");
    if (this.blob) throw new Error("Opaque type is not mutable");
    return this.native;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.serialize(this.getBlob());
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize<K extends KoinosClass>(
    ClassT: KoinosClassBuilder<K>,
    vb: VariableBlob
  ): Opaque<K> {
    const blob = vb.deserializeVariableBlob();
    return new Opaque<K>(ClassT, blob);
  }

  calcSerializedSize(): number {
    return this.isUnboxed()
      ? this.native.calcSerializedSize()
      : this.blob.length();
  }

  toJSON(): unknown {
    if (!this.isUnboxed()) throw new Error("Opaque type not unboxed");
    return this.native.toJSON();
  }
}

export default Opaque;
