import * as React from 'react';

type Props = {
  variant?: 'primary' | 'secondary';
} & React.HTMLProps<HTMLButtonElement>;

const Button: React.FC<Props> = ({ children }) => {
  return <button>{children}</button>;
};

export default Button;
