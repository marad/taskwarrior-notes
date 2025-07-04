#!/bin/bash

read TASK_BEFORE
read TASK
echo "$TASK"

UUID=$(echo "$TASK" | yq .uuid)
STATUS=$(echo "$TASK" | yq .status)

NOTE_PATH=$(twn path --task "$TASK")

if [ ! -f "$NOTE_PATH" ]; then
  echo "Task without a note. Not syncing."
  exit 0
fi

LOGS=$(twn sync --task "$TASK")
RESULT=$?

if [ $RESULT -ne 0 ]; then 
  LOG_PATH="/tmp/task-note-sync.$UUID.log"
  echo "Task sync failed. See $LOG_PATH"
else 
  echo "Task note synced."
fi

