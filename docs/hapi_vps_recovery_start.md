## hapi vps recovery start

Start recovery mode

### Synopsis

This endpoint initiates the recovery mode for a specified virtual machine. 
Recovery mode is a special state that allows users to perform system rescue operations, such as repairing file systems, 
recovering data, or troubleshooting issues that prevent the virtual machine from booting normally.

Virtual machine will boot recovery disk image and original disk image will be mounted in /mnt directory.

```
hapi vps recovery start <virtual machine ID> [flags]
```

### Options

```
  -h, --help              help for start
      --password string   Temporary root password for recovery mode
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps recovery](hapi_vps_recovery.md)	 - Recovery mode management

