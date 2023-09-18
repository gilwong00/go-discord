'use client';

import { FormInput } from '@/components/FormInput';
import { theme } from '@/theme/theme';
import React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import styled from 'styled-components';

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

  const handleFormSubmit: SubmitHandler<LoginForm> = async data => {};

  return (
    <Container>
      <HeaderContainer>
        <h2>Welcome back!</h2>
        <p>We&apos;re so excited to see you again!</p>
      </HeaderContainer>
      <FormContainer>
        <form onSubmit={handleSubmit(handleFormSubmit)}>
          <FormInput
            {...register('username', { required: true })}
            label='Username'
            error={errors.username}
            type='text'
          />
          <FormInput
            {...register('password', { required: true })}
            label='Password'
            error={errors.password}
            type='text'
          />
          <button type='submit'>Login</button>
        </form>
      </FormContainer>
    </Container>
  );
};

export default LoginPage;
