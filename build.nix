{ buildGoModule, stdenv }:
buildGoModule rec {
  src = ./.;

  name = "roblox-account-value-api";
  vendorHash = if stdenv.hostPlatform.isDarwin
               then "sha256-IqyWj5oZ+yIsJ3DHQGneyQiuXgRxuV9Dx7W8SKpJdQs="
               else "sha256-nXQf45F/EcUcXYMDJRztYNvr8lkDoBHuOwEx88orD6U=";

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
