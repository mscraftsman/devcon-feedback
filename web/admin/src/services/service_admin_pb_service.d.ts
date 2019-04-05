// package: admin
// file: service_admin.proto

import * as service_admin_pb from "./service_admin_pb";
import {grpc} from "grpc-web-client";

type AdminCreateFeedback = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.CreateFeedbackRequest;
  readonly responseType: typeof service_admin_pb.CreateFeedbackResponse;
};

type AdminGetFeedback = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.GetRequest;
  readonly responseType: typeof service_admin_pb.GetFeedbackResponse;
};

type AdminListFeedbacks = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.ListRequest;
  readonly responseType: typeof service_admin_pb.ListFeedbacksResponse;
};

type AdminCountFeedbacks = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.ListRequest;
  readonly responseType: typeof service_admin_pb.CountResponse;
};

type AdminUpdateFeedback = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.UpdateFeedbackRequest;
  readonly responseType: typeof service_admin_pb.UpdateFeedbackResponse;
};

type AdminDeleteFeedback = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.DeleteRequest;
  readonly responseType: typeof service_admin_pb.DeleteResponse;
};

type AdminLookupFeedbacks = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.LookupRequest;
  readonly responseType: typeof service_admin_pb.LookupResponse;
};

type AdminCreateVisitor = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.CreateVisitorRequest;
  readonly responseType: typeof service_admin_pb.CreateVisitorResponse;
};

type AdminGetVisitor = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.GetRequest;
  readonly responseType: typeof service_admin_pb.GetVisitorResponse;
};

type AdminListVisitors = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.ListRequest;
  readonly responseType: typeof service_admin_pb.ListVisitorsResponse;
};

type AdminCountVisitors = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.ListRequest;
  readonly responseType: typeof service_admin_pb.CountResponse;
};

type AdminUpdateVisitor = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.UpdateVisitorRequest;
  readonly responseType: typeof service_admin_pb.UpdateVisitorResponse;
};

type AdminDeleteVisitor = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.DeleteRequest;
  readonly responseType: typeof service_admin_pb.DeleteResponse;
};

type AdminLookupVisitors = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.LookupRequest;
  readonly responseType: typeof service_admin_pb.LookupResponse;
};

type AdminCreateRating = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.CreateRatingRequest;
  readonly responseType: typeof service_admin_pb.CreateRatingResponse;
};

type AdminGetRating = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.GetRequest;
  readonly responseType: typeof service_admin_pb.GetRatingResponse;
};

type AdminListRatings = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.ListRequest;
  readonly responseType: typeof service_admin_pb.ListRatingsResponse;
};

type AdminCountRatings = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.ListRequest;
  readonly responseType: typeof service_admin_pb.CountResponse;
};

type AdminUpdateRating = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.UpdateRatingRequest;
  readonly responseType: typeof service_admin_pb.UpdateRatingResponse;
};

type AdminDeleteRating = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.DeleteRequest;
  readonly responseType: typeof service_admin_pb.DeleteResponse;
};

type AdminLookupRatings = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_admin_pb.LookupRequest;
  readonly responseType: typeof service_admin_pb.LookupResponse;
};

