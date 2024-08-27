{ pkgs
, nix2container
, lib
,
}:
let
  toml = pkgs.formats.toml { };
  config = toml.generate "config.toml" {
    port = "80";
    robux_per_euro = 60;
  };

  # Needed on non linux systems since docker runs a linux vm
  dockerCallPackage =
    if pkgs.stdenv.isLinux
    then pkgs.callPackage
    else pkgs.pkgsCross."${pkgs.stdenv.hostPlatform.uname.processor}-multiplatform".callPackage;

  default = dockerCallPackage ./build.nix { };
in
nix2container.packages.${pkgs.system}.nix2container.buildImage {
  name = "ghcr.io/batteredbunny/roblox-account-value-api";
  tag = "latest";

  copyToRoot = pkgs.cacert;

  config = {
    entrypoint = [ "${lib.getExe default}" "-c" "${config}" ];
    exposed_ports = [ 80 ];
  };
}
