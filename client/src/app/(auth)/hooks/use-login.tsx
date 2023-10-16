'use client';

import { useMutation } from '@tanstack/react-query';
import { login } from '../../../../gen/proto/v1/user/user-UserService_connectquery';
import { useRouter } from 'next/navigation';
import { LoginRequest } from '../../../../gen/proto/v1/user/user_pb';

export const useLogin = () => {
  const router = useRouter();
  const baseMutation = login.useMutation();

  const loginUseryMutation = useMutation({
    ...baseMutation,
    mutationFn: (
      payload: Pick<LoginRequest, 'usernameOrEmail' | 'password'>
    ) => {
      return baseMutation.mutationFn(payload);
    },
    onSuccess: () => {
      router.push('/home');
    }
  });

  return {
    loginUseryMutation
  };
};
