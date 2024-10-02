{
  inputs = {
    nixpkgs.url = "github:cachix/devenv-nixpkgs/rolling";
    systems.url = "github:nix-systems/default";
    devenv.url = "github:cachix/devenv";
    devenv.inputs.nixpkgs.follows = "nixpkgs";
  };

  nixConfig = {
    extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
    extra-substituters = "https://devenv.cachix.org";
  };

  outputs =
    {
      self,
      nixpkgs,
      devenv,
      systems,
      ...
    }@inputs:
    let
      forEachSystem = nixpkgs.lib.genAttrs (import systems);
    in
    {
      packages = forEachSystem (system: {
        devenv-up = self.devShells.${system}.default.config.procfileScript;
      });

      devShells = forEachSystem (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = devenv.lib.mkShell {
            inherit inputs pkgs;
            modules = [
              {
                packages = with pkgs; [
                  git
                  docker
                  docker-compose
                  go-migrate
                ];

                dotenv.enable = true;

                # Main service language
                languages.go = {
                  enable = true;
                  package = pkgs.go_1_22;
                };

                # Main database
                # services.postgres = {
                #   enable = true;
                #   package = pkgs.postgresql_15;
                #   initialDatabases = [ { name = "SongsLib"; } ];
                # };

                # Web-proxy
                # services.nginx.enable = true;

                # Cache for service
                # services.redis.enable = true;

                scripts.hello.exec = ''
                  echo hello from $GREET
                '';

                enterShell = # bash
                  ''
                    echo "Service: Songs library"
                    go version
                  '';

                enterTest = # bash
                  ''
                    echo "Running tests"
                    git --version | grep --color=auto "${pkgs.git.version}"
                  '';
              }
            ];
          };
        }
      );
    };
}
