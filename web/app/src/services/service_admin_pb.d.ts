// package: admin
// file: service_admin.proto

import * as jspb from "google-protobuf";
import * as models_pb from "./models_pb";

export class Filter extends jspb.Message {
  getField(): string;
  setField(value: string): void;

  getOperation(): string;
  setOperation(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Filter.AsObject;
  static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Filter;
  static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
}

export namespace Filter {
  export type AsObject = {
    field: string,
    operation: string,
    value: string,
  }
}

export class GetRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRequest): GetRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRequest;
  static deserializeBinaryFromReader(message: GetRequest, reader: jspb.BinaryReader): GetRequest;
}

export namespace GetRequest {
  export type AsObject = {
    key: string,
    id: string,
  }
}

export class ListRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  hasListopts(): boolean;
  clearListopts(): void;
  getListopts(): models_pb.ListOpts | undefined;
  setListopts(value?: models_pb.ListOpts): void;

  clearFiltersList(): void;
  getFiltersList(): Array<Filter>;
  setFiltersList(value: Array<Filter>): void;
  addFilters(value?: Filter, index?: number): Filter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRequest): ListRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRequest;
  static deserializeBinaryFromReader(message: ListRequest, reader: jspb.BinaryReader): ListRequest;
}

export namespace ListRequest {
  export type AsObject = {
    key: string,
    listopts?: models_pb.ListOpts.AsObject,
    filtersList: Array<Filter.AsObject>,
  }
}

export class CountResponse extends jspb.Message {
  getCount(): number;
  setCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CountResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CountResponse): CountResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CountResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CountResponse;
  static deserializeBinaryFromReader(message: CountResponse, reader: jspb.BinaryReader): CountResponse;
}

export namespace CountResponse {
  export type AsObject = {
    count: number,
  }
}

export class DeleteRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRequest): DeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRequest;
  static deserializeBinaryFromReader(message: DeleteRequest, reader: jspb.BinaryReader): DeleteRequest;
}

export namespace DeleteRequest {
  export type AsObject = {
    key: string,
    id: string,
  }
}

export class DeleteResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteResponse): DeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteResponse;
  static deserializeBinaryFromReader(message: DeleteResponse, reader: jspb.BinaryReader): DeleteResponse;
}

export namespace DeleteResponse {
  export type AsObject = {
  }
}

export class UploadOpts extends jspb.Message {
  getWatermark(): boolean;
  setWatermark(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UploadOpts.AsObject;
  static toObject(includeInstance: boolean, msg: UploadOpts): UploadOpts.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UploadOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UploadOpts;
  static deserializeBinaryFromReader(message: UploadOpts, reader: jspb.BinaryReader): UploadOpts;
}

export namespace UploadOpts {
  export type AsObject = {
    watermark: boolean,
  }
}

export class UploadRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getId(): string;
  setId(value: string): void;

  getField(): string;
  setField(value: string): void;

  getFilename(): string;
  setFilename(value: string): void;

  getContent(): Uint8Array | string;
  getContent_asU8(): Uint8Array;
  getContent_asB64(): string;
  setContent(value: Uint8Array | string): void;

  hasUploadoptions(): boolean;
  clearUploadoptions(): void;
  getUploadoptions(): UploadOpts | undefined;
  setUploadoptions(value?: UploadOpts): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UploadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UploadRequest): UploadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UploadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UploadRequest;
  static deserializeBinaryFromReader(message: UploadRequest, reader: jspb.BinaryReader): UploadRequest;
}

export namespace UploadRequest {
  export type AsObject = {
    key: string,
    id: string,
    field: string,
    filename: string,
    content: Uint8Array | string,
    uploadoptions?: UploadOpts.AsObject,
  }
}

export class UploadResponse extends jspb.Message {
  getUri(): string;
  setUri(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UploadResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UploadResponse): UploadResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UploadResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UploadResponse;
  static deserializeBinaryFromReader(message: UploadResponse, reader: jspb.BinaryReader): UploadResponse;
}

export namespace UploadResponse {
  export type AsObject = {
    uri: string,
  }
}

