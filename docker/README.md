#Scripts for managing docker image/containers and etcd

##Rebuilding the docker image of NLAN agent
You need to rebuild nlan/agent:ver0.1 image everytime you modify github.com/araobp/go-nlan/nlan/agent.

Just execute the following shell script to rebuild the image and run Linux containers from the image: 
```
$ ./restart.sh
```

##Starting etcd
Make etcd bind 0.0.0.0:
```
$ ./etcd.sh
```

