// package: admin
// file: service_admin.proto

var service_admin_pb = require("./service_admin_pb");
var grpc = require("grpc-web-client").grpc;

var Admin = (function () {
  function Admin() {}
  Admin.serviceName = "admin.Admin";
  return Admin;
}());

Admin.CreateFeedback = {
  methodName: "CreateFeedback",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.CreateFeedbackRequest,
  responseType: service_admin_pb.CreateFeedbackResponse
};

Admin.GetFeedback = {
  methodName: "GetFeedback",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.GetRequest,
  responseType: service_admin_pb.GetFeedbackResponse
};

Admin.ListFeedbacks = {
  methodName: "ListFeedbacks",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.ListRequest,
  responseType: service_admin_pb.ListFeedbacksResponse
};

Admin.CountFeedbacks = {
  methodName: "CountFeedbacks",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.ListRequest,
  responseType: service_admin_pb.CountResponse
};

Admin.UpdateFeedback = {
  methodName: "UpdateFeedback",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.UpdateFeedbackRequest,
  responseType: service_admin_pb.UpdateFeedbackResponse
};

Admin.DeleteFeedback = {
  methodName: "DeleteFeedback",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.DeleteRequest,
  responseType: service_admin_pb.DeleteResponse
};

Admin.LookupFeedbacks = {
  methodName: "LookupFeedbacks",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.LookupRequest,
  responseType: service_admin_pb.LookupResponse
};

Admin.CreateVisitor = {
  methodName: "CreateVisitor",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.CreateVisitorRequest,
  responseType: service_admin_pb.CreateVisitorResponse
};

Admin.GetVisitor = {
  methodName: "GetVisitor",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.GetRequest,
  responseType: service_admin_pb.GetVisitorResponse
};

Admin.ListVisitors = {
  methodName: "ListVisitors",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.ListRequest,
  responseType: service_admin_pb.ListVisitorsResponse
};

Admin.CountVisitors = {
  methodName: "CountVisitors",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.ListRequest,
  responseType: service_admin_pb.CountResponse
};

Admin.UpdateVisitor = {
  methodName: "UpdateVisitor",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.UpdateVisitorRequest,
  responseType: service_admin_pb.UpdateVisitorResponse
};

Admin.DeleteVisitor = {
  methodName: "DeleteVisitor",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.DeleteRequest,
  responseType: service_admin_pb.DeleteResponse
};

Admin.LookupVisitors = {
  methodName: "LookupVisitors",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.LookupRequest,
  responseType: service_admin_pb.LookupResponse
};

Admin.CreateRating = {
  methodName: "CreateRating",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.CreateRatingRequest,
  responseType: service_admin_pb.CreateRatingResponse
};

Admin.GetRating = {
  methodName: "GetRating",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.GetRequest,
  responseType: service_admin_pb.GetRatingResponse
};

Admin.ListRatings = {
  methodName: "ListRatings",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.ListRequest,
  responseType: service_admin_pb.ListRatingsResponse
};

Admin.CountRatings = {
  methodName: "CountRatings",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.ListRequest,
  responseType: service_admin_pb.CountResponse
};

Admin.UpdateRating = {
  methodName: "UpdateRating",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.UpdateRatingRequest,
  responseType: service_admin_pb.UpdateRatingResponse
};

Admin.DeleteRating = {
  methodName: "DeleteRating",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.DeleteRequest,
  responseType: service_admin_pb.DeleteResponse
};

Admin.LookupRatings = {
  methodName: "LookupRatings",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: service_admin_pb.LookupRequest,
  responseType: service_admin_pb.LookupResponse
};

exports.Admin = Admin;

function AdminClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

AdminClient.prototype.createFeedback = function createFeedback(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.CreateFeedback, {
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

AdminClient.prototype.getFeedback = function getFeedback(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.GetFeedback, {
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

AdminClient.prototype.listFeedbacks = function listFeedbacks(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.ListFeedbacks, {
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

AdminClient.prototype.countFeedbacks = function countFeedbacks(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.CountFeedbacks, {
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

AdminClient.prototype.updateFeedback = function updateFeedback(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.UpdateFeedback, {
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

AdminClient.prototype.deleteFeedback = function deleteFeedback(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.DeleteFeedback, {
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

AdminClient.prototype.lookupFeedbacks = function lookupFeedbacks(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.LookupFeedbacks, {
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

AdminClient.prototype.createVisitor = function createVisitor(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.CreateVisitor, {
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

AdminClient.prototype.getVisitor = function getVisitor(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.GetVisitor, {
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

AdminClient.prototype.listVisitors = function listVisitors(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.ListVisitors, {
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

AdminClient.prototype.countVisitors = function countVisitors(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.CountVisitors, {
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

AdminClient.prototype.updateVisitor = function updateVisitor(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.UpdateVisitor, {
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

AdminClient.prototype.deleteVisitor = function deleteVisitor(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.DeleteVisitor, {
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

AdminClient.prototype.lookupVisitors = function lookupVisitors(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.LookupVisitors, {
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

AdminClient.prototype.createRating = function createRating(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.CreateRating, {
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

AdminClient.prototype.getRating = function getRating(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.GetRating, {
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

AdminClient.prototype.listRatings = function listRatings(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.ListRatings, {
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

AdminClient.prototype.countRatings = function countRatings(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.CountRatings, {
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

AdminClient.prototype.updateRating = function updateRating(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.UpdateRating, {
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

AdminClient.prototype.deleteRating = function deleteRating(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.DeleteRating, {
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

AdminClient.prototype.lookupRatings = function lookupRatings(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.LookupRatings, {
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

exports.AdminClient = AdminClient;

