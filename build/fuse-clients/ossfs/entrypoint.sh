#!/bin/bash
SIGTERM_DELAY="${SIGTERM_DELAY:-60}"
delay_term() {
    echo "Received SIGTERM signal. Delaying termination for ${SIGTERM_DELAY} seconds..."
    while true; do
        if ! kill -0 $INTERNAL_PROCESS_PID 2>/dev/null; then
            echo "ossfs exited. Exiting..."
            exit 0
        fi
        if [ ${SIGTERM_DELAY} -le 0 ]; then
            echo "Timeout reached. Sending SIGTERM to ossfs process..."
            kill -TERM "$INTERNAL_PROCESS_PID"
            wait "$INTERNAL_PROCESS_PID"
            exit 0
        fi
        sleep 1
        ((SIGTERM_DELAY--))
    done
}
sigterm_handler() {
    if [ -z $(findmnt -t fuse.ossfs -n -o TARGET) ]; then
        echo "could not find ossfs mountpoint"
        exit 1
    else
        delay_term
    fi
}
trap 'sigterm_handler' SIGTERM
ossfs "$@" &
INTERNAL_PROCESS_PID="$!"
wait "$INTERNAL_PROCESS_PID"

