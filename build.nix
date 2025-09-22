{ buildGoModule }:
buildGoModule rec {
  src = ./.;

  name = "roblox-account-value-api";
  vendorHash = "sha256-H81yxAe1sJAtCAywqwhT3t32VNKUocEDEcL4GwxN1dc=";

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
