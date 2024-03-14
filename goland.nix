with import <nixpkgs> {};

stdenv.mkDerivation {

  name = "core";

  buildInputs = with pkgs; [
    jetbrains.goland
  ];

  shellHook = ''
    goland .
    exit
  '';
}

