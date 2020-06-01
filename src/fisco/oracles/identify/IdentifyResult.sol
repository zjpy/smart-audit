pragma solidity ^0.4.24;


library IdentifyResult {
    // 调用预言机服务的人脸比对接口返回的结果
    struct IdentifyCompare {
        // 人脸比对结果，0表示请求成功，非0会有响应的错误码意义
        uint32 Result;
        // 相似度评分，在0~100之间取值，越大表示越相似
        uint32 Score;
        // 会话序号
        string Sequence;
    }
}
