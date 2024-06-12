## NGINX 리버스 프록시 예제 코드

### 구성
- url path를 사용한 구분
    - https://localhost/app1, https://localhost/app2 뒤의 app1, app2로 구분
    - 각 앱에서 엔드포인트 별로 baseUrl을 Prefix 추가해줘야 함 (다른 방법있는지 확인 필요)
- subdomain을 사용한 구분
    - app1.localhost, app2.localhost 처럼 사용
    - 테스트는 성공했지만 ngrok 사용 시 서브 도메인 사용이 제한됨(유료플랜 필요)

### Commands
```bash
# 실행
docker-compose up --build

# 종료
docker-compose down
```
