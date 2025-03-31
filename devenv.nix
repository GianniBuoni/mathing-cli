{pkgs, ...}: {
  packages = with pkgs; [
    sqlc
  ];
  languages.go.enable = true;
}
