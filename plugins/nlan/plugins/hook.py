import tega.tree
import tega.subscriber

HOOK_PATH_REGEX = r'operational-(\w*)\.ip\.hook'
HOOK_ROUTE_PATH = 'operational-{}.ip.hook.route'
HOOK_ADDR_PATH = 'operational-{}.ip.hook.addr'

HOOK_STATS_REGEX = r'stats-(\w*)\.hook'
HOOK_NETSTAT_PATH = 'stats-{}.hook.netstat'

class Hook(tega.subscriber.PlugIn):
    '''
    Calls hook functions to reflesh operational trees 
    '''

    def __init__(self):
        super().__init__()
        plugins = tega.tree.Cont('plugins')
        with self.tx() as t:
            plugins.hook = self.func(self._hook)  # Attached to plugins.hook
            t.put(plugins.hook, ephemeral=True)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    def _hook(self):
        '''
        Kicks off hook functions in a batch
        '''
        hooks = self.get(path=HOOK_PATH_REGEX, regex_flag=True)
        for l in hooks:
            router = l[2][0][0]
            self.rpc(path=HOOK_ROUTE_PATH.format(router))
            self.rpc(path=HOOK_ADDR_PATH.format(router))

        hooks = self.get(path=HOOK_STATS_REGEX, regex_flag=True)
        for l in hooks:
            router = l[2][0][0]
            self.rpc(path=HOOK_NETSTAT_PATH.format(router))
