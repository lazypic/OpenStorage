# OpenStorage

디스크 슬롯이 충분하다면, 대규모 분산 스토리지로 TrueNAS를 설치하는 것이 좋아보이지만 만약 충분하지 않다면 FreeBSD를 이용해서 raidz2 기술을 활용하는 것이 더 효율적이다.
이 때 중요한것은 여러 스토리지의 상태 모니터링이다. 스토리지 서버에 단순한 웹서버를 Go를 이용해(의존성 최소화) 만들고 스토리지의 상태를 json으로 출력되도록 정보를 구성하고 추후 모니터링에서 활용할 수 있도록 구성한다.

## Port

기본적으로 9090을 사용한다.

## FreeBSD 서비스 설치

1. 설치 스크립트 실행:
```bash
sudo ./install.sh
```

설치 스크립트는 다음 작업을 수행합니다:
- Go 프로그램 빌드
- 바이너리를 `/usr/local/bin`에 설치
- rc.d 스크립트를 `/usr/local/etc/rc.d`에 설치
- 서비스 활성화 및 시작

## 서비스 관리

서비스 관리 명령어:
```bash
# 서비스 시작
sudo service openstorage start

# 서비스 중지
sudo service openstorage stop

# 서비스 재시작
sudo service openstorage restart

# 서비스 상태 확인
sudo service openstorage status
```

## 포트 변경

서비스의 포트를 변경하려면 `/usr/local/etc/rc.d/openstorage` 파일을 수정하세요:
```bash
sudo vi /usr/local/etc/rc.d/openstorage
```

`command_args` 값을 원하는 포트로 변경:
```bash
command_args="-port 8080"  # 원하는 포트로 변경
```

변경 후 서비스를 재시작:
```bash
sudo service openstorage restart
```

## 명령어 및 서비스

```bash
curl "http://localhost:9090"
```

## 상태 가지고오기

```bash
curl -s http://localhost:9090 | jq -r '"Storage Status: \(.errors) | Pool: \(.pool_name) | State: \(.status)"'
```
