{
  description = "Developer Shell";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    hardware.url = "github:nixos/nixos-hardware";
    sops-nix.url = "github:mic92/sops-nix";
    colmena.url = "github:zhaofengli/colmena";
  };

  outputs = {
    self,
    nixpkgs,
    colmena,
    ...
  } @ inputs: {
    colmena = import ./hosts/colmena.nix {inherit nixpkgs inputs;};
    devShell.x86_64-linux = let
      pkgs = nixpkgs.legacyPackages.x86_64-linux;
    in
      pkgs.mkShell {
        packages = with pkgs; [
          go
          pulumi
          pulumiPackages.pulumi-language-go
        ];
      };
  };
}
