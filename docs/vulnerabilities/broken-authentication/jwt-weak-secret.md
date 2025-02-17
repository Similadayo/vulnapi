---
description: A vulnerability occurs when a JSON Web Token (JWT) is signed with a common, a well-known, or a weak secret. In this scenario, the token lacks proper cryptographic protection, making it susceptible to manipulation.
---

# JWT Weak Secret

<table>
    <tr>
        <th>Severity</th>
        <td>High</td>
    </tr>
    <tr>
        <th>CVEs</th>
        <td>
            <ul>
                <li><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-27172">CVE-2023-27172</a></li>
                <li><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-46943">CVE-2023-46943</a></li>
            </ul>
        </td>
    </tr>
    <tr>
        <th>Classifications</th>
        <td>
            <ul>
                <li><a href="https://cwe.mitre.org/data/definitions/287.html">CWE-287: Improper Authentication</a></li>
                <li><a href="https://cwe.mitre.org/data/definitions/307.html">CWE-307: Improper Restriction of Excessive Authentication Attempts</a></li>
                <li><a href="https://cwe.mitre.org/data/definitions/798.html">CWE-798: Use of Hard-coded Credentials</a></li>
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

A vulnerability occurs when a JSON Web Token (JWT) is signed with a common, a well-known, or a weak secret. In this scenario, the token lacks proper cryptographic protection, making it susceptible to manipulation. Attackers can find the secret then modify the token's claims and content without detection, potentially leading to unauthorized access and data tampering.

## What are the different scenarios?

- **Common Secret**: The secret key used to sign the JWT is a common value, such as `secret`, `password`, or `123456`. Attackers can easily guess (brute-force) the secret key.
- **Well-Known Secret**: The secret key used to sign the JWT is a well-known value or a default value This can happen when you use a default secret key provided by a product, library or framework.. Attackers can find the secret key in public repositories, forums, or documentation.
- **Weak Secret**: The secret key used to sign the JWT is a weak value, such as a short guessable string such as `security2024`. Attackers can use dictionary attacks, rainbow tables, or other brute-force techniques to find the secret key.

## Example

Here is a valid JWT signed with HS256 algorithm and a robust secret key:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDI2MjIsImlhdCI6MTUxNjIzOTAyMiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMmNiMzA3YmEtYmI0Ni00MTk0LTg1NGYtNDc3NDA0NmQ5YzliIn0.ZuwZrXpLRj17vDjOLoOOJ7pr1CN5DnE8Clgn4y-fjNs
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

The following JWT is signed with `secret` secret:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDI2MjIsImlhdCI6MTUxNjIzOTAyMiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMmNiMzA3YmEtYmI0Ni00MTk0LTg1NGYtNDc3NDA0NmQ5YzliIn0.gTgBr6lotpAxs4M46PgUXrjhIN5-gYG4HffKSEIB6Ys
```

# What is the impact?

The impact of using a weak secret key to sign a JWT is significant. Attackers can easily find the secret key and modify the token's claims and content without detection. This can lead to unauthorized access, data tampering, and other security risks.

## How to test?

TODO: VulnAPI Command

## How to remediate?

Ensure to change the secret key to a strong and unique value. Use a secure random generator to create the secret key and store it securely. Rotate the secret key periodically to mitigate the risk of unauthorized access and data tampering.
