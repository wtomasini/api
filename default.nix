with import <nixpkgs> {};

stdenv.mkDerivation {

  name = "api";

  buildInputs = with pkgs; [
    zstd
    go
    gnumake
  ];

  shellHook = ''
    export GOPATH=$PWD/.go
    export PATH=$PATH:$GOPATH/bin
    export GOPRIVATE=gitlab.authwise.io/*
  '';

  hardeningDisable = [ "fortify" ];

}


