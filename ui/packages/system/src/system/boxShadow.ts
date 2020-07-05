import { system, Config } from 'styled-system';
import { Prop } from '..';
import { BoxShadow } from '../theme';

export type BoxShadowProps = {
  boxShadow?: Prop<BoxShadow>;
};

const config: Config = {
  boxShadow: {
    property: 'boxShadow',
    scale: 'shadows',
  },
};

export const boxShadow = system(config);
