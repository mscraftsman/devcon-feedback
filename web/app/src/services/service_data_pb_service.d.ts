// package: data
// file: service_data.proto

import * as service_data_pb from "./service_data_pb";
import {grpc} from "grpc-web-client";

type DataGetFeedbackByID = {
  readonly methodName: string;
  readonly service: typeof Data;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_data_pb.GetByIDRequest;
  readonly responseType: typeof service_data_pb.FeedbackResponse;
};

type DataListFeedbacks = {
  readonly methodName: string;
  readonly service: typeof Data;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_data_pb.ListRequest;
  readonly responseType: typeof service_data_pb.ListFeedbacksResponse;
};

type DataGetVisitorByID = {
  readonly methodName: string;
  readonly service: typeof Data;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_data_pb.GetByIDRequest;
  readonly responseType: typeof service_data_pb.VisitorResponse;
};

type DataListVisitors = {
  readonly methodName: string;
  readonly service: typeof Data;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_data_pb.ListRequest;
  readonly responseType: typeof service_data_pb.ListVisitorsResponse;
};

type DataGetRatingByID = {
  readonly methodName: string;
  readonly service: typeof Data;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_data_pb.GetByIDRequest;
  readonly responseType: typeof service_data_pb.RatingResponse;
};

type DataGetModifiedRatingByID = {
  readonly methodName: string;
  readonly service: typeof Data;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_data_pb.GetByIDRequest;
  readonly responseType: typeof service_data_pb.TimestampResponse;
};

type DataListRatings = {
  readonly methodName: string;
  readonly service: typeof Data;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_data_pb.ListRequest;
  readonly responseType: typeof service_data_pb.ListRatingsResponse;
};

export class Data {
  static readonly serviceName: string;
  static readonly GetFeedbackByID: DataGetFeedbackByID;
  static readonly ListFeedbacks: DataListFeedbacks;
  static readonly GetVisitorByID: DataGetVisitorByID;
  static readonly ListVisitors: DataListVisitors;
  static readonly GetRatingByID: DataGetRatingByID;
  static readonly GetModifiedRatingByID: DataGetModifiedRatingByID;
  static readonly ListRatings: DataListRatings;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: () => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: () => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class DataClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getFeedbackByID(
    requestMessage: service_data_pb.GetByIDRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.FeedbackResponse|null) => void
  ): UnaryResponse;
  getFeedbackByID(
    requestMessage: service_data_pb.GetByIDRequest,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.FeedbackResponse|null) => void
  ): UnaryResponse;
  listFeedbacks(
    requestMessage: service_data_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.ListFeedbacksResponse|null) => void
  ): UnaryResponse;
  listFeedbacks(
    requestMessage: service_data_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.ListFeedbacksResponse|null) => void
  ): UnaryResponse;
  getVisitorByID(
    requestMessage: service_data_pb.GetByIDRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.VisitorResponse|null) => void
  ): UnaryResponse;
  getVisitorByID(
    requestMessage: service_data_pb.GetByIDRequest,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.VisitorResponse|null) => void
  ): UnaryResponse;
  listVisitors(
    requestMessage: service_data_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.ListVisitorsResponse|null) => void
  ): UnaryResponse;
  listVisitors(
    requestMessage: service_data_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.ListVisitorsResponse|null) => void
  ): UnaryResponse;
  getRatingByID(
    requestMessage: service_data_pb.GetByIDRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.RatingResponse|null) => void
  ): UnaryResponse;
  getRatingByID(
    requestMessage: service_data_pb.GetByIDRequest,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.RatingResponse|null) => void
  ): UnaryResponse;
  getModifiedRatingByID(
    requestMessage: service_data_pb.GetByIDRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.TimestampResponse|null) => void
  ): UnaryResponse;
  getModifiedRatingByID(
    requestMessage: service_data_pb.GetByIDRequest,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.TimestampResponse|null) => void
  ): UnaryResponse;
  listRatings(
    requestMessage: service_data_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.ListRatingsResponse|null) => void
  ): UnaryResponse;
  listRatings(
    requestMessage: service_data_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: service_data_pb.ListRatingsResponse|null) => void
  ): UnaryResponse;
}

