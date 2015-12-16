#Simple database

The package "github.com/araobp/nlan/util" includes command execution utilities: cmd() and cmpd(). When rebooting a Docker container, NLAN agent needs to load NLAN state from a local database to resume NLAN state by executing cmd(). On the other hand, cmdp() just skips its execution when restarting, since its state survives over reboots.

I considered using [tiedot](https://github.com/HouzuoGuo/tiedot) as a document-oriented database, but it seems too heavy for NLAN, since I am planning to run NLAN on Raspberry Pi.

I need to develop a light-weight document-oriented database on my own, reducing dependencies on other go libraries as much as possible.

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

