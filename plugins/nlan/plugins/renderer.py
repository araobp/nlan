import tega.tree
import tega.subscriber
from tega.util import dict2cont

from mako.template import Template
import yaml

class Render(tega.subscriber.PlugIn):
    '''
    State renderer
    '''
    def __init__(self):
        super().__init__()

    def initialize(self):
        self.nlan = tega.tree.Cont('nlan')
        with self.tx() as t:
            self.nlan.render = self.func(self._state)  # Attached to nlan.render
            t.put(self.nlan.render)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    def _state(self, template):
        '''
        NLAN state registartion with tega db
        '''
        with open(template) as f:
            temp = f.read()
            print(temp)

            with self.tx() as t:
                routers = t.get('nlan.ip')
                r = {}
                for k, v in routers.items():
                    r[k] = v.split('/')[0]
                state_yaml = Template(temp).render(**r)
                state = yaml.load(state_yaml)
                state_dict = dict(state=yaml.load(state_yaml))
                self.nlan.state = dict2cont(state_dict)
                t.put(self.nlan.state)

