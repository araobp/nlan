#
# cmdutil.py: command execution utilities. 
# Usage example: python cmd.py 'ls -l' 'check_output'
# Refer to Python documentation: http://docs.python.org/2/library/subprocess.html
#
# 2014/1/29

import subprocess

CalledProcessError = subprocess.CalledProcessError

		
def _cmd(check, persist, *args):

    cmd_args = []
    args = list(args)
    for l in args:
        for ll in l.split():
            cmd_args.append(ll)
    return _cmd2(check=check, persist=persist, args=cmd_args)


# If type(args) is list, use this function.
def _cmd2(check, persist, args):
   
    logger = None
    init = False
    try:
        if 'logger' in __n__:
            logger = __n__['logger']
        if 'init' in __n__ and __n__['init'] == 'start':
            init = True
    except:
        pass

    argstring = ' '.join(args)
    logstr = 'cmd: ' + argstring 

    if persist:
        if init:
            logstr = logstr + ' [SKIPPED...]'
            if logger:
                logger.debug(logstr)
            return

    out = None
    returncode = 0
    def log():
        if logger:
            if out:
                logger.debug('{}\n{}'.format(logstr, out))
            else:
                logger.debug(logstr)
    
    try:
        out = subprocess.check_output(args, stderr=subprocess.STDOUT)
    except CalledProcessError as e:
        if check == 'call':
            returncode = e.returncode 
        else:
            log()
            raise CmdError(argstring, e.returncode, out)
    except Exception as e:
        if check == 'call':
            returncode = e.returncode 
        else:
            log()
            raise CmdError(argstring, 1)

    log()

    if check == 'call':
        return returncode
    elif check == 'check_call':
        return 0
    else:
        return out


class CmdError(Exception):

    def __init__(self, command, returncode, out=None):

        self.message = "Command execution error" 
        self.command = command
        self.returncode = returncode
        self.out = out

    def __str__(self):

        message = ''
        return self.message

# If you can ignore error condition, use this function.
def cmd(*args):
	return _cmd('call', False, *args)

def cmd2(args):
        return _cmd2('call', False, args)

def cmdp(*args):
	return _cmd('call', True, *args)

def cmd2p(args):
        return _cmd2('call', True, args)

# If you want the program to stop in case of error, use this function.	
def check_cmd(*args):
	return _cmd('check_call', False, *args)

def check_cmd2(args):
        return _cmd2('check_call', False, args)

def check_cmdp(*args):
	return _cmd('check_call', True, *args)

def check_cmd2p(args):
        return _cmd2('check_call', True, args)

# If you want to get command output, use this function.
def output_cmd(*args):
	return _cmd('check_output', False, *args)

def output_cmd2(args):
        return _cmd2('check_output', False, args)
	
def output_cmdp(*args):
	return _cmd('check_output', True, *args)

def output_cmd2p(args):
        return _cmd2('check_output', True, args)

