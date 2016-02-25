import tega.tree
import tega.util
import tega.subscriber
from tega.subscriber import SCOPE
from tega.idb import OPE

import json
import logging

CONFIG_PATH = r'config-.*'

class Topo(tega.subscriber.PlugIn):
    '''
    "nlan.state to network graph" transformation.
    '''
    def __init__(self):
        super().__init__()

    def initialize(self):
        self.plugins = tega.tree.Cont('plugins')
        self.subscribe(CONFIG_PATH, SCOPE.GLOBAL, regex_flag=True)

    def on_notify(self, notifications):
        '''
        nlan.state --> vertexs/edges transformation
        '''
        n_put = []
        n_delete = []
        logging.debug(notifications)
        for n in notifications:
            state = notification['instance']
            ope = notification['ope']
            path = notification['path']
            if ope == 'PUT':
                n_put.append([path, state])
            elif ope == 'DELETE':
                # TODO: implementation
                pass

        nodes = [] 
        edges = [] 

        for router, model in n_put:
            m = model['Ptn']['Networks']
            seq = 0
            local_ip=''
            for n in m:
                id_ = n['Id']
                links = n['Links']
                local_ip = links['LocalIp']
                remote_ips = links['RemoteIps']
                for remote_ip in remote_ips:
                    edges.append(dict(id='n'+str(seq), source=local_ip, target=remote_ip))
                    seq += 1
            nodes.append(dict(id=local_ip))

        # Puts the transformed data on tega db
        with self.tx() as t:
            t.put(path='topo', instance=dict(nodes=nodes, edges=edges))
            logging.info('topo put')

    def on_message(self, channel, tega_id, message):
        pass

