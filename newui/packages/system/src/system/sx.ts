import css, { SystemCssProperties } from '@styled-system/css';

import { Prop } from '..';
import { Color } from '../theme';
import { LiteralUnion } from 'type-fest';

type Sx = Exclude<SystemCssProperties, 'color'> & {
  textColor?: Prop<LiteralUnion<Color, string>>;
}

export type SxProp = {
  sx?: Sx;
  theme?: Object;
};

export const sx = ({ sx, theme }: SxProp) => css(sx)(theme);
