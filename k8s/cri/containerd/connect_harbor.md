# Harbor 연동 이슈!!

## 이슈

- TLS/SSL 적용이 된 Harbor 에서 이미지를 pull 받지 못하는 이슈 발생

## 원인 분석

- Harbor login 가능 여부 확인

  - 가능은 하였으나 사용하는 쿠버네티스 클러스가 docker 컨테이너로 생성한 kind 였음!!
  - 그렇기 때문에 kind 컨테이너 내부로 들어가서 확인했어야 했음
  - 확인 결과 Harbor login 실패

- CRI 확인!
  - containerd
- CLI 확인!
  - crictl 설치되어 있었음!

## containerd private registry 설정

1. /etc/hosts 파일에 도메인 추가

- ex) 192.168.130.101 harbor.ac2dia.com

2. containerd 설정 파일 확인

- /etc/containerd/config.toml 이용
  - Configure Image Registry 설정!
  - harbor registry 등록
  - harbor 접속 정보 입력
  - harbor TLS/SSL ca, cert, key 파일

## CRI 아키텍처 학습

![CRI 아키텍처](./cri_architecture.png)

## 참고문헌

[1] 02. ContainerD private Registry image pull, https://ikcoo.tistory.com/230
[2] Configure Image Registry, https://github.com/containerd/containerd/blob/main/docs/cri/registry.md
[3] Architecture of The CRI Plugin, https://github.com/containerd/containerd/blob/main/docs/cri/architecture.md
