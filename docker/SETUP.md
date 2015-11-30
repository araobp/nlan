#Prerequisites

##Open vSwitch installation

```
$wget http://openvswitch.org/releases/openvswitch-2.4.0.tar.gz
```
Follow the instructions included in the archive.

NLAN requires "dkms", "common" and "switch" only. Use dpkg command (dpkg -i) to install the deb packages.

##Working with Docker

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

[Step3] Allow containers to run tcpdump
```
$ mv /usr/sbin/tcpdump /usr/bin/tcpdump
```
(to avoid Permission Denied error)

[Step4] Commit the image
```
$ docker commit <image name> image
```

