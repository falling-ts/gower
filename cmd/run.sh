#!/bin/bash

echo "# rm -rf /etc/systemd/system/gower.service"
systemctl stop gower
systemctl disable gower
rm -rf /etc/systemd/system/gower.service
systemctl daemon-reload
systemctl reset-failed
echo

echo "# cp gower.service /etc/systemd/system/"
chmod +x gower
cp gower.service /etc/systemd/system/
echo

echo "# systemctl restart gower"
systemctl daemon-reload
systemctl restart gower
systemctl enable gower
