import tega.tree
import tega.subscriber

class IpAddressManagement(tega.subscriber.PlugIn):
    '''
    Manages IP addresses of Linux containers
    '''

    def __init__(self):
        super().__init__()

    def initialize(self):
        nlan = tega.tree.Cont('nlan')
        with self.tx() as t:
            nlan.ipam = self.func(self._gen)  # Attached to nlan.ipam
            t.put(nlan.ipam, ephemeral=True)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    MASK = '24'

    def _gen(self, ip, *routers):
        '''
        IP address generation
        '''
        nlan = tega.tree.Cont('nlan')
        with self.tx() as t:
            abcd = ip.split('.')
            d = abcd[3]
            abc = abcd[:2]
            for r in routers:
                nextip = abcd[:3]
                nextip.append(d)
                nlan.ip[r] = '{}/{}'.format('.'.join(nextip), self.MASK)
                t.put(nlan.ip[r])
                d = str(int(d) + 1)

