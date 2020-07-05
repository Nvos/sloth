import { system, Config } from 'styled-system';
import { Prop } from '../';
import { Space } from '../theme';

export type BackgroundColorProps = {
  backgroundColor?: Prop<Space>;
  bg?: Prop<Space>;
};

const config: Config = {
  backgroundColor: {
    property: 'backgroundColor',
    scale: 'colors',
  },
};

config.bg = config.backgroundColor;

export const backgroundColor = system(config);
