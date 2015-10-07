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

##What I have noticed
1. 'goyang --format=proto' does not add 'syntax = "proto2";' at the beginning of proto file.

##What I have not tested yet
1. Test if goyang supports YANG "rpc" statement.
2. goyang with gRPC
