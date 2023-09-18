// @generated by protoc-gen-connect-query v0.4.4 with parameter "target=ts"
// @generated from file proto/v1/member/member.proto (package proto.member.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateServerMemberRequest, CreateServerMemberResponse } from "./member_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { createQueryService } from "@connectrpc/connect-query";

export const typeName = "proto.member.v1.MemberService";

/**
 * @generated from service proto.member.v1.MemberService
 */
export const MemberService = {
  typeName: "proto.member.v1.MemberService",
  methods: {
    /**
     * @generated from rpc proto.member.v1.MemberService.CreateServerMember
     */
    createServerMember: {
      name: "CreateServerMember",
      I: CreateServerMemberRequest,
      O: CreateServerMemberResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * @generated from rpc proto.member.v1.MemberService.CreateServerMember
 */
export const createServerMember = createQueryService({
  service: MemberService,
}).createServerMember;
