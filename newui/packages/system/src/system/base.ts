import css, {
  SystemCssProperties,
  SystemStyleObject,
} from '@styled-system/css';
import { LiteralUnion } from 'type-fest';
import { Prop } from '..';
import { Color } from '../theme';

type Base = SystemStyleObject;

export type BaseProp = {
  __base?: Base;
  theme?: Object;
};

export const base = ({ __base, theme }: BaseProp) => css(__base)(theme);
