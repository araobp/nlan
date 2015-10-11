#!/usr/bin/env python

import json
import yaml
import sys 

if __name__ == "__main__":
    file_in = sys.argv[1]
    with open(file_in) as f:
        print(json.dumps(yaml.load(f)))

