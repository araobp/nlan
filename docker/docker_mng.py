#!/usr/bin/env python3.4

import sys
import subprocess

IMAGE = 'nlan/agent:ver0.1'  # Docker image name
NODES = ['pe1', 'pe2', 'pe3', 'pe4', 'rr', 'ce1', 'ce2', 'ce3', 'ce4']

#name = lambda prefix, i: '{}{}'.format(prefix, str(i))
run = lambda node: ['docker', 'run', '-i', '-t', '-d', '-v',
                        '/tmp:/var/volume:rw', '--privileged',
                        '--name', node, '--env',
                        'HOSTNAME='+node, IMAGE]
stop = lambda node: ['docker', 'stop', node]
start = lambda node: ['docker', 'start', node]
rm = lambda node: ['docker', 'rm', node]

cmds = dict(run=run, stop=stop, start=start, rm=rm)

def build_command(command):
    '''
    Closure to build a func for executing a command to manage Docker.
    '''
    def exec():
        for node in NODES:
            subprocess.call(command(node))
    return exec

if __name__ == '__main__':

    ope = sys.argv[1]
    build_command(cmds[ope])()
