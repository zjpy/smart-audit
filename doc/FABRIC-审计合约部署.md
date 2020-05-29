#	审计合约部署(FABRIC)

###启动测试链

1.修改github.com/hyperledger/fabric-samples/chaincode-docker-devmode/docker-compose-simple.yaml

在chaincode及cli里修正environment以及volumes：

```dockerfile
environment:
		CORE_PEER_ADDRESSAUTODETECT=true

volumes:
		- ./../chaincode/smart_audit/fabric:/opt/gopath/src/fabric
		- ./../chaincode/smart_audit/core:/opt/gopath/src/core
		- ./../chaincode/smart_audit/oracles:/opt/gopath/src/oracles
```

2.拷贝smart-audit源码至chaincode目录, 启动docker compose

```shell
cd github.com/smart-audit/src
cp -r fabric github.com/hyperledger/fabric-samples/chaincode/smart_audit
cp -r contract github.com/hyperledger/fabric-samples/chaincode/smart_audit
cp -r oracles github.com/hyperledger/fabric-samples/chaincode/smart_audit

cd github.com/hyperledger/fabric-samples/chaincode-docker-devmode
docker-compose -f docker-compose-simple.yaml up
```



###启动时间服务合约

初始及实例化时间服务合约

```shell
docker exec cli peer chaincode install -p chaincodedev/chaincode/smart_audit/fabric/time -n Time -v 0
docker exec cli peer chaincode instantiate -n Time -v 0 -c '{"Args":[]}' -C myc
```



###启动人脸识别服务合约

初始化及实例化人脸识别服务合约

```shell
docker exec cli peer chaincode install -p chaincodedev/chaincode/smart_audit/fabric/face -n FaceRecognize -v 0
docker exec cli peer chaincode instantiate -n FaceRecognize -v 0 -c '{"Args":[]}' -C myc
```



###启动物体识别合约

初始及实例化服务识别服务合约

```shell
docker exec cli peer chaincode install -p chaincodedev/chaincode/smart_audit/fabric/identify -n ObjectRecognize -v 0
docker exec cli peer chaincode instantiate -n ObjectRecognize -v 0 -c '{"Args":[]}' -C myc
```



###启动位置服务合约

初始及实例化位置服务合约

```shell
docker exec cli peer chaincode install -p chaincodedev/chaincode/smart_audit/fabric/location -n Location -v 0
docker exec cli peer chaincode instantiate -n Location -v 0 -c '{"Args":[]}' -C myc
```



###启动审计合约

初始及实例化审计合约

```shell
docker exec -it cli bash
peer chaincode install -p chaincodedev/chaincode/smart_audit/fabric/audit -n audit -v 0
peer chaincode instantiate -n audit -v 0 -c '{"Args":["init","maintainer1","maintainer2"]}' -C myc
```



### 调用及查询

1.查询是否审计维护人员初始化成功

```shell
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["getMaintainers"]}'
```

2审计当事人录入

```shell
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["registerAuditee", "ZhanSan"]}'
```

3.审计规则录入

```shell
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["registerRule", "AND", "Time", "(>= 9) AND (<= 18)", "Location", "IN(39.9 116.3 1000)", "FaceRecognize", "", "ObjectRecognize", ""]}'
```

4.项目录入

```shell
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["registerProject", "POS Audit", "This is a bank project, used by bank employees to check if they did check the POS related bussiness themselfs within the specified time and location","0", "0"]}'
```

5.审计事件新增

 ```shell
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["addEvent", "0", "0", "0", "1589532423", "Time", "2020-05-27T15:04:05.000Z", "Location", "39.901 116.299", "FaceRecognize", "/9j/4SMF...", "ObjectRecognize", "iVBORw0..."]}'
 ```

6询审计当事人

```shell
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["getAuditee", "0"]}'
```

7.查询项目

```shell
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["getProject", "0"]}'
```

8.查询规则

```shell
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["getRule", "0"]}'
```



###快速启动所有合约

```shell
docker exec cli peer chaincode install -p chaincodedev/chaincode/smart_audit/fabric/time -n Time -v 0
docker exec cli peer chaincode instantiate -n Time -v 0 -c '{"Args":[]}' -C myc
docker exec cli peer chaincode install -p chaincodedev/chaincode/smart_audit/fabric/face -n FaceRecognize -v 0
docker exec cli peer chaincode instantiate -n FaceRecognize -v 0 -c '{"Args":[]}' -C myc
docker exec cli peer chaincode install -p chaincodedev/chaincode/smart_audit/fabric/identify -n ObjectRecognize -v 0
docker exec cli peer chaincode instantiate -n ObjectRecognize -v 0 -c '{"Args":[]}' -C myc
docker exec cli peer chaincode install -p chaincodedev/chaincode/smart_audit/fabric/location -n Location -v 0
docker exec cli peer chaincode instantiate -n Location -v 0 -c '{"Args":[]}' -C myc
docker exec cli peer chaincode install -p chaincodedev/chaincode/smart_audit/fabric/audit -n audit -v 0
docker exec cli peer chaincode instantiate -n audit -v 0 -c '{"Args":["init","maintainer1","maintainer2"]}' -C myc

```

### 快速录入数据

```shell
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["registerAuditee", "ZhanSan"]}'
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["registerRule", "AND", "Time", "(>= 9) AND (<= 18)", "Location", "IN(39.9 116.3 1000)", "FaceRecognize", "", "ObjectRecognize", ""]}'
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["registerProject", "POS Audit", "This is location","0", "0"]}'
docker exec cli peer chaincode invoke -n audit -C myc -c '{"Args": ["addEvent", "0", "0", "0", "1589532423", "Time", "2020-05-26T15:04:05.000Z", "Location", "39.901 116.299", "FaceRecognize", "/9j/4SMF...", "ObjectRecognize", "iVBORw0..."]}'
```

