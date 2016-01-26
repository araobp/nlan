##Tega plugins for NLAN

###IP Address Management
```
Function Name: nlan.ipam
Arguments: ip, *routers
Usage example:
[tega: 1] nlan.ipam('10.10.10.1','pe1','pe2','pe3','pe4','rr','ce1','ce2','ce3','ce4')
[tega: 2] get nlan.ip
{ce1: 10.10.10.6/24, ce2: 10.10.10.7/24, ce3: 10.10.10.8/24, ce4: 10.10.10.9/24, pe1: 10.10.10.1/24,
  pe2: 10.10.10.2/24, pe3: 10.10.10.3/24, pe4: 10.10.10.4/24, rr: 10.10.10.5/24}

```
###Template
```
Function Name: nlan.template
Arguments: filename
Usage example:
[tega: 1] nlan.template('ptn-bgp.yaml')
```

###Topo (topology processing)
This plugin subscribes nlan.state and converts nlan.state into network graph (vertexes/edges).
```
'nlan.state' data change notification --> transformation -> vertexes/edges
```
