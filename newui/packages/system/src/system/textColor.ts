import { system, Config } from 'styled-system';
import { Prop } from '../';
import { Color } from '../theme';

export type TextColorProps = {
  textColor?: Prop<Color>;
};

export const config: Config = {
  textColor: {
    property: 'color',
    scale: 'colors',
  },
};

export const textColor = system(config);
