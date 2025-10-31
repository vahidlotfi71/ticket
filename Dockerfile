# استفاده از ایمیج رسمی Go
FROM golang:latest

# تعیین مسیر کاری داخل کانتینر
WORKDIR /app

# استفاده از Go Proxy جایگزین برای دانلود پکیج‌ها
ENV GOPROXY=https://goproxy.io,direct
ENV GO111MODULE=on

# کپی فایل‌های go.mod و go.sum برای استفاده از کش
COPY go.mod go.sum ./

# دانلود ماژول‌ها (استفاده از کش اگر قبلاً دانلود شده باشد)
RUN go mod download

# کپی بقیه فایل‌های پروژه
COPY . .

# ساخت باینری پروژه
RUN go build -o main .

# باز کردن پورت 8000
EXPOSE 8000

# اجرای فایل ساخته شده
CMD ["./main"]
