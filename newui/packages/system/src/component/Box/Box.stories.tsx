import React from 'react';
import { ThemeProvider } from 'emotion-theming';
import { action } from '@storybook/addon-actions';
import Box from './Box';

export default {
  title: 'Box',
  component: Box,
};

export const SxProp = () => (
  <ThemeProvider
    theme={{
      space: [0, 2, 4, 8, 16, 32, 64, 128],
    }}
  >
    <Box sx={{ p: 4, m: 4, color: 'primary.100',  }} onClick={action('clicked')}>
      Sx prop styling
    </Box>
  </ThemeProvider>
);
