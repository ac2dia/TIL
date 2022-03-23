# PaaS-TA 배포

## Bosh login
``` shell
cd ~/workspace/paasta-5.5/deployment/bosh-deployment
source login.sh

update cloud config (bosh-lite-cloud-config.yml 수정 필요)
cd ~/workspace/paasta-5.5/deployment/cloud-config
vi bosh-lite-cloud-config.yml
===========================================
* disk_size 가 30720 인 라인에서 공백이 포함되어 있어 제거 필요

as is
  - disk_size: 30720

to be
- disk_size: 30720
===========================================
bosh -e micro-bosh update-cloud-config bosh-lite-cloud-config.yml
```

## runtime config 등록
``` shell
cd ~/workspace/paasta-5.5/deployment/bosh-deployment
./update-runtime-config.sh
bosh -e micro-bosh configs
```

## stemcell 등록
``` shell
cd ~/workspace/paasta-5.5/stemcell/paasta
bosh -e micro-bosh upload-stemcell bosh-stemcell-315.64-warden-boshlite-ubuntu-xenial-go_agent.tgz
bosh -e micro-bosh stemcells
```

## PaaS-TA 배포
``` shell
cd ~/workspace/paasta-5.5/deployment/paasta-deployment
chmod +x ./*.sh
vi deploy-bosh-lite.sh
===========================================
* bosh 명령어 사용시 -e 옵션을 통해 environment 지정 필요

as-is
bosh -d paasta -n deploy paasta-deployment.yml

to-be
bosh -e micro-bosh -d paasta -n deploy paasta-deployment.yml
===========================================
./deploy-bosh-lite.sh
```

## IP route 설정 (bosh-lite를 이용할때)
``` shell
sudo ip route add   10.244.0.0/16 via 10.0.1.6
```
