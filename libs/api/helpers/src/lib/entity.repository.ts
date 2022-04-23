import {
  Document,
  FilterQuery,
  Model,
  QueryOptions,
  LeanDocument,
  HydratedDocument,
} from 'mongoose';

import { NotFoundException } from '@nestjs/common';
// import { handleError } from '@mussia14/backend/errors';

import { BadRequestException } from '@nestjs/common';

export function handleError(err: Error): never {
  throw new BadRequestException(err.message);
}


export abstract class EntityRepository<
  T extends Document,
  CreateDto,
  UpdateDto
> {
  constructor(protected readonly model: Model<T>) {}

  findById(
    id: string,
    projection?: Record<string, unknown>
    // config: QueryOptions
  ): Promise<HydratedDocument<T, any, void>> {
    return (
      this.model
        .findById(id, projection)
        // .lean() // todo failing types
        .then((res) => {
          if (!res) {
            throw new NotFoundException(`resource with id ${id} not found`);
          }
          return res;
        })
        .catch(handleError)
    );
  }

  findAll(
    query: FilterQuery<T>,
    projection: any,
    config: QueryOptions
  ): Promise<LeanDocument<HydratedDocument<T>>[]> {
    return this.model.find(query, projection, config).lean().catch(handleError);
  }

  create(createEntityData: CreateDto): Promise<T> {
    const entity = new this.model(createEntityData);
    return entity.save().catch(handleError);
  }

  findOneAndUpdate(
    id: string,
    updateEntityData: UpdateDto
  ): Promise<HydratedDocument<T, any, any> | null> {
    return this.model
      .findByIdAndUpdate(id, updateEntityData, {
        new: true,
        useFindAndModify: false,
      })
      .catch(handleError);
  }

  deleteMany(entityFilterQuery: FilterQuery<T>): Promise<boolean> {
    return this.model
      .deleteMany(entityFilterQuery)
      .then((deleteResult) => {
        return deleteResult.deletedCount >= 1;
      })
      .catch(handleError);
  }

  deleteOne(id: string): Promise<string> {
    return this.model
      .findByIdAndDelete(id)
      .then((res) => {
        if (!res) {
          throw new NotFoundException('Not found item');
        }
        return res._id;
      })
      .catch(handleError);
  }

  createMock() {
    // return new this.model({});
  }
}
