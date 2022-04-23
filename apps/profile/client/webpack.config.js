const reactWebpackConfig = require('@nrwl/react/plugins/webpack');
const path = require('path');
// const { WebpackManifestPlugin } = require('webpack-manifest-plugin');
// const { ESBuildMinifyPlugin } = require('esbuild-loader');
const { ModuleFederationPlugin } = require('webpack').container;
// const BundleAnalyzerPlugin =
//   require('webpack-bundle-analyzer').BundleAnalyzerPlugin;

function addEsBuildAndRemoveBabel(config) {
  config.module.rules.splice(1, 1, {
    test: /\.tsx?$/,
    loader: 'esbuild-loader',
    options: {
      loader: 'tsx',
      target: 'es2015',
    },
  });
}

function getWebpackConfig(config) {
  config = reactWebpackConfig(config);
  // console.log( 'config', config.module.rules);
  // config.output.publicPath = 'http://localhost:3001';
  config.module.rules.splice(1, 1, {
    test: /\.tsx?$/,
    loader: 'esbuild-loader',
    options: {
      loader: 'tsx',
      target: 'es2015',
    },
  });

  // addEsBuildAndRemoveBabel(config);

  config.plugins.push(
    new ModuleFederationPlugin({
      // runtime: 'my-run-time',
      name: 'react_app',
      library: { type: 'var', name: 'react_app' },
      filename: 'remoteEntry.js',
      remotes: {
        // nav: 'nav@http://localhost:3003/remoteEntry.js',
        // // Header SecondHeader
        // home: 'home@http://localhost:3001/remoteEntry.js',
        // // Stam fruit ProductCard
        // 'react-app': 'react-app@http://localhost:3339/remoteEntry.js',
      },
      exposes: {
        './Button': path.join(__dirname, 'src/app/button.tsx'),
      },
      shared: {
        react: {
          eager: true,
        },
        'react-dom': {
          eager: true,
        },
        // 'react-router-dom': {
        //   // eager: true,
        // },
        // 'react-query': {
        //   // eager: true,
        // },
      },
      // shared: [{ 'react': {
      //
      //   } }, 'react-dom', 'react-router-dom'],
    })
    // new BundleAnalyzerPlugin({
    //   analyzerMode: 'static',
    //   reportFilename: 'bundles-report/index.html',
    // })
  );

  config.mode = 'development';

  config.optimization = {
    ...config.optimization,
    // runtimeChunk: false,
    minimizer: [
      // new ESBuildMinifyPlugin({
      //   target: 'es2015',
      // }),
    ],
    splitChunks: {
      chunks(chunk) {
        return false;
      },
    },
  };

  // https://github.com/facebook/create-react-app/blob/bb64e31a81eb12d688c14713dce812143688750a/packages/react-scripts/config/webpack.config.js#L705
  // config.plugins.push(
  //   new WebpackManifestPlugin({
  //     fileName: 'manifest.json',
  //     publicPath: '/',
  //     generate: (seed, files, entrypoints) => {
  //       const manifestFiles = files.reduce((manifest, file) => {
  //         manifest[file.name] = file.path;
  //         return manifest;
  //       }, seed);
  //       const entrypointFiles = entrypoints.main.filter(
  //         (fileName) => !fileName.endsWith('.map')
  //       );
  //
  //       return {
  //         files: manifestFiles,
  //         entrypoints: entrypointFiles,
  //       };
  //     },
  //   })
  // );

  return config;
}

module.exports = getWebpackConfig;
