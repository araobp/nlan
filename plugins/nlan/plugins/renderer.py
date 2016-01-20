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
        PTN-BGP state
        '''
        with open(template) as f:
            temp = f.read()
            print(temp)

            with self.tx() as t:
                routers = t.get('nlan.ip')
                state_yaml = Template(temp).render(**routers)
                print(state_yaml)
                state = yaml.load(state_yaml)
                print(state)
                print('-----dict2cont-----')
                for r, s in state.items():
                    s_cont = dict2cont(s)
                    self.nlan.state[r] = s_cont
                    t.put(s_cont)

