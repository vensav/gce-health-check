
# Usage: ./create_svc.sh <path to gce-health-check binary> <port-to-run>
# Example: ./scripts/create_svc.sh "./build" 1000

# Copy the executable to /usr/local/bin
mv ${1}/gce-health-check-linux* /usr/local/bin/gce-health-check
chmod +x /usr/local/bin/gce-health-check

# Create the service file
echo "[Unit]\n" > /etc/systemd/system/gce-health-check.service
echo "Description=gce health check service\n" >> /etc/systemd/system/gce-health-check.service
echo "After=network-online.target\n\n" >> /etc/systemd/system/gce-health-check.service

echo "[Service]\n" >> /etc/systemd/system/gce-health-check.service
echo "ExecStart=/usr/local/bin/gce-health-check ${2} \n\n" >> /etc/systemd/system/gce-health-check.service

echo "[Install]\n" >> /etc/systemd/system/gce-health-check.service
echo "WantedBy=multi-user.target\n\n" >> /etc/systemd/system/gce-health-check.service

# Load systemd and start the service
systemctl daemon-reload
systemctl enable gce-health-check.service
systemctl start gce-health-check.service
systemctl status gce-health-check.service

