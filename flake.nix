# flake.nix
{
    description = "Simple flake for Orb";

    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
        utils.url = "github:numtide/flake-utils";
    };

    outputs = { self, nixpkgs, utils }:
        utils.lib.eachDefaultSystem (
            system:
        let
            # pkgs = nixpkgs.legacyPackages.${system};
            pkgs = import nixpkgs { inherit system; };
            gotools = with pkgs; [
                go
                gopls
                go-tools
                # delve
            ];
        in {
            devShell = pkgs.mkShell {
                packages = gotools;
                # shellHook = ''
                    #   export MY_VAR="hello world"
                # '';
            };
        }
    );
}
