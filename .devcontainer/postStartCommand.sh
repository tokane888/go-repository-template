#!/bin/bash

USER=$(whoami)

# host ~/.ssh is mounted at /tmp/.ssh; copy it to the HOME directory.
# mounting directly into HOME fails because the directory already exists.
cp -r /tmp/.ssh "${HOME}"
chown -R "${USER}":"${USER}" "${HOME}"/.ssh
chmod 600 "${HOME}"/.ssh/*
