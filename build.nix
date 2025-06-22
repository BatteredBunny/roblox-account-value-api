{ buildGoModule }:
buildGoModule rec {
  src = ./.;

  name = "roblox-account-value-api";
  vendorHash = "sha256-h/izD5yFfGaM54CxSgq2dJim24xz2YrLMP/NEq0wiDY=";

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
