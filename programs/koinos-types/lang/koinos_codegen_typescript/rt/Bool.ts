import { VariableBlob } from "./VariableBlob";

export type BooleanLike = boolean | Bool;

export class Bool {
  public bool: boolean;

  constructor(bool: BooleanLike = false) {
    this.bool = bool instanceof Bool ? bool.bool : bool;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    vb.writeUint8(this.bool ? 1 : 0);
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): Bool {
    const value = vb.readUint8();
    if (value !== 0 && value !== 1) throw new Error("Boolean must be 0 or 1");
    return new Bool(!!value);
  }

  calcSerializedSize(): number {
    return 1;
  }

  toBoolean(): boolean {
    return this.bool;
  }

  toJSON(): boolean {
    return this.bool;
  }

  toString(): string {
    return this.bool.toString();
  }
}

export default Bool;
