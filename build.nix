{ buildGoModule }:
buildGoModule rec {
  src = ./.;

  name = "roblox-account-value-api";
  vendorHash = "sha256-sz1cfnhwDqIyy1RHxLWvXxBMGR+NX6FhvHrdWtwUOq4=";

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
