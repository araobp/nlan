import tega.tree
import tega.subscriber

PATH_REGEX = r'stats-(\w*)\.netstat'
GRAPH_SERVERS_CLIENTS_PATH = 'graph.server_client.{}'

class ServerClient(tega.subscriber.PlugIn):
    '''
    Graph of servers-clients
    '''

    def __init__(self):
        super().__init__()
        plugins = tega.tree.Cont('plugins')
        with self.tx() as t:
            plugins.server_client = self.func(self._server_client) 
            t.put(plugins.server_client, ephemeral=True)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    def _server_client(self, port):
        '''
        Collects data from every stats tree and put graph of servers-clientss
        onto tega db.
        '''
        data = self.get(PATH_REGEX, regex_flag=True)
        server_client = []

        for v in data:
            router = v[2][0][0]
            netstat = v[1]
            for elm in netstat.tcp:
                port_ = elm['ForeignPort']
                if port == port_:
                    server_client.append((elm['Ip'], elm['ForeignIp']))

        with self.tx() as t:
            graph = tega.idb.Cont('graph')
            subtree = graph.server_client
            subtree[str(port)] = server_client
            t.put(subtree[str(port)])

