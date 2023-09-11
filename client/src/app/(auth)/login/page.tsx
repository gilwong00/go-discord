'use client';

import { theme } from '@/theme/theme';
import React from 'react';
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

const LoginPage = () => {
  return (
    <Container>
      <HeaderContainer>
        <h2>Welcome back!</h2>
        <p>We&apos;re so excited to see you again!</p>
      </HeaderContainer>
      <FormContainer></FormContainer>
    </Container>
  );
};

export default LoginPage;
