
{ pkgs }: {
  deps = [
    pkgs.go
    pkgs.gotools
    pkgs.go-tools
    pkgs.nodejs_18
    pkgs.nodePackages.npm
    pkgs.postgresql
    pkgs.git
  ];
}
