import { apiProducts } from './api-products';

describe('apiProducts', () => {
  it('should work', () => {
    expect(apiProducts()).toEqual('api-products');
  });
});
