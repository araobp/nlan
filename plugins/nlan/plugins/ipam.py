import tega.tree
import tega.subscriber

class IpAddressManagement(tega.subscriber.PlugIn):
    '''
    Manages IP addresses of Linux containers
    '''

    def __init__(self):
        super().__init__()

    def initialize(self):
        plugins = tega.tree.Cont('plugins')
        with self.tx() as t:
            plugins.ipam = self.func(self._gen)  # Attached to plugins.ipam
            t.put(plugins.ipam, ephemeral=True)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    MASK = '24'

    def _gen(self, addr, *routers):
        '''
        IP address generation
        '''
        ip = tega.tree.Cont('ip')
        with self.tx() as t:
            abcd = addr.split('.')
            d = abcd[3]
            abc = abcd[:2]
            for r in routers:
                nextip = abcd[:3]
                nextip.append(d)
                ip[r] = '{}/{}'.format('.'.join(nextip), self.MASK)
                d = str(int(d) + 1)
            t.put(ip)

