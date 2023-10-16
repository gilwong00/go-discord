// @generated by protoc-gen-connect-query v0.5.3 with parameter "target=ts"
// @generated from file proto/v1/user/user.proto (package user.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import {
  CreateUserRequest,
  CreateUserResponse,
  LoginRequest,
  LoginResponse
} from './user_pb.ts';
import { MethodKind } from '@bufbuild/protobuf';
import {
  createQueryService,
  createUnaryHooks,
  UnaryFunctionsWithHooks
} from '@connectrpc/connect-query';

export const typeName = 'user.v1.UserService';

/**
 * @generated from service user.v1.UserService
 */
export const UserService = {
  typeName: 'user.v1.UserService',
  methods: {
    /**
     * @generated from rpc user.v1.UserService.CreateUser
     */
    createUser: {
      name: 'CreateUser',
      I: CreateUserRequest,
      O: CreateUserResponse,
      kind: MethodKind.Unary
    },
    /**
     * @generated from rpc user.v1.UserService.Login
     */
    login: {
      name: 'Login',
      I: LoginRequest,
      O: LoginResponse,
      kind: MethodKind.Unary
    }
  }
} as const;

const $queryService = createQueryService({ service: UserService });

/**
 * @generated from rpc user.v1.UserService.CreateUser
 */
export const createUser: UnaryFunctionsWithHooks<
  CreateUserRequest,
  CreateUserResponse
> = {
  ...$queryService.createUser,
  ...createUnaryHooks($queryService.createUser)
};

/**
 * @generated from rpc user.v1.UserService.Login
 */
export const login: UnaryFunctionsWithHooks<LoginRequest, LoginResponse> = {
  ...$queryService.login,
  ...createUnaryHooks($queryService.login)
};
