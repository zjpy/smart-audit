pragma solidity 0.4.24;
pragma experimental ABIEncoderV2;

contract Audit {
    uint256 numAuditees;
    mapping (uint256 => string) auditees;

    struct Rule {
        string time;
        string location;
        string faceRecognize;
        string objectRecognize;
    }
    uint256 numRules;
    mapping (uint256 => Rule) rules;
    mapping (uint256 => string) ruleNames;

    struct Project {
        string detail;
        uint256 auditeeId;
        uint256 ruleId;
    }
    uint256 numProjects;
    mapping (uint256 => Project) projects;

    mapping (uint256 => mapping (uint256 => mapping (uint256 => Rule))) events;

    function registerAuditee(string name) public {
        numAuditees++;
        auditees[numAuditees] = name;
    }

    function registerRule(string name, string time, string location, string faceRec, string objectRec) public {
        numRules++;
        rules[numRules] = Rule(time, location, faceRec, objectRec);
        ruleNames[numRules] = name;
    }

    function registerProject(string detail, uint256 auditeeId, uint256 ruleId) public {
        numProjects++;
        projects[numProjects] = Project(detail, auditeeId, ruleId);
    }

    function addEvent(uint256 auditeeId, uint256 projectId, uint256 ruleId,
        string time, string location, string faceRec, string objectRec) public {
        events[auditeeId][projectId][ruleId] = Rule(time, location, faceRec, objectRec);
    }

    function getAuditee(uint256 auditeeId) public view returns(string) {
        return auditees[auditeeId];
    }

    function getRule(uint256 ruleId)
        public
        view
        returns(Rule)
    {
        return rules[ruleId];
    }

    function getProject(uint256 projectId) public view returns(Project) {
        return projects[projectId];
    }

    function getMaintainers() public {

    }

    function queryEvents(uint256 auditeeId, uint256 projectId, uint256 ruleId) public view returns(Rule) {
       return events[auditeeId][projectId][ruleId];
    }
}