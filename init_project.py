"""
See source code for cookiecutter here
https://github.com/cookiecutter/cookiecutter/blob/master/cookiecutter/main.py
"""

from cookiecutter.main import cookiecutter

# Global variables
charts_contexts = []



#Prompt users to add their startup configs here
project_name = raw_input("What's the project name?: ")
need_database = raw_input("Do you need a database? [y/n]: ")

database_type = ""
if need_database:
    database_type = raw_input("Choose one [mysql/cassandra]: ")


### Create main service structure

# Create dict context to pass to cookiecutter
extra_context={'project_name': project_name}

# Create Manifest file to define hook resources
# Todo

# Create project from the cookiecutter-pypackage/ template
# Ask for user input to put into context
cookiecutter('./hello-cookie-cutter/',
            no_input=True,
            overwrite_if_exists=True,
            extra_context=extra_context)

### Create database module
if need_database:
    # Ask for user input to put into context
    db_context = {
        'module_name': database_type,
        'file_name': '{}.py'.format(database_type),
        'filecontent': 'Some content from init',
    }
    # Create database from the cookiecutter-pypackage/ template
    cookiecutter('./{}/'.format(database_type),
                no_input=True,
                output_dir=project_name,
                overwrite_if_exists=True,
                extra_context=db_context)

    db_charts_context = { }
    charts_contexts.append(db_charts_context)

### Create whatever module
whatever_module = False
whatever_module_name = "whatever_module"
if whatever_module:
    # Ask for user input to put into context
    whatever_context = {}
    # Create database from the cookiecutter-pypackage/ template
    cookiecutter('./{}/'.format(whatever_module_name),
                no_input=True,
                output_dir=project_name,
                overwrite_if_exists=True,
                extra_context=whatever_context)

    whatever_charts_context = {}




### Create charts

# Todo merge charts contexts
charts_context = {
    'module_name': '.charts',
    'kubernetes_namespace': project_name,
    'file_name':'{}.yaml'.format("development"),
    'database': database_type,
    'database_deployment': "\nSome content from file: \n\tsubcontent: value1",
    }
charts_folder = "charts_folder"
# Create charts from the cookiecutter-pypackage/ template
cookiecutter('./{}/'.format(charts_folder),
            no_input=True,
            output_dir=project_name,
            overwrite_if_exists=True,
            extra_context=charts_context)
