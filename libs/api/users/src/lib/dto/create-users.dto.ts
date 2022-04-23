import { OmitType } from '@nestjs/swagger';
import { User } from '../users.entity';

export class CreateUserDto extends OmitType(User, ['_id'] as const) {}
