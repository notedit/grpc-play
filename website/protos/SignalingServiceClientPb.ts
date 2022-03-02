/**
 * @fileoverview gRPC-Web generated client stub for signaling
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as signaling_pb from './signaling_pb';


export class SignalingClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorJoin = new grpcWeb.MethodDescriptor(
    '/signaling.Signaling/Join',
    grpcWeb.MethodType.UNARY,
    signaling_pb.JoinRequest,
    signaling_pb.JoinResponse,
    (request: signaling_pb.JoinRequest) => {
      return request.serializeBinary();
    },
    signaling_pb.JoinResponse.deserializeBinary
  );

  join(
    request: signaling_pb.JoinRequest,
    metadata: grpcWeb.Metadata | null): Promise<signaling_pb.JoinResponse>;

  join(
    request: signaling_pb.JoinRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: signaling_pb.JoinResponse) => void): grpcWeb.ClientReadableStream<signaling_pb.JoinResponse>;

  join(
    request: signaling_pb.JoinRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: signaling_pb.JoinResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/signaling.Signaling/Join',
        request,
        metadata || {},
        this.methodDescriptorJoin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/signaling.Signaling/Join',
    request,
    metadata || {},
    this.methodDescriptorJoin);
  }

  methodDescriptorLeave = new grpcWeb.MethodDescriptor(
    '/signaling.Signaling/Leave',
    grpcWeb.MethodType.UNARY,
    signaling_pb.LeaveRequest,
    signaling_pb.LeaveResponse,
    (request: signaling_pb.LeaveRequest) => {
      return request.serializeBinary();
    },
    signaling_pb.LeaveResponse.deserializeBinary
  );

  leave(
    request: signaling_pb.LeaveRequest,
    metadata: grpcWeb.Metadata | null): Promise<signaling_pb.LeaveResponse>;

  leave(
    request: signaling_pb.LeaveRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: signaling_pb.LeaveResponse) => void): grpcWeb.ClientReadableStream<signaling_pb.LeaveResponse>;

  leave(
    request: signaling_pb.LeaveRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: signaling_pb.LeaveResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/signaling.Signaling/Leave',
        request,
        metadata || {},
        this.methodDescriptorLeave,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/signaling.Signaling/Leave',
    request,
    metadata || {},
    this.methodDescriptorLeave);
  }

}

