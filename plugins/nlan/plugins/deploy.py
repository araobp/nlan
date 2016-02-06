import tega.tree
import tega.subscriber

import os
import subprocess

class Deployment(tega.subscriber.PlugIn):
    '''
    Deployment of NLAN services
    '''


    def __init__(self):
        super().__init__()

    def initialize(self):
        try:
            self.script = os.environ['SETUP_SCRIPT']
        except:
            self.script = 'setup.sh'
        self.scriptdir = os.path.join(os.environ['GOPATH'], 'src/github.com/araobp/nlan/')
        nlan = tega.tree.Cont('nlan')
        with self.tx() as t:
            nlan.deploy = self.func(self._deploy)  # Attached to nlan.deploy
            t.put(nlan.deploy, ephemeral=True)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    def _deploy(self):
        '''
        Deployment
        '''
        os.chdir(self.scriptdir)
        self.process = subprocess.Popen(
                [os.path.join(self.scriptdir, self.script),
                    self.scriptdir],
                preexec_fn=os.setsid)

