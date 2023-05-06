const defaultViteBody = /<body>/;
const nextViteBody = `
  <body>
  	<script src="https://static.zdassets.com/zendesk_app_framework_sdk/2.0/zaf_sdk.min.js"></script>
`;

export const injectZafHtmlPlugin = () => {
  return {
    name: "inject-zaf-html-plugin",
    transformIndexHtml: (html) => {
      return html.replace(defaultViteBody, nextViteBody);
    },
  };
};
