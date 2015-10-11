#neutron-lan rewrite
This is a directory for rewriting nlan agent.

##NLAN config modules
1. NLAN agent receives a data change notification from etcd
2. NLAN agent calls a NLAN config module

##NLAN rpc modules
RPC modules registers their RPC services with gRPC
