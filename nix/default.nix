{
  pkgs ? import <nixpkgs> { },
}:
let
  pkg = pkgs.callPackage ./pkg.nix { inherit pkgs; };
in
{
  horcrux = pkg;
  default = pkg;
}
