{ buildGoModule }:
buildGoModule rec {
  src = ./.;

  name = "roblox-account-value-api";
  vendorHash = "sha256-0upzjZoKiAOoChAeSLEwZbHw73le/qKrE/+8CP6D2Qs=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    homepage = "https://github.com/BatteredBunny/roblox-account-value-api";
    description = "Webapp that calculates account value in robux and euro";
    mainProgram = name;
  };
}
