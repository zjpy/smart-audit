## 环境准备

### 系统环境准备

目前测试环境为：

- Mac OSX、Ubuntu 16.04
- JDK 1.8+

### 启动链

进入到FISCO BCOS链运行所在目录（以下该目录简称为`$fisco`），执行如下命令：

1. 通过如下命令下载用于准备链环境的脚本：

   `curl -LO https://github.com/FISCO-BCOS/FISCO-BCOS/releases/download/v2.3.0/build_chain.sh && chmod u+x build_chain.sh`

2. 搭建单群组4节点联盟链：

   ```shell
   ./build_chain.sh -l "127.0.0.1:4" -p 30300,20200,8545
   ```

   执行成功后会在`nodes`子目录下生成所有相关文件

3. 启动链

   ```shell
   nodes/127.0.0.1/start_all.sh
   ```

4. 成功启动后可以通过以下命令查看以启动的节点进程：

   ```shell
   ps -ef | grep -v grep | grep fisco-bcos
   ```

### 准备控制台

进入`$fisco`目录并执行以下操作：

1. 通过如下命令获取控制台：

   ```shell
   nodes/127.0.0.1/download_console.sh
   ```

2. 通过如下命令拷贝控制台配置文件：

   ```shell
   cp -n console/conf/applicationContext-sample.xml console/conf/applicationContext.xml
   ```

3. 配置控制台证书：

   ```shell
   cp nodes/127.0.0.1/sdk/* console/conf/
   ```

4. 至此完成控制台的启动工作，通过如下命令启动控制台：

   ```shell
   console/start.sh
   ```

5. 在命令行中输入如下命令以查询当前块高：

   ```shell
   getBlockNumber
   ```

   由于在FISCO BCOS中只有新交易到达才会增长块高，因此这里返回的结果为`0`

### 拷贝合约

1. 进入到smart-audit源码文件下，然后进入到`src/fisco`子目录，通过如下命令拷贝所有合约文件到控制台所维护的目录下：

   ```shell
   cp -r audit/*.sol audit/interface $fisco/console/contracts/solidity/
   ```

2. 通过如下命令拷贝人脸识别预言机服务相关的智能合约到控制台所维护的目录下：

   ```shell
   cp oracles/face/*.sol $fisco/console/contracts/solidity/
   ```

3. 通过如下命令拷贝物体识别预言机服务相关的智能合约到控制台所维护的目录下：

   ```shell
   cp oracles/identify/*.sol $fisco/console/contracts/solidity/
   ```

4. 通过如下命令拷贝时间服务预言机服务相关的智能合约到控制台所维护的目录下：

   ```shell
   cp oracles/time/*.sol $fisco/console/contracts/solidity/
   ```

5. 通过如下命令拷贝定位服务预言机服务相关的智能合约到控制台所维护的目录下：

   ```shell
   cp oracles/location/*.sol $fisco/console/contracts/solidity/
   ```

合约拷贝完成后即可在控制台命令行中通过合约名找到相应的智能合约，后续所有操作都将在控制台命令行中执行

## 合约调用

### 部署合约

#### 部署时间合约

1. 通过如下命令部署合约

   ```shell
   deployByCNS DummyTimeService 1.0
   ```

   如果部署合约成功会得到如下响应消息：

   ```shell
   contract address: 0x7edac4c1fd59d55130806dbe537b718594189374
   ```
   
2. 可以通过如下命令查询时间服务的部署合约：

   ```shell
   queryCNS DummyTimeService
   ```

   若步骤1部署成功会得到如下响应：

   ```shell
   ---------------------------------------------------------------------------------------------
   |                   version                   |                   address                   |
   |                     1.0                     | 0x7edac4c1fd59d55130806dbe537b718594189374  |
   ---------------------------------------------------------------------------------------------
   ```

   可以看到这里查询到的合约地址与步骤1的返回结果一致

3. 通过如下命令可以对时间合约的验证方法进行测试：

   ```shell
   callByCNS DummyTimeService:1.0 validate 0 [""]
   ```

   目前验证方法有90%的概率通过验证，如果成功会返回相应的交易hash，响应消息如下：

   ```shell
   transaction hash: 0xb27a1be3394c382bbafbbaea951799292b4e6e64e484e4f0eda68030ffe5b808
   ```

   如果验证失败则返回如下响应消息：

   ```shell
   时间超出正常工作范围
   ```

