
# Usage: ./create_svc.sh <GOOS_GOARCH> <tag> <port-to-run>
# Example: ./scripts/create_svc.sh " linux_amd64  1.0.5 1000

# Downlod the binary
wget https://github.com/vensav/gce-health-check/releases/download/v${2}/gce-health-check_${2}_${1}.tar.gz
tar -xzf gce-health-check_*.tar.gz

# Copy the executable to /usr/local/bin
mv gce-health-check /usr/local/bin
chmod +x /usr/local/bin/gce-health-check

# Create the service file
echo "[Unit]\n" > /etc/systemd/system/gce-health-check.service
echo "Description=gce health check service\n" >> /etc/systemd/system/gce-health-check.service
echo "After=network-online.target\n\n" >> /etc/systemd/system/gce-health-check.service

echo "[Service]\n" >> /etc/systemd/system/gce-health-check.service
echo "ExecStart=/usr/local/bin/gce-health-check ${3} \n\n" >> /etc/systemd/system/gce-health-check.service

echo "[Install]\n" >> /etc/systemd/system/gce-health-check.service
echo "WantedBy=multi-user.target\n\n" >> /etc/systemd/system/gce-health-check.service

# Load systemd and start the service
systemctl daemon-reload
systemctl enable gce-health-check.service
systemctl start gce-health-check.service
systemctl status gce-health-check.service

# Cleanup
rm -f gce-health-check_*_linux_amd64.tar.gz
