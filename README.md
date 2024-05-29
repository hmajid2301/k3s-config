# Home Lab

This is a repository containing all the configuration related to my Home Lab.
- 4x Raspberry Pi 4
- 1x Framework 12th Gen Intel Main Board
- 1x Mini Forum

You can read more about how it is all setup on my [blog](https://haseebmajid.dev/series/setup-raspberry-pi-cluster-with-k3s-and-nixos/).
You can find a list of how these devices are [setup using Nix](https://github.com/hmajid2301/dotfiles).

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

## Apps

We have deployed the following apps in our home-lab:

- jellyfin (managed via Nix)
- cert-manager
- kubernetes-dashboard
- gitlab runner


## Images

![Pi Cluster](docs/images/pi-cluster-front.jpeg)
![Dashboard](docs/images/dashboard.png)

## More Information

- Hosts Manged by Nix: https://gitlab.com/hmajid2301/dotfiles
- Blog Post: https://haseebmajid.dev/posts/2023-11-18-how-i-setup-my-raspberry-pi-cluster-with-nixos/
