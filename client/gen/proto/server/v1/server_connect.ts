// @generated by protoc-gen-connect-es v0.13.2 with parameter "target=ts"
// @generated from file proto/server/v1/server.proto (package server.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateServerRequest, CreateServerResponse } from "./server_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service server.v1.ServerService
 */
export const ServerService = {
  typeName: "server.v1.ServerService",
  methods: {
    /**
     * @generated from rpc server.v1.ServerService.CreateServer
     */
    createServer: {
      name: "CreateServer",
      I: CreateServerRequest,
      O: CreateServerResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

