import { VariableBlob } from "./VariableBlob";

export type BooleanLike = boolean | Bool;

export class Bool {
  public bool: boolean;

  constructor(bool: BooleanLike = false) {
    this.bool = bool instanceof Bool ? bool.bool : bool;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(1);
    vb.buffer.writeByte(this.bool ? 1 : 0);
    if (!blob) vb.flip();
    return vb;
  }

  static deserialize(vb: VariableBlob): Bool {
    if (vb.buffer.limit === 0) throw new Error("Unexpected EOF");
    const value = vb.buffer.readByte();
    if (value !== 0 && value !== 1) throw new Error("Boolean must be 0 or 1");
    return new Bool(!!value);
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
