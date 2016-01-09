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
     [ tega(tornado)   ] 
       |      |      |
  tega API (HTTP/WebSocket)
       |      |      |
       V      V      V
   [Agent] [Agent] [Agent]
```

