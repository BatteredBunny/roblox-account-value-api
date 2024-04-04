inputs: {
  pkgs,
  config ? pkgs.config,
  lib ? pkgs.lib,
  system,
  self,
  ...
}: let
  cfg = config.services.roblox-account-value-api;

  toml = pkgs.formats.toml {};
  tomlSetting = toml.generate "config.toml" cfg.settings;
in {
  options.services.roblox-account-value-api = {
    enable = lib.mkEnableOption "roblox-account-value-api";

    package = lib.mkOption {
      description = "package to use";
      default = inputs.self.packages.${system}.default;
    };

    settings = {
      port = lib.mkOption {
        type = lib.types.int;
        apply = toString;
        description = "port to run http api on";
      };

      robux_per_euro = lib.mkOption {
        type = lib.types.int;
        description = "configure the conversion rate of robux to euro";
        default = 60;
      };
    };
  };

  config = lib.mkIf cfg.enable {
    systemd.services.roblox-account-value-api = {
      enable = true;
      serviceConfig = {
        DynamicUser = true;
        ProtectSystem = "full";
        ProtectHome = "yes";
        DeviceAllow = [""];
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
      wantedBy = ["default.target"];
    };
  };
}
