# อ่านก่อนเริ่มทำงาน
## Commit Message
- [Created] สร้างไฟล์ ชื่อไฟล์
- [Created] สร้างฟังก์ชัน ชื่อฟังก์ชัน
- [Edited] แก้ไขไฟล์ ชื่อไฟล์
- [Edited] แก้ไขฟังก์ชัน ชื่อฟังก์ชัน
- [Removed] ลบไฟล์ ชื่อไฟล์ เพราะ?
- [Removed] ลบฟังก์ชัน ชื่อฟังก์ชัน เพราะ?

## ข้อตกลงในการทำงานร่วมกัน
- **ให้ commit (production code + test code) รวมกันได้เลย**
- git pull บ่อย ๆ
- ถ้าเกิด Conflict ให้แก้ไขและรัน Acceptance Test ให้ผ่านก่อน
- ใช้ Merge Message เดิมได้
- ถ้าจะสั่ง git push ต้องรัน Acceptance Test ในเครื่องตัวเองให้ผ่านก่อนแล้วบอกเพื่อนในทีมด้วย

## วิธีการติดตั้งโปรเจคสำหรับใช้ทำงานร่วมกัน
```bash
# clone project ลงไปในเครื่องก่อน
git clone https://github.com/SCK-SEAL-TEAM-One/chatbot-api.git

# เข้าไปยัง directory project
cd chatbot-api

# รันคำสั่งนี้ก่อน
./install_project.sh

# ติดตั้ง Visual Studio Code ก่อนทำคำสั่งนี้
./open_vscode_with_gopath.sh

```

## Acceptance Test
| Acceptance Test     |                      |                             |                           |                             |
|---------------------|----------------------|-----------------------------|---------------------------|-----------------------------|
| company_id (string) | message (string)     | timestamp (string)          | answer (string)           | timestamp (string)          |
| "1234"              | "หอพักย่านลาดพร้าวซอย1" | "2018-08-12T 03:59:58.855Z" | "ละออเพลส,เรือนสุดสุข,ละอ่อน" | "2018-08-12T 03:59:59.855Z" |
| "1234"              | "หอพักย่านลาดพร้าวซอย2" | "2018-08-12T 04:00:00.000Z" | "ไม่พบข้อมูล"                | "2018-08-12T 04:00:01.000Z" |
| "1234"              | ""                   | "2018-08-12T 10:00:00.000Z" | "โปรดกรอกข้อมูล"            | "2018-08-12T 10:00:01.000Z" |