export class LookupRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getTerm(): string;
  setTerm(value: string): void;

  hasListopts(): boolean;
  clearListopts(): void;
  getListopts(): models_pb.ListOpts | undefined;
  setListopts(value?: models_pb.ListOpts): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LookupRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LookupRequest): LookupRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LookupRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LookupRequest;
  static deserializeBinaryFromReader(message: LookupRequest, reader: jspb.BinaryReader): LookupRequest;
}

export namespace LookupRequest {
  export type AsObject = {
    key: string,
    term: string,
    listopts?: models_pb.ListOpts.AsObject,
  }
}

export class LookupResult extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getLabel(): string;
  setLabel(value: string): void;

  getStatus(): string;
  setStatus(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LookupResult.AsObject;
  static toObject(includeInstance: boolean, msg: LookupResult): LookupResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LookupResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LookupResult;
  static deserializeBinaryFromReader(message: LookupResult, reader: jspb.BinaryReader): LookupResult;
}

export namespace LookupResult {
  export type AsObject = {
    id: string,
    label: string,
    status: string,
  }
}

export class LookupResponse extends jspb.Message {
  clearResultList(): void;
  getResultList(): Array<LookupResult>;
  setResultList(value: Array<LookupResult>): void;
  addResult(value?: LookupResult, index?: number): LookupResult;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LookupResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LookupResponse): LookupResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LookupResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LookupResponse;
  static deserializeBinaryFromReader(message: LookupResponse, reader: jspb.BinaryReader): LookupResponse;
}

export namespace LookupResponse {
  export type AsObject = {
    resultList: Array<LookupResult.AsObject>,
  }
}

export class CreateFeedbackRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  hasFeedback(): boolean;
  clearFeedback(): void;
  getFeedback(): models_pb.Feedback | undefined;
  setFeedback(value?: models_pb.Feedback): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateFeedbackRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateFeedbackRequest): CreateFeedbackRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateFeedbackRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateFeedbackRequest;
  static deserializeBinaryFromReader(message: CreateFeedbackRequest, reader: jspb.BinaryReader): CreateFeedbackRequest;
}

export namespace CreateFeedbackRequest {
  export type AsObject = {
    key: string,
    feedback?: models_pb.Feedback.AsObject,
  }
}

export class CreateFeedbackResponse extends jspb.Message {
  hasFeedback(): boolean;
  clearFeedback(): void;
  getFeedback(): models_pb.Feedback | undefined;
  setFeedback(value?: models_pb.Feedback): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateFeedbackResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateFeedbackResponse): CreateFeedbackResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateFeedbackResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateFeedbackResponse;
  static deserializeBinaryFromReader(message: CreateFeedbackResponse, reader: jspb.BinaryReader): CreateFeedbackResponse;
}

export namespace CreateFeedbackResponse {
  export type AsObject = {
    feedback?: models_pb.Feedback.AsObject,
  }
}

export class GetFeedbackResponse extends jspb.Message {
  hasFeedback(): boolean;
  clearFeedback(): void;
  getFeedback(): models_pb.Feedback | undefined;
  setFeedback(value?: models_pb.Feedback): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFeedbackResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetFeedbackResponse): GetFeedbackResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetFeedbackResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFeedbackResponse;
  static deserializeBinaryFromReader(message: GetFeedbackResponse, reader: jspb.BinaryReader): GetFeedbackResponse;
}

export namespace GetFeedbackResponse {
  export type AsObject = {
    feedback?: models_pb.Feedback.AsObject,
  }
}

export class ListFeedbacksResponse extends jspb.Message {
  clearFeedbacksList(): void;
  getFeedbacksList(): Array<models_pb.Feedback>;
  setFeedbacksList(value: Array<models_pb.Feedback>): void;
  addFeedbacks(value?: models_pb.Feedback, index?: number): models_pb.Feedback;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFeedbacksResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListFeedbacksResponse): ListFeedbacksResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListFeedbacksResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFeedbacksResponse;
  static deserializeBinaryFromReader(message: ListFeedbacksResponse, reader: jspb.BinaryReader): ListFeedbacksResponse;
}

export namespace ListFeedbacksResponse {
  export type AsObject = {
    feedbacksList: Array<models_pb.Feedback.AsObject>,
  }
}

