import { Body, Delete, Param, Query, Req, Post } from '@nestjs/common';
import { Request } from 'express';
import { OmitType, PartialType } from '@nestjs/swagger';
import {
  SwaggerDeleteByIdDecorators,
  SwaggerGetByIdDecorators,
  SwaggerPostDecorators,
  SwaggerPutDecorators,
  ControllerDecorators,
  SwaggerGetDecorators,
} from '@nx-go-playground/api/helpers';

import { UsersService } from './users.service';
import { CreateUserDto } from './dto/create-users.dto';
import { UpdateUserDto } from './dto/update-users.dto';
import { User } from './users.entity';

enum Projection {
  name = 'name',
  role = 'role',
  email = 'email',
  tenantId = 'tenantId',
  provider = 'provider',
}

// @ApiCookieAuth()
// @ApiOAuth2(['users:write'])
// @ApiBearerAuth()
// @ApiBasicAuth()
@ControllerDecorators('users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @SwaggerPostDecorators(User)
  post(@Body() createItemDto: CreateUserDto): Promise<User> {
    return this.usersService.create(createItemDto);
  }

  @SwaggerGetDecorators<User>(
    Projection,
    User,
    PartialType(OmitType(User, ['_id', 'password']))
  )
  getData(
    @Req() request: Request,
    @Query('projection') projection: Projection | [Projection] | null,
    @Query('limit') limit = 0
    // @Query('search') search: Partial<Omit<User, '_id' | 'password'>>
  ): Promise<Array<User>> {
    delete request.query['limit'];
    delete request.query['projection'];
    // console.log('search', search);
    return this.usersService.findAll(request.query, projection, {
      limit,
      // page: Number(skip),
    });
  }

  @SwaggerGetByIdDecorators(Projection, User)
  findById(
    @Query('projection') projection: any,
    @Param('id') id: string // did not work with id of type User['_id'] or custom new
  ): Promise<User> {
    return this.usersService.findById(id, projection);
  }

  @SwaggerPutDecorators(User)
  update(@Body() body: UpdateUserDto, @Param('id') id: string) {
    // Promise<User>
    return this.usersService.update(id, body);
  }

  @SwaggerDeleteByIdDecorators()
  delete(@Param('id') id: string): Promise<string> {
    return this.usersService.delete(id);
  }

  @Delete()
  deleteAll() {
    return this.usersService.deleteMany();
  }
}
