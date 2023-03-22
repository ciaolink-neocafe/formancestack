/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import * as shared from "../shared";
import { AxiosResponse } from "axios";

export class GetManyConfigsRequest extends SpeakeasyBase {
  /**
   * Optional filter by endpoint URL
   */
  @SpeakeasyMetadata({
    data: "queryParam, style=form;explode=true;name=endpoint",
  })
  endpoint?: string;

  /**
   * Optional filter by Config ID
   */
  @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=id" })
  id?: string;
}

export class GetManyConfigsResponse extends SpeakeasyBase {
  /**
   * OK
   */
  @SpeakeasyMetadata()
  configsResponse?: shared.ConfigsResponse;

  @SpeakeasyMetadata()
  contentType: string;

  /**
   * Error
   */
  @SpeakeasyMetadata()
  errorResponse?: shared.ErrorResponse;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  rawResponse?: AxiosResponse;
}
