import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_UINT128 = BigInt("0x" + "F".repeat(32));
export class UInt128 extends KoinosBigInt {
    constructor(number: bigint | string | number = 0) {
        super(number, 128, MAX_UINT128);
      }
}

export default UInt128;