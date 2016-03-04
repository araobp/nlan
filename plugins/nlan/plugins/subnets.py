import tega.tree
import tega.subscriber

PATH_REGEX = r'operational-(\w*)\.ip'
GRAPH_SUBNETS_PATH = 'graph.subnets'

class Subnets(tega.subscriber.PlugIn):
    '''
    Graph of subnets
    '''

    def __init__(self):
        super().__init__()
        plugins = tega.tree.Cont('plugins')
        with self.tx() as t:
            plugins.subnets = self.func(self._subnets) 
            t.put(plugins.subnets, ephemeral=True)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    def _subnets(self):
        '''
        Collects data from every operational tree and put graph of subnets
        onto tega db.
        '''
        data = self.get(PATH_REGEX, regex_flag=True)
        addrs = {}
        subnets = {}

        for v in data:
            router = v[2][0][0]
            ip = v[1]
            for addr in ip.addr:
                addrs[addr] = router

        for v in data:
            router = v[2][0][0]
            ip = v[1]
            for subnet, route in ip.route.items():
                if subnet == 'default':
                    continue
                next_hop = route.Via
                if not next_hop:
                    continue
                edge = (router, addrs[next_hop])
                if not subnet in subnets:
                    subnets[subnet] = [edge]
                else:
                    subnets[subnet].append(edge)

        with self.tx() as t:
            t.put(path=GRAPH_SUBNETS_PATH, instance=subnets)

