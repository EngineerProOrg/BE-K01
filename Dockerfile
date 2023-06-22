# Sử dụng ảnh golang:1.16.0-alpine làm base image
FROM golang:latest

# Sao chép mã nguồn vào thư mục /app trong container
COPY Dockerfile docker-compose.yaml go.mod main.go /app/

# Đặt thư mục làm việc mặc định trong container
WORKDIR /app

RUN go get github.com/EngineerProOrg/BE-K01
# Cài đặt dependencies và biên dịch ứng dụng
RUN go mod download && go build -o main .

# Chạy ứng dụng
CMD ["./main"]