# kk

Command shortcuts from yaml.
They work in every shell so no ALIAS or fish functions for simple stuff.

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
kk gp
kk pg pull
kk ts
```
