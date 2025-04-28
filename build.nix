{ buildGoModule, stdenv }:
buildGoModule rec {
  src = ./.;

  name = "roblox-account-value-api";
  vendorHash = if stdenv.hostPlatform.isDarwin
               then "sha256-A2BgXQsJ1njVsYtw86KUu4YIpiTspX7SqeZYXmFSWps="
               else "sha256-A2BgXQsJ1njVsYtw86KUu4YIpiTspX7SqeZYXmFSWps=";

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
