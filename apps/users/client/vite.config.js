// import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';
// import yargs from 'yargs';
// const argv = yargs.argv;
// const { port, open } = argv;
// console.log(argv);
// export default defineConfig(async () => {
//   return {
//     build: {
//       // manifest: true,
//     },
//     server: {
//       port,
//       open,
//     },
//     plugins: [solidPlugin()],
//   };
// });
import { defineConfig } from 'vite';
// import react from '@vitejs/plugin-react';
// import federation from '@originjs/vite-plugin-federation';
import { argv } from 'yargs';
// import path from 'path';

const { port, open } = argv;

// https://vitejs.dev/config/
export default defineConfig((c, o) => {
  console.log(c, o);
  console.log('argv', argv);
  console.log('port', port);
  console.log('open', open);
  console.log('open', __dirname);
  console.log('open', process.cwd());
  console.log('open', __filename);
  return {
    // output: {
    //   // format: 'esm',
    //   dir: 'dist',
    // },
    // build: {
    //   lib: {
    //     entry: 'src/index.tsx',
    //     // name: 'MyLib',
    //     formats: ['system'],
    //     // fileName: (format) => `my-lib.${format}.js`
    //   },
    //   rollupOptions: {
    //     // make sure to externalize deps that shouldn't be bundled
    //     // into your library
    //     external: ['solid-js'],
    //     // output: {
    //     //   dir: 'dist',
    //     //   // Provide global variables to use in the UMD build
    //     //   // for externalized deps
    //     //   globals: {
    //     //     vue: 'react'
    //     //   }
    //     // },
    //   },
    // },

    plugins: [
      solidPlugin(),
      //   federation({
      //     // remoteType: 'var',
      //     filename: 'remoteEntry.js',
      //     name: 'aris',
      //     exposes: {
      //       './Button': path.join(__dirname, 'src/button.tsx'),
      //       // './Button': './src/button.tsx',
      //     },
      //     remotes: {
      //       // home: 'home@http://localhost:3001/remoteEntry.js',
      //     },
      //     shared: ['solid-js'],
      //   }),
      //   // react(),
    ],
    server: {
      // port: 9999,
      open,
    },
    // plugins: [solidPlugin()],
    // plugins: [
    //   react(),
    //   federation({
    //     name: 'app',
    //     remoteType: 'var',
    //     filename: 'remoteEntry.js',
    //     exposes: {
    //       './Button': './src/Button',
    //     },
    //     remotes: {
    //       nav: 'nav',
    //     },
    //     shared: ['react', 'react-dom'],
    //   }),
    // ],
  };
});
