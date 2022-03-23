# BOSH VM (inception) 배포

## PaaS-TA 설치 파일 다운로드

```shell
sudo apt install curl -y
mkdir workspace && cd workspace
curl -Lo paasta-5.5.zip https://nextcloud.paas-ta.org/index.php/s/ARmoNN83XGgt2gg/download
unzip paasta-5.5.zip
rm paasta-5.5.zip  #(zip 파일이 약 8.4GB 용량을 차지하므로 향후 파스-타 설치 후 vboxmange 로 save 정상적 수행 및 bosh 와 paasta 가 기동시 깨지지 않도록 하기 위함)
```

## Bosh 설치

```shell
curl -Lo ./bosh https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-6.1.0-linux-amd64
chmod +x ./bosh
sudo mv ./bosh /usr/local/bin/bosh
bosh -v
```

## 종속성 파일 설치

```shell
sudo apt-get install -y build-essential zlibc zlib1g-dev ruby ruby-dev openssl libxslt1-dev libxml2-dev libssl-dev libreadline7 libreadline-dev libyaml-dev libsqlite3-dev sqlite3
```

## VirtualBox 6.0 설치

```shell
sudo apt update
sudo apt upgrade
wget -q https://www.virtualbox.org/download/oracle_vbox_2016.asc -O- | sudo apt-key add -
wget -q https://www.virtualbox.org/download/oracle_vbox.asc -O- | sudo apt-key add -
sudo add-apt-repository "deb http://download.virtualbox.org/virtualbox/debian bionic contrib"
sudo apt update
sudo apt install virtualbox-6.0
VBoxManage --version
```

## Deploy Bosh VM

```shell
cd ~/workspace/paasta-5.5/deployment/bosh-deployment
chmod 755 *.sh
./deploy-bosh-lite.sh
```

## Alias 설정

```shell
bosh alias-env micro-bosh -e 10.0.1.6 --ca-cert <(bosh int warden/creds.yml --path /director_ssl/ca)
```

## Bosh login 쉘 스크립트 생성

```shell
vi login.sh
#!/usr/bin

export BOSH_CA_CERT=$(bosh int ./warden/creds.yml --path /director_ssl/ca)
export BOSH_CLIENT=admin
export BOSH_CLIENT_SECRET=$(bosh int ./warden/creds.yml --path /admin_password)
```

## Bosh login 쉘 적용

```shell
source login.sh
```

## Bosh login 확인

```shell
bosh -e micro-bosh env
```

## jumpbox key 생성 (Bosh VM 로그인 용)

```shell
bosh int warden/creds.yml --path /jumpbox_ssh/private_key > jumpbox.key
chmod 600 jumpbox.key
ssh jumpbox@10.0.1.6 -i jumpbox.key
```

## Bosh로 배포된 프로그램 process 확인 (bosh vm or paas-ta vm ssh 접속 후 이용)

```shell
sudo su
monit summary
```

## credhub cli 설치

```shell
cd ~
wget https://github.com/cloudfoundry-incubator/credhub-cli/releases/download/2.0.0/credhub-linux-2.0.0.tgz
tar -xvf credhub-linux-2.0.0.tgz
chmod +x credhub
sudo mv credhub /usr/local/bin/credhub
credhub --version
```

## credhub shell 생성

```shell
cd ~/workspace/paasta-5.5/deployment/bosh-deployment

vi credhub_login.sh
#!/usr/bin

export CREDHUB_CLIENT=credhub-admin
export CREDHUB_SECRET=$(bosh int --path /credhub_admin_client_secret warden/creds.yml)
export CREDHUB_CA_CERT=$(bosh int --path /credhub_tls/ca warden/creds.yml)
```

## credhub shell 쉘 적용

```shell
cd ~/workspace/paasta-5.5/deployment/bosh-deployment
source credhub_login.sh
```

## credhub 로그인

```shell
credhub login -s https://10.0.1.6:8844 --skip-tls-validation
```

## Virtual Box vm 환경 저장

```shell
vboxmanage controlvm $(bosh int warden/state.json --path /current_vm_cid) savestate
```

## Virtual Box vm 복구

```shell
vboxmanage startvm $(bosh int warden/state.json --path /current_vm_cid) --type headless
```
