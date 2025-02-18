/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose } from "class-transformer";

export enum ErrorErrorCode {
  Validation = "VALIDATION",
  NotFound = "NOT_FOUND",
}

/**
 * General error
 */
export class ErrorT extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "errorCode" })
  errorCode: ErrorErrorCode;

  @SpeakeasyMetadata()
  @Expose({ name: "errorMessage" })
  errorMessage: string;
}
