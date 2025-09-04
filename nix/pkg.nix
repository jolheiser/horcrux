{
  pkgs ? import <nixpkgs> { },
}:
let
  name = "horcrux";
in
pkgs.buildGoModule {
  pname = name;
  version = "main";
  src = pkgs.nix-gitignore.gitignoreSource [ ] (
    builtins.path {
      inherit name;
      path = ../.;
    }
  );
  vendorHash = pkgs.lib.fileContents ../go.mod.sri;
  env.CGO_ENABLED = 0;
  flags = [ "-trimpath" ];
  ldflags = [
    "-s"
    "-w"
    "-extldflags -static"
  ];
  meta = {
    description = "Mirror git repositories across multiple remotes";
    homepage = "https://git.jolheiser.com/horcrux";
    mainProgram = "horcrux";
  };
}
