#!/bin/bash
SIGTERM_DELAY="${SIGTERM_DELAY:-60}"
sigterm_handler() {
    echo "Received SIGTERM signal. Delaying termination for ${SIGTERM_DELAY} seconds..."
    sleep $SIGTERM_DELAY
    kill -TERM "$INTERNAL_PROCESS_PID"
    wait "$INTERNAL_PROCESS_PID"
}
trap 'sigterm_handler' SIGTERM
ossfs "$@" &
INTERNAL_PROCESS_PID="$!"
wait "$INTERNAL_PROCESS_PID"

