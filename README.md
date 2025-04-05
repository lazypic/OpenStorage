# OpenStorage

디스크 슬롯이 충분하다면, TrueNAS를 설치하는 것이 좋아보이지만 만약 충분하지 않다면 FreeBSD를 이용해서 raidz2 기술을 활용하는 것이 더 좋아보인다.
이 때 중요한것은 디스크의 모니터링이다. 스토리지 서버에 단순한 웹서버를 Go를 이용해(의존성 최소화) 만들고 스토리지의 상태를 json으로 출력되도록 정보를 구성하고 추후 모니터링에서 활용할 수 있도록 구성한다.

## Port

기본적으로 9090을 사용한다.

## 명령어 및 서비스

```
curl "http://localhost:9090/run/zpool/list?args=-Hp"
curl "http://localhost:9090/run/zfs/list?args=-Hp"
```
