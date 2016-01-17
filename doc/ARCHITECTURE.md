##Current architecture

```
Tega plugins
[nlan.ipam] <----> [     tega(tornado)       ] <-- (HTTP/WebSocket) --> tega CLI
                     ^          ^          ^
                     |          |          |
                         (HTTP/WebSocket)
                     |          |          |
              . . . . . . . . . . . . . . . . . . .
              .      V          V          V      .
              .   [Agent]    [Agent]    [Agent]   .
              .   Container  Container  Container .
              . . . . . . . . . . . . . . . . . . .
              Docker containers (or physical routers)

tega plugins:
- nlan.ipam: IP Address Management
    :
```
