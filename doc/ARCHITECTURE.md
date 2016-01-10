##Architecture

####Current
```
          etc/*.yaml
              |
     [      master     ]----[etcd]
       |      |      |        |
         gRPC                 |
       |      |      |        |
       V      V      V        |
   [Agent] [Agent] [Agent] ---+
```

####Next
I am going to add "DCN with two-phase commit" feature to tega, and use tega CLI as NLAN CLI.

```
[tega plugins] --- [ tega(tornado)   ] --- tega CLI
                     |      |      |
                tega API (HTTP/WebSocket)
                     |      |      |
                     V      V      V
                  [Agent] [Agent] [Agent]
```

New features to be added to tega:
- DCN invoked from CLI
- Ephemeral nodes to register IP addresses of each containers (as DNS): I don't use mDNS, because I want to deploy NLAN across real WAN
- DCN with two phase commit


[DATABASE](./DATABASE.md)

####Use of protobuf
I am going to use protobuf as a schema language for tega:
- To generate code for go struct <-> JSON mapping
- To validate JSON data

tega is a schema-less db and protobuf is just for data mapping and data validation.
