# DevOps for networking containers with VXLAN

This project re-uses outputs from my other project "[neutron-lan](https://github.com/araobp/neutron-lan)".

##Background and motivation

As my hobby, I just want to develop a very simple DevOps framework for networking containers for several use cases.

##NLAN services
- PTN: Packet Transport Network (Layer 1 and Layer 2)
- DVR: Distributed Virtual Switch and Distributed Virtual Router (Layer 2 and Layer 3)
- vHosts: netns-based virtual hosts
- Router: Quagga configuration

##Target use cases

### Use case 1: Network simulation

This use case makes use of NLAN's PTN, vHosts and Router services.
![WAN simulation](https://docs.google.com/drawings/d/1VKfKlwnzWQ2-ImfXeB5uNegGBK0BnaGU_4lS8h4Qpcw/pub?w=640&h=480)

####Declarative state representations:
- [ptn-bgp.yaml](./etc/ptn-bgp.yaml)
- [ptn-ospf.yaml](./etc/ptn-ospf.yaml)

### Use case 2: SOHO NFV (Network Functions Virtualization)
![SONO-NFV](https://docs.google.com/drawings/d/11fJUimZVrGxqAdq-hJK4abDu0ZThkfHGtbl_94zW0rQ/pub?w=640&h=480)

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
- [YANG model](./model/nlan/nlan.yang)
- [protobuf model](./model/nlan/nlan.proto)
- [rpc model](./model/nlan/rpc.proto)
- [Go gRPC stub](./model/nlan/nlan.pb.go)

##Key technologies used for this project
- Golang
- Open vSwitch, OVSDB and JSON-RPC(RFC7047)
- VXLAN
- Docker
- etcd
- YANG/goyang
- JSON/YAML
- Protocol buffers and gRPC
- Quagga

#NLAN installation

[Step 1] Make a Docker image named "router" following the instruction [here](./docker/SETUP.md).

[Step 2] Run the following shell script to build Docker image with NLAN agent embedded and to start the containers:
```
$ ./setup.sh
```
[Step 3]
Try this to deploy "PTN/Vhost/Router" services:
```
$ ./master.sh
```
The script sets up [this network](https://camo.githubusercontent.com/3f15c9634b2491185ec680fa5bb7d19f6f01146b/68747470733a2f2f646f63732e676f6f676c652e636f6d2f64726177696e67732f642f31564b664b6c776e7a5751322d496d6658654235754e656747424b30426e6147555f346c53386834517063772f7075623f773d39363026683d373230).

[Step 4]
Open ssh session to the containers:
```
$ cd docker
$ ./ssh.sh pe1
       :
$ ./ssh.sh ce1
       :
```
The password is "root".

Then you can do whatever you try.

#Development environment setup

##Building Golang and protobuf for 32bit Linux
I use a very old PC with a 32bit CPU at home, so I need to build 32bit binary from source codes:
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
