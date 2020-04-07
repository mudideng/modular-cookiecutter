#!/usr/bin/env python

import sys
import subprocess

module_name = '{{ cookiecutter.name }}'

def main(args):
    {% if cookiecutter.git_init | int %}
    try:
        p = subprocess.Popen(["git", "init"])
        p.communicate()

        p = subprocess.Popen(["git", "add", "."])
        p.communicate()

        p = subprocess.Popen(["git", "commit", "-m", "chore: Create '{}' from cookiecutter template".format(module_name)])
        p.communicate()

        p = subprocess.Popen(["git", "remote", "add", "upstream", "git@github.com:tink-ab/{}.git".format(module_name)])
        p.communicate()
    except:
        print("Unexpected error:", sys.exc_info()[0])
        return 1
    {% endif %}
    return 0


if __name__ == '__main__':
    sys.exit(main(sys.argv))
