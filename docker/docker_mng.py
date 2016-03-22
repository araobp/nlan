#!/usr/bin/env python3.5

import os
from subprocess import run, PIPE
import sys

IMAGE = 'nlan/agent:ver0.1'  # Docker image name

def _set_symbolic_link(node):
    run(['mkdir', '-p', '/var/run/netns'])
    pid = run(['docker', 'inspect', '--format', '{{.State.Pid}}', node], stdout=PIPE, universal_newlines=True).stdout.rstrip()
    target = '/var/run/netns/' + node
    if os.path.islink(target):
        os.remove(target)
    run(['ln', '-s', '/proc/{}/ns/net'.format(pid), target])
    print('pid ' + pid)

def _run(node):
    run(['docker', 'run', '-i', '-t', '-d', '-v', '/tmp:/var/volume:rw', '--privileged', '--name', node, '--env', 'HOSTNAME='+node, IMAGE])
    _set_symbolic_link(node)

def _rm(node):
    run(['docker', 'rm', node])

def _start(node):
    run(['docker', 'start', node])
    _set_symbolic_link(node)

def _stop(node):
    run(['docker', 'stop', node])

if __name__ == '__main__':
    ope = sys.argv[1]
    nodes = sys.argv[2:]
    for node in nodes:
        globals()['_'+ope](node)

