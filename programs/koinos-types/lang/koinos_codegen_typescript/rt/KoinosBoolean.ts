import { VariableBlob } from "./VariableBlob";

export class KoinosBoolean {
  public bool: boolean;

  constructor(bool = false) {
    this.bool = bool;
  }

  serialize(vb: VariableBlob): VariableBlob {
    vb.buffer.writeByte(this.bool ? 1 : 0);
    return vb;
  }

  static deserialize(vb: VariableBlob): KoinosBoolean {
    if (vb.buffer.limit === 0) throw new Error("Unexpected EOF");
    const value = vb.buffer.readByte();
    if (value !== 0 && value !== 1) throw new Error("Boolean must be 0 or 1");
    return new KoinosBoolean(!!value);
  }

  toBoolean(): boolean {
    return this.bool;
  }

  toString(): string {
    return this.bool.toString();
  }
}

export default KoinosBoolean;
