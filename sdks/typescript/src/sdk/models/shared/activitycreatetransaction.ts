/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { PostTransaction } from "./posttransaction";
import { Expose, Type } from "class-transformer";

export class ActivityCreateTransaction extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "data" })
  @Type(() => PostTransaction)
  data?: PostTransaction;

  @SpeakeasyMetadata()
  @Expose({ name: "ledger" })
  ledger?: string;
}
