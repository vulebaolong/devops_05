```bash
# tạo máy
docker run -d -it --name ubuntu-demo ubuntu:26.04

# truy cập
docker exec -it ubuntu-demo bash

ls
ls -l 
# -l: list danh sách

ls -la
# a: all, hiển thị cả những file ẩn (dấu chấm đăng trước)

cd
# change directory: di chuyển
# không quy định nơi tới chỉ gõ mỗi cd: tự thêm default "~"
# ~ thư mục home/desktop của user hiên tại

pwd
# print working directory

mkdir
# make directory
# tạo folder/directory (đường dẫn)

rm -rf ten_file/folder
# remove: dùng để xoá folder hoặc file
# -r: recursive: xoá thư mục và toàn bộ file/thư mục bên trong
# -f: force: ép xoá, không cần hỏi xác nhận, và bỏ qua lỗi file không tồn tại

cat, vim, nano
# công cụ editor

cp
# copy

# copy file
cp ten_file_cu ten_file_moi

# copy folder
cp -r ten_folder_cu ten_folder_moi
# -r: recursive copy toàn bộ thư mục con và file bên trong


mv
# move
# dùng để di chuyển hoặc đổi tên file/folder
mv ten_file/folder_cu ten_file/folder_moi

find noi_bat_dau -name *.config
# tìm kiếm file/folder

df -h
# disk free coi dung lượng của máy
# -h: human-readable hiển thị dễ dọc hơn

du -sh *
# disk useage
# -s: summary tóm tắt

free -h
# Theo dõi RAM

top
# Theo dõi realtime
# nhấm phím "m" chuyển đổi giao diện xem RAM

kill number_PID
# tắt đàng hoàng, tuần tự

kill -9 number_PID
# ép buộc kill

Ctrl + R
# Tìm kiếm (reverse-i-search) tìm kiếm đã gõ những lệnh nào

grep -i
# tìm từ khoá, sử dụng regular (biểu thức chính quy)
# -i: không phân biệt hoa thường
# -n: hiển thị số dòng

history
# show lại tất cả lịch sử gõ lệnh

!100
# 100: số thư tự của history

!!
# chạy lệnh cũ

tail -f
# xem log
# -f: folow theo dõi realtime

# /dev/null
# thùng rác, hố đen

apt update && apt install -y vim
# -y: tự động chọn yes
# ERROR: permission deined => thêm sudo ở đầu "sudo apt update && sudo apt install vim"

vim TEN_FILE

:wq
# w: write (save)
# q: quit

:q!
# thoát và không sửa

# xoá nhanh tất cả
gg: đưa trỏ chuột lên trên đầu
dG: xoá tất cả đằng sau chỏ chuột
```