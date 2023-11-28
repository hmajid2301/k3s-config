{
  description = "Developer Shell";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }: {
    devShell.x86_64-linux =
      let
        pkgs = nixpkgs.legacyPackages.x86_64-linux;
      in
      pkgs.mkShell {
        packages = with pkgs;[
					go
					pulumi 
					pulumiPackages.pulumi-language-go
        ];
      };
  };
}
