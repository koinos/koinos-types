import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_UINT160 = BigInt("0x" + "F".repeat(40));
export class UInt160 extends KoinosBigInt {
  constructor(number: bigint | string | number = 0) {
    super(number, 160, MAX_UINT160);
  }
}

export default UInt160;
