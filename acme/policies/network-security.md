name: Network Security Policy
acronym: NSP
satisfies:
  TSC:
    - CC6.6
    - CC6.7
  ISO27001:
    - A.8.20
    - A.8.21
    - A.8.22
    - A.8.23
majorRevisions:
  - date: Jan 1 2024
    comment: Initial document
---

# Purpose

This policy establishes requirements for the security of {{.Name}}'s networks and network services, including segregation, web filtering, and the protection of information in transit.

# Scope

This policy applies to all network infrastructure, services, and connections owned, operated, or used by {{.Name}}, including corporate networks, cloud environments, and remote access infrastructure.

# Network Security

{{.Name}}'s networks are managed and controlled to protect information in systems and applications. Network security controls include firewalls, intrusion detection and prevention systems, and network access controls. Network activity is monitored and logged to detect anomalous behaviour. Network security configurations are reviewed at least annually and following significant changes. Responsibility for network security is clearly assigned.

# Security of Network Services

All network services, whether provided internally or by external suppliers, are subject to defined security requirements. Service level agreements and contracts with network service providers include security obligations. The security features of network services, including authentication, encryption, and access control, are specified and verified prior to use. Third-party network services are reviewed periodically to ensure compliance with security requirements.

# Segregation of Networks

Networks are segmented based on the sensitivity of information processed, user groups, and system function. Production, development, and test environments are separated. Administrative and management networks are isolated from user-facing networks. Network segmentation is enforced through firewalls, VLANs, and access control lists. The network segmentation design is documented and reviewed at least annually.

# Web Filtering

Access to internet content from {{.Name}} networks is controlled through a web filtering solution. Categories of content that are blocked include malicious sites, known threat infrastructure, and content that violates the Acceptable Use Policy. Web filtering rules are reviewed and updated regularly. Exceptions to the web filtering policy require documented approval and are time-limited where possible. Web access logs are retained and reviewed for anomalous activity.
