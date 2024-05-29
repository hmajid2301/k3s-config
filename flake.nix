{
  description = "Developer Shell";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = {nixpkgs, ...}: {
    devShell.x86_64-linux = let
      pkgs = nixpkgs.legacyPackages.x86_64-linux;
    in
      pkgs.mkShell {
        shellHook = ''
          export KUBECONFIG=~/.kube/config.personal
        '';
        packages = with pkgs; [
          fluxcd
          kubernetes-helm
          k9s
          kubectl
          sops

          go-task
          mkdocs
          python311Packages.mkdocs-material
        ];
      };
  };
}
