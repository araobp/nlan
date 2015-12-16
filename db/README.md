#Simple database

##Idea
gRPC uses protobuf as a wire format. I develop a gRPC intercepter to extract CRUD data from gRPC and save it onto a file:

```
--- protobuf []byte --> gRPC server --> gRPC methods --+--> RPC processing
                                                       |
                                                 codec.Marshal()
                                                       |
                                                       V
                                                      file
```

##Format
I refer to my previous database project "[tega](https://github.com/araobp/tega)".

The format takes a CommitLog-style like this:
```
ADD {"a": 1, "b": 2}
UPDATE {"a": 2}
DELETE {"b": null}
```

