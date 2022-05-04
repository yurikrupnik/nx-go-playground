const { merge } = require('webpack-merge');
const path = require('path');
const singleSpaDefaults = require('webpack-config-single-spa-react-ts');

module.exports = (webpackConfigEnv, argv) => {
  const defaultConfig = singleSpaDefaults({
    orgName: 'aris',
    projectName: 'nira',
    webpackConfigEnv,
    argv,
  });

  console.log('defaultConfig', defaultConfig);
  return merge(defaultConfig, {
    context: path.resolve(process.cwd(), 'tsconfig.base.json'),
    // modify the webpack config however you'd like to by adding to this object
  });
};
