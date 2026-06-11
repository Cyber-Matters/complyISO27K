name: Cryptographic Key Management Policy
acronym: CKMP
satisfies:
  TSC:
    - CC9.9
    - CC6.7
  ISO27001:
    - A.8.24
majorRevisions:
  - date: Jan 1 2024
    comment: Initial document
---

# Purpose and Scope

a. This policy defines the requirements for the generation, storage, distribution, rotation, revocation, and destruction of cryptographic keys used to protect the organisation's information assets.

a. This policy applies to all cryptographic keys used within the organisation, including keys used for data encryption, TLS certificates, code signing, SSH access, and API authentication.

a. This policy applies to all employees, contractors, and system administrators who manage or have access to cryptographic keys.

# Background

a. Cryptographic keys are the foundation of data confidentiality and integrity controls. Compromise of a cryptographic key can render encryption controls ineffective. This policy establishes standards to protect keys throughout their lifecycle.

# Policy

a. *Key generation*

    i. Cryptographic keys must be generated using cryptographically secure random number generators.

    i. Key lengths and algorithms must meet current industry standards. At minimum:

        1. Symmetric encryption: AES-256 or equivalent.

        1. Asymmetric encryption: RSA-2048 or higher, or equivalent elliptic curve algorithms (e.g. ECDSA P-256).

        1. Hashing: SHA-256 or higher.

    i. Keys must never be generated on systems that are not under the organisation's control.

a. *Key storage*

    i. Private keys must be stored in a hardware security module (HSM), secrets management system, or other approved key store. Storing private keys in plaintext files or source code repositories is prohibited.

    i. Access to key stores must be restricted to authorised personnel and service accounts, and must be protected by multi-factor authentication.

a. *Key distribution*

    i. Keys must only be distributed to authorised recipients using encrypted channels.

    i. Key material must never be transmitted via unencrypted email, chat, or similar channels.

a. *Key rotation*

    i. Encryption keys must be rotated in accordance with the following schedule, or sooner if compromise is suspected:

        1. TLS certificates: at least annually, or upon expiry.

        1. Data encryption keys: at least annually for keys protecting high-value data.

        1. API keys and service credentials: at least annually.

    i. Automated rotation must be implemented wherever technically feasible.

a. *Key revocation and destruction*

    i. Keys must be revoked immediately upon suspected or confirmed compromise, or when the associated system or service is decommissioned.

    i. Revoked or expired keys must be securely destroyed in a manner that prevents recovery.

    i. Key revocation events must be documented and reviewed by the Information Security team.

a. *Audit and oversight*

    i. The lifecycle of all production cryptographic keys must be documented and auditable.

    i. The Information Security team must review key management practices at least annually.
