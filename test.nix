{ self, testers }:
testers.nixosTest {
  name = "roblox-account-value-api";

  interactive.nodes.machine = {
    virtualisation.forwardPorts = [
      {
        from = "host";
        host.port = 8888;
        guest.port = 8888;
      }
    ];
  };

  nodes.machine =
    { ... }:
    {
      imports = [
        self.nixosModules.default
      ];

      nixpkgs.overlays = [
        self.overlays.default
      ];

      services.roblox-account-value-api = {
        enable = true;
        openFirewall = true;
        settings = {
          port = 8888;
          robux_per_euro = 123;
        };
      };
    };

  testScript =
    { nodes, ... }:
    let
      port = toString nodes.machine.services.roblox-account-value-api.settings.port;
      robux_per_euro = toString nodes.machine.services.roblox-account-value-api.settings.robux_per_euro;
    in
    ''
      start_all()
      machine.wait_for_unit("roblox-account-value-api.service")
      machine.wait_for_open_port(${port})
      machine.succeed("curl -f http://localhost:${port}/")
      machine.succeed("curl -f http://localhost:${port}/api/exchange-rate | grep ${robux_per_euro}")
    '';
}
