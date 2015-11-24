#Creating a Docker image for NLAN

##Open vSwitch installation

```
$wget http://openvswitch.org/releases/openvswitch-2.4.0.tar.gz
```
Follow the instructions included in the archive.

NLAN requires "dkms", "common" and "switch" only. Use dpkg command (dpkg -i) to install the deb packages.

##Image creation

![working_with_docker](https://docs.google.com/drawings/d/161Bn80w8JZKQ7BXmIo0br7xQ4kqEdBc_XZ254zuORSU/pub?w=680&h=400)

[Step1] Create an image of Debian/Ubuntu with Open vSwitch installed

You need to copy the following deb packages to the Docker containers:
- openvswitch-switch_*.deb
- openvswitch-common_*.deb

Then "dpkg -i" to install them.

[Step2] Allow ssh root login to the Docker container
```
/etc/ssh/ssh_config

#PermitRootLogin wihtout-password
PermitRootLogin yes
```

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

