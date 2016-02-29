import tega.tree
import tega.subscriber
from tega.util import dict2cont

import mako.template
import os
import yaml

IP_PATH = 'ip'
PLUGINS_PATH = 'plugins'
CONFIG_PATH = 'config-{}'

class Template(tega.subscriber.PlugIn):
    '''
    State renderer 
    '''
    def __init__(self):
        super().__init__()

    def initialize(self):
        self.etcdir = os.path.join(os.environ['GOPATH'], 'src/github.com/araobp/nlan/etc/')
        plugins = tega.tree.Cont('plugins')
        with self.tx() as t:
            plugins.template = self.func(self._render)  # Attached to nlan.template
            t.put(plugins.template, ephemeral=True)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    def _render(self, filename):
        '''
        NLAN state registartion with tega db
        '''
        with open(os.path.join(self.etcdir, filename)) as f:
            temp = f.read()

            with self.tx() as t:
                routers = self.get(IP_PATH)
                r = {k: v.split('/')[0] for k, v in routers.items()}
                state_yaml = mako.template.Template(temp).render(**r)
                state = yaml.load(state_yaml)
                config = {CONFIG_PATH.format(r): v for r, v in state.items()}
                for root_oid, c in config.items():
                    t.put(dict2cont({root_oid: c}))

