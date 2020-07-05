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
} from '../../system';
import { compose } from 'styled-system';
import { createShouldForwardProp } from '@styled-system/should-forward-prop';

type Props = PaddingProps &
  MarginProps &
  TextColorProps &
  BackgroundColorProps &
  SxProp;

const shouldForwardProp = createShouldForwardProp([
  ...padding.propNames,
  ...margin.propNames,
  ...textColor.propNames,
  ...backgroundColor.propNames,
]);

const composed = compose(padding, margin, textColor, backgroundColor);

const Box = styled('div', {
  shouldForwardProp,
})<Props>({}, composed, sx);

export default Box;
