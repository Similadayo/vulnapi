---
description: A vulnerability occurs when a JSON Web Token (JWT) is signed with an empty secret. In this scenario, the token lacks proper cryptographic protection, making it susceptible to manipulation.
---

# JWT Blank Secret

<table>
    <tr>
        <th>Severity</th>
        <td>High</td>
    </tr>
    <tr>
        <th>CVEs</th>
        <td>
            <ul>
                <li><a href="https://www.cve.org/CVERecord?id=CVE-2019-20933">CVE-2019-20933</a></li>
                <li><a href="https://www.cve.org/CVERecord?id=CVE-2020-28637">CVE-2020-28637</a></li>
            </ul>
        </td>
    </tr>
    <tr>
        <th>Classifications</th>
        <td>
            <ul>
                <a href="https://cwe.mitre.org/data/definitions/287.html">CWE-287: Improper Authentication</a>
            </ul>
        </td>
    </tr>
    <tr>
        <th>OWASP Category</th>
        <td>
            <a href="https://owasp.org/API-Security/editions/2023/en/0xa2-broken-authentication/">OWASP API2:2023 Broken Authentication</a>
        </td>
    </tr>
</table>

A vulnerability occurs when a JSON Web Token (JWT) is signed with an empty secret. In this scenario, the token lacks proper cryptographic protection, making it susceptible to manipulation. Attackers can modify the token's claims and content without detection, potentially leading to unauthorized access and data tampering.

## Example

Here is a valid JWT signed with HS256 algorithm:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDI2MjIsImlhdCI6MTUxNjIzOTAyMiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMmNiMzA3YmEtYmI0Ni00MTk0LTg1NGYtNDc3NDA0NmQ5YzliIn0.SCC35SSgMSMr0kV1i_TuPAhiSGtsC1cFGCfvaus5GyU
```

This decoded JWT contains, this parts:

```json:header
{
  "alg": "HS256",
  "typ": "JWT"
}
```

```json:payload
{
  "iat": 1516239022,
  "exp": 1516242622,
  "name": "John Doe",
  "sub": "2cb307ba-bb46-4194-854f-4774046d9c9b"
}
```

The following JWT is signed with an empty secret:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDI2MjIsImlhdCI6MTUxNjIzOTAyMiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMmNiMzA3YmEtYmI0Ni00MTk0LTg1NGYtNDc3NDA0NmQ5YzliIn0.SCC35SSgMSMr0kV1i_TuPAhiSGtsC1cFGCfvaus5GyU
```

# What is the impact?

Signing a JWT with a blank secret has a significant impact on the security of the token. A blank secret means that there is no secret key used to sign the token, making it vulnerable to tampering and unauthorized access.

By signing a JWT with a blank secret, anyone with access to the token can modify its contents without detection. This can lead to various security risks, such as impersonation, data tampering, and unauthorized access to protected resources.

## How to test?

TODO: VulnAPI Command

## How to remediate?

To remediate the JWT blank secret vulnerability, ensure that all JWTs are signed with a secure secret key. Use strong cryptographic algorithms and keep the secret key confidential to prevent unauthorized access and tampering of the tokens.
