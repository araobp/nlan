##DATABASE

etcd is not suitable for my project: etcd is not light-weight.

tega runs much lighter than etcd and it supports an interactive CLI for data manipulation.

I consider migrating from etcd to tega.

###tega installation

```
$ cd
$ git clone http://github.com/araobp/tega
$ pip3 install tornado
$ pip3 install httplib2
$ pip3 install pyyaml
$ mkdir ~/tega/var
```

Append the following line to your ~/.bashrc
```
export PYTHONPATH=$PYTHONPATH:$HOME/tega
```

Start tega server like this:
```
$ cd ~/tega/scripts
$ ./global
```

Python3.4 users also require the following package:
```
$ pip3 install backports_abc
```

Test tega CLI:
```
$ cd ~/tega/scripts
$ ./cli
```
