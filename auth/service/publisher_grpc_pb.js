// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var publisher_pb = require('./publisher_pb.js');

function serialize_publisher_GetRequest(arg) {
  if (!(arg instanceof publisher_pb.GetRequest)) {
    throw new Error('Expected argument of type publisher.GetRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_publisher_GetRequest(buffer_arg) {
  return publisher_pb.GetRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_publisher_GetResponse(arg) {
  if (!(arg instanceof publisher_pb.GetResponse)) {
    throw new Error('Expected argument of type publisher.GetResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_publisher_GetResponse(buffer_arg) {
  return publisher_pb.GetResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ServiceService = exports.ServiceService = {
  uploadPicture: {
    path: '/publisher.Service/UploadPicture',
    requestStream: false,
    responseStream: false,
    requestType: publisher_pb.GetRequest,
    responseType: publisher_pb.GetResponse,
    requestSerialize: serialize_publisher_GetRequest,
    requestDeserialize: deserialize_publisher_GetRequest,
    responseSerialize: serialize_publisher_GetResponse,
    responseDeserialize: deserialize_publisher_GetResponse,
  },
};

exports.ServiceClient = grpc.makeGenericClientConstructor(ServiceService);
