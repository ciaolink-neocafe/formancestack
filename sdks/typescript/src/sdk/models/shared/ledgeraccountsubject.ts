/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose } from "class-transformer";

export class LedgerAccountSubject extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "identifier" })
  identifier: string;

  @SpeakeasyMetadata()
  @Expose({ name: "type" })
  type: string;
}
