import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_UINT64  = BigInt("0x" + "F".repeat(16));
export class UInt64 extends KoinosBigInt {
    constructor(number: bigint | string | number = 0) {
        super(number, 64, MAX_UINT64);
      }
}

export default UInt64;