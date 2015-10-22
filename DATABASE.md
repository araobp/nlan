##WHICH DATABASE I SHOULD USE

- MongoDB can be a database to store Go struct objects serialized into JSON.
- Or use encoding/json package to marshal/unmarshal Go struct objects stored on etcd.
- Or use ovsdb with some libraries I wrote for [neutron-lan](https://github.com/araobp/neutron-lan).
- [tega](https://github.com/araobp/tega) is the most suitable database to satisfy my requirements, but it was written in Python (consumes a lot of memory, dependent on Tornado) and it has not been completed yet...
- [mergo](https://github.com/imdario/mergo) is a very interesting Golang package, but it is just for default config.
- OpenDaylight's MD-SAL datastore is perfect, but it is for Java only.

##Conclusion

- ovsdb just to fetch a port number of VXLAN tunnel (read-only).
<<<<<<< HEAD
- use etcd instead of ovsdb.
- run etcd on both master and agents
- use gRPC for CRUD and RPC

###Schema generation
```
[YANG schema]--->[protobuf]--->[pb.go(go struct)]
```

###Data flow
```
1. PUT/DELETE

          Master                                Agent                    
--------------------------              ----------------------------------------
<<<<<<< HEAD
[JSON/YAML]--->[Go struct]--->[gRPC]--->[Go struct]--->[JSON/YAML]--->[etcd]
    |                                        |
    |                                        V
  [etcd]                                  CLI/API


2. AGENT RESTART

          Master                                Agent                    
--------------------------              ----------------------------------------
                                        [Go struct]<---[JSON/YAML]<---[etcd]
                                             |
                                             V
                                          CLI/API
```
