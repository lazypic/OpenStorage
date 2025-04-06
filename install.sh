#!/bin/sh

# 설치 디렉토리 설정
INSTALL_DIR="/usr/local/bin"
RC_DIR="/usr/local/etc/rc.d"

# Go 프로그램 빌드
echo "프로그램 빌드 중..."
go build -o openstorage main.go

if [ $? -ne 0 ]; then
    echo "빌드 실패"
    exit 1
fi

# 프로그램 설치
echo "프로그램 설치 중..."
sudo install -m 755 openstorage $INSTALL_DIR/

# rc.d 스크립트 설치
echo "서비스 스크립트 설치 중..."
sudo install -m 555 openstorage $RC_DIR/

# 서비스 활성화
echo "서비스 활성화 중..."
if ! grep -q 'openstorage_enable="YES"' /etc/rc.conf; then
    echo 'openstorage_enable="YES"' | sudo tee -a /etc/rc.conf
fi

# 서비스 시작
echo "서비스 시작 중..."
sudo service openstorage start

echo "설치가 완료되었습니다."
echo "서비스 상태 확인: sudo service openstorage status" 