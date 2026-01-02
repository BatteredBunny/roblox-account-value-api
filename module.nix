{
  pkgs,
  config ? pkgs.config,
  lib ? pkgs.lib,
  ...
}:
let
  cfg = config.services.roblox-account-value-api;

  toml = pkgs.formats.toml { };
  tomlSetting = toml.generate "config.toml" cfg.settings;
in
{
  options.services.roblox-account-value-api = {
    enable = lib.mkEnableOption "roblox-account-value-api";

    package = lib.mkOption {
      description = "package to use";
      default = pkgs.callPackage ./build.nix { };
    };

    openFirewall = lib.mkOption {
      type = lib.types.bool;
      default = false;
      description = "Whether to open the firewall for the Roblox Account Value API port.";
    };

    settings = {
      port = lib.mkOption {
        type = lib.types.int;
        apply = toString;
        description = "port to run http api on";
      };

      behindReverseProxy =
        lib.mkEnableOption "Enable if setting up the service behind a reverse proxy"
        // {
          default = false;
        };

      robux_per_euro = lib.mkOption {
        type = lib.types.int;
        description = "configure the conversion rate of robux to euro";
        default = 60;
      };
    };
  };

  config = lib.mkIf cfg.enable {
    networking.firewall.allowedTCPPorts = lib.mkIf cfg.openFirewall [
      (lib.strings.toIntBase10 cfg.settings.port) # TODO: kinda dumb refactor me and above settings to toml part
    ];

    systemd.services.roblox-account-value-api = {
      enable = true;
      serviceConfig = {
        DynamicUser = true;
        ProtectSystem = "full";
        ProtectHome = "yes";
        DeviceAllow = [ "" ];
        LockPersonality = true;
        MemoryDenyWriteExecute = true;
        PrivateDevices = true;
        ProtectClock = true;
        ProtectControlGroups = true;
        ProtectHostname = true;
        ProtectKernelLogs = true;
        ProtectKernelModules = true;
        ProtectKernelTunables = true;
        ProtectProc = "invisible";
        RestrictNamespaces = true;
        RestrictRealtime = true;
        RestrictSUIDSGID = true;
        SystemCallArchitectures = "native";
        PrivateUsers = true;
        ExecStart = "${lib.getExe cfg.package} -c=${tomlSetting}";
        Restart = "always";
      };

      environment.GIN_MODE = "release";
      wantedBy = [ "default.target" ];
    };
  };
}
