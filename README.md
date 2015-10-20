# Rewrite [neutron-lan](https://github.com/araobp/neutron-lan) with Go lang and Docker

##Goal of this project
The goal of this project is to study how DevOps tool for networking can be developed with Golang and Docker containers.

##Policy
- Write codes!
- Powerpoint documents, no thanks.

##Use case
Network simulation, especially simulated WAN to test routing daemons (such as quagga/zebra) and other SDN-related "go-something".

##Comparison
- [Software Defined Networking at Scale](http://files.meetup.com/8218762/Bikash_Koley%20SDN_meetup%20May%202015.pdf)
- [neutron-lan](https://github.com/araobp/neutron-lan) -- golang/gRPC/OpenConfig/YANG/protobuf/etcd
- [My ODL-app project](https://github.com/araobp/odl-app) -- Java/YANG/MD-SAL

##Architecture
```
     [    Master   ] --- Global DB
        |       |
      gRPC    gRPC ...
        |       |
        V       V
    [Agent]   [Agent] --- Local DB
```
##Rewrite works overview
- Write NLAN network service models in YANG
- Use gRPC instead of "Python OrderedDict over SSH"
- Rewirte NLAN Master and Agent (config/rpc) in Go lang
- Remove OVSDB-dependency as much as possible
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
I use a very old PC with a 32bit CPU, so I need to build 32bit binary from source codes:
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

###Docker-related
- [Docker](https://github.com/docker/docker)
- [Kubernetes](https://github.com/kubernetes/kubernetes)

###Database-related
- [etcd](https://github.com/coreos/etcd)
- [mergo](https://github.com/imdario/mergo)
- [ovsdb client](https://github.com/socketplane/libovsdb)
- [yaml](https://github.com/go-yaml/yaml)
- [goyang](https://github.com/openconfig/goyang)
- [gocql](https://github.com/gocql/gocql)

###RPC and protocols
- [gRPC](https://github.com/grpc/grpc-go/)
- [gobgp](https://github.com/osrg/gobgp)
