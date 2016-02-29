import tega.tree
import tega.subscriber

class Workflows(tega.subscriber.PlugIn):
    '''
    Manages IP addresses of Linux containers
    '''

    def __init__(self):
        super().__init__()

    def initialize(self):
        workflows = tega.tree.Cont('workflows')
        with self.tx() as t:
            workflows.ptnbgp = self.func(self._ptnbgp)
            workflows.ptnospf = self.func(self._ptnospf)
            t.put(workflows, ephemeral=True)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    def _ipam(self):
        ipam_args = ['10.10.10.1','pe1','pe2','pe3','pe4','rr','ce1','ce2','ce3','ce4']
        self.get('plugins.ipam')(*ipam_args)

    def _ptnbgp(self):
        self._ipam()
        self.get('plugins.template')('ptn-bgp.yaml')
        self.get('plugins.deploy')()

    def _ptnospf(self):
        self._ipam()
        self.get('plugins.template')('ptn-ospf.yaml')
        self.get('plugins.deploy')()
