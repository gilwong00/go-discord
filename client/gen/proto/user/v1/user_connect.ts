// @generated by protoc-gen-connect-es v0.13.2 with parameter "target=ts"
// @generated from file proto/user/v1/user.proto (package user.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateUserRequest, CreateUserResponse } from "./user_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service user.v1.UserService
 */
export const UserService = {
  typeName: "user.v1.UserService",
  methods: {
    /**
     * @generated from rpc user.v1.UserService.CreateUser
     */
    createUser: {
      name: "CreateUser",
      I: CreateUserRequest,
      O: CreateUserResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
