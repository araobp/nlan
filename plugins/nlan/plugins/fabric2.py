import tega.tree
import tega.subscriber

class Fabric2(tega.subscriber.PlugIn):
    
    def __init__(self):
        super().__init__()
        
        plugins = tega.tree.Cont('plugins')
        with self.tx() as t:
            plugins.fabric2 = self.func(self._state)
            t.put(plugins.fabric2, ephemeral=True)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    def _state(self):

        routers_leaf = ['lf1', 'lf2', 'lf3', 'lf4', 'lf5', 'lf6']
        routers_spine = ['sp1', 'sp2']
        routers = []
        routers.extend(routers_leaf)
        routers.extend(routers_spine)

        self.rpc('plugins.ipam', '10.10.1.1', *routers)
        routers_ = self.get('ip')
        ip = {} # ip addresses
        r = {}  # roots
        for router, ip_ in routers_.items():
            ip[router] = ip_.split('/')[0]

        # Roots
        for router in routers:
            r[router] = tega.tree.Cont('config-'+router)

        # Loopback address
        i = 0
        for router in routers:
            i += 1
            r[router].Router.Loopback = '10.1.1.{}/32'.format(i)

        # Interfaces
        def _params(mode, local, remote, address):
            return dict(Mode=mode, Local=local, Remote=remote, Address=address)
        i = 0
        for leaf in routers_leaf:
            i += 1
            j = 0
            for spine in routers_spine:
                j += 1
                address = '10.20{}.{}.1/24'.format(i, j)
                dev = '{}-{}'.format(leaf, spine)
                r[leaf].Interfaces[dev] = _params("gre", ip[leaf], ip[spine], address)
                address = '10.20{}.{}.2/24'.format(i, j)
                dev = '{}-{}'.format(spine, leaf)
                r[spine].Interfaces[dev] = _params("gre", ip[spine], ip[leaf], address)

        # BGP
        i = 0
        na = []
        nb = []
        for router in routers_leaf:
            i += 1
            n1 = dict(Peer='10.20{}.1.2'.format(i), RemoteAs=1000)
            n2 = dict(Peer='10.20{}.2.2'.format(i), RemoteAs=1000)
            r[router].Router.Bgp['10{}'.format(i)].Neighbors = [n1, n2]
            na.append(dict(Peer='10.20{}.1.1'.format(i), RemoteAs=100+i, NextHopSelf=True))
            nb.append(dict(Peer='10.20{}.2.1'.format(i), RemoteAs=100+i, NextHopSelf=True))
        r['sp1'].Router.Bgp['1000'].Neighbors = na
        r['sp2'].Router.Bgp['1000'].Neighbors = nb

        # Vhosts
        def _vhosts(*args):
            return [dict(Network=n, Vhosts=2) for n in args]
        i = 0
        for router in routers_leaf:
            i += 1
            r[router].Vhosts.VhostProps = _vhosts('172.21.{}.1/24'.format(i), '172.22.{}.1/24'.format(i))

        # Commit
        with self.tx() as t:
            for router in routers:
                t.put(r[router])
