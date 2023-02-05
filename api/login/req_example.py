import requests
from http.cookies import SimpleCookie

# Perform Login
login_url = "http://localhost:8080/api/v1/verify"
login_packet = {
    "message":"localhost:8080 wants you to sign in with your Ethereum account:\n0x9D85ca56217D2bb651b00f15e694EB7E713637D4\n\nSign in with Ethereum to the app.\n\nURI: http://localhost:8080\nVersion: 1\nChain ID: 1\nNonce: spAsCWHwxsQzLcMzi\nIssued At: 2022-01-29T03:22:26.716Z",
    "signature":"0xe117ad63b517e7b6823e472bf42691c28a4663801c6ad37f7249a1fe56aa54b35bfce93b1e9fa82da7d55bbf0d75ca497843b0702b9dfb7ca9d9c6edb25574c51c"
}
login_response = requests.post(login_url, json=login_packet)

# Server will give you sessionID (sID) in header
ck = SimpleCookie()
ck.load(login_response.headers['Set-Cookie'])
print(ck)

ck.items()
my_session_id = ck.get('sID').value


# Perform action that needs valid session ID
url = "http://localhost:8080/api/test"
cookie = {
    "sID": my_session_id
}

response = requests.get(url, cookies=cookie)
print(response.content)  # {"message": "testing point"}

