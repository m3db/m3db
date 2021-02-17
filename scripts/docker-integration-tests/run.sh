#!/usr/bin/env bash

set -ex

TESTS=(
	scripts/docker-integration-tests/prometheus/test.sh
)

# Some systems, including our default Buildkite hosts, don't come with netcat
# installed and we may not have perms to install it. "Install" it in the worst
# possible way.
if ! command -v nc && [[ "$BUILDKITE" == "true" ]]; then
	echo "installing netcat"
	NCDIR="$(mktemp -d)"

	yumdownloader --destdir "$NCDIR" --resolve nc
	(
		cd "$NCDIR"
		RPM=$(find . -maxdepth 1 -name '*.rpm' | tail -n1)
		rpm2cpio "$RPM" | cpio -id
	)

	export PATH="$PATH:$NCDIR/usr/bin"

	function cleanup_nc() {
		rm -rf "$NCDIR"
	}

	trap cleanup_nc EXIT
fi

if [[ -z "$SKIP_SETUP" ]] || [[ "$SKIP_SETUP" == "false" ]]; then
	scripts/docker-integration-tests/setup.sh
fi

NUM_TESTS=${#TESTS[@]}
MIN_IDX=$((NUM_TESTS*BUILDKITE_PARALLEL_JOB/BUILDKITE_PARALLEL_JOB_COUNT))
MAX_IDX=$(((NUM_TESTS*(BUILDKITE_PARALLEL_JOB+1)/BUILDKITE_PARALLEL_JOB_COUNT)-1))

ITER=0
for test in "${TESTS[@]}"; do
	if [[ $ITER -ge $MIN_IDX && $ITER -le $MAX_IDX ]]; then
		# Ensure all docker containers have been stopped so we don't run into issues
		# trying to bind ports.
		docker rm -f $(docker ps -aq) 2>/dev/null || true
		echo "----------------------------------------------"
		echo "running $test"
		if ! (export M3_PATH=$(pwd) && $test); then
			echo "--- :bk-status-failed: $test FAILED"
			exit 1
		fi
	fi
	ITER="$((ITER+1))"
done
