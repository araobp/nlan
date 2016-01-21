## Using gobgp command

Use [gobgp.sh](../scripts/gobgp.sh) to issue gobgp commands to "rr" container.

### global rib
```
$ ./gobgp.sh global rib
    Network             Next Hop             AS_PATH              Age        Attrs
*>  10.1.1.1/32         10.200.1.101                              00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.1.1.2/32         10.200.1.102                              00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.1.1.3/32         10.200.1.103                              00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.1.1.4/32         10.200.1.104                              00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.1.2.1/32         10.200.1.101         1001                 00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.1.2.1/32         10.200.1.102         1001                 00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.1.2.2/32         10.200.1.101         1002                 00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.1.2.2/32         10.200.1.102         1002                 00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.1.2.3/32         10.200.1.103         1003                 00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.1.2.3/32         10.200.1.104         1003                 00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.1.2.4/32         10.200.1.103         1004                 00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.1.2.4/32         10.200.1.104         1004                 00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.10.10.0/24       10.200.1.101                              00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.10.10.0/24       10.200.1.102                              00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.10.10.0/24       10.200.1.103                              00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.10.10.0/24       10.200.1.104                              00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.200.1.0/24       10.200.1.101                              00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.200.1.0/24       10.200.1.102                              00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.200.1.0/24       10.200.1.103                              00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.200.1.0/24       10.200.1.104                              00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.200.2.0/24       10.200.1.101                              00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.200.2.0/24       10.200.1.102                              00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.200.2.0/24       10.200.1.103                              00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   10.200.2.0/24       10.200.1.104                              00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.201.11.0/24      10.200.1.101                              00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.201.12.0/24      10.200.1.101                              00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.202.11.0/24      10.200.1.102                              00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.202.12.0/24      10.200.1.102                              00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.203.13.0/24      10.200.1.103                              00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.203.14.0/24      10.200.1.103                              00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.204.13.0/24      10.200.1.104                              00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  10.204.14.0/24      10.200.1.104                              00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  172.17.0.0/16       10.200.1.101                              00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   172.17.0.0/16       10.200.1.102                              00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   172.17.0.0/16       10.200.1.103                              00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   172.17.0.0/16       10.200.1.104                              00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  172.21.1.0/24       10.200.1.101         1001                 00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   172.21.1.0/24       10.200.1.102         1001                 00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  172.21.2.0/24       10.200.1.101         1002                 00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   172.21.2.0/24       10.200.1.102         1002                 00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  172.21.3.0/24       10.200.1.103         1003                 00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   172.21.3.0/24       10.200.1.104         1003                 00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  172.21.4.0/24       10.200.1.103         1004                 00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   172.21.4.0/24       10.200.1.104         1004                 00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  172.22.1.0/24       10.200.1.101         1001                 00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   172.22.1.0/24       10.200.1.102         1001                 00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  172.22.2.0/24       10.200.1.101         1002                 00:19:06   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   172.22.2.0/24       10.200.1.102         1002                 00:19:35   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  172.22.3.0/24       10.200.1.103         1003                 00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   172.22.3.0/24       10.200.1.104         1003                 00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*>  172.22.4.0/24       10.200.1.103         1004                 00:18:57   [{Origin: ?} {Med: 0} {LocalPref: 100}]
*   172.22.4.0/24       10.200.1.104         1004                 00:19:08   [{Origin: ?} {Med: 0} {LocalPref: 100}]
```

### neighbor
```
$ ./gobgp.sh neighbor
Peer              AS  Up/Down State       |#Advertised Received Accepted
10.200.1.101     100 00:20:27 Establ      |         27       13       13
10.200.1.102     100 00:20:56 Establ      |         25       13       13
10.200.1.103     100 00:20:18 Establ      |         27       13       13
10.200.1.104     100 00:20:29 Establ      |         25       13       13

$ ./gobgp.sh neighbor 10.200.1.101
BGP neighbor is 10.200.1.101, remote AS 100
  BGP version 4, remote router ID 10.1.1.1
  BGP state = BGP_FSM_ESTABLISHED, up for 00:20:50
  BGP OutQ = 0, Flops = 0
  Hold time is 0, keepalive interval is 30 seconds
  Configured hold time is 90, keepalive interval is 30 seconds
  Neighbor capabilities:
    BGP_CAP_MULTIPROTOCOL:
        RF_IPv4_UC:     advertised and received
    BGP_CAP_ROUTE_REFRESH:      advertised and received
    BGP_CAP_FOUR_OCTET_AS_NUMBER:       advertised and received
    BGP_CAP_ROUTE_REFRESH_CISCO:        received
  Message statistics:
                         Sent       Rcvd
    Opens:                  1          1
    Notifications:          0          0
    Updates:                9          3
    Keepalives:            42         43
    Route Refesh:           0          0
    Discarded:              0          0
    Total:                 52         47
  Route statistics:
    Advertised:            27
    Received:              13
    Accepted:              13
```

