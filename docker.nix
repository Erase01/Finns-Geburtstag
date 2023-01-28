{ pkgs ? import <nixpkgs> {} }:

pkgs.dockerTools.buildImage {
  name = "finn-geburtstag";
  tag = "latest";

  copyToRoot = pkgs.buildEnv {
    name = "finn-geburtstag";
    paths = [ pkgs.go ];
  };

  config = {
    Cmd = [ "go" "run" "." ];
    ExposedPorts = {
      "8080/tcp" = {};
    };
  };
}
