// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file proto/v1/message/message.proto (package proto.message.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { SendDirectMessageRequest, SendDirectMessageResponse } from "./message_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service proto.message.v1.ServerService
 */
export const ServerService = {
  typeName: "proto.message.v1.ServerService",
  methods: {
    /**
     * @generated from rpc proto.message.v1.ServerService.SendDirectMessage
     */
    sendDirectMessage: {
      name: "SendDirectMessage",
      I: SendDirectMessageRequest,
      O: SendDirectMessageResponse,
      kind: MethodKind.ServerStreaming,
    },
  }
} as const;

