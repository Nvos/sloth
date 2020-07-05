import babel from "@rollup/plugin-babel";
import * as reactPlugin from "vite-plugin-react";
import type { UserConfig } from "vite";

const config: UserConfig = {
  jsx: "react",
  plugins: [reactPlugin],
  optimizeDeps: {
    link: ["@sloth/system"],
  },
  rollupInputOptions: {
    plugins: [babel({ babelHelpers: "bundled", })],
  },
};

export default config;
