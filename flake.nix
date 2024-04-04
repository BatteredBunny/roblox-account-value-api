{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    nix2container.url = "github:nlewo/nix2container";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    nix2container,
    ...
  } @ inputs:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {
          inherit system;
        };

        # Needed on non linux systems since docker runs a linux vm
        dockerCallPackage =
          if pkgs.stdenv.isLinux
          then pkgs.callPackage
          else pkgs.pkgsCross."${pkgs.stdenv.hostPlatform.uname.processor}-multiplatform".callPackage;
      in
        with pkgs; {
          nixosModules.default = import ./module.nix inputs;

          devShells.default = mkShell {
            buildInputs = [
              gnumake
              wire
              go
            ];
          };

          packages = {
            default = callPackage ./build.nix {};
            docker = callPackage ./docker.nix {
              default = dockerCallPackage ./build.nix {};
              inherit nix2container;
            };
          };
        }
    );
}
