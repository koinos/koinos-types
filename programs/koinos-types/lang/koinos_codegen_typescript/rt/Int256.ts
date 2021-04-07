import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_INT256 = BigInt("0x7" + "F".repeat(63));
export const MIN_INT256 = -BigInt("0x8" + "0".repeat(63));
export class Int256 extends KoinosBigInt {
    constructor(number: bigint | string | number = 0) {
        super(number, 256, MAX_INT256, MIN_INT256);
      }
}

export default Int256;