```bash
lsblk

# kiểm tra ổ cứng có định dạng chưa
lsblk -f /dev/nvme1n1

# format ổ cứng
sudo mkfs.ext4 /dev/nvme1n1

# mount ổ cứng vào 1 folder
sudo mount /dev/nvme1n1 /data
```

## setting khi reset (reboot) không bị mất mount
- /etc/fstab

```bash
sudo mount /dev/nvme1n1 /data

# lấy uuid
lsblk -o NAME,SIZE,UUID

sudo vim fstab

UUID=5461f479-eb7d-4861-b3b8-8b00ffffa078 /data ext4 defaults,nofail 0 2

sudo systemctl daemon-reload
```

- cập nhật lại size khi tăng size cho volume
```bash
sudo resize2fs /dev/nvme1n1

df -h /data
```

## giảm
- 4GB -> 2GB

- kiểm tra dung lượng ổ đạt yêu cầu để giảm
```bash
df -h /data
# Filesystem      Size  Used Avail Use% Mounted on
# /dev/nvme1n1    3.9G   32K  3.7G   1% /data
```

- đi tạo volume 2GB
- format ổ 2GB
```bash
sudo mkfs.ext4 /dev/nvme2n1

sudo umount /data
```