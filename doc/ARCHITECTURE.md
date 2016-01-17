##Current architecture

```                          
                                [Tega CLI]
                Tega plugin          ^
                [nlan.ipam]          |
                     ^        (HTTP/WebSocket)
                     |               |          <= Northbound API (Python APIs or REST/WebSocket)
                     V               V
                   [     tega(tornado)       ]  <= Sort of MD-SAL
                     ^          ^          ^
                     |          |          |    <= Southbound API (REST/WebSocket)
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
