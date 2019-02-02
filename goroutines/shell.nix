{ pkgs ? import <nixpkgs> {} }:
with pkgs; mkShell {
    name = "Go";
    buildInputs = [ go_1_11
                    tmux
                  ];
    shellHook = ''
        export GOPATH=`pwd`

        gofmts() {
            gofmt -w -s -e $1

            if (( $? == 0 )); then
                awk '{ gsub(/\t/, "    "); print }' < $1 > tmp
                cat tmp > $1
                rm tmp
            fi
        }

        export -f gofmts
    '';
}