#### 部署定位合约

1. 通过如下命令部署合约

   ```shell
   deployByCNS DummyLocationService 1.0
   ```

   如果部署合约成功会得到如下响应消息：

   ```shell
   contract address: 0x66038422a9700ca5af804e1bb53a77a20e806de2
   ```

2. 可以通过如下命令查询时间服务的部署合约：

   ```shell
   queryCNS DummyTimeService
   ```

   若步骤1部署成功会得到如下响应：

   ```shell
   ---------------------------------------------------------------------------------------------
   |                   version                   |                   address                   |
   |                     1.0                     | 0x66038422a9700ca5af804e1bb53a77a20e806de2  |
   ---------------------------------------------------------------------------------------------
   ```

   可以看到这里查询到的合约地址与步骤1的返回结果一致

3. 通过如下命令可以对时间合约的验证方法进行测试：

   ```shell
   callByCNS DummyLocationService:1.0 validate 0 [""]
   ```

   目前验证方法有90%的概率通过验证，如果成功会返回相应的交易hash，响应消息如下：

   ```shell
   transaction hash: 0x76bb610b0384cef67325c66ea1b26cd482cc1932ffe5f0cd804b93ecc6515e0e
   ```

   如果验证失败则返回如下响应消息：

   ```shell
   地理位置超出正常工作范围
   ```

#### 部署人脸识别合约

1. 通过如下命令部署合约

   ```shell
   deployByCNS DummyFaceService 1.0
   ```

   如果部署合约成功会得到如下响应消息：

   ```shell
   contract address: 0x943eb039fe84f35e7fb8c09535f46441449a76c9
   ```

2. 可以通过如下命令查询时间服务的部署合约：

   ```shell
   queryCNS DummyFaceService
   ```

   若步骤1部署成功会得到如下响应：

   ```shell
   ---------------------------------------------------------------------------------------------
   |                   version                   |                   address                   |
   |                     1.0                     | 0x943eb039fe84f35e7fb8c09535f46441449a76c9  |
   ---------------------------------------------------------------------------------------------
   ```

   可以看到这里查询到的合约地址与步骤1的返回结果一致

3. 通过如下命令可以对时间合约的验证方法进行测试：

   ```shell
   callByCNS DummyFaceService:1.0 validate 0 [""]
   ```

   目前验证方法有>80%的概率通过验证，如果成功会返回相应的交易hash，响应消息如下：

   ```shell
   transaction hash: 0xe61d943eead331b9dba26e8ea6fecc00f5dc0f6e026b1ccb5a175b3209066b06
   ```

   如果验证失败则返回如下响应消息：

   ```shell
   非本人操作
   ```

#### 部署物体识别合约

1. 通过如下命令部署合约

   ```shell
   deployByCNS DummyIdentifyService 1.0
   ```

   如果部署合约成功会得到如下响应消息：

   ```shell
   contract address: 0xa348d87bb437b88dacd551be17e6111ec0b1f745
   ```

2. 可以通过如下命令查询时间服务的部署合约：

   ```shell
   queryCNS DummyIdentifyService
   ```

   若步骤1部署成功会得到如下响应：

   ```shell
   ---------------------------------------------------------------------------------------------
   |                   version                   |                   address                   |
   |                     1.0                     | 0xa348d87bb437b88dacd551be17e6111ec0b1f745  |
   ---------------------------------------------------------------------------------------------
   ```

   可以看到这里查询到的合约地址与步骤1的返回结果一致

3. 通过如下命令可以对时间合约的验证方法进行测试：

   ```shell
   callByCNS DummyIdentifyService:1.0 validate 0 [""]
   ```

   目前验证方法有>80%的概率通过验证，如果成功会返回相应的交易hash，响应消息如下：

   ```shell
   transaction hash: 0x3a72e3d2cd7476d35767ff89cc01566f118ee9e343f27bc99532f345e9b9ed81
   ```

   如果验证失败则返回如下响应消息：

   ```shell
   缺少需要识别物体的数据
   ```

#### 部署审计业务合约



### 调用合约

#### 合约维护人员查询



#### 录入审计当事人



#### 录入规则



#### 录入项目



#### 新增审计事件

1. 


## 参考链接

1. [FISCO BCOS官方节点搭建向导](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/installation.html)
2. [控制台命令官方文档说明](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/manual/console.html#id20)

