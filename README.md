# Rewrite [neutron-lan](https://github.com/araobp/neutron-lan) with Go lang, YANG and Docker

##Goal of this project
The goal of this project is to study how DevOps tool for networking can be developed with Golang, YANG and Docker.

Current status (as of November 30th, 2015): PTN simulation is working with multiple containers.

##Background and motivation
- The classical definition of SDN is becoming obsolete for most of cases (excluding cases where dynamic control is required).
- The next definition of SDN is similar to those of Cloud Management System (such as OpenStack) and PaaS (such as Kubernetes/Docker).

##Policy
- Write codes! Powerpoint documents, no thanks.
- Respect CLI! Think "CLI for SDN"! (YANG/NETCONF is sort of "CLI for SDN", but we may re-invent another form of "CLI for SDN" supporting transaction/rollback)

##Use cases
- Network simulation, especially simulated WAN to test routing daemons (such as quagga/zebra) and other SDN-related "go-something"
- Network simulation for open source SDN controllers such as OpenDaylight.
- DevOps for networking with whitebox routers/switches.

##Future use cases
If some SoC for SOHO routers supports VXLAN off-loading in future, I will develop SDN for LAN/WAN with this code (nlan/go-nlan). A sort of "(wired, not wireless) LAN controller".

##Network simulation with Linux containers
I use Linux containers as virtual routers, and this tool will set up virtual links (L0/L1) and virtual switches (L2) over the containers. Then I will run Quagga/Zebra(L3) daemons over the virtual routers to study how legacy routing protocols work.
- [An example of such a network](https://camo.githubusercontent.com/3f15c9634b2491185ec680fa5bb7d19f6f01146b/68747470733a2f2f646f63732e676f6f676c652e636f6d2f64726177696e67732f642f31564b664b6c776e7a5751322d496d6658654235754e656747424b30426e6147555f346c53386834517063772f7075623f773d39363026683d373230)
- [Working with Docker for network simulation](https://camo.githubusercontent.com/77cf473ea9499432e57b06a951f5f5248419f9e1/68747470733a2f2f646f63732e676f6f676c652e636f6d2f64726177696e67732f642f313631426e383077384a5a4b513742586d496f306272377851346b71456442635f585a3235347a754f5253552f7075623f773d36383026683d343030)

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

##Declarative state representation
See [ptn.yaml](nlan/etc/ptn.yaml) as a declarative state representation of simulated Packet Transport Network.

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
