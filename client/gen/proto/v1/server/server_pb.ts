// @generated by protoc-gen-es v1.3.3 with parameter "target=ts"
// @generated from file proto/v1/server/server.proto (package proto.server.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message proto.server.v1.Server
 */
export class Server extends Message<Server> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string logo_url = 3;
   */
  logoUrl = "";

  /**
   * @generated from field: string invite_code = 4;
   */
  inviteCode = "";

  constructor(data?: PartialMessage<Server>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.server.v1.Server";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "logo_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "invite_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Server {
    return new Server().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Server {
    return new Server().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Server {
    return new Server().fromJsonString(jsonString, options);
  }

  static equals(a: Server | PlainMessage<Server> | undefined, b: Server | PlainMessage<Server> | undefined): boolean {
    return proto3.util.equals(Server, a, b);
  }
}

/**
 * @generated from message proto.server.v1.CreateServerRequest
 */
export class CreateServerRequest extends Message<CreateServerRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: optional string logo_url = 2;
   */
  logoUrl?: string;

  constructor(data?: PartialMessage<CreateServerRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.server.v1.CreateServerRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "logo_url", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateServerRequest {
    return new CreateServerRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateServerRequest {
    return new CreateServerRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateServerRequest {
    return new CreateServerRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateServerRequest | PlainMessage<CreateServerRequest> | undefined, b: CreateServerRequest | PlainMessage<CreateServerRequest> | undefined): boolean {
    return proto3.util.equals(CreateServerRequest, a, b);
  }
}

/**
 * @generated from message proto.server.v1.CreateServerResponse
 */
export class CreateServerResponse extends Message<CreateServerResponse> {
  /**
   * @generated from field: proto.server.v1.Server server = 1;
   */
  server?: Server;

  constructor(data?: PartialMessage<CreateServerResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto.server.v1.CreateServerResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "server", kind: "message", T: Server },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateServerResponse {
    return new CreateServerResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateServerResponse {
    return new CreateServerResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateServerResponse {
    return new CreateServerResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateServerResponse | PlainMessage<CreateServerResponse> | undefined, b: CreateServerResponse | PlainMessage<CreateServerResponse> | undefined): boolean {
    return proto3.util.equals(CreateServerResponse, a, b);
  }
}

