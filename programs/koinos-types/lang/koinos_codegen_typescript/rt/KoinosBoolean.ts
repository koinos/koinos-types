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

  toBoolean(): boolean {
    return this.bool;
  }

  toString(): string {
    return this.bool.toString();
  }
}

export default KoinosBoolean;
