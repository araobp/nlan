##gRPC test

Models defined in the protocol buffers schema:
- bridges
- vxlan
- subnets

Add/Update/Delete APIs with one of the models.

```
  client             server(NE)
     |                  |
     |------ Add ------>|
     |<-----------------|
     |                  |
     |----- Update ---->|
     |<-----------------|
     |                  |
     |----- Delete ---->|
     |<-----------------|
     |                  |
```

