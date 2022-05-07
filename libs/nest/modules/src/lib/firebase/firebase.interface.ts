// import type { Type } from '@nestjs/common';
// import type { ModuleMetadata } from '@nestjs/common/interfaces';
// import { AppOptions, app } from 'firebase-admin';
//
// export type FirebaseModuleOptions = AppOptions;
//
// export type FirebaseModuleAsyncOptions = {
//   useClass?: Type<FirebaseModuleOptionsFactory>;
//   useFactory?: (...args: any[]) => Promise<AppOptions> | AppOptions;
//   inject?: any[];
//   useExisting?: Type<FirebaseModuleOptionsFactory>;
// } & Pick<ModuleMetadata, 'imports'>;
//
// export interface FirebaseModuleOptionsFactory {
//   createFirebaseModuleOptions(): Promise<AppOptions> | AppOptions;
// }
//
// export type FirebaseAdmin = app.App;
