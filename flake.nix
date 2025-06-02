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
      pkgs = nixpkgs.legacyPackages.${system};

      goCore = with pkgs; [
        go
	gopls
	go-tools
        # delve
      ];
    in {
      devShell = pkgs.mkShell {
        packages = goCore;
        # shellHook = ''
        #   export MI_VARIABLE="hola mundo"
        # '';
      };
    }
  );
}
