 #!/bin/bash
 printf "\nRegenerating gqlgen files\n"

 time go run -v github.com/99designs/gqlgen $1

 printf "\nDone.\n\n"
