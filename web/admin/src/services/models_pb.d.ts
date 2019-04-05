// package: models
// file: models.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class ListOpts extends jspb.Message {
  getOffset(): number;
  setOffset(value: number): void;

  getLimit(): number;
  setLimit(value: number): void;

  clearSortList(): void;
  getSortList(): Array<ListSortOpts>;
  setSortList(value: Array<ListSortOpts>): void;
  addSort(value?: ListSortOpts, index?: number): ListSortOpts;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListOpts.AsObject;
  static toObject(includeInstance: boolean, msg: ListOpts): ListOpts.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListOpts;
  static deserializeBinaryFromReader(message: ListOpts, reader: jspb.BinaryReader): ListOpts;
}

export namespace ListOpts {
  export type AsObject = {
    offset: number,
    limit: number,
    sortList: Array<ListSortOpts.AsObject>,
  }
}

export class ListSortOpts extends jspb.Message {
  getField(): string;
  setField(value: string): void;

  getAscending(): boolean;
  setAscending(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListSortOpts.AsObject;
  static toObject(includeInstance: boolean, msg: ListSortOpts): ListSortOpts.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListSortOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListSortOpts;
  static deserializeBinaryFromReader(message: ListSortOpts, reader: jspb.BinaryReader): ListSortOpts;
}

export namespace ListSortOpts {
  export type AsObject = {
    field: string,
    ascending: boolean,
  }
}

export class Feedback extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getStatus(): string;
  setStatus(value: string): void;

  getUserid(): string;
  setUserid(value: string): void;

  hasCreatedat(): boolean;
  clearCreatedat(): void;
  getCreatedat(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedat(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getSessionid(): string;
  setSessionid(value: string): void;

  getReaction1(): string;
  setReaction1(value: string): void;

  getReaction2(): string;
  setReaction2(value: string): void;

  getReaction3(): string;
  setReaction3(value: string): void;

  getReaction4(): string;
  setReaction4(value: string): void;

  getVisitorid(): string;
  setVisitorid(value: string): void;

  hasVisitor(): boolean;
  clearVisitor(): void;
  getVisitor(): Visitor | undefined;
  setVisitor(value?: Visitor): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Feedback.AsObject;
  static toObject(includeInstance: boolean, msg: Feedback): Feedback.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Feedback, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Feedback;
  static deserializeBinaryFromReader(message: Feedback, reader: jspb.BinaryReader): Feedback;
}

export namespace Feedback {
  export type AsObject = {
    id: string,
    status: string,
    userid: string,
    createdat?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    sessionid: string,
    reaction1: string,
    reaction2: string,
    reaction3: string,
    reaction4: string,
    visitorid: string,
    visitor?: Visitor.AsObject,
  }
}

export class Visitor extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getStatus(): string;
  setStatus(value: string): void;

  getUserid(): string;
  setUserid(value: string): void;

  hasCreatedat(): boolean;
  clearCreatedat(): void;
  getCreatedat(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedat(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getName(): string;
  setName(value: string): void;

  getBookmark(): string;
  setBookmark(value: string): void;

  getPhotolink(): string;
  setPhotolink(value: string): void;

  clearFeedbacksList(): void;
  getFeedbacksList(): Array<Feedback>;
  setFeedbacksList(value: Array<Feedback>): void;
  addFeedbacks(value?: Feedback, index?: number): Feedback;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Visitor.AsObject;
  static toObject(includeInstance: boolean, msg: Visitor): Visitor.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Visitor, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Visitor;
  static deserializeBinaryFromReader(message: Visitor, reader: jspb.BinaryReader): Visitor;
}

export namespace Visitor {
  export type AsObject = {
    id: string,
    status: string,
    userid: string,
    createdat?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    name: string,
    bookmark: string,
    photolink: string,
    feedbacksList: Array<Feedback.AsObject>,
  }
}

export class Rating extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getStatus(): string;
  setStatus(value: string): void;

  getUserid(): string;
  setUserid(value: string): void;

  hasCreatedat(): boolean;
  clearCreatedat(): void;
  getCreatedat(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedat(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedat(): boolean;
  clearUpdatedat(): void;
  getUpdatedat(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedat(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getRatingcount(): string;
  setRatingcount(value: string): void;

  getScore(): number;
  setScore(value: number): void;

  getReactionsummary(): string;
  setReactionsummary(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Rating.AsObject;
  static toObject(includeInstance: boolean, msg: Rating): Rating.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Rating, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Rating;
  static deserializeBinaryFromReader(message: Rating, reader: jspb.BinaryReader): Rating;
}

export namespace Rating {
  export type AsObject = {
    id: string,
    status: string,
    userid: string,
    createdat?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedat?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    ratingcount: string,
    score: number,
    reactionsummary: string,
  }
}

