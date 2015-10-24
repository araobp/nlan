# Rewrite [neutron-lan](https://github.com/araobp/neutron-lan) with Go lang and Docker

##Goal of this project
The goal of this project is to study how DevOps tool for networking can be developed with Golang and Docker containers.

##Policy
- Write codes!
- Powerpoint documents, no thanks.

##Use case
- Network simulation, especially simulated WAN to test routing daemons (such as quagga/zebra) and other SDN-related "go-something"
- Network simulation for open source SDN controllers (ODENOS, OpenDaylight and ONOS)

##Network simulation with Linux containers
I use Linux containers as virtual routers, and this tool will set up virtual links (L0/L1) and virtual switches (L2) over the containers. Then I will run quagga/zebra(L3) daemons over the virtual routers to study how legacy routing protocols work.
- [An example of such a network](https://camo.githubusercontent.com/3f15c9634b2491185ec680fa5bb7d19f6f01146b/68747470733a2f2f646f63732e676f6f676c652e636f6d2f64726177696e67732f642f31564b664b6c776e7a5751322d496d6658654235754e656747424b30426e6147555f346c53386834517063772f7075623f773d39363026683d373230)

##Comparison
- [Software Defined Networking at Scale](http://files.meetup.com/8218762/Bikash_Koley%20SDN_meetup%20May%202015.pdf)
- [odenos project](https://github.com/o3project/odenos) -- Java/MessagePack/ODENOS-network-component/Redis/ZooKeeper
- [My "odl-app" project](https://github.com/araobp/odl-app) -- Java/YANG/OpenDaylight-MD-SAL
- [My "onos-app" project](https://github.com/araobp/onos-app) -- Java/kryo/ONOS-ECMAP/ONOS-RAFT
- [My "neutron-lan" project](https://github.com/araobp/neutron-lan) -- Python/OrderdDict-over-SSH/OVSDB-schema/MIME/ovsdb
- This "golan" project -- Golang/gRPC/OpenConfig/protobuf/etcd

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
