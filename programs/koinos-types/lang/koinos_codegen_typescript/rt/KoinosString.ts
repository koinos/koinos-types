import * as ByteBuffer from "bytebuffer";
import { VariableBlob } from "./VariableBlob";

export class KoinosString {
  
    public str: string;
  
    constructor(str: string = "") {
      this.str = str;
    }
    
    serialize(vb: VariableBlob): VariableBlob {
      const buffer = ByteBuffer.fromUTF8(this.str);
      return vb.buffer
        .writeVarint64(buffer.limit)
        .append(buffer);
    }
  
    toString(): string {
      return this.str;
    }
  }

export default KoinosString;