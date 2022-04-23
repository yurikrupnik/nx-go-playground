import {
  Document,
  FilterQuery,
  QueryOptions,
  LeanDocument,
  HydratedDocument,
} from 'mongoose';
import { EntityRepository } from './entity.repository';

export class CrudApiService<
  T extends Document,
  CreateDto,
  UpdateDto,
  Repository extends EntityRepository<T, CreateDto, UpdateDto>
> {
  constructor(private readonly repository: Repository) {}

  findAll(
    query: FilterQuery<T>,
    projection: any,
    config: QueryOptions
  ): Promise<LeanDocument<HydratedDocument<T>>[]> {
    return this.repository.findAll(query, projection, config);
  }

  findById(id: string, projection: Partial<T>): Promise<T> {
    return this.repository.findById(id, projection);
  }

  create(body: CreateDto): Promise<T> {
    return this.repository.create(body);
  }

  update(id: string, body: UpdateDto) {
    return this.repository.findOneAndUpdate(id, body);
  }

  delete(id: string): Promise<string> {
    return this.repository.deleteOne(id);
  }

  deleteMany() {
    // this.repository.pago
    return this.repository.deleteMany({});
  }
}
