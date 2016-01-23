import tega.tree
import tega.subscriber
from tega.subscriber import SCOPE

class SigmaJs(tega.subscriber.PlugIn):
    '''
    "nlan.state to network graph" transformation for Sigma js

    [Reference] http://sigmajs.org/
    '''
    def __init__(self):
        super().__init__()

    def initialize(self):
        self.nlan = tega.tree.Cont('nlan')
        self.subscribe('nlan.state', SCOPE.GLOBAL)

    def on_notify(self, notifications):
        # You write the data transformation script here.
        print(notifications)

    def on_message(self, channel, tega_id, message):
        pass

