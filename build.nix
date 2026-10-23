{ buildGoLatestModule }:
buildGoLatestModule {
  src = ./.;

  name = "roblox-account-value-api";
  vendorHash = "sha256-8Cs2ejOTQnkJ4kHcVSytEJVgY2W7qGk2451RWHU3nQw=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    homepage = "https://github.com/BatteredBunny/roblox-account-value-api";
    description = "Webapp that calculates account value in robux and euro";
    mainProgram = "roblox-account-value-api";
  };
}
