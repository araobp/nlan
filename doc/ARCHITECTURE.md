##Current architecture

```
[tega plugins] --- [ tega(tornado)   ] --- tega CLI
                     |      |      |
                tega API (HTTP/WebSocket)
                     |      |      |
                     V      V      V
                  [Agent] [Agent] [Agent]

tega plugins:
- nlan.ipam: IP Address Management
    :
```
