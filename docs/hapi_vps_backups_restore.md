## hapi vps backups restore

Restore backup

### Synopsis

This endpoint restores a backup for a specified virtual machine.
The system will then initiate the restore process, which may take some time depending on the size of the backup.

All data on the virtual machine will be overwritten with the data from the backup.

```
hapi vps backups restore <virtual machine ID> <backup ID> [flags]
```

### Options

```
  -h, --help   help for restore
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps backups](hapi_vps_backups.md)	 - VM backup management

