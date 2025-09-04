{
  config,
  lib,
  pkgs,
  ...
}:
let
  cfg = config.services.horcrux;
  pkg = pkgs.callPackage ./pkg.nix { inherit pkgs; };
  jsonFormat = pkgs.formats.json { };
in
{
  options.services.horcrux = {
    enable = lib.mkEnableOption "Horcrux git mirroring";
    package = lib.mkOption {
      type = lib.types.package;
      default = pkg;
      description = "The horcrux package";
    };
    json = lib.mkOption {
      type = lib.types.bool;
      default = false;
      description = "JSON logging";
    };
    debug = lib.mkOption {
      type = lib.types.bool;
      default = false;
      description = "Debug logging";
    };
    storage = lib.mkOption {
      type = lib.types.str;
      default = "/var/lib/horcrux";
      description = "Storage for the mirror repos";
    };
    config = lib.mkOption {
      type = lib.types.attrs;
      default = { };
      description = "Config contents";
    };
  };
  config = lib.mkIf cfg.enable {
    systemd.services.horcrux = {
      description = "Horcrux git mirroring";
      after = [ "network.target" ];
      wantedBy = [ "multi-user.target" ];
      path = [
        pkgs.openssh
        pkgs.git
      ];
      serviceConfig = {
        ExecStart =
          let
            finalConfig = cfg.config // {
              inherit (cfg) storage;
            };
            configFile = pkgs.writeText "horcrux-config.jsonnet" (
              builtins.readFile (jsonFormat.generate "horcrux-config" finalConfig)
            );
            args = [
              "--config=${configFile}"
              (lib.optionalString cfg.json "--json")
              (lib.optionalString cfg.debug "--debug")
            ];
          in
          "${lib.getExe cfg.package} ${lib.concatStringsSep " " args}";
        Restart = "always";
        User = "horcrux";
        Group = "horcrux";
        StateDirectory = "horcrux";
      };
    };
    users = {
      users.horcrux = {
        isSystemUser = true;
        group = "horcrux";
      };
      groups.horcrux = { };
    };
  };
}
