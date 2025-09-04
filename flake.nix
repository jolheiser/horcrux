{
  description = "Mirror git repositories across multiple remotes";
  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
  outputs =
    { nixpkgs, ... }:
    let
      systems = [
        "aarch64-darwin"
        "aarch64-linux"
        "x86_64-darwin"
        "x86_64-linux"
      ];
      forAllSystems = f: nixpkgs.lib.genAttrs systems f;
    in
    {
      packages = forAllSystems (system: import ./nix { pkgs = import nixpkgs { inherit system; }; });
      nixosModules.default = import ./nix/module.nix;
    };
}
