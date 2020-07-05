import React, { useState } from 'react';
import { Box, Button, Table, theme, TextField } from '@sloth/system';
import { ThemeProvider } from 'emotion-theming';
import { Provider } from 'urql';
import { ActivityTable } from './ActivityTable';
import { client } from './client';

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
            color: 'red',
          }}
        >
          <ActivityTable />
        </Box>
      </ThemeProvider>
    </Provider>
  );
}

export default App;
