#Setting up the software on Raspberry Pi

2015/12/27

![Raspberry Pi](https://raw.github.com/araobp/neutron-lan/master/misc/rpi.png)

Raspberry Pi 1 Model B

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
export PATH=$PATH:$GOROOT/bin
export GOPATH=$PATH/work
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
```

Note: it took hours...

## [Step5] YANG

```
$ pip install pyang
$ go get github.com/openconfig/goyang
```

## [Step6] Networking modules

### ip-command-related capabilities
netns has already been supported on this kernel, so I do not need to reconfigure the kernel to add netns.

### Linux Bridge

### Open vSwitch

### Quagga

### GoBGP

