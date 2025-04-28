{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    nix2container.url = "github:nlewo/nix2container";
  };

  outputs =
    { self
    , nixpkgs
    , nix2container
    , ...
    }:
    let
      inherit (nixpkgs) lib;

      systems = lib.systems.flakeExposed;

      forAllSystems = lib.genAttrs systems;

      nixpkgsFor = forAllSystems (system: import nixpkgs {
        inherit system;
      });
    in
    {
      overlays.default = final: prev: {
        roblox-account-value-api = self.packages.${final.stdenv.system}.roblox-account-value-api;
      };

      nixosModules.default = import ./module.nix;

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              gnumake
              wire
              go
            ];
          };
        });

      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        rec {
          roblox-account-value-api = default;
          default = pkgs.callPackage ./build.nix { };
          docker = pkgs.callPackage ./docker.nix { inherit nix2container; };
        }
      );
    };
}
