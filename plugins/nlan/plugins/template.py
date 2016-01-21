import tega.tree
import tega.subscriber
from tega.util import dict2cont

import mako.template
import os
import yaml

class Template(tega.subscriber.PlugIn):
    '''
    State renderer 
    '''
    def __init__(self):
        super().__init__()

    def initialize(self):
        self.etcdir = os.path.join(os.environ['GOPATH'], 'src/github.com/araobp/nlan/etc/')
        self.nlan = tega.tree.Cont('nlan')
        with self.tx() as t:
            self.nlan.template = self.func(self._render)  # Attached to nlan.template
            t.put(self.nlan.template)

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
                routers = t.get('nlan.ip')
                r = {}
                for k, v in routers.items():
                    r[k] = v.split('/')[0]
                state_yaml = mako.template.Template(temp).render(**r)
                state = yaml.load(state_yaml)
                state_dict = dict(state=yaml.load(state_yaml))
                self.nlan.state = dict2cont(state_dict)
                t.put(self.nlan.state)

