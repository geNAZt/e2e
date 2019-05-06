#!/bin/bash
Xvfb :0 -screen 0 1920x1080x24 &> xvfb.log &
export DISPLAY=:0
sleep 1
xauth generate :0 . trusted
xfce4-session &
x11vnc &


# Start the e2e tester
