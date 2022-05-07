import { apiUsers } from './api-users';

describe('apiUsers', () => {
  it('should work', () => {
    expect(apiUsers()).toEqual('api-users');
  });
});
