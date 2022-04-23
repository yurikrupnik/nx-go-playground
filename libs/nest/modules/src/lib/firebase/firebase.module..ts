// import { DynamicModule, Global, Module, Provider } from '@nestjs/common';
// import { ClassProvider } from '@nestjs/common/interfaces';
// import { FirebaseConstants } from './firebase.constants';
// import {
//   FirebaseModuleAsyncOptions,
//   FirebaseModuleOptions,
//   FirebaseModuleOptionsFactory,
// } from './firebase.interface';
// import { getFirebaseAdmin } from './utils';
// import { AppOptions, app } from 'firebase-admin';
//
// @Global()
// @Module({})
// export class FirebaseModule {
//   public static forRoot(options: AppOptions): DynamicModule {
//     const provider: Provider<app.App> = {
//       provide: FirebaseConstants.FIREBASE_TOKEN,
//       useValue: getFirebaseAdmin(options),
//     };
//
//     return {
//       exports: [provider],
//       module: FirebaseModule,
//       providers: [provider],
//     };
//   }
//
//   public static forRootAsync(
//     options: FirebaseModuleAsyncOptions
//   ): DynamicModule {
//     const firebaseProvider: Provider = {
//       inject: [FirebaseConstants.FIREBASE_MODULE],
//       provide: FirebaseConstants.FIREBASE_TOKEN,
//       useFactory: (options: FirebaseModuleOptions) => getFirebaseAdmin(options),
//     };
//
//     const asyncProviders = this.createAsyncProviders(options);
//     return {
//       module: FirebaseModule,
//       imports: [...(options.imports || [])],
//       providers: [...asyncProviders, firebaseProvider],
//       exports: [firebaseProvider],
//     };
//   }
//
//   private static createAsyncProviders(
//     options: FirebaseModuleAsyncOptions
//   ): Provider[] {
//     if (options.useFactory || options.useExisting) {
//       return [this.createAsyncOptionsProvider(options)];
//     }
//     return [
//       this.createAsyncOptionsProvider(options),
//       {
//         provide: options.useClass,
//         useClass: options.useClass,
//         inject: [options.inject || []],
//       } as ClassProvider,
//     ];
//   }
//
//   private static createAsyncOptionsProvider(
//     options: FirebaseModuleAsyncOptions
//   ): Provider {
//     if (options.useFactory) {
//       return {
//         provide: FirebaseConstants.FIREBASE_MODULE,
//         useFactory: options.useFactory,
//         inject: options.inject || [],
//       };
//     }
//     return {
//       provide: FirebaseConstants.FIREBASE_MODULE,
//       useFactory: async (
//         optionsFactory: FirebaseModuleOptionsFactory
//       ): Promise<FirebaseModuleOptions> =>
//         await optionsFactory.createFirebaseModuleOptions(),
//       inject: options.useClass ? [options.useClass] : [],
//     };
//   }
// }
