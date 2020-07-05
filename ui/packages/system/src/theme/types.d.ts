import CSS from 'csstype';
import { LiteralUnion } from 'type-fest';

export type FontSize = LiteralUnion<
  0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8,
  CSS.FontSizeProperty<number>
>;

export type Space = LiteralUnion<
  0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12,
  CSS.PaddingProperty<number>
>;

export type Color =
  | 'primary.100'
  | 'primary.200'
  | 'primary.300'
  | 'primary.400'
  | 'primary.500'
  | 'primary.600'
  | 'primary.700'
  | 'primary.800'
  | 'primary.900'
  | 'neutral.100'
  | 'neutral.200'
  | 'neutral.300'
  | 'neutral.400'
  | 'neutral.500'
  | 'neutral.600'
  | 'neutral.700'
  | 'neutral.800'
  | 'neutral.900'
  | 'info.100'
  | 'info.200'
  | 'info.300'
  | 'info.400'
  | 'info.500'
  | 'info.600'
  | 'info.700'
  | 'info.800'
  | 'info.900'
  | 'warn.100'
  | 'warn.200'
  | 'warn.300'
  | 'warn.400'
  | 'warn.500'
  | 'warn.600'
  | 'warn.700'
  | 'warn.800'
  | 'warn.900'
  | 'danger.100'
  | 'danger.200'
  | 'danger.300'
  | 'danger.400'
  | 'danger.500'
  | 'danger.600'
  | 'danger.700'
  | 'danger.800'
  | 'danger.900'
  | 'success.100'
  | 'success.200'
  | 'success.300'
  | 'success.400'
  | 'success.500'
  | 'success.600'
  | 'success.700'
  | 'success.800'
  | 'success.900';

export type BoxShadow =
  | 'elevation.0'
  | 'elevation.1'
  | 'elevation.2'
  | 'elevation.3';
