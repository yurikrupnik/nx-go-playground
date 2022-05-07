// import { PassportStrategy } from '@nestjs/passport';
// import { Injectable, UnauthorizedException } from '@nestjs/common';
// import { Strategy, ExtractJwt } from 'passport-firebase-jwt';
// import { InjectFirebaseAdmin, FirebaseAdmin } from '@nx-go-playground/nest/modules';
//
// @Injectable()
// export class FirebaseAuthStrategy extends PassportStrategy(Strategy) {
//   constructor(@InjectFirebaseAdmin() private readonly firebase: FirebaseAdmin) {
//     super({
//       jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
//     });
//   }
//
//   validate(token: string) {
//     return this.firebase
//       .auth()
//       .verifyIdToken(token) // todo check what was the second parapm of true
//       .catch((err: Error) => {
//         console.log(err);
//         throw new UnauthorizedException();
//       });
//   }
// }
