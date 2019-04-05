// package: data
// file: service_data.proto

import * as jspb from "google-protobuf";
import * as models_pb from "./models_pb";
import * as service_admin_pb from "./service_admin_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class GetByIDRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetByIDRequest): GetByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByIDRequest;
  static deserializeBinaryFromReader(message: GetByIDRequest, reader: jspb.BinaryReader): GetByIDRequest;
}

export namespace GetByIDRequest {
  export type AsObject = {
    key: string,
    id: string,
  }
}

export class GetBySlugRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getSlug(): string;
  setSlug(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetBySlugRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetBySlugRequest): GetBySlugRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetBySlugRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetBySlugRequest;
  static deserializeBinaryFromReader(message: GetBySlugRequest, reader: jspb.BinaryReader): GetBySlugRequest;
}

export namespace GetBySlugRequest {
  export type AsObject = {
    key: string,
    slug: string,
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
  getFiltersList(): Array<service_admin_pb.Filter>;
  setFiltersList(value: Array<service_admin_pb.Filter>): void;
  addFilters(value?: service_admin_pb.Filter, index?: number): service_admin_pb.Filter;

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
    filtersList: Array<service_admin_pb.Filter.AsObject>,
  }
}

export class TimestampResponse extends jspb.Message {
  hasModifiedat(): boolean;
  clearModifiedat(): void;
  getModifiedat(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setModifiedat(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TimestampResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TimestampResponse): TimestampResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TimestampResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TimestampResponse;
  static deserializeBinaryFromReader(message: TimestampResponse, reader: jspb.BinaryReader): TimestampResponse;
}

export namespace TimestampResponse {
  export type AsObject = {
    modifiedat?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class FeedbackResponse extends jspb.Message {
  hasFeedback(): boolean;
  clearFeedback(): void;
  getFeedback(): models_pb.Feedback | undefined;
  setFeedback(value?: models_pb.Feedback): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FeedbackResponse.AsObject;
  static toObject(includeInstance: boolean, msg: FeedbackResponse): FeedbackResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: FeedbackResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FeedbackResponse;
  static deserializeBinaryFromReader(message: FeedbackResponse, reader: jspb.BinaryReader): FeedbackResponse;
}

export namespace FeedbackResponse {
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

export class VisitorResponse extends jspb.Message {
  hasVisitor(): boolean;
  clearVisitor(): void;
  getVisitor(): models_pb.Visitor | undefined;
  setVisitor(value?: models_pb.Visitor): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VisitorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: VisitorResponse): VisitorResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VisitorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VisitorResponse;
  static deserializeBinaryFromReader(message: VisitorResponse, reader: jspb.BinaryReader): VisitorResponse;
}

export namespace VisitorResponse {
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

export class RatingResponse extends jspb.Message {
  hasRating(): boolean;
  clearRating(): void;
  getRating(): models_pb.Rating | undefined;
  setRating(value?: models_pb.Rating): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RatingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RatingResponse): RatingResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RatingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RatingResponse;
  static deserializeBinaryFromReader(message: RatingResponse, reader: jspb.BinaryReader): RatingResponse;
}

export namespace RatingResponse {
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

export class DeckRequest extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  clearDecksList(): void;
  getDecksList(): Array<string>;
  setDecksList(value: Array<string>): void;
  addDecks(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeckRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeckRequest): DeckRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeckRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeckRequest;
  static deserializeBinaryFromReader(message: DeckRequest, reader: jspb.BinaryReader): DeckRequest;
}

export namespace DeckRequest {
  export type AsObject = {
    key: string,
    decksList: Array<string>,
  }
}

export class DeckContentResponse extends jspb.Message {
  clearDeckcontentsList(): void;
  getDeckcontentsList(): Array<DeckContent>;
  setDeckcontentsList(value: Array<DeckContent>): void;
  addDeckcontents(value?: DeckContent, index?: number): DeckContent;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeckContentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeckContentResponse): DeckContentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeckContentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeckContentResponse;
  static deserializeBinaryFromReader(message: DeckContentResponse, reader: jspb.BinaryReader): DeckContentResponse;
}

export namespace DeckContentResponse {
  export type AsObject = {
    deckcontentsList: Array<DeckContent.AsObject>,
  }
}

export class DeckContent extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  clearCardrefsList(): void;
  getCardrefsList(): Array<CardReference>;
  setCardrefsList(value: Array<CardReference>): void;
  addCardrefs(value?: CardReference, index?: number): CardReference;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeckContent.AsObject;
  static toObject(includeInstance: boolean, msg: DeckContent): DeckContent.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeckContent, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeckContent;
  static deserializeBinaryFromReader(message: DeckContent, reader: jspb.BinaryReader): DeckContent;
}

export namespace DeckContent {
  export type AsObject = {
    name: string,
    cardrefsList: Array<CardReference.AsObject>,
  }
}

export class CardReference extends jspb.Message {
  getType(): string;
  setType(value: string): void;

  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CardReference.AsObject;
  static toObject(includeInstance: boolean, msg: CardReference): CardReference.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CardReference, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CardReference;
  static deserializeBinaryFromReader(message: CardReference, reader: jspb.BinaryReader): CardReference;
}

export namespace CardReference {
  export type AsObject = {
    type: string,
    id: string,
  }
}

