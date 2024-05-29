# Cluster

## Nodes

These are all the nodes in my home lab, all of them are running NixOS you can find their specific config
[here](https://gitlab.com/hmajid2301/dotfiles).

|   Hostname                                         |            Board                                          |               CPU                                    |  RAM        |         Primary GPU
| :---------:                                        | :-------------------------:                               | :----------------------------:                       | :---:       | :-------------------------:
| `FrameworkedUp`                                    | Framework 12th Gen Intel                                  |  i7-1280P                                            | 32GB        | Intel Iris Graphics
| `pi-server-1`                                      | Rapsberry Pi 4B                                           |  Quad core Cortex-A72                                | 4GB         | VideoCore VI @ 500 MHz
| `pi-server-2`                                      | Rapsberry Pi 4B                                           |  Quad core Cortex-A72                                | 4GB         | VideoCore VI @ 500 MHz
| `pi-server-3`                                      | Rapsberry Pi 4B                                           |  Quad core Cortex-A72                                | 8GB         | VideoCore VI @ 500 MHz
| `pi-server-4`                                      | Rapsberry Pi 4B                                           |  Quad core Cortex-A72                                | 8GB         | VideoCore VI @ 500 MHz

## Setup

This repository uses fluxcd for GitOps, to automatically deploy changes made to this repository on the main branch.

## Usage

> Note I copied the kube config file from the cluster and called it config.personal

```bash
git clone git@gitlab.com:hmajid2301/home-lab.git
cd home-lab
direnv allow

# To force a reconcile with the git repo
flux reconcile  source git flux-system
```


## More Information

- Hosts Manged by Nix: https://gitlab.com/hmajid2301/dotfiles
- Blog Post: https://haseebmajid.dev/posts/2023-11-18-how-i-setup-my-raspberry-pi-cluster-with-nixos/
