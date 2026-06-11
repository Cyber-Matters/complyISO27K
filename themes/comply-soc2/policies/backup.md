name: Backup and Recovery Policy
acronym: BRP
satisfies:
  TSC:
    - A1.2
    - A1.3
    - CC9.1
  ISO27001:
    - A.8.13
    - A.5.29
    - A.5.30
majorRevisions:
  - date: Jan 1 2024
    comment: Initial document
---

# Purpose and Scope

a. This policy defines the requirements for backing up organisational data and information systems, and for verifying that data can be successfully restored.

a. This policy applies to all production systems, databases, and data repositories that contain business-critical or confidential information.

a. This policy applies to system owners, administrators, and any third-party service providers responsible for data hosting on behalf of the organisation.

# Background

a. Data loss can result from hardware failure, software errors, accidental deletion, ransomware, or other disruptive events. Reliable, tested backups are a fundamental control for ensuring the availability and resilience of the organisation's information assets.

# Policy

a. *Backup requirements*

    i. All production data must be backed up at a frequency commensurate with its criticality:

        1. Critical systems and databases: daily incremental and weekly full backups at minimum.

        1. Supporting systems: weekly full backups at minimum.

    i. Backups must be stored in a physically separate location from the primary system, or in a geographically separate cloud region.

    i. Backup media and storage must be protected using encryption consistent with the Encryption Policy.

    i. Access to backup systems and media must be restricted to authorised personnel only.

a. *Retention*

    i. Backups must be retained for a minimum of ninety (90) days, unless legal, regulatory, or contractual obligations require a longer retention period.

    i. Retention periods for specific data types are defined in the Data Retention Policy.

a. *Testing and verification*

    i. Restoration tests must be performed at least quarterly for critical systems to confirm that backup data is complete and recoverable.

    i. Results of restoration tests must be documented and reviewed by the system owner.

    i. Any failures or anomalies identified during testing must be remediated promptly and re-tested.

a. *Roles and responsibilities*

    i. System owners are responsible for ensuring that backups are configured, monitored, and tested in accordance with this policy.

    i. The Information Security team is responsible for reviewing backup compliance at least annually.

a. *Incident response*

    i. In the event of data loss, the incident must be reported in accordance with the Security Incident Response Policy. Recovery from backup must be initiated by the system owner as soon as practicable.
