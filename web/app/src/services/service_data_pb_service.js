// package: data
// file: service_data.proto

var service_data_pb = require("./service_data_pb");
var grpc = require("grpc-web-client").grpc;

var Data = (function () {
  function Data() {}
  Data.serviceName = "data.Data";
  return Data;
}());

Data.GetFeedbackByID = {
  methodName: "GetFeedbackByID",
  service: Data,
  requestStream: false,
  responseStream: false,
  requestType: service_data_pb.GetByIDRequest,
  responseType: service_data_pb.FeedbackResponse
};

Data.ListFeedbacks = {
  methodName: "ListFeedbacks",
  service: Data,
  requestStream: false,
  responseStream: false,
  requestType: service_data_pb.ListRequest,
  responseType: service_data_pb.ListFeedbacksResponse
};

Data.GetVisitorByID = {
  methodName: "GetVisitorByID",
  service: Data,
  requestStream: false,
  responseStream: false,
  requestType: service_data_pb.GetByIDRequest,
  responseType: service_data_pb.VisitorResponse
};

Data.ListVisitors = {
  methodName: "ListVisitors",
  service: Data,
  requestStream: false,
  responseStream: false,
  requestType: service_data_pb.ListRequest,
  responseType: service_data_pb.ListVisitorsResponse
};

Data.GetRatingByID = {
  methodName: "GetRatingByID",
  service: Data,
  requestStream: false,
  responseStream: false,
  requestType: service_data_pb.GetByIDRequest,
  responseType: service_data_pb.RatingResponse
};

Data.GetModifiedRatingByID = {
  methodName: "GetModifiedRatingByID",
  service: Data,
  requestStream: false,
  responseStream: false,
  requestType: service_data_pb.GetByIDRequest,
  responseType: service_data_pb.TimestampResponse
};

Data.ListRatings = {
  methodName: "ListRatings",
  service: Data,
  requestStream: false,
  responseStream: false,
  requestType: service_data_pb.ListRequest,
  responseType: service_data_pb.ListRatingsResponse
};

exports.Data = Data;

function DataClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

DataClient.prototype.getFeedbackByID = function getFeedbackByID(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Data.GetFeedbackByID, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

DataClient.prototype.listFeedbacks = function listFeedbacks(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Data.ListFeedbacks, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

DataClient.prototype.getVisitorByID = function getVisitorByID(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Data.GetVisitorByID, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

DataClient.prototype.listVisitors = function listVisitors(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Data.ListVisitors, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

DataClient.prototype.getRatingByID = function getRatingByID(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Data.GetRatingByID, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

DataClient.prototype.getModifiedRatingByID = function getModifiedRatingByID(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Data.GetModifiedRatingByID, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

DataClient.prototype.listRatings = function listRatings(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Data.ListRatings, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.DataClient = DataClient;

