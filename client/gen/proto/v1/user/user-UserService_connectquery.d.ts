// @generated by protoc-gen-connect-query v0.5.3 with parameter "target=ts+dts+js,import_extension=none,ts_nocheck=false"
// @generated from file proto/v1/user/user.proto (package user.v1, syntax proto3)
/* eslint-disable */

import { CreateUserRequest, CreateUserResponse, LoginRequest, LoginResponse } from "./user_pb";
import { MethodKind } from "@bufbuild/protobuf";
import { UnaryFunctionsWithHooks } from "@connectrpc/connect-query";

/**
 * @generated from service user.v1.UserService
 */
export declare const UserService: {
  readonly typeName: "user.v1.UserService",
  readonly methods: {
    /**
     * @generated from rpc user.v1.UserService.CreateUser
     */
    readonly createUser: {
      readonly name: "CreateUser",
      readonly I: typeof CreateUserRequest,
      readonly O: typeof CreateUserResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc user.v1.UserService.Login
     */
    readonly login: {
      readonly name: "Login",
      readonly I: typeof LoginRequest,
      readonly O: typeof LoginResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

export const createUser: UnaryFunctionsWithHooks<CreateUserRequest, CreateUserResponse>;
export const login: UnaryFunctionsWithHooks<LoginRequest, LoginResponse>;
