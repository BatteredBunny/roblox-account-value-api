{
  inputs = {
    nixpkgs.url = "https://channels.nixos.org/nixos-unstable/nixexprs.tar.zst";
  };

  outputs =
    {
      self,
      nixpkgs,
      ...
    }:
    let
      inherit (nixpkgs) lib;

      systems = lib.systems.flakeExposed;

      forAllSystems = lib.genAttrs systems;

      nixpkgsFor = forAllSystems (
        system:
        import nixpkgs {
          inherit system;
        }
      );
    in
    {
      overlays.default = final: prev: {
        roblox-account-value-api = final.callPackage ./build.nix { };
      };

      nixosModules.default = import ./module.nix;

      devShells = forAllSystems (
        system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              wire
              go_latest
            ];
          };
        }
      );

      packages = forAllSystems (
        system:
        let
          pkgs = nixpkgsFor.${system};
          overlay = lib.makeScope pkgs.newScope (final: self.overlays.default final pkgs);
        in
        {
          inherit (overlay) roblox-account-value-api;
          default = overlay.roblox-account-value-api;
        }
      );

      checks = forAllSystems (
        system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          test-service = pkgs.callPackage ./test.nix {
            inherit self;
          };
        }
      );
    };
}
