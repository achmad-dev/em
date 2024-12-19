### Notes architecture are in excalidraw


### Library used
backend:
- fiber
- sqlx (to connect to db)
- hmac (to secure connection with client)
- bcrypt
frontend:
- react
- vite
- crypto-js (hmac) (to secure connection with backend)
- shadcn
development/deployment:
-docker
-docker compose
-make
-bash

### Notes for app
the users can only see it's own events
for vendor user can accept or reject it's event that proposed to them


### Notes to prevent mitm
To really protect against man-in-the-middle attacks, you have to:

apply some kind of integrity check on all exchanges between client and server;
enforce at least one-way authentication (the server must authenticate the client, or the client must authenticate the server).
Basically, doing what SSL does. In a Web context with plain HTTP, the client is stupid and won't do the necessary things. To make the client "intelligent", you have to include some code on the client side, i.e. Javascript -- but if you don't download that Javascript over HTTPS, you lose.

So my guess is that your interlocutor is misguided, or deliberately lies to you, or both.

4. Use HMAC (Hash-based Message Authentication Code) for Integrity Checking:
HMAC ensures that messages haven't been tampered with during transmission. It uses a shared secret key between the sender and receiver to hash the message. If the hash doesn’t match at the receiver’s end, it indicates that the message has been altered.

Example of HMAC in C#:

csharp

using System.Security.Cryptography;
using System.Text;
public string ComputeHMACSHA256(string message, string key) {
  using(HMACSHA256 hmac = new HMACSHA256(Encoding.UTF8.GetBytes(key))) {
    byte[] hashMessage = hmac.ComputeHash(Encoding.UTF8.GetBytes(message));
    return Convert.ToBase64String(hashMessage);
  }
}
Monitoring and Logging for Suspicious Activity:
Add logging and monitoring to your middleware to track unusual patterns, such as repeated access attempts or unexpected payloads, which could indicate an ongoing MITM attack.

Example:

csharp

public async Task InvokeAsync(HttpContext context) {
  var ipAddress = context.Connection.RemoteIpAddress.ToString();
  Console.WriteLine($"Request from IP: {ipAddress}");
  await _next(context);
}



### example of X header
x-ktbs-request-id:
- 3804cd45-2a60-4ba9-b307-1dda4e762830
x-ktbs-signature:
- 440e2a5322925ddd65839e9ae0f46dc6c73cb979ef633e37cf3caf50083cb261
x-ktbs-time:
- 1734440030



