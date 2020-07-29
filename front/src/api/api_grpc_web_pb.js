/**
 * @fileoverview gRPC-Web generated client stub for api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.api = require('./api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.api.PingClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.api.PingPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.PingMessage,
 *   !proto.api.PingMessage>}
 */
const methodDescriptor_Ping_SayHello = new grpc.web.MethodDescriptor(
  '/api.Ping/SayHello',
  grpc.web.MethodType.UNARY,
  proto.api.PingMessage,
  proto.api.PingMessage,
  /**
   * @param {!proto.api.PingMessage} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.PingMessage.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.PingMessage,
 *   !proto.api.PingMessage>}
 */
const methodInfo_Ping_SayHello = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.PingMessage,
  /**
   * @param {!proto.api.PingMessage} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.PingMessage.deserializeBinary
);


/**
 * @param {!proto.api.PingMessage} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.PingMessage)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.PingMessage>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.PingClient.prototype.sayHello =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.Ping/SayHello',
      request,
      metadata || {},
      methodDescriptor_Ping_SayHello,
      callback);
};


/**
 * @param {!proto.api.PingMessage} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.PingMessage>}
 *     A native promise that resolves to the response
 */
proto.api.PingPromiseClient.prototype.sayHello =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.Ping/SayHello',
      request,
      metadata || {},
      methodDescriptor_Ping_SayHello);
};


module.exports = proto.api;

