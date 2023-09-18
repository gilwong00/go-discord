'use client';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { TransportProvider } from '@connectrpc/connect-query';
import { useState } from 'react';
import { createConnectTransport } from '@connectrpc/connect-web';

const transport = createConnectTransport({
  baseUrl: 'localhost:5000'
});

export const Providers = ({ children }: { children: React.ReactNode }) => {
  const [client] = useState(new QueryClient());

  return (
    <TransportProvider transport={transport}>
      <QueryClientProvider client={client}>{children}</QueryClientProvider>
    </TransportProvider>
  );
};
