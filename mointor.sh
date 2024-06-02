#!/bin/sh

WATCH_DIR=${WATCH_DIR}
JUNIT_FILE="junit.xml"
OUTPUT_FILE=${OUTPUT_FILE}

echo "Monitoring directory: $WATCH_DIR for JUnit file: $JUNIT_FILE"

inotifywait -m -e create "$WATCH_DIR" | while read path action file; do
    if [ "$file" = "$JUNIT_FILE" ]; then
        echo "Detected $JUNIT_FILE creation. Processing $OUTPUT_FILE."

        if [ -f "$OUTPUT_FILE" ]; then
          ./resolve

            echo "Processing completed."
        else
            echo "Output file $OUTPUT_FILE not found!"
        fi

        exit 0
    fi
done