### adj-rin-in/adj-rib-out
```
$ ./gobgp.sh neighbor 10.200.1.101 adj-in
    Network             Next Hop             AS_PATH              Age        Attrs
    10.1.1.1/32         10.200.1.101                              00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    10.1.2.1/32         10.200.1.101         1001                 00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    10.1.2.2/32         10.200.1.101         1002                 00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    10.10.10.0/24       10.200.1.101                              00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    10.200.1.0/24       10.200.1.101                              00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    10.200.2.0/24       10.200.1.101                              00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    10.201.11.0/24      10.200.1.101                              00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    10.201.12.0/24      10.200.1.101                              00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    172.17.0.0/16       10.200.1.101                              00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    172.21.1.0/24       10.200.1.101         1001                 00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    172.21.2.0/24       10.200.1.101         1002                 00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    172.22.1.0/24       10.200.1.101         1001                 00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]
    172.22.2.0/24       10.200.1.101         1002                 00:22:53   [{Origin: ?} {Med: 0} {LocalPref: 100}]

$ ./gobgp.sh neighbor 10.200.1.101 adj-out
    Network             Next Hop             AS_PATH              Attrs
    10.1.1.2/32         10.200.1.102                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    10.1.1.3/32         10.200.1.103                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.3} {ClusterList: [0.0.0.0]}]
    10.1.1.4/32         10.200.1.104                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.4} {ClusterList: [0.0.0.0]}]
    10.1.2.1/32         10.200.1.102         1001                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    10.1.2.2/32         10.200.1.102         1002                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    10.1.2.3/32         10.200.1.103         1003                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.3} {ClusterList: [0.0.0.0]}]
    10.1.2.4/32         10.200.1.103         1004                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.3} {ClusterList: [0.0.0.0]}]
    10.10.10.0/24       10.200.1.102                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    10.200.1.0/24       10.200.1.102                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    10.200.2.0/24       10.200.1.102                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    10.201.11.0/24      10.200.1.102         1001                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    10.201.12.0/24      10.200.1.102         1002                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    10.202.11.0/24      10.200.1.102                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    10.202.12.0/24      10.200.1.102                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    10.203.13.0/24      10.200.1.103                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.3} {ClusterList: [0.0.0.0]}]
    10.203.14.0/24      10.200.1.103                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.3} {ClusterList: [0.0.0.0]}]
    10.204.13.0/24      10.200.1.104                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.4} {ClusterList: [0.0.0.0]}]
    10.204.14.0/24      10.200.1.104                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.4} {ClusterList: [0.0.0.0]}]
    172.17.0.0/16       10.200.1.102                              [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    172.21.1.0/24       10.200.1.102         1001                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    172.21.2.0/24       10.200.1.102         1002                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    172.21.3.0/24       10.200.1.103         1003                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.3} {ClusterList: [0.0.0.0]}]
    172.21.4.0/24       10.200.1.103         1004                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.3} {ClusterList: [0.0.0.0]}]
    172.22.1.0/24       10.200.1.102         1001                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    172.22.2.0/24       10.200.1.102         1002                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.2} {ClusterList: [0.0.0.0]}]
    172.22.3.0/24       10.200.1.103         1003                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.3} {ClusterList: [0.0.0.0]}]
    172.22.4.0/24       10.200.1.103         1004                 [{Origin: ?} {Med: 0} {LocalPref: 100} {Originator: 10.1.1.3} {ClusterList: [0.0.0.0]}]
```
