FROM nginx:latest

# SSL 인증서 디렉토리 생성
RUN mkdir -p /etc/nginx/ssl

# SSL 인증서 생성
RUN openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/nginx/ssl/nginx.key -out /etc/nginx/ssl/nginx.crt -subj "/CN=localhost"

# Nginx 설정 파일 복사
COPY nginx.conf /etc/nginx/nginx.conf
