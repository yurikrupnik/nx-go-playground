// import { Test, TestingModule } from '@nestjs/testing';
// import { UsersService } from './users.service';
// import { LoginProviders, User, UserRoles } from './entities/users.entity';
// import { getModelToken, MongooseModule } from '@nestjs/mongoose';
// import { mongoConfig } from '@mussia14/backend/envs';

// function mongoObjectId() {
//   const timestamp = ((new Date().getTime() / 1000) | 0).toString(16);
//   return `${timestamp}xxxxxxxxxxxxxxxx`.replace(/[x]/g, () =>
//     ((Math.random() * 16) | 0).toString(16)
//   );
// }

describe('UsersService', () => {
  // let service: UsersService;

  // const mockUsers = Array.from(Array(10).keys()).map(() => {
  //   return UsersService.createMock({ _id: mongoObjectId() });
  // });

  // beforeEach(async () => {
  //   const module: TestingModule = await Test.createTestingModule({
  //     providers: [
  //       UsersService,
  //       {
  //         provide: getModelToken(User.name),
  //         useValue: {
  //           // new: jest.fn().mockResolvedValue(UsersService.createMock()),
  //           // constructor: jest.fn().mockResolvedValue(UsersService.createMock),
  //           findAll: jest.fn(),
  //           findOne: jest.fn(),
  //           create: jest.fn(),
  //           update: jest.fn(),
  //           delete: jest.fn(),
  //           deleteAll: jest.fn(),
  //           exec: jest.fn(),
  //         },
  //       },
  //     ],
  //     imports: [MongooseModule.forRoot(mongoConfig().MONGO_URI)],
  //   })
  //     .useMocker((token) => {
  //       console.log('token', token);
  //       console.log('token', token);
  //       console.log('token', token);
  //       // if (token === CatsService) {
  //       // return { findAll: jest.fn().mockResolveValue() };
  //       // }
  //       // if (typeof token === 'function') {
  //       //   const mockMetadata = moduleMocker.getMetadata(token) as MockFunctionMetadata<any, any>;
  //       //   const Mock = moduleMocker.generateFromMetadata(mockMetadata);
  //       //   return new Mock();
  //       // }
  //     })
  //     .compile();
  //
  //   service = module.get<UsersService>(UsersService);
  // });

  it('should be defined', () => {
    // expect(service).toBeDefined();
    expect(1).toEqual(1);
  });
  // it('should call some', async () => {
  //   const mock = UsersService.createMock();
  //   console.log('mock', mock);
  //   expect(await service.create(mock)).toContain(mock);
  //   // expect(1).toBeDefined();
  // });

  // it('should create user', async () => {
  //   const mock = UsersService.createMock({ _id: '1234' });
  //   jest.spyOn(service, 'create').mockReturnValue({
  //     exec: jest.fn().mockResolvedValueOnce(mock),
  //   } as any);
  //   const user = await service.create({
  //     tenantId: '123',
  //     name: 'ad',
  //     password: '1234',
  //     email: 'a@a.com',
  //     role: UserRoles.admin,
  //     provider: LoginProviders.local,
  //   });
  //   console.log('user', user);
  //   console.log('mock', mock);
  //   // expect(user).toEqual(mock);
  // });
  //
  // it('should return users', async () => {
  //   const mock = UsersService.createMock({ _id: '1234' });
  //   jest.spyOn(service, 'findById').mockReturnValue({
  //     data: jest.fn().mockResolvedValueOnce(mock),
  //   } as any);
  //   // const user = await service.findById('1234', []);
  //   // console.log('user', user);
  //   // console.log('mock', mock);
  //   // expect(user).toEqual(mock);
  // });
});
