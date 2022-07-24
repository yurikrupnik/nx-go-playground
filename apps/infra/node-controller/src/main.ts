import Operator, { ResourceEventType, ResourceEvent } from '@dot-i/k8s-operator';

export default class MyOperator extends Operator {
  protected async init() {
    await this.watchResource('', 'v1', 'namespaces', async (e) => {
      const object = e.object;
      const metadata = object.metadata;
      console.log('metadata', metadata)
      switch (e.type) {
        case ResourceEventType.Added:
          // do something useful here
          break;
        case ResourceEventType.Modified:
          // do something useful here
          break;
        case ResourceEventType.Deleted:
          // do something useful here
          break;
      }
    });
  }
}
