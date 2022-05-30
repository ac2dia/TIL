# Loadbalancer (Octavia) 설치

- 오픈스택에서는 로드 밸런서 기능을 위해 Octavia 라는 컴포넌트를 제공합니다.
- 해당 문서의 경우 kolla-ansible 을 이용한 오픈스택 배포이며, 버전은 Xena 를 사용하였습니다.

## Enabling Octavia

```yaml
enable_octavia: 'yes'
```

- globals.yml 설정

## Certificates

```yaml
octavia_certs_country: KO
octavia_certs_state: Seoul
octavia_certs_organization: OpenStack
octavia_certs_organizational_unit: Octavia
```

- 인증서 자동 생성 방식 채택

```shell
$ kolla-ansible octavia-certificates
```

- 해당 명령어 입력시 /etc/kolla/config/octavia 경로에 인증서 파일들이 생성됩니다.

## Registering Openstack Resources

### Customize Amphora flavor¶

```yaml
octavia_amp_flavor:
  name: 'amphora'
  is_public: yes
  vcpus: 2
  ram: 2048
  disk: 10
```

- 커스텀하게 설정 가능

### Customise network and subnet¶

```yaml
octavia_amp_network:
  name: lb-mgmt-net
  provider_network_type: vxlan
  external: false
  shared: true
  subnet:
    name: lb-mgmt-subnet
    cidr: '10.1.2.0/24'
    allocation_pool_start: '10.1.2.100'
    allocation_pool_end: '10.1.2.200'
    gateway_ip: '10.1.2.1'
    enable_dhcp: yes
```

```shell
$ kolla-ansible -i <inventory> deploy --tags common,horizon,octavia
```

- 오픈스택 자원들에 대한 설정 후 다시 deploy 를 통해 재배포를 합니다.

## Amphora image¶

```shell
$ sudo apt -y install debootstrap qemu-utils git kpartx
$ git clone https://opendev.org/openstack/octavia -b stable/xena
$ python3 -m venv dib-venv
$ source dib-venv/bin/activate
$ pip install diskimage-builder

$ cd octavia/diskimage-create
$ ./diskimage-create.sh

$ . /etc/kolla/octavia-openrc.sh

$ openstack image create amphora-x64-haproxy.qcow2 --container-format bare --disk-format qcow2 --public --tag amphora --file amphora-x64-haproxy.qcow2
```

## 결론

- 해당 과정들을 마치면 horizon 을 통해 네트워크 > 로드밸런서 탭이 생성됩니다.
- 해당 로드밸런서의 경우 ( 로드 밸런서, 리스너, 풀, 풀 멤버 ) 로 구성되어 있습니다.
- 자동 방식 / 수동 방식이 있는데 편한 방법으로 사용하면 될것 같습니다.
- 배포 중 실패하는 경우가 잦은데 그러한 경우 octavia-worker log 를 참고하시면 트러블 슈팅에 도움이 될 것 같습니다.

# Reference

[1] Octavia Xena, https://docs.openstack.org/kolla-ansible/xena/reference/networking/octavia.html
[2] OpenStack Stein 설치 / Kolla-Ansible 이용 / Ubuntu 18.04, ODROID-H2 Cluster 환경, https://ssup2.github.io/record/OpenStack_Stein_%EC%84%A4%EC%B9%98_Kolla-Ansible_Ubuntu_18.04_ODROID-H2_Cluster/#15-octavia-flavor-keypair-security-group-%EC%84%A4%EC%A0%95-%EB%B0%8F-octavia-%EB%B0%B0%ED%8F%AC