export class UpdateFeedbackRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  hasFeedback(): boolean;
  clearFeedback(): void;
  getFeedback(): models_pb.Feedback | undefined;
  setFeedback(value?: models_pb.Feedback): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateFeedbackRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateFeedbackRequest): UpdateFeedbackRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateFeedbackRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateFeedbackRequest;
  static deserializeBinaryFromReader(message: UpdateFeedbackRequest, reader: jspb.BinaryReader): UpdateFeedbackRequest;
}

export namespace UpdateFeedbackRequest {
  export type AsObject = {
    key: string,
    feedback?: models_pb.Feedback.AsObject,
  }
}

export class UpdateFeedbackResponse extends jspb.Message {
  hasFeedback(): boolean;
  clearFeedback(): void;
  getFeedback(): models_pb.Feedback | undefined;
  setFeedback(value?: models_pb.Feedback): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateFeedbackResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateFeedbackResponse): UpdateFeedbackResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateFeedbackResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateFeedbackResponse;
  static deserializeBinaryFromReader(message: UpdateFeedbackResponse, reader: jspb.BinaryReader): UpdateFeedbackResponse;
}

export namespace UpdateFeedbackResponse {
  export type AsObject = {
    feedback?: models_pb.Feedback.AsObject,
  }
}

export class CreateVisitorRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  hasVisitor(): boolean;
  clearVisitor(): void;
  getVisitor(): models_pb.Visitor | undefined;
  setVisitor(value?: models_pb.Visitor): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateVisitorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateVisitorRequest): CreateVisitorRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateVisitorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateVisitorRequest;
  static deserializeBinaryFromReader(message: CreateVisitorRequest, reader: jspb.BinaryReader): CreateVisitorRequest;
}

export namespace CreateVisitorRequest {
  export type AsObject = {
    key: string,
    visitor?: models_pb.Visitor.AsObject,
  }
}

export class CreateVisitorResponse extends jspb.Message {
  hasVisitor(): boolean;
  clearVisitor(): void;
  getVisitor(): models_pb.Visitor | undefined;
  setVisitor(value?: models_pb.Visitor): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateVisitorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateVisitorResponse): CreateVisitorResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateVisitorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateVisitorResponse;
  static deserializeBinaryFromReader(message: CreateVisitorResponse, reader: jspb.BinaryReader): CreateVisitorResponse;
}

export namespace CreateVisitorResponse {
  export type AsObject = {
    visitor?: models_pb.Visitor.AsObject,
  }
}

export class GetVisitorResponse extends jspb.Message {
  hasVisitor(): boolean;
  clearVisitor(): void;
  getVisitor(): models_pb.Visitor | undefined;
  setVisitor(value?: models_pb.Visitor): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetVisitorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetVisitorResponse): GetVisitorResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetVisitorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetVisitorResponse;
  static deserializeBinaryFromReader(message: GetVisitorResponse, reader: jspb.BinaryReader): GetVisitorResponse;
}

export namespace GetVisitorResponse {
  export type AsObject = {
    visitor?: models_pb.Visitor.AsObject,
  }
}

export class ListVisitorsResponse extends jspb.Message {
  clearVisitorsList(): void;
  getVisitorsList(): Array<models_pb.Visitor>;
  setVisitorsList(value: Array<models_pb.Visitor>): void;
  addVisitors(value?: models_pb.Visitor, index?: number): models_pb.Visitor;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListVisitorsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListVisitorsResponse): ListVisitorsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListVisitorsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListVisitorsResponse;
  static deserializeBinaryFromReader(message: ListVisitorsResponse, reader: jspb.BinaryReader): ListVisitorsResponse;
}

export namespace ListVisitorsResponse {
  export type AsObject = {
    visitorsList: Array<models_pb.Visitor.AsObject>,
  }
}

export class UpdateVisitorRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  hasVisitor(): boolean;
  clearVisitor(): void;
  getVisitor(): models_pb.Visitor | undefined;
  setVisitor(value?: models_pb.Visitor): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateVisitorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateVisitorRequest): UpdateVisitorRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateVisitorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateVisitorRequest;
  static deserializeBinaryFromReader(message: UpdateVisitorRequest, reader: jspb.BinaryReader): UpdateVisitorRequest;
}

