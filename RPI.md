#Setting up the software on Raspberry Pi

2015/12/27

![Raspberry Pi](https://raw.github.com/araobp/neutron-lan/master/misc/rpi.png)

My Raspberry Pi 1 Model B that I bought in Akihabara in Feb 2014.

## [Step1] Hypriot (Debian Linux with docker pre-installed)
I used 8Gbytes SD memory card. zenmap is to find IP address that DHCP server (on my home gateway) assigned to my Raspberry Pi.
- [Hypriot Docker image](http://blog.hypriot.com/downloads/)
- [Win32 Disk Imager](http://sourceforge.net/projects/win32diskimager/)
- [zenmap](https://nmap.org/)

## [Step2] Python
python2.7 is required for pyang.
```
$ apt-get update
$ apt-get install python2.7
$ apt-get install python-pip
$ apt-get install python3.4
$ apt-get install python3-pip
```
Note: python3.4 and pip3 are optional.

## [Step3] Go
```
$ cd $HOME
$ apt-get install gcc
$ apt-get install bzip2
$ curl http://dave.cheney.net/paste/go-linux-arm-bootstrap-c788a8e.tbz | tar xj
$ curl https://storage.googleapis.com/golang/go1.5.2.src.tar.gz | tar xz
$ export GOROOT_BOOTSTRAP=/root/go-linux-arm-bootstrap
$ cd $HOME/go/src
$ ./make.bash
$ ulimit -s 1024
```
Then append the following two lines to $HOME/.bashrc:
```
export GOROOT=$HOME/go
export GOPATH=$HOME/work
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

Note: it took two hours to complete the building processes.

[Reference] http://dave.cheney.net/2015/09/04/building-go-1-5-on-the-raspberry-pi

## [Step4] protocol buffers
```
$ apt-get clone https://github.com/google/protobuf
$ apt-get install autoconf
$ apt-get install unzip
$ apt-get install libtool
$ apt-get install g++
$ apt-get install make
$ cd protobuf
$ ./autogen.sh
$ ./configure
$ make
$ make install
$ go get github.com/golang/protobuf/proto
$ go get github.com/golang/protobuf/protoc-gen-go
$ cp ~/work/src/github.com/golang/protobuf/protoc-gen-go/protoc-gen-go ~/work/bin
```
Note: it took hours...

Do not forget to append the following line to your .bashrc:
```
export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARLY_PATH
```

## [Step5] YANG

```
$ pip install pyang
$ go get github.com/openconfig/goyang
```

## [Step6] Networking modules

### ip-command-related capabilities
netns has already been supported on this kernel, so I do not need to reconfigure the kernel to add netns.

### Linux Bridge
```
$ apt-get install bridge-utils
```
Confirm that docker0 has already been created:
```
$ brctl show
bridge name     bridge id               STP enabled     interfaces
docker0         8000.024244da82d8       no
```
### GoBGP
```
$ go get github.com/osrg/gpbgp/gobgpd
$ go get github.com/osrg/gpbgp/gobgp
```
Note: GoBGP is optinal -- you may run gobgpd instead of quagga/bgp on Route Reflector container.

### Open vSwitch
Compile and build deb packages:
```
$ wget http://openvswitch.org/releases/openvswitch-2.4.0.tar.gz
$ tar zxvf openvswitch-2.4.0.tar.gz
$ cd openvswitch-2.4.0
$ apt-get install build-essential fakeroot
$ apt-get install debhelper autoconf automake libssl-dev bzip2 openssl graphviz python-all procps python-qt4 python-zopeinterface python-twisted-conch libtool
$ `DEB_BUILD_OPTIONS='parallel=8 nocheck' fakeroot debian/rules binary`
```

Confirm that deb packages have been created:
```
$ cd
$ ls -F
      :
openvswitch-2.4.0/
openvswitch-common_2.4.0-1_armhf.deb
openvswitch-datapath-dkms_2.4.0-1_all.deb
openvswitch-datapath-source_2.4.0-1_all.deb
openvswitch-dbg_2.4.0-1_armhf.deb
openvswitch-ipsec_2.4.0-1_armhf.deb
openvswitch-pki_2.4.0-1_all.deb
openvswitch-switch_2.4.0-1_armhf.deb
openvswitch-test_2.4.0-1_all.deb
openvswitch-testcontroller_2.4.0-1_armhf.deb
openvswitch-vtep_2.4.0-1_armhf.deb
      :
```
Then install part of the deb packages:
```
$ cd
$ apt-get install dkms uuid-runtime
$ dpkg -i openvswitch-common_2.4.0-1_armhf.deb
$ dpkg -i openvswitch-switch_2.4.0-1_armhf.deb
$ dpkg -i openvswitch-datapath-dkms_2.4.0-1_all.deb
```
## [Step7] etcd
```
$ go get github.co
$ cd ~/work/src/github.com/coreos/etcd
$ cp etcd ~/work/bin
$ cd etcdctl
$ go build -o etcdctl
$ cp etcdctl ~/work/bin
```

## [Step8] Pulling rpi-raspbian docker image
```
$ docker pull resin/rpi-raspbian
```

## [Step9] Creating "router" container

### Installing required utilities
```
$ docker run --name base -i -t resin/rpi-raspbian /bin/bash
root@dce29feab2aa:/# apt-get update
root@dce29feab2aa:/# apt-get install ssh
root@dce29feab2aa:/# apt-get install bridge-utils
root@dce29feab2aa:/# apt-get install quagga
root@dce29feab2aa:/# apt-get install vim
root@dce29feab2aa:/# cd
root@dce29feab2aa:/# mkdir bin
```
### Allowing SSH root loging
Append the following to /etc/ssh/sshd_config to allow ssh root login to the Docker container:
```
#PermitRootLogin wihtout-password
PermitRootLogin yes
```
Then
```
$ /etc/init.d/ssh start
```
### Setting up Quagga
```
$ cd /etc/quagga
$ touch zebra.conf
$ touch ospfd.conf
$ touch bgpd.conf
```

Then edit "/etc/quagga/daemons" as follows:
```
zebra=yes
bgpd=yes
ospfd=yes
ospf6d=no
ripd=no
ripngd=no
isisd=no
babeld=no
```

### Copying additional packages and binaries to the container
Copy ovs packages and gobgp to the container:
```
$ ip addr show dev eth0
18: eth0@if19: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.2/16 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 fe80::42:acff:fe11:2/64 scope link
       valid_lft forever preferred_lft forever
```

At the docker host,
```
$ cd
$ scp openvswitch-common_2.4.0-1_armhf.deb root@172.17.0.2:~
$ scp openvswitch-switch_2.4.0-1_armhf.deb root@172.17.0.2:~
$ cd ~/work/bin
$ scp gobgp root@172.17.0.2:~/bin
$ scp gobgpd root@172.17.0.2:~/bin
```

At the container,
```
$ dpkg -i openvswitch-common_2.4.0-1_armhf.deb
$ dpkg -i openvswitch-switch_2.4.0-1_armhf.deb
```
If you encounter dependency problems, try:
```
$ apt-get -f install
```
### Commit the change
```
$ docker commit base router
$ docker images
REPOSITORY           TAG                 IMAGE ID            CREATED             VIRTUAL SIZE
router               latest              47057103372d        6 minutes ago       165.1 MB
resin/rpi-raspbian   latest              e97a8531a526        5 days ago          80.28 MB
hypriot/rpi-swarm    latest              039c550f6208        7 weeks ago         10.92 MB
```
## [Step10] nlan installation
```
$ go get github.com/araobp/nlan
```

## [Step11] Setting up NLAN and creating router containers
```
$ cd ~/work/src/github.com/araobp/nlan
$ ./setup.sh
```

## Dependency hell...
I took a lot of steps to setup the software, because I needed to resolve dependency problems.

I know the reason why a lot of people are migrating to Go and Docker...
