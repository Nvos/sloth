import React, { useState } from 'react';
import { theme } from '@sloth/system';
import { ThemeProvider } from 'emotion-theming';
import { Provider } from 'urql';
import { ActivityTable } from './ActivityTable';
import { client } from './client';
import { Box } from '@chakra-ui/core';

function App() {
  return (
    <Provider value={client}>
      <ThemeProvider theme={theme}>
        <Box
          sx={{
            height: '100vh',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
          }}
        >
          <ActivityTable />
        </Box>
      </ThemeProvider>
    </Provider>
  );
}

export default App;
