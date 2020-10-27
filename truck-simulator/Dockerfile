FROM golang:latest
RUN go get github.com/ericchiang/pup
ENV CURRENT_JOB_SELECTOR ".current-jobs .job > .title text{}"
ENV CUSTOM_MESSAGE "Someone is delivering %s."
ENV PROFILE_URL "https://www.worldoftrucks.com/en/profile/4992950"
COPY entrypoint.sh /
ENTRYPOINT /entrypoint.sh
