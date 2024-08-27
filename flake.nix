{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    nix2container.url = "github:nlewo/nix2container";
  };

  outputs =
    { self
    , nixpkgs
    , flake-utils
    , nix2container
    , ...
    } @ inputs: {
      overlays.default = final: prev: {
        roblox-account-value-api = self.packages.default.${final.system};
      };

      nixosModules.default = import ./module.nix;
    }
    //
    (flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      with pkgs; {
        devShells.default = mkShell {
          buildInputs = [
            gnumake
            wire
            go
          ];
        };

        packages = {
          default = callPackage ./build.nix { };
          docker = callPackage ./docker.nix { inherit nix2container; };
        };
      }
    ));
}
