import { system, Config } from 'styled-system';
import { Prop } from '../';
import { Space } from '../theme';

export type PaddingProps = {
  p?: Prop<Space>;
  padding?: Prop<Space>;
  pt?: Prop<Space>;
  paddingTop?: Prop<Space>;
  pr?: Prop<Space>;
  paddingRight?: Prop<Space>;
  pb?: Prop<Space>;
  paddingBottom?: Prop<Space>;
  pl?: Prop<Space>;
  paddingLeft?: Prop<Space>;
  px?: Prop<Space>;
  paddingX?: Prop<Space>;
  py?: Prop<Space>;
  paddingY?: Prop<Space>;
};

const config: Config = {
  padding: {
    property: 'padding',
    scale: 'space',
  },
  paddingTop: {
    property: 'paddingTop',
    scale: 'space',
  },
  paddingBottom: {
    property: 'paddingBottom',
    scale: 'space',
  },
  paddingRight: {
    property: 'paddingRight',
    scale: 'space',
  },
  paddingLeft: {
    property: 'paddingLeft',
    scale: 'space',
  },
  paddingX: {
    properties: ['paddingLeft', 'paddingRight'],
    scale: 'space',
  },
  paddingY: {
    properties: ['paddingTop', 'paddingBottom'],
    scale: 'space',
  },
};

config.p = config.padding;
config.pt = config.paddingTop;
config.pr = config.paddingRight;
config.pb = config.paddingBottom;
config.pl = config.paddingLeft;
config.px = config.paddingX;
config.py = config.paddingY;

export const padding = system(config);
