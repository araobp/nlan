# Tunnel setup sequence diagram

Issue#16: https://github.com/araobp/nlan/issues/16

This sequence is to set up tunnels among containers via macvlan interfaces.

This uses Tega's pubsub feature for coordination among plugin.deploy and NLAN agents.

```
                 
plugin.deploy    tega                                CLL or Jupyter notebook
  |               |                                     |
  |<- REQ RPC ----|<-- REQ RPC -------------------------|
  |-------------->|-- RES RPC ------------------------->|
  |               |                       
  |---- put ----->|(LAUNCHED)     
  |  state.c1     |
  |               |                       
launch containers |                   container w/ NLAN agent
  |               |                       |
netns symbolic    |                       |
link              |                       |
  |               |                       |
  |     (STARTING)|<-- put state.c1 ------|
  |<-- notify ----|     w/ command        |
  |   state.c1    |                       |
  |   command     |                       |
macvlan setup     |                       |
  |               |                       |
  |---- put ----->|(CONFIG)               |
  |               |--- notify state.c1 -->|
  |               |                       V
  |               |                 start config
  |               |                       |
  |               |                   ip add tunnel tun0 mode gre ...
  |               |                       |
```
