## hapi vps snapshots create

Create snapshot

### Synopsis

This endpoint creates a snapshot of a specified virtual machine. A snapshot captures the state and data of 
the virtual machine at a specific point in time, allowing users to restore the virtual machine to that state if needed. 
This operation is useful for backup purposes, system recovery, and testing changes without affecting the current state
of the virtual machine.

Creating new snapshot will overwrite the existing snapshot!

```
hapi vps snapshots create <virtual machine ID> [flags]
```

### Options

```
  -h, --help   help for create
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps snapshots](hapi_vps_snapshots.md)	 - Snapshot management

