"""
See source code for cookiecutter here
https://github.com/cookiecutter/cookiecutter/blob/master/cookiecutter/main.py
"""

from cookiecutter.main import cookiecutter

#Prompt users to add their startup configs here
project_name = raw_input("What's the project name?: ")

# Create dict context to pass to cookiecutter
extra_context={'project_name': project_name}

# Create Manifest file to define hook resources
# Todo





# Create project from the cookiecutter-pypackage/ template
cookiecutter('./hello-cookie-cutter/',
            # no_input=True,
            overwrite_if_exists=True,
             extra_context=extra_context)

