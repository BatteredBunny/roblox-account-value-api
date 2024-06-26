{
  pkgs,
  default,
  nix2container,
  lib,
}: let
  toml = pkgs.formats.toml {};
  config = toml.generate "config.toml" {
    port = "80";
    robux_per_euro = 60;
  };
in
  nix2container.packages.${pkgs.system}.nix2container.buildImage {
    name = "ghcr.io/batteredbunny/roblox-account-value-api";
    tag = "latest";

    copyToRoot = pkgs.cacert;

    config = {
      entrypoint = ["${lib.getExe default}" "-c" "${config}"];
      exposed_ports = [80];
    };
  }
