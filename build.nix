{buildGoModule}:
buildGoModule rec {
  src = ./.;

  name = "roblox-account-value-api";
  vendorHash = "sha256-IqyWj5oZ+yIsJ3DHQGneyQiuXgRxuV9Dx7W8SKpJdQs=";

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
