##Testing pubsub pattern with etcd

"etcdctl" does not seem to support "CreateInOrder" that is necessary to append/remove an item to/from a list on etcd.

This sample program is to test such a operation by using APIs defined in keys.go

```
[provider] --> etcd --> [watcher]
```
