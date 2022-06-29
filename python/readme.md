Working with pipenv

```bash
# show deps
pipenv graph

# install deps from lockfile
pipenv install --ignore-pipfile

# see outdated deps
pipenv update --outdated

# update all packages
pipenv update

# update specific package
pipenv update <specific_package>
```
