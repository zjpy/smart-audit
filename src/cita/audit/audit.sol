pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./Rules.sol";

// Smart Audit 合约
contract Audit is Rules {

    /// @dev 规则合约的构造方法，这里实现Rules.
    /// @param face 人脸识别服务合约地址.
    /// @param identify 物体识别服务合约地址.
    /// @param time 时间服务合约地址.
    /// @param location 定位服务合约地址.
    constructor(
        address face,
        address identify,
        address time,
        address location
    ) Rules (face, identify, time, location) {

    }

    /// @dev 注册一个验证人的事件.
    /// @param AuditeeID Auditee ID.
    event registerAuditeeEvent(uint32 AuditeeID);

    /// @dev 注册一个项目的事件.
    /// @param ProjectID Project ID.
    event registerProjectEvent(uint32 ProjectID);

    uint32 numAuditees;
    mapping (uint32 => string) auditees;

    struct Project {
        string detail;
        uint32 auditeeID;
        uint32 relationID;
    }
    uint32 numProjects;
    mapping (uint32 => Project) projects;

    struct Event {
        uint32 auditeeID;
        uint32 projectID;
        string[] value;
    }
    uint32 numEvents;
    mapping (uint32 => Event) events;

    function registerAuditee(string name) public {
        numAuditees++;
        auditees[numAuditees] = name;

        emit registerAuditeeEvent(numAuditees);
    }

    function registerProject(string detail, uint32 auditeeID, uint32 relationID) public {
        string storage auditee = auditees[auditeeID];
 //       require(auditee, "auditee 不存在");

        ValidationRelationship storage relation = relationMap[relationID];
 //       require(relation, "relation 不存在");

        numProjects++;
        projects[numProjects] = Project(detail, auditeeID, relationID);

        emit registerProjectEvent(numProjects);
    }

    function addEvent(uint32 auditeeID, uint32 projectID, string[] value) public {
        string storage auditee = auditees[auditeeID];
 //       require(auditee, "auditee 不存在");

        Project storage project = projects[projectID];
 //       require(project, "项目不存在");

        validateRules(project.relationID, value);

        numEvents++;
        events[numEvents] = Event(auditeeID, projectID, value);
    }

    function getAuditee(uint32 auditeeID) public view returns(string) {
        return auditees[auditeeID];
    }

    function getRule(uint32 relationID)
        public
        view
        returns(ValidationRelationship)
    {
        return relationMap[relationID];
    }

    function getProject(uint32 projectID) public view returns(Project) {
        return projects[projectID];
    }

    function getMaintainers(uint32 projectID) public view returns(uint32) {
        Project storage project = projects[projectID];
 //       require(project, "项目不存在");

        return project.auditeeID;
    }

    function queryEvents(uint32 eventID) public view returns(Event) {
       return events[eventID];
    }
}