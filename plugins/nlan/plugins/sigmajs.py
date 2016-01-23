import tega.tree
import tega.subscriber
from tega.subscriber import SCOPE

class SigmaJs(tega.subscriber.PlugIn):
    '''
    "nlan.state to network graph" transformation for Sigma js

    [Reference] http://sigmajs.org/
    '''
    def __init__(self):
        super().__init__()

    def initialize(self):
        self.nlan = tega.tree.Cont('nlan')
        self.subscribe('nlan.state', SCOPE.GLOBAL)

    def on_notify(self, notifications):
        # You write the data transformation script here.
        state = notifications[0]['instance']
        for router, model in state.items():
            print('')
            print('router: {}'.format(router))
            m = model['Ptn']['Networks']
            for n in m:
                id_ = n['Id']
                links = n['Links']
                local_ip = links['LocalIp']
                remote_ips = links['RemoteIps']
                print('local_ip: {}'.format(local_ip))
                print('remote_ips: {}'.format(remote_ips))

    def on_message(self, channel, tega_id, message):
        pass

