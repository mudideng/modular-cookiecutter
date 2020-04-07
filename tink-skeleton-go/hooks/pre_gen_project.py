#!/usr/bin/env python

import re
import sys

NAME_REGEX = r'^tink-[-_a-zA-Z0-9]+$'

module_name = '{{ cookiecutter.name }}'

def main(args):
    if not re.match(NAME_REGEX, module_name):
        print('ERROR: name "%s" does not start with "tink-"' % module_name)
        return 1

    return 0


if __name__ == '__main__':
    sys.exit(main(sys.argv))
