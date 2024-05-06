{
  nixpkgs,
  inputs,
  ...
}: {
  meta = {
    nixpkgs = import nixpkgs {
      system = "x86_64-linux";
    };
    specialArgs = inputs;
  };

  defaults = {pkgs, ...}: {
    imports = [
      inputs.hardware.nixosModules.raspberry-pi-4
      inputs.sops-nix.nixosModules.sops
      ./rpis/common.nix
    ];
  };

  one = {
    imports = [
      ./rpis/one.nix
    ];

    nixpkgs.system = "aarch64-linux";
    deployment = {
      buildOnTarget = true;
      targetHost = "one";
      targetUser = "one";
      tags = ["rpi"];
    };
  };

  two = {
    imports = [
      ./rpis/two.nix
    ];

    nixpkgs.system = "aarch64-linux";
    deployment = {
      buildOnTarget = true;
      targetHost = "two";
      targetUser = "two";
      tags = ["infra" "rpi"];
    };
  };

  three = {
    imports = [
      ./rpis/three.nix
    ];

    nixpkgs.system = "aarch64-linux";
    deployment = {
      buildOnTarget = true;
      targetHost = "three";
      targetUser = "three";
      tags = ["infra" "rpi"];
    };
  };

  four = {
    imports = [
      ./rpis/four.nix
    ];

    nixpkgs.system = "aarch64-linux";
    deployment = {
      targetHost = "four";
      targetUser = "four";
      tags = ["infra" "rpi"];
    };
  };
}
