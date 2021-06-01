#!/bin/bash
# makes a curl request to http://localhost:8080/ping
# in the event it fails, it retries twice more after a short wait period

# initial wait to give api time to start up
sleep 1

# set variables
MAX_RETRY_COUNT=5
RETRY_COUNT=0

# make curl request
until curl http://localhost:8080/ping
    do
        ((RETRY_COUNT++))
        echo "Failed attempt #$RETRY_COUNT"

        # if retry count is above or equal maximum; abort with failure
        if (( $RETRY_COUNT >= $MAX_RETRY_COUNT ))
            then
                echo "Failed $MAX_RETRY_COUNT times, aborting."
                exit 1
        fi

        sleep $(expr $RETRY_COUNT \* $RETRY_COUNT)
done

# success message
echo "Successfully pinged api service."
exit 0
