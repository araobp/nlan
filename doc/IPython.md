#Using IPython to interact with NLAN

This is experimental -- I do not have a concrete idea on how I use IPython in this project...

##Set up
[Step 1] Install sqlite3, since jupyter notbook depends on sqlite3.
```
$ apt-get install sqlite3
```
[Step 2] Download Python3.4 and build/install it:
```
$ wget https://www.python.org/ftp/python/3.4.4/Python-3.4.4.tgz
$ tar zxvf Python-3.4.4.tgz
$ cd Python-3.4.4
$ ./configure
$ make
$ make install
```
[Step 3] Install IPython and Jupyter
```
$ pip3 install ipython
$ pip3 install jupyter
```

##Reference
- [IPython](http://ipython.org/)
- [Jupyter documentation](http://jupyter.readthedocs.org/en/latest/)

