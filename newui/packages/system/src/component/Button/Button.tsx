import React, { forwardRef, FC } from 'react';
import Box, { BoxProps } from '../Box/Box';

type Props = {
  variant?: 'primary' | 'secondary';
} & React.HTMLProps<HTMLButtonElement> &
  BoxProps;

export type Ref = HTMLButtonElement;

const Button = forwardRef<Ref, Props>((props, ref) => {
  return (
    <Box
      ref={ref as any}
      as="button"
      {...(props as any)}
      __base={{
        appearance: 'none',
        textDecoration: 'none',
        textAlign: 'center',
        border: 'none',
        background: 'none',
        '&:hover': {
          cursor: 'pointer',
        },
      }}
    />
  );
});

export default Button;
