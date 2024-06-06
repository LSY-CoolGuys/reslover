#!/bin/sh

WATCH_DIR=${WATCH_DIR}
JUNIT_FILE="junit.xml"
OUTPUT_FILE=${OUTPUT_FILE}


echo "WATCH_DIR= $WATCH_DIR "

if [ ! -d "$WATCH_DIR" ]; then
    mkdir -p "$WATCH_DIR"
    echo "$WATCH_DIR is created"
fi

echo "Monitoring directory: $WATCH_DIR for JUnit file: $JUNIT_FILE"

inotifywait -m -e create "$WATCH_DIR" | while read path action file; do
    if [ "$file" = "$JUNIT_FILE" ]; then
        echo "Detected $JUNIT_FILE creation. Processing $OUTPUT_FILE."
        /bin/sh -c "./resolve"
        exit 0
    fi
done