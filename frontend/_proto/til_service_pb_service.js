// package: til
// file: til_service.proto

var til_service_pb = require("./til_service_pb");
var grpc = require("grpc-web-client").grpc;

var TilService = (function () {
  function TilService() {}
  TilService.serviceName = "til.TilService";
  return TilService;
}());

TilService.Ping = {
  methodName: "Ping",
  service: TilService,
  requestStream: false,
  responseStream: false,
  requestType: til_service_pb.PingRequest,
  responseType: til_service_pb.PingResponse
};

exports.TilService = TilService;

function TilServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

TilServiceClient.prototype.ping = function ping(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(TilService.Ping, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

exports.TilServiceClient = TilServiceClient;

