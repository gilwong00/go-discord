import { useMutation } from '@tanstack/react-query';
import { login } from '../../../../gen/proto/v1/user/user-UserService_connectquery';
import { LoginRequest } from '../../../../gen/proto/v1/user/user_pb';
import { useRouter } from 'next/router';

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
