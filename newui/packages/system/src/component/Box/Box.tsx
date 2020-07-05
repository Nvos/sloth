import styled from '@emotion/styled';
import {
  padding,
  margin,
  textColor,
  backgroundColor,
  PaddingProps,
  MarginProps,
  TextColorProps,
  BackgroundColorProps,
  SxProp,
  sx,
  base,
  BaseProp,
} from '../../system';
import { compose } from 'styled-system';
import { createShouldForwardProp } from '@styled-system/should-forward-prop';

export type BoxProps = PaddingProps &
  MarginProps &
  TextColorProps &
  BackgroundColorProps &
  SxProp &
  BaseProp & { as?: React.ElementType };

const shouldForwardProp = createShouldForwardProp([
  ...(padding.propNames ?? []),
  ...(margin.propNames ?? []),
  ...(textColor.propNames ?? []),
  ...(backgroundColor.propNames ?? []),
]);

const composed = compose(padding, margin, textColor, backgroundColor);

const Box = styled('div', {
  shouldForwardProp,
})<BoxProps>(
  {
    boxSizing: 'border-box',
    margin: 0,
    minWidth: 0,
  },
  composed,
  base,
  sx
);

export default Box;
