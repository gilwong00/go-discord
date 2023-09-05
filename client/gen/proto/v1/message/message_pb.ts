// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file proto/v1/message/message.proto (package proto.message.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message as Message$1, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message proto.message.v1.Message
 */
export class Message extends Message$1<Message> {
  /**
   * @generated from field: string content = 1;
   */
  content = "";

  /**
   * @generated from field: optional string file_url = 2;
   */
  fileUrl?: string;

  constructor(data?: PartialMessage<Message>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.message.v1.Message";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "content", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "file_url", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Message {
    return new Message().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Message {
    return new Message().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Message {
    return new Message().fromJsonString(jsonString, options);
  }

  static equals(a: Message | PlainMessage<Message> | undefined, b: Message | PlainMessage<Message> | undefined): boolean {
    return proto3.util.equals(Message, a, b);
  }
}

/**
 * @generated from message proto.message.v1.SendDirectMessageRequest
 */
export class SendDirectMessageRequest extends Message$1<SendDirectMessageRequest> {
  /**
   * @generated from field: proto.message.v1.Message message = 1;
   */
  message?: Message;

  /**
   * @generated from field: string from_member_id = 2;
   */
  fromMemberId = "";

  /**
   * @generated from field: string to_member_id = 3;
   */
  toMemberId = "";

  constructor(data?: PartialMessage<SendDirectMessageRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.message.v1.SendDirectMessageRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "message", T: Message },
    { no: 2, name: "from_member_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "to_member_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendDirectMessageRequest {
    return new SendDirectMessageRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendDirectMessageRequest {
    return new SendDirectMessageRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendDirectMessageRequest {
    return new SendDirectMessageRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SendDirectMessageRequest | PlainMessage<SendDirectMessageRequest> | undefined, b: SendDirectMessageRequest | PlainMessage<SendDirectMessageRequest> | undefined): boolean {
    return proto3.util.equals(SendDirectMessageRequest, a, b);
  }
}

/**
 * @generated from message proto.message.v1.SendDirectMessageResponse
 */
export class SendDirectMessageResponse extends Message$1<SendDirectMessageResponse> {
  /**
   * @generated from field: proto.message.v1.Message message = 1;
   */
  message?: Message;

  /**
   * @generated from field: google.protobuf.Timestamp sent_at = 2;
   */
  sentAt?: Timestamp;

  constructor(data?: PartialMessage<SendDirectMessageResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.message.v1.SendDirectMessageResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "message", T: Message },
    { no: 2, name: "sent_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendDirectMessageResponse {
    return new SendDirectMessageResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendDirectMessageResponse {
    return new SendDirectMessageResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendDirectMessageResponse {
    return new SendDirectMessageResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SendDirectMessageResponse | PlainMessage<SendDirectMessageResponse> | undefined, b: SendDirectMessageResponse | PlainMessage<SendDirectMessageResponse> | undefined): boolean {
    return proto3.util.equals(SendDirectMessageResponse, a, b);
  }
}

