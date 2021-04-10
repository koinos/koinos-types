import { VariableBlob, KoinosClass, KoinosClassBuilder } from "./VariableBlob";

export class Opaque<T extends KoinosClass> {
  private native: T;

  private blob: VariableBlob;

  private ClassT: KoinosClassBuilder<T>;

  constructor(ClassT: KoinosClassBuilder<T>, nativeOrBlob?: T | VariableBlob) {
    this.ClassT = ClassT;
    if (nativeOrBlob instanceof ClassT) {
      this.native = nativeOrBlob;
    } else if (nativeOrBlob instanceof VariableBlob) {
      this.blob = nativeOrBlob;
    } else {
      this.native = new ClassT();
    }
  }

  unbox(): void {
    if (!this.native && this.blob)
      this.native = this.blob.deserialize(this.ClassT);
  }

  box(): void {
    if (this.native) {
      if (!this.blob) this.serialize();
      this.native = null;
    }
  }

  isUnboxed(): boolean {
    return !!this.native;
  }

  makeImmutable(): void {
    if (this.native && !this.blob) this.serialize();
  }

  makeMutable(): void {
    if (!this.native) this.unbox();
    if (this.native && this.blob) this.blob = null;
  }

  isMutable(): boolean {
    return !!this.native && !this.blob;
  }

  getBlob(): VariableBlob {
    if (this.native && !this.blob) this.serialize();
    return this.blob;
  }

  getNative(): T {
    if (!this.native) throw new Error("Opaque type not unboxed");
    if (this.blob) throw new Error("Opaque type is not mutable");
    return this.native;
  }

  private serialize(): void {
    this.blob = this.native.serialize();
  }
}

export default Opaque;
