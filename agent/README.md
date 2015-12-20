#NLAN Agent

NLAN agent runs on each Linxu container.

##Persistency

|command  | persistent | database  |util|
|---------|------------|-----------|----|
|ip       | N          |           |cmd |
|brctl    | N          |           |cmd |
|ovs-vsctl| Y          | ovsdb     |cmdp|
|vtysh    | Y          |/etc/quagga|cmdp|


