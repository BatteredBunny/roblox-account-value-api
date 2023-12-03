{
  pkgs,
  default,
  nix2container,
  lib,
}: let
  config = pkgs.writeText "config.toml" ''
    port = "80"
    robux_per_euro = 60
  '';
in
  nix2container.packages.${pkgs.system}.nix2container.buildImage {
    name = "ghcr.io/ayes-web/roblox-account-value-api";
    tag = "latest";

    copyToRoot = pkgs.cacert;

    config = {
      entrypoint = ["${lib.getExe default}" "-c" "${config}"];
      exposed_ports = [80];
    };
  }
