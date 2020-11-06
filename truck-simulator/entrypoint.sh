#!/usr/bin/env bash

JOB_NAME=`curl --silent ${PROFILE_URL} | pup ${CURRENT_JOB_SELECTOR}`

if [[ ! -z "${JOB_NAME}" ]]
then
  SLACK_MESSAGE=`printf "${CUSTOM_MESSAGE}" "${PROFILE_URL}" "${JOB_NAME}"`
  curl -X POST -H 'Content-type: application/json' --data "{\"text\":\"${SLACK_MESSAGE}\"}" ${SLACK_WEBHOOK_URL}
fi
