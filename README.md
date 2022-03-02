# SideEcho
> echo를 사용한 기본 서버 아키텍처를 작성합니다.

## Get Started
```bash
// swagger 업데이트
$ swag i

// 서버 실행
$ go run . -c config/config.toml
```

## 기본 구조

```
client - server - handler - controller or manager
```

* `api/server.go`에서 기본적인 서버 설정을 다룹니다.
* `api/`의 하위 폴더들에서는 routing 별로 폴더를 나누어 api(`server.go`)와 핸들러(`handler.go`)를 작성합니다.
* client의 요청이 들어오면 `server.go`의 routing에 맞게 핸들러의 함수가 실행됩니다.
  이후 핸들러에서 내부적으로 가지고 있는 controller의 함수를 실행시킵니다.
* 메인 서비스 동작은 controller 또는 manager에서 구현하며 핸들러는 요청을 받아 manager 또는 controller로 전달하는 작업까지만 실행합니다.

## Handler, Context

---

* 핸들러는 interface를 넘겨 routing 별 실행 함수를 미리 정하고, 이후 핸들러 struct를 구현합니다. 
* 핸들러는 api 실행에 필요한 controller를 가지고 있으며 요청이 들어올 떄 controller의 함수를 실행시킨다.
* custom context는 request를 처리할 떄 필요한 상태값을 필드로 가집니다.
    * 핸들러는 서버 로직에서 필요한 필드
    * context는 요청의 흐름에서 필요한 값들
    
## Todo

- [ ] Database 관련 모듈 작성
  - [ ] DB 추상화 모듈 작성
  - [ ] ORM 모듈 작성
- [X] swagger 작성
- [ ] Dockerfile 작성
- [ ] test 추가
  - [ ] `*_test.go` 코드 추가
  - [ ] 통합 테스트 추가
  - [ ] CI 적용
