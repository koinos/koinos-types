import { VariableBlob } from "./VariableBlob";

export class KoinosBoolean {
  
    public bool: boolean;
  
    constructor(bool: boolean = false) {
      this.bool = bool;
    }
  
    serialize(vb: VariableBlob): VariableBlob {
      return vb.buffer.writeByte(this.bool ? 1: 0);
    }
  
    toBoolean(): boolean {
      return this.bool;
    }
  
    toString(): string {
      return this.bool.toString();
    }
  }

export default KoinosBoolean;
