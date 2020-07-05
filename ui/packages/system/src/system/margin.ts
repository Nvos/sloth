import { system, Config } from 'styled-system';
import { Prop } from '../';
import { Space } from '../theme';

export type MarginProps = {
  m?: Prop<Space>;
  margin?: Prop<Space>;
  mt?: Prop<Space>;
  marginTop?: Prop<Space>;
  mr?: Prop<Space>;
  marginRight?: Prop<Space>;
  mb?: Prop<Space>;
  marginBottom?: Prop<Space>;
  ml?: Prop<Space>;
  marginLeft?: Prop<Space>;
  mv?: Prop<Space>;
  marginVertical?: Prop<Space>;
  mh?: Prop<Space>;
  marginHorizontal?: Prop<Space>;
};

const config: Config = {
  margin: {
    property: 'margin',
    scale: 'space',
  },
  marginTop: {
    property: 'marginTop',
    scale: 'space',
  },
  marginBottom: {
    property: 'marginBottom',
    scale: 'space',
  },
  marginRight: {
    property: 'marginRight',
    scale: 'space',
  },
  marginLeft: {
    property: 'marginLeft',
    scale: 'space',
  },
  marginVertical: {
    properties: ['marginLeft', 'marginRight'],
    scale: 'space',
  },
  marginHorizontal: {
    properties: ['marginTop', 'marginBottom'],
    scale: 'space',
  },
};

config.m = config.margin;
config.mt = config.marginTop;
config.mr = config.marginRight;
config.mb = config.marginBottom;
config.ml = config.marginLeft;
config.mv = config.marginVertical;
config.mh = config.marginHorizontal;

export const margin = system(config);
