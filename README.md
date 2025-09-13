# kvm-vm-iplookup

`kvm-vm-iplookup` is a command-line tool written in Go that retrieves the IP address of a KVM (Kernel-based Virtual Machine) virtual machine based on its name. It reads from a status file typically used by libvirt's dnsmasq.

## Features

- Lookup IP addresses of KVM virtual machines by their names
- Use a default status file location or specify a custom one

## Installation

Build the tool:

```
$ go build
```

## Usage

### Basic usage:

```
$ kvm-vm-iplookup <vm_name>
```

This will look up the IP address of the specified VM using the default status file (`/var/lib/libvirt/dnsmasq/virbr0.status`).

### Specify a custom status file:

```
$ kvm-vm-iplookup -f /path/to/custom/status/file <vm_name>
```

## Example

```
$ kvm-vm-iplookup vm_name
192.168.2.100
```

## License

This project is licensed under the [MIT License](./LICENSE).
