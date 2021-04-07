import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_INT160 = BigInt("0x7" + "F".repeat(39));
export const MIN_INT160 = -BigInt("0x8" + "0".repeat(39));
export class Int160 extends KoinosBigInt {
    constructor(number: bigint | string | number = 0) {
        super(number, 160, MAX_INT160, MIN_INT160);
      }
}

export default Int160;