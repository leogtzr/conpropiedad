#!/bin/bash
set -o xtrace
set -o nounset
set -o pipefail

readonly work_dir="$(dirname "$(readlink --canonicalize-existing "${0}")")"
readonly target_dir="${work_dir}/target"

get_jar_to_run() {
    ls -1 "${target_dir}" | grep --extended-regexp 'jar$'
}

readonly jar_file="${target_dir}/$(get_jar_to_run)"
readonly opts="-Dspring.profiles.active=development"
readonly jar_opts="--input=/tmp/input.txt"

java ${opts} -jar "${jar_file}" ${jar_opts}

exit
