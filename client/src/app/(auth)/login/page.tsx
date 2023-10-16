'use client';

import { FormInput } from '@/components/FormInput';
import { theme } from '@/theme/theme';
import React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import styled from 'styled-components';
import Link from 'next/link';
import { useLogin } from '../hooks/use-login';

const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 500px;
  width: 800px;
  background-color: ${theme.colors.blue01};
  border-radius: ${theme.borderRadius.md}px;
  padding: 32px;
  justify-content: space-between;
`;

const HeaderContainer = styled.div`
  text-align: center;

  p {
    margin-top: 10px;
    color: ${theme.colors.gray01};
    font-size: 16px;
  }
`;

const FormContainer = styled.div`
  height: 350px;
  width: 80%;
`;

const Form = styled.form`
  height: 90%;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
`;

const HyperLink = styled.span`
  color: ${theme.colors.blue02};
  cursor: pointer;
  font-size: 14px;
  margin-top: 5px;
  font-weight: 500;
  letter-spacing: normal;
`;

const PageLink = styled(Link)`
  color: ${theme.colors.blue02};
  cursor: pointer;
  font-size: 14px;
  margin-top: 5px;
  font-weight: 500;
  letter-spacing: normal;
`;

const LoginButtonContainer = styled.div`
  width: 100%;
`;

const LoginButton = styled.button`
  height: 44px;
  width: 100%;
  border: none;
  border-radius: 4px;
  background-color: ${theme.colors.primary};
  cursor: pointer;

  &:hover {
    opacity: 50%;
  }
`;

const RegisterContainer = styled.div`
  margin-top: 5px;
`;

const NeedAccount = styled.span`
  color: ${theme.colors.gray02};
  font-size: 14px;
`;

interface LoginForm {
  username: string;
  password: string;
}

const LoginPage = () => {
  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm<LoginForm>();
  const { loginUseryMutation } = useLogin();

  const handleFormSubmit: SubmitHandler<LoginForm> = async data => {
    void loginUseryMutation.mutateAsync({
      usernameOrEmail: data.username,
      password: data.password
    });
  };

  return (
    <Container>
      <HeaderContainer>
        <h2>Welcome back!</h2>
        <p>We&apos;re so excited to see you again!</p>
      </HeaderContainer>
      <FormContainer>
        <Form onSubmit={handleSubmit(handleFormSubmit)}>
          <FormInput
            {...register('username', { required: 'Username is required' })}
            label='Username'
            isRequired
            backgroundColor={theme.colors.black02}
            error={errors.username?.message ?? ''}
            type='text'
          />
          <div>
            <FormInput
              {...register('password', { required: 'Password is required' })}
              label='Password'
              isRequired
              backgroundColor={theme.colors.black02}
              error={errors.password?.message ?? ''}
              type='text'
            />
            <HyperLink>Forgot your password?</HyperLink>
          </div>
          <LoginButtonContainer>
            <LoginButton type='submit'>Login</LoginButton>
            <RegisterContainer>
              <NeedAccount>Need an account?</NeedAccount>{' '}
              <PageLink href='/signup'>Register</PageLink>
            </RegisterContainer>
          </LoginButtonContainer>
        </Form>
      </FormContainer>
    </Container>
  );
};

export default LoginPage;
