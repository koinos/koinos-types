import { VariableBlob } from "./VariableBlob";
import { VarInt } from "./VarInt";

export type StringLike = string | Str;

export class Str {
  public str: string;

  constructor(str: StringLike = "") {
    this.str = str instanceof Str ? str.str : str;
  }

  serialize(blob?: VariableBlob): VariableBlob {
    const vb = blob || new VariableBlob(this.calcSerializedSize());
    const bytes = new TextEncoder().encode(this.str);
    vb.serialize(new VarInt(bytes.length));
    vb.write(bytes);
    if (!blob) vb.resetCursor();
    return vb;
  }

  static deserialize(vb: VariableBlob): Str {
    const size = vb.deserialize(VarInt).toNumber();
    const subBuffer = vb.read(size);
    const str = new TextDecoder().decode(subBuffer);
    return new Str(str);
  }

  calcSerializedSize(): number {
    const size = new TextEncoder().encode(this.str).length;
    const header = Math.ceil(Math.log2(size + 1) / 7);
    return header + size;
  }

  toString(): string {
    return this.str;
  }

  toJSON(): string {
    return this.str;
  }
}

export default Str;
