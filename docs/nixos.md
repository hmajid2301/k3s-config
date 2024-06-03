# NixOS

All of my servers are running NixOS and their config is managed in a separated [repository](https://gitlab.com/hmajid2301/dotfiles).
This includes everything that needs to be managed by the nodes such as how to partition the disks. What file system to
use. This also deploys k3s on each host.

## Why

NixOS is great, I am definitely a fanboy. It is great as it makes roll back very easy if something goes wrong. In fact
deploy-rs will roll back our config if a deployment makes the machine inaccessible i.e. cannot ssh to it.

It also allows us to have our OS config in code, so it can be easily shared across my devices. I don't have to worry
about upgrades, as updates are never done in place.

There are other featues I like such as the cool tooling, `nixos-anywhere`, `disko` and `deploy-rs` to make managing the
servers a lot easier.

## Installation

Will install NixOS, assuming you can ssh to this machine.

```bash
nixos-anywhere --flake '.#mainboard' nixos@192.168.1.6
```
## Deploy

We use [ deploy-rs ](https://github.com/serokell/deploy-rs) to deploy changes to our servers like so:
```bash
deploy .#um790 --hostname um790 --ssh-user nixos
```
