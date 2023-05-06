/**
 * @type {import('vite').UserConfig}
 */
export const viteConfigShared = {
  resolve: {
    alias: {
      "@": "/src",
      "@components": "ui/src/components",
    },
  },
};
