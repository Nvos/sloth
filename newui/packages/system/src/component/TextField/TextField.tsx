import React, { FC } from 'react';
import styled from '@emotion/styled';
import css from '@styled-system/css';

const Label = styled('label')`
  display: block;
`;

const Field = styled('input')((props) =>
  css({
    height: 40,
    borderRadius: 1,
    display: 'block',
    borderWidth: 1,
    borderColor: 'primary.500',
    boxSizing: 'border-box',
  })(props.theme)
);

const Root = styled('div')``;

type Props = React.HTMLProps<HTMLInputElement>;

const TextField: FC<Props> = ({ name, id, ...props }) => {
  return (
    <Root>
      <Label htmlFor={id}>{name}</Label>
      <Field id={id} name={name} {...props} />
    </Root>
  );
};

export default TextField;
