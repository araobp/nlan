# Rewrite [neutron-lan](https://github.com/araobp/neutron-lan) with Go lang and etcd

##Goal of this project
The goal of this project is to study how DevOps tool for networking can be developed with Golang and Docker containers.

##Policy
- Write codes!
- Powerpoint documents, no thanks.

##Architecture
```
     [    Master   ] --- Global DB (etcd)
        |       |
      gRPC    gRPC ...
        |       |
        V       V
    [Agent]   [Agent] --- Local DB (etcd and ovsdb(read-only))
```
##Rewrite works overview
- Write NLAN network service models in YANG
- Use gRPC instead of "Python OrderedDict over SSH"
- Rewirte NLAN Master and Agent (config/rpc) in Go lang
- Remove OVSDB-dependency, use etcd instead
- Bi-directional RPCs (any libraries available on github?)

##NLAN model in YANG and protobuf
Go stub generation
```
                             ___________  
                            /rpc model / proto3
                            ~~~~~~~~~~~
                                 |
                               merge
                                 |
                                 V
 ___________                _______________                _____________
/YANG model/ == goyang ==> /Protobuf model/ == protoc ==> /Go gRPC stub/
~~~~~~~~~~~                ~~~~~~~~~~~~~~~                ~~~~~~~~~~~~~
                               proto3
```
- [YANG model](./nlan/model/nlan/nlan.yang)
- [protobuf model](./nlan/model/nlan/nlan.proto)
- [rpc model](./nlan/model/nlan/rpc.proto)
- [Go gRPC stub](./nlan/model/nlan/nlan.pb.go)

##JSON database with YANG as a schema language for JSON

I used ovsdb as a JSON database with OVSDB-schema in my [neutron-lan](http://github.com/araobp/neutron-lan) project.

[Some thoughts](./DATABASE.md) for this project.

##Preparations
- Go lang installation: https://golang.org/dl/
- Protobuf build and installation: https://github.com/google/protobuf/blob/master/INSTALL.txt
```
$ autoconf
$ ./autogen.sh
$ ./configure
$ make
$ make install
```
- etcd installation: https://github.com/coreos/etcd
```
$ ./build
``` 
- Add /usr/local/lib to LD_LIBRARY_PATH
```
$ export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARLY_PATH

```

##Go plugin for vim

Install [vim-go](https://github.com/fatih/vim-go) to your vim.

##Interesting libraries written in Go lang
- [gRPC](https://github.com/grpc/grpc-go/)
- [etcd](https://github.com/coreos/etcd)
- [godoc etcd client](https://godoc.org/github.com/coreos/etcd/client)
- [OVSDB client](https://github.com/socketplane/libovsdb)
- [YAML](https://github.com/go-yaml/yaml)
- [YANG](https://github.com/openconfig/goyang)
- [gocql](https://github.com/gocql/gocql)
- [Docker](https://github.com/docker/docker)
- [Kubernetes](https://github.com/kubernetes/kubernetes)

##Reference
- [Software Defined 
Networking at Scale, Bikash Koley, Anees Shaikh on behalf of Google Technical Infrastructure
, 5/12/2015](http://files.meetup.com/8218762/Bikash_Koley%20SDN_meetup%20May%202015.pdf)
