# kk

Command shortcuts from yaml.
They work in every shell so no alias or fish functions for simple stuff.

## Installation

```[bash]
go install -ldflags="-s -w"
```

## Configuration

```[yaml] ~/.kk.yaml
shortcuts:
  gp:
    description: "git pull"
    command: git
    arguments:
      - pull
  pg:
    description: "git on the password store"
    command: git
    chdir: ~/.password-store/
  ts:
    description: "task sync"
    command: task
    arguments:
      - sync
```

## Usage

```[bash]
# list shortcuts
kk

# git pull
kk gp

# pull new passwords
kk pg pull

# sync tasks
kk ts
```

## TODO

- craft some advanced usage examples
- print a message when there are no shortcuts defined
- provide shortcut completion (how?)
