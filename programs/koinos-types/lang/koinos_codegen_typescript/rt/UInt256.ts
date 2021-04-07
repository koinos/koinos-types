import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_UINT256 = BigInt("0x" + "F".repeat(64));
export class UInt256 extends KoinosBigInt {
    constructor(number: bigint | string | number = 0) {
        super(number, 256, MAX_UINT256);
      }
}

export default UInt256;