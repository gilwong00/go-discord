// @generated by protoc-gen-connect-query v0.4.4 with parameter "target=ts"
// @generated from file proto/v1/user/user.proto (package user.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateUserRequest, CreateUserResponse, LoginRequest, LoginResponse } from "./user_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { createQueryService } from "@connectrpc/connect-query";

export const typeName = "user.v1.UserService";

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
    /**
     * @generated from rpc user.v1.UserService.Login
     */
    login: {
      name: "Login",
      I: LoginRequest,
      O: LoginResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * @generated from rpc user.v1.UserService.CreateUser
 */
export const createUser = createQueryService({
  service: UserService,
}).createUser;

/**
 * @generated from rpc user.v1.UserService.Login
 */
export const login = createQueryService({
  service: UserService,
}).login;