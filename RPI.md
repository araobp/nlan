#Setting up the software on Raspberry Pi

![Raspberry Pi](https://raw.github.com/araobp/neutron-lan/master/misc/rpi.png)

Raspberry Pi 1 Model B

## [Step1] Hypriot
- [Hypriot Docker image](http://blog.hypriot.com/downloads/)
- [Win32 Disk Imager](http://sourceforge.net/projects/win32diskimager/)
- [zenmap](https://nmap.org/)

## [Step2] Python
python2.7 is required for pyang.
```
$ apt-get install python2.7
$ apt-get install python3.4
```
Note: python3.4 is optional.

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
```

Note: it took two hours to complete the building processes.

[Reference] http://dave.cheney.net/2015/09/04/building-go-1-5-on-the-raspberry-pi

## [Step4] protocol buffers
```
$ apt-get clone https://github.com/google/protobuf
$ apt-get install autoconf
$ apt-get install unzip
$ apt-get install libtool
$ cd protobuf
$ ./autogen.sh
$ ./configure
$ make
$ make install
```

## [Step5] Networking modules

### ip-command-related capabilities
netns has already been supported on this kernel, so I do not need to reconfigure the kernel to add netns.

### Linux Bridge

### Open vSwitch

### Quagga

### GoBGP

