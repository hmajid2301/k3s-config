# Home Lab

This is a repository containing all the configuration related to my Home Lab. At the moment this manages a
4x RPI K3S cluster using pulumi.

You can rad more about how it is all setup on my [blog](https://haseebmajid.dev/series/setup-raspberry-pi-cluster-with-k3s-and-nixos/).

## Usage

> Note I copied the kube config file from the cluster and called it config.personal

To apply changes to the cluster.

```bash
export KUBECONFIG=$HOME/.kube/config.personal
pulumi up
```

Teardown and destroy the cluster, at least all the kube resources created by pulumi.

```bash
pulumi down
```

## Images

![Pi Cluster](docs/pi-cluster-front.jpeg)
![Dashboard](docs/dashboard.png)


## More Information

- Hosts Manged by Nix: https://gitlab.com/hmajid2301/dotfiles
- Blog Post: https://haseebmajid.dev/posts/2023-11-18-how-i-setup-my-raspberry-pi-cluster-with-nixos/
