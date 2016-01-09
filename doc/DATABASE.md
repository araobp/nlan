##DATABASE

2016/01/09

I happend to find "IPython" at a book store, and I consider using IPython as a frontend for nlan.

I also developed a document-oriented database in Python last year (2015). The database also support an interactive CLI for data manipulation.

What can I realize with the combinaion of nlan, IPython and tega? 

Note: this README is still experimental...

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

###IPython/Jypyter intallation
You may install ipython and jupyter on any linux PC (on which nlan is not running)
```
$ pip3 install ipython
$ pip3 install jupyter
```
Start jupyter notebook
```
$ jupyter notebook
```
