##goyang test

Prerequisites:
- pyang
- goyang

pyang installation:
```
$ pip install pyang
```
goyang installation:
```
$ go get github.com/openconfig/goyang
$ cd <workspace>/src/github.com/openconfig/goyang
$ go build
```

##What I have tested
1. Models were defined in YANG.
2. "goyang" generated proto files from the YANG models.
3. "protoc" generated pb.go files from the proto files.
4. I used the generated pb.go files to make data structure, marshal the data into the wire format ([]byte) and unmarshal the []byte data back to the original data strucutre.

```
YANG models => proto files => pb.go files

Original data == (marshal) ==> wire format == (unmarshal) ==> Original data
```

##Console output
```
$ go build yang-test.go
$ ./yang-test
--- NLAN bridges module ---
Before encoding: OvsBridges:true
Wire format(encoded): [8 1]
After decoding: OvsBridges:true

--- NLAN vxlan module ---
Before encoding: LocalIp:"10.0.0.1" RemoteIps:"192.168.56.101" RemoteIps:"192.168.56.102"
Wire format(encoded): [10 8 49 48 46 48 46 48 46 49 18 14 49 57 50 46 49 54 56 46 53 54 46 49 48 49 18 14 49 57 50 46 49 54 56 46 53 54 46 49 48 50]
After decoding: LocalIp:"10.0.0.1" RemoteIps:"192.168.56.101" RemoteIps:"192.168.56.102"

--- NLAN subnets module ---
Before encoding: IpDvr:<Addr:"172.16.42.1" Mode:"dvr" > Vid:10 Vni:100
Wire format(encoded): [10 18 10 11 49 55 50 46 49 54 46 52 50 46 49 26 3 100 118 114 32 10 40 100]
After decoding: IpDvr:<Addr:"172.16.42.1" Mode:"dvr" > Vid:10 Vni:100
```

##What I have noticed
1. 'goyang --format=proto' does not add 'syntax = "proto2";' at the beginning of proto file.

##What I have not tested yet
1. Test if goyang supports YANG "rpc" statement.
2. goyang with gRPC
