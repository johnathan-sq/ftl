controller: watchexec -r -e go -i "examples/**" -- ftl-controller --key C01H5BRT09Y07547SETZ4HWRA09 --bind http://localhost:8892
regenerate: watchexec -e yaml -e sql -e proto -i "examples/**" --debounce 1s -- make generate
runner0: watchexec -r -e go -i "examples/**" -- ftl-runner --key R01H5BTS6ABP1EHGZSAGJMBV50A --language go --bind http://localhost:8894
runner1: watchexec -r -e go -i "examples/**" -- ftl-runner --key R01H5BTSGKQ8AZ9S22N9N1SM9HV --language go --bind http://localhost:8895