export class Admin {
  static readonly serviceName: string;
  static readonly CreateFeedback: AdminCreateFeedback;
  static readonly GetFeedback: AdminGetFeedback;
  static readonly ListFeedbacks: AdminListFeedbacks;
  static readonly CountFeedbacks: AdminCountFeedbacks;
  static readonly UpdateFeedback: AdminUpdateFeedback;
  static readonly DeleteFeedback: AdminDeleteFeedback;
  static readonly LookupFeedbacks: AdminLookupFeedbacks;
  static readonly CreateVisitor: AdminCreateVisitor;
  static readonly GetVisitor: AdminGetVisitor;
  static readonly ListVisitors: AdminListVisitors;
  static readonly CountVisitors: AdminCountVisitors;
  static readonly UpdateVisitor: AdminUpdateVisitor;
  static readonly DeleteVisitor: AdminDeleteVisitor;
  static readonly LookupVisitors: AdminLookupVisitors;
  static readonly CreateRating: AdminCreateRating;
  static readonly GetRating: AdminGetRating;
  static readonly ListRatings: AdminListRatings;
  static readonly CountRatings: AdminCountRatings;
  static readonly UpdateRating: AdminUpdateRating;
  static readonly DeleteRating: AdminDeleteRating;
  static readonly LookupRatings: AdminLookupRatings;
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

export class AdminClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createFeedback(
    requestMessage: service_admin_pb.CreateFeedbackRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CreateFeedbackResponse|null) => void
  ): UnaryResponse;
  createFeedback(
    requestMessage: service_admin_pb.CreateFeedbackRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CreateFeedbackResponse|null) => void
  ): UnaryResponse;
  getFeedback(
    requestMessage: service_admin_pb.GetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.GetFeedbackResponse|null) => void
  ): UnaryResponse;
  getFeedback(
    requestMessage: service_admin_pb.GetRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.GetFeedbackResponse|null) => void
  ): UnaryResponse;
  listFeedbacks(
    requestMessage: service_admin_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.ListFeedbacksResponse|null) => void
  ): UnaryResponse;
  listFeedbacks(
    requestMessage: service_admin_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.ListFeedbacksResponse|null) => void
  ): UnaryResponse;
  countFeedbacks(
    requestMessage: service_admin_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CountResponse|null) => void
  ): UnaryResponse;
  countFeedbacks(
    requestMessage: service_admin_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CountResponse|null) => void
  ): UnaryResponse;
  updateFeedback(
    requestMessage: service_admin_pb.UpdateFeedbackRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.UpdateFeedbackResponse|null) => void
  ): UnaryResponse;
  updateFeedback(
    requestMessage: service_admin_pb.UpdateFeedbackRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.UpdateFeedbackResponse|null) => void
  ): UnaryResponse;
  deleteFeedback(
    requestMessage: service_admin_pb.DeleteRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.DeleteResponse|null) => void
  ): UnaryResponse;
  deleteFeedback(
    requestMessage: service_admin_pb.DeleteRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.DeleteResponse|null) => void
  ): UnaryResponse;
  lookupFeedbacks(
    requestMessage: service_admin_pb.LookupRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.LookupResponse|null) => void
  ): UnaryResponse;
  lookupFeedbacks(
    requestMessage: service_admin_pb.LookupRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.LookupResponse|null) => void
  ): UnaryResponse;
  createVisitor(
    requestMessage: service_admin_pb.CreateVisitorRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CreateVisitorResponse|null) => void
  ): UnaryResponse;
  createVisitor(
    requestMessage: service_admin_pb.CreateVisitorRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CreateVisitorResponse|null) => void
  ): UnaryResponse;
  getVisitor(
    requestMessage: service_admin_pb.GetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.GetVisitorResponse|null) => void
  ): UnaryResponse;
  getVisitor(
    requestMessage: service_admin_pb.GetRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.GetVisitorResponse|null) => void
  ): UnaryResponse;
  listVisitors(
    requestMessage: service_admin_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.ListVisitorsResponse|null) => void
  ): UnaryResponse;
  listVisitors(
    requestMessage: service_admin_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.ListVisitorsResponse|null) => void
  ): UnaryResponse;
  countVisitors(
    requestMessage: service_admin_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CountResponse|null) => void
  ): UnaryResponse;
  countVisitors(
    requestMessage: service_admin_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CountResponse|null) => void
  ): UnaryResponse;
  updateVisitor(
    requestMessage: service_admin_pb.UpdateVisitorRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.UpdateVisitorResponse|null) => void
  ): UnaryResponse;
  updateVisitor(
    requestMessage: service_admin_pb.UpdateVisitorRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.UpdateVisitorResponse|null) => void
  ): UnaryResponse;
  deleteVisitor(
    requestMessage: service_admin_pb.DeleteRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.DeleteResponse|null) => void
  ): UnaryResponse;
  deleteVisitor(
    requestMessage: service_admin_pb.DeleteRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.DeleteResponse|null) => void
  ): UnaryResponse;
  lookupVisitors(
    requestMessage: service_admin_pb.LookupRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.LookupResponse|null) => void
  ): UnaryResponse;
  lookupVisitors(
    requestMessage: service_admin_pb.LookupRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.LookupResponse|null) => void
  ): UnaryResponse;
  createRating(
    requestMessage: service_admin_pb.CreateRatingRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CreateRatingResponse|null) => void
  ): UnaryResponse;
  createRating(
    requestMessage: service_admin_pb.CreateRatingRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CreateRatingResponse|null) => void
  ): UnaryResponse;
  getRating(
    requestMessage: service_admin_pb.GetRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.GetRatingResponse|null) => void
  ): UnaryResponse;
  getRating(
    requestMessage: service_admin_pb.GetRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.GetRatingResponse|null) => void
  ): UnaryResponse;
  listRatings(
    requestMessage: service_admin_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.ListRatingsResponse|null) => void
  ): UnaryResponse;
  listRatings(
    requestMessage: service_admin_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.ListRatingsResponse|null) => void
  ): UnaryResponse;
  countRatings(
    requestMessage: service_admin_pb.ListRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CountResponse|null) => void
  ): UnaryResponse;
  countRatings(
    requestMessage: service_admin_pb.ListRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.CountResponse|null) => void
  ): UnaryResponse;
  updateRating(
    requestMessage: service_admin_pb.UpdateRatingRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.UpdateRatingResponse|null) => void
  ): UnaryResponse;
  updateRating(
    requestMessage: service_admin_pb.UpdateRatingRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.UpdateRatingResponse|null) => void
  ): UnaryResponse;
  deleteRating(
    requestMessage: service_admin_pb.DeleteRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.DeleteResponse|null) => void
  ): UnaryResponse;
  deleteRating(
    requestMessage: service_admin_pb.DeleteRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.DeleteResponse|null) => void
  ): UnaryResponse;
  lookupRatings(
    requestMessage: service_admin_pb.LookupRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.LookupResponse|null) => void
  ): UnaryResponse;
  lookupRatings(
    requestMessage: service_admin_pb.LookupRequest,
    callback: (error: ServiceError|null, responseMessage: service_admin_pb.LookupResponse|null) => void
  ): UnaryResponse;
}

