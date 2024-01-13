# build/setup
as long as your config files **do not use relative paths**, this repo should be fully portable on your computer (can be moved anywhere).

## 1. server
in `dev-scripts`, run `build-all.bat`

## 2. web
in `eh-system-web`, follow readme to build website (don't need to do anything with nodejs server)

# usage
```
eh_system.exe -c <config name>
```

config name must reference the name of a file in the config folder, **without the file extension**

recommended to create run scripts like in `scripts` folder

# configuration
config files must be placed in config folder and have file extension `.yml`

see examples for how to edit

# development
see `dev-scripts` for scripts to build and run