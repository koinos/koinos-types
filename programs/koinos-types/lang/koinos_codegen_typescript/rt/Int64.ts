import { KoinosBigInt } from "./KoinosBigInt";

export const MAX_INT64 = BigInt("0x7" + "F".repeat(15));
export const MIN_INT64 = -BigInt("0x8" + "0".repeat(15));
export class Int64 extends KoinosBigInt {
  constructor(number: bigint | string | number = 0) {
    super(number, 64, MAX_INT64, MIN_INT64);
  }
}

export default Int64;
