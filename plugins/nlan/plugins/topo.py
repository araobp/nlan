import tega.tree
import tega.util
import tega.subscriber
from tega.subscriber import SCOPE
from tega.idb import OPE

import json
import logging

class Topo(tega.subscriber.PlugIn):
    '''
    "nlan.state to network graph" transformation.
    '''
    def __init__(self):
        super().__init__()

    def initialize(self):
        self.nlan = tega.tree.Cont('nlan')
        self.subscribe('nlan.state', SCOPE.GLOBAL)

    def on_notify(self, notifications):
        '''
        nlan.state --> vertexs/edges transformation
        '''
        print(notifications)
        notification = notifications[0]
        state = notification['instance']
        ope = notification['ope']
        path = notification['path']

        print(ope)
        print(path)
        if ope == OPE.DELETE.name and path == 'nlan.state': 
            with self.tx() as t:
                t.delete('nlan.topo')
                logging.info('nlan.topo deleted')
        else: 
            if path != 'nlan.state':
                state = idb.get(path='nlan.state')
                logging.info('nlan.state got')
            nodes = [] 
            edges = [] 
            network = {'nodes': nodes, 'edges': edges}

            for router, model in state.items():
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
            cont = tega.util.subtree('nlan.topo', network)
            with self.tx() as t:
                t.put(cont)
                logging.info('nlan.topo put')

    def on_message(self, channel, tega_id, message):
        pass

