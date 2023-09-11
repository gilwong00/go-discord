'use client';

import styled from 'styled-components';
import { theme } from '@/theme/theme';

const Layout = styled.div`
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: ${theme.colors.primary};
`;

const AuthLayout = ({ children }: { children: React.ReactNode }) => {
  return <Layout>{children}</Layout>;
};

export default AuthLayout;
