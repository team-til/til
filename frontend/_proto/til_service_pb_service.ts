// package: til
// file: til_service.proto

import * as til_service_pb from "./til_service_pb";
export class TilService {
  static serviceName = "til.TilService";
}
export namespace TilService {
  export class Ping {
    static methodName = "Ping";
    static service = TilService;
    static requestStream = false;
    static responseStream = false;
    static requestType = til_service_pb.PingRequest;
    static responseType = til_service_pb.PingResponse;
  }
}
