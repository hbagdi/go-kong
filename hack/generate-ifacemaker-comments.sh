#!/bin/bash

extractStructName() {
  local filename=$1
  echo "$(grep -Eow "type (.*) service" "$filename" | cut -d' ' -f2)"
}

generateInterfaceFilename() {
  local filename=$1
  output=${filename//.go/}
  echo "${output}_interface.go"
}

processFiles() {
  local filenames=$1

  for filename in $filenames
  do
    ## Get the necessary names
    structName=$(extractStructName "$filename")
    interfaceFilename=$(generateInterfaceFilename "$filename")

    ## Create the 'go generate' comment
    goGenerate="//go:generate ifacemaker -f $filename -s $structName -i ${structName}Interface -p kong -c \"DO NOT EDIT: Auto generated\" -o ${interfaceFilename}"

    ## a blank line for readability
    sed -i "2a\ " "$filename"

    ## Inject the go:generate comment into the file
    sed -i "3i${goGenerate}" "$filename"
  done
}

main() {
  pushd ../kong >> /dev/null || exit

  local filenames
  filenames=$(ls -b *service.go)

  processFiles "$filenames"

  popd >> /dev/null || exit
}

main
