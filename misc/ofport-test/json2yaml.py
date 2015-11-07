#!/usr/bin/env python3.4
import json
import yaml
import sys

if __name__ == '__main__':
    with open(sys.argv[1]) as f:
        data = json.load(f)
        print(yaml.dump(data))
