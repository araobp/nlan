##Findings so far (2016/01/10)

In this project, I originally used etcd, goyang and gRPC, but those have some issues.

In summary:
- etcd is just a KVS, neither document-oriented DB nor column-oriented DB, no good for storing config for networking.
- etcd does not support "ephemeral nodes".
- etcd uses a certain ammount of CPU time and memory.
- etcd's APIs are not user-friendly.
- etcd is not easy to manage.
- goyang does not support part of YANG spec.
- YANG is incompatible with JSON, especially YANG's list type.
- YANG is not for DevOps: most of DevOps frameworks use JSON/YAML format.
- hard to debug gRPC messaging, since it's encoded in a binary format.

So I gave up using those.

Other findings:
- [I have got that etcd (schema-less, KVS) is not suitable for this project](https://github.com/araobp/nlan/issues/12), so I am going to use [tega](https://github.com/araobp/tega) instead -- Tornado/Python is very good for a single-core 32bit CPU, and hash-table-based python dict is useful for data manipulation in some cases.
- The combination of Docker and Golang is OK.
- protobuf and gRPC are useful. But I would use pubsub (subscribe/notify) rather than gRPC, for managing containers: something like ZooKeeper.
- YANG is not compatible with JSON/YAML, which is the reason why a lot of people dislike YANG.
- OVSDB is just a read-only database in my project.
- Processing overhead issues (Linux/OVS-bridges and VXLAN): I need to consider using macvlan-vepa with a smart physical switch with VLAN and SNMP support.
- Cumulus Linux uses [netlink](https://tools.ietf.org/html/rfc3549) for controlling/managing both Linux switching/routing tables and hardware(ASIC), which seems very interesting: I have also tried out [tenus](https://github.com/milosgajdos83/tenus).

##YANG-JSON mapping(IETF)

[Reference] https://tools.ietf.org/html/draft-ietf-netmod-yang-json

