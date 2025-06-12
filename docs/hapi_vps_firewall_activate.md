## hapi vps firewall activate

Activate firewall

### Synopsis

This endpoint activates a firewall for a specified virtual machine.

Only one firewall can be active for a virtual machine at a time.

```
hapi vps firewall activate <firewall ID> <virtual machine ID>  [flags]
```

### Options

```
  -h, --help   help for activate
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hapi.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hapi vps firewall](hapi_vps_firewall.md)	 - Firewall management

