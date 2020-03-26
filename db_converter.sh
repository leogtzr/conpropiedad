#!/bin/bash
set -o xtrace
set -o nounset
set -o pipefail

readonly work_dir="$(dirname "$(readlink --canonicalize-existing "${0}")")"
readonly error_words_file_missing=80
readonly words_file="${work_dir}/words.txt"
readonly input_file="${work_dir}/input.txt"

if [[ ! -f "${words_file}" ]]; then
    echo "words file: ${words_file} not found" >&2
    exit ${error_words_file_missing}
fi

if grep --fixed-strings '@' "${words_file}" | \
    sed --expression 's/:/;/g' --expression 's/; /;/g' --expression 's/@/;/g' > \
    "${input_file}"; then

    cp "${input_file}" '/tmp'
fi

exit