export namespace UpdateVisitorRequest {
  export type AsObject = {
    key: string,
    visitor?: models_pb.Visitor.AsObject,
  }
}

export class UpdateVisitorResponse extends jspb.Message {
  hasVisitor(): boolean;
  clearVisitor(): void;
  getVisitor(): models_pb.Visitor | undefined;
  setVisitor(value?: models_pb.Visitor): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateVisitorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateVisitorResponse): UpdateVisitorResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateVisitorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateVisitorResponse;
  static deserializeBinaryFromReader(message: UpdateVisitorResponse, reader: jspb.BinaryReader): UpdateVisitorResponse;
}

export namespace UpdateVisitorResponse {
  export type AsObject = {
    visitor?: models_pb.Visitor.AsObject,
  }
}

export class CreateRatingRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  hasRating(): boolean;
  clearRating(): void;
  getRating(): models_pb.Rating | undefined;
  setRating(value?: models_pb.Rating): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRatingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRatingRequest): CreateRatingRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateRatingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRatingRequest;
  static deserializeBinaryFromReader(message: CreateRatingRequest, reader: jspb.BinaryReader): CreateRatingRequest;
}

export namespace CreateRatingRequest {
  export type AsObject = {
    key: string,
    rating?: models_pb.Rating.AsObject,
  }
}

export class CreateRatingResponse extends jspb.Message {
  hasRating(): boolean;
  clearRating(): void;
  getRating(): models_pb.Rating | undefined;
  setRating(value?: models_pb.Rating): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRatingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRatingResponse): CreateRatingResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateRatingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRatingResponse;
  static deserializeBinaryFromReader(message: CreateRatingResponse, reader: jspb.BinaryReader): CreateRatingResponse;
}

export namespace CreateRatingResponse {
  export type AsObject = {
    rating?: models_pb.Rating.AsObject,
  }
}

export class GetRatingResponse extends jspb.Message {
  hasRating(): boolean;
  clearRating(): void;
  getRating(): models_pb.Rating | undefined;
  setRating(value?: models_pb.Rating): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRatingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetRatingResponse): GetRatingResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetRatingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRatingResponse;
  static deserializeBinaryFromReader(message: GetRatingResponse, reader: jspb.BinaryReader): GetRatingResponse;
}

export namespace GetRatingResponse {
  export type AsObject = {
    rating?: models_pb.Rating.AsObject,
  }
}

export class ListRatingsResponse extends jspb.Message {
  clearRatingsList(): void;
  getRatingsList(): Array<models_pb.Rating>;
  setRatingsList(value: Array<models_pb.Rating>): void;
  addRatings(value?: models_pb.Rating, index?: number): models_pb.Rating;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRatingsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRatingsResponse): ListRatingsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListRatingsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRatingsResponse;
  static deserializeBinaryFromReader(message: ListRatingsResponse, reader: jspb.BinaryReader): ListRatingsResponse;
}

export namespace ListRatingsResponse {
  export type AsObject = {
    ratingsList: Array<models_pb.Rating.AsObject>,
  }
}

export class UpdateRatingRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  hasRating(): boolean;
  clearRating(): void;
  getRating(): models_pb.Rating | undefined;
  setRating(value?: models_pb.Rating): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRatingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRatingRequest): UpdateRatingRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateRatingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRatingRequest;
  static deserializeBinaryFromReader(message: UpdateRatingRequest, reader: jspb.BinaryReader): UpdateRatingRequest;
}

export namespace UpdateRatingRequest {
  export type AsObject = {
    key: string,
    rating?: models_pb.Rating.AsObject,
  }
}

export class UpdateRatingResponse extends jspb.Message {
  hasRating(): boolean;
  clearRating(): void;
  getRating(): models_pb.Rating | undefined;
  setRating(value?: models_pb.Rating): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRatingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRatingResponse): UpdateRatingResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateRatingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRatingResponse;
  static deserializeBinaryFromReader(message: UpdateRatingResponse, reader: jspb.BinaryReader): UpdateRatingResponse;
}

export namespace UpdateRatingResponse {
  export type AsObject = {
    rating?: models_pb.Rating.AsObject,
  }
}

