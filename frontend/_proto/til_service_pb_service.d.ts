// package: til
// file: til_service.proto

import * as til_service_pb from "./til_service_pb";
import {grpc} from "grpc-web-client";

type TilServicePing = {
  readonly methodName: string;
  readonly service: typeof TilService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof til_service_pb.PingRequest;
  readonly responseType: typeof til_service_pb.PingResponse;
};

export class TilService {
  static readonly serviceName: string;
  static readonly Ping: TilServicePing;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor }

interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}

export class TilServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  ping(
    requestMessage: til_service_pb.PingRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: til_service_pb.PingResponse|null) => void
  ): void;
  ping(
    requestMessage: til_service_pb.PingRequest,
    callback: (error: ServiceError, responseMessage: til_service_pb.PingResponse|null) => void
  ): void;
}

