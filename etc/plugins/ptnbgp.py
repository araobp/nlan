import tega.tree
import tega.subscriber

class PtnBgp(tega.subscriber.PlugIn):

    def __init__(self):
        super().__init__()

        plugins = tega.tree.Cont('plugins')
        with self.tx() as t:
            plugins.ptnbgp = self.func(self._state)
            t.put(plugins.ptnbgp, ephemeral=True)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    def _state(self):

        _routers = ['10.10.1.1', 'pe1', 'pe2', 'pe3', 'pe4', 'rr', 'ce1', 'ce2', 'ce3', 'ce4']

        self.rpc('plugins.ipam', *_routers)
        routers = self.get('ip')
        g = globals()
        for router, ip in routers.items():
            g[router] = ip.split('/')[0]
            plugins = tega.tree.Cont('plugins')
        
        # Roots 
        _pe1 = tega.tree.Cont('config-pe1')
        _pe2 = tega.tree.Cont('config-pe2')
        _pe3 = tega.tree.Cont('config-pe3')
        _pe4 = tega.tree.Cont('config-pe4')
        _rr  = tega.tree.Cont('config-rr')
        _ce1 = tega.tree.Cont('config-ce1')
        _ce2 = tega.tree.Cont('config-ce2')
        _ce3 = tega.tree.Cont('config-ce3')
        _ce4 = tega.tree.Cont('config-ce4')

        # Loopback address
        _pe1.Router.Loopback = '10.1.1.1/32'
        _pe2.Router.Loopback = '10.1.1.2/32'
        _pe3.Router.Loopback = '10.1.1.3/32'
        _pe4.Router.Loopback = '10.1.1.4/32'
        _rr.Router.Loopback  = '10.1.1.5/32'
        _ce1.Router.Loopback = '10.1.2.1/32'
        _ce2.Router.Loopback = '10.1.2.2/32'
        _ce3.Router.Loopback = '10.1.2.3/32'
        _ce4.Router.Loopback = '10.1.2.4/32'

        # Nodes (vertices) 
        def _nodes(ptn, l2sw):
            return dict(Ptn=ptn, L2sw=l2sw)
        _pe1.Ptn.wan.Nodes = _nodes('ptpe1w', 'l2spe1w')
        _pe2.Ptn.wan.Nodes = _nodes('ptpe2w', 'l2spe2w')
        _pe3.Ptn.wan.Nodes = _nodes('ptpe3w', 'l2spe3w')
        _pe4.Ptn.wan.Nodes = _nodes('ptpe4w', 'l2spe4w')
        _rr.Ptn.wan.Nodes  = _nodes('ptrrw' , 'l2srrw')
        _pe1.Ptn.access.Nodes = _nodes('ptpe1a', 'l2spe1a')
        _pe2.Ptn.access.Nodes = _nodes('ptpe2a', 'l2spe2a')
        _pe3.Ptn.access.Nodes = _nodes('ptpe3a', 'l2spe3a')
        _pe4.Ptn.access.Nodes = _nodes('ptpe4a', 'l2spe4a')
        _ce1.Ptn.access.Nodes = _nodes('ptce1a', 'l2sce1a')
        _ce2.Ptn.access.Nodes = _nodes('ptce2a', 'l2sce2a')
        _ce3.Ptn.access.Nodes = _nodes('ptce3a', 'l2sce3a')
        _ce4.Ptn.access.Nodes = _nodes('ptce4a', 'l2sce4a')

        # Links (edges)
        def _links(localIp, remoteIps):
            return dict(LocalIp=localIp, RemoteIps=remoteIps)
        _pe1.Ptn.wan.Links = _links(pe1, [pe2, pe3, pe4, rr])
        _pe2.Ptn.wan.Links = _links(pe2, [pe1, pe3, pe4, rr])
        _pe3.Ptn.wan.Links = _links(pe3, [pe1, pe2, pe4, rr])
        _pe4.Ptn.wan.Links = _links(pe4, [pe1, pe2, pe3, rr])
        _rr.Ptn.wan.Links  = _links(rr,  [pe1, pe2, pe3, pe4])
        _pe1.Ptn.access.Links = _links(pe1, [ce1, ce2])
        _pe2.Ptn.access.Links = _links(pe2, [ce1, ce2])
        _pe3.Ptn.access.Links = _links(pe3, [ce3, ce4])
        _pe4.Ptn.access.Links = _links(pe4, [ce3, ce4])
        _ce1.Ptn.access.Links = _links(ce1, [pe1, pe2])
        _ce2.Ptn.access.Links = _links(ce2, [pe1, pe2])
        _ce3.Ptn.access.Links = _links(ce3, [pe3, pe4])
        _ce4.Ptn.access.Links = _links(ce4, [pe3, pe4])
        
        # VPN 
        def _vpn(vid, vni, peers, ip):
            return dict(Vid=vid, Vni=vni, Peers=peers, Ip=ip)
        vpn101 = _vpn(101, 1, [pe2, pe3, pe4, rr],  '10.200.1.101/24')
        vpn102 = _vpn(102, 2, [pe2, pe3, pe4, rr],  '10.200.2.101/24')
        _pe1.Ptn.wan.L2Vpn = [vpn101, vpn102]
        vpn101 = _vpn(101, 1, [pe1, pe3, pe4, rr],  '10.200.1.102/24')
        vpn102 = _vpn(102, 2, [pe1, pe3, pe4, rr],  '10.200.2.102/24')
        _pe2.Ptn.wan.L2Vpn = [vpn101, vpn102]
        vpn101 = _vpn(101, 1, [pe1, pe2, pe4, rr],  '10.200.1.103/24')
        vpn102 = _vpn(102, 2, [pe1, pe2, pe4, rr],  '10.200.2.103/24')
        _pe3.Ptn.wan.L2Vpn = [vpn101, vpn102]
        vpn101 = _vpn(101, 1, [pe1, pe2, pe3, rr],  '10.200.1.104/24')
        vpn102 = _vpn(102, 2, [pe1, pe2, pe3, rr],  '10.200.2.104/24')
        _pe4.Ptn.wan.L2Vpn = [vpn101, vpn102]
        vpn101 = _vpn(101, 1, [pe1, pe2, pe3, pe4], '10.200.1.105/24')
        vpn102 = _vpn(102, 2, [pe1, pe2, pe3, pe4], '10.200.2.105/24')
        _rr.Ptn.wan.L2Vpn = [vpn101, vpn102]
        vpn11 = _vpn(11, 111, [ce1], '10.201.11.1/24')
        vpn12 = _vpn(12, 112, [ce2], '10.201.12.1/24')
        _pe1.Ptn.access.L2Vpn = [vpn11, vpn12]
        vpn11 = _vpn(11, 211, [ce1], '10.202.11.1/24')
        vpn12 = _vpn(12, 212, [ce2], '10.202.12.1/24')
        _pe2.Ptn.access.L2Vpn = [vpn11, vpn12]
        vpn13 = _vpn(13, 313, [ce3], '10.203.13.1/24')
        vpn14 = _vpn(14, 314, [ce4], '10.203.14.1/24')
        _pe3.Ptn.access.L2Vpn = [vpn13, vpn14]
        vpn13 = _vpn(13, 413, [ce3], '10.204.13.1/24')
        vpn14 = _vpn(14, 414, [ce4], '10.204.14.1/24')
        _pe4.Ptn.access.L2Vpn = [vpn13, vpn14]
        vpn111 = _vpn(1, 111, [pe1], '10.201.11.2/24')
        vpn211 = _vpn(2, 211, [pe2], '10.202.11.2/24')
        _ce1.Ptn.access.L2Vpn = [vpn111, vpn211]
        vpn112 = _vpn(1, 112, [pe1], '10.201.12.2/24')
        vpn212 = _vpn(2, 212, [pe2], '10.202.12.2/24')
        _ce2.Ptn.access.L2Vpn = [vpn112, vpn212]
        vpn313 = _vpn(3, 313, [pe3], '10.203.13.2/24')
        vpn413 = _vpn(4, 413, [pe4], '10.204.13.2/24')
        _ce3.Ptn.access.L2Vpn = [vpn313, vpn413]
        vpn314 = _vpn(3, 314, [pe3], '10.203.14.2/24')
        vpn414 = _vpn(4, 414, [pe4], '10.204.14.2/24')
        _ce4.Ptn.access.L2Vpn = [vpn314, vpn414]

        # BGP
        def _neigh(peer, remoteAs, *, nextHopSelf=False, routeReflectorClient=False):
            if nextHopSelf:
                return dict(Peer=peer, RemoteAs=remoteAs, NextHopSelf=True)
            elif routeReflectorClient is True:
                return dict(Peer=peer, RemoteAs=remoteAs, RouteReflectorClient=True)
            else:
                return dict(Peer=peer, RemoteAs=remoteAs)
        n100 =  _neigh('10.200.1.105', 100, nextHopSelf=True)
        n1001 = _neigh('10.201.11.2', 1001)
        n1002 = _neigh('10.201.12.2', 1002)
        _pe1.Router.Bgp['100'].Neighbors = [n100, n1001, n1002]
        n100 =  _neigh('10.200.1.105', 100, nextHopSelf=True)
        n1001 = _neigh('10.202.11.2', 1001)
        n1002 = _neigh('10.202.12.2', 1002)
        _pe2.Router.Bgp['100'].Neighbors = [n100, n1001, n1002]
        n100 =  _neigh('10.200.1.105', 100, nextHopSelf=True)
        n1003 = _neigh('10.203.13.2', 1003)
        n1004 = _neigh('10.203.14.2', 1004)
        _pe3.Router.Bgp['100'].Neighbors = [n100, n1003, n1004]
        n100 =  _neigh('10.200.1.105', 100, nextHopSelf=True)
        n1003 = _neigh('10.204.13.2', 1003)
        n1004 = _neigh('10.204.14.2', 1004)
        _pe4.Router.Bgp['100'].Neighbors = [n100, n1003, n1004]
        n100_101 = _neigh('10.200.1.101', 100, routeReflectorClient=True)
        n100_102 = _neigh('10.200.1.102', 100, routeReflectorClient=True)
        n100_103 = _neigh('10.200.1.103', 100, routeReflectorClient=True)
        n100_104 = _neigh('10.200.1.104', 100, routeReflectorClient=True)
        _rr.Router.Bgp['100'].Neighbors = [n100_101, n100_102, n100_103, n100_104]
        _rr.Router.EmbeddedBgp = True 
        n201 = _neigh('10.201.11.1', 100)
        n202 = _neigh('10.202.11.1', 100)
        _ce1.Router.Bgp['1001'].Neighbors = [n201, n202]
        n201 = _neigh('10.201.12.1', 100)
        n202 = _neigh('10.202.12.1', 100)
        _ce2.Router.Bgp['1002'].Neighbors = [n201, n202]
        n203 = _neigh('10.203.13.1', 100)
        n204 = _neigh('10.204.13.1', 100)
        _ce3.Router.Bgp['1003'].Neighbors = [n203, n204]
        n203 = _neigh('10.203.14.1', 100)
        n204 = _neigh('10.204.14.1', 100)
        _ce4.Router.Bgp['1004'].Neighbors = [n203, n204]

        # Vhosts 
        def _vhosts(*args):
            return [dict(Network=n, Vhosts=2) for n in args]
        _ce1.Vhosts.VhostProps = _vhosts('172.21.1.1/24', '172.22.1.1/24')
        _ce2.Vhosts.VhostProps = _vhosts('172.21.2.1/24', '172.22.2.1/24')
        _ce3.Vhosts.VhostProps = _vhosts('172.21.3.1/24', '172.22.3.1/24') 
        _ce4.Vhosts.VhostProps = _vhosts('172.21.4.1/24', '172.22.4.1/24')

        with self.tx() as t:
           t.put(_pe1)
           t.put(_pe2)
           t.put(_pe3)
           t.put(_pe4)
           t.put(_rr)
           t.put(_ce1)
           t.put(_ce2)
           t.put(_ce3)
           t.put(_ce4)
