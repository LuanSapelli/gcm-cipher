# golang-gcm-cipher 🧙🏽‍♂️

Simple API to encrypt and decrypt strings using the AES/GCM method.

### how-to

##### Executing the API
- Set the enviroment vars (`.env.example`), these vars are encryption keys, you can change them as you wish.
- Execute `go run main.go`.
- In postman or browser, use the `health-check` endpoint (`http://localhost/gcm-cipher/api/health-check`) to check if the service is running. 

##### Encrypt Endpoint 🔒
- URL: `http://localhost/gcm-cipher/api/encrypt-data`
- METHOD: `POST`
- REQUEST BODY: `{"string": "string to encrypt"}`
- RESPONSE BODY: `{"encrypted_string": "uAeAKXYqkrlOVnTTbaF4qw==$wztsL+JtOfbVsqqTKOVEQ0c="}`

##### Decrypt Endpoint 🔓
- URL: `http://localhost/gcm-cipher/api/decrypt-data`
- METHOD: `POST`
- REQUEST BODY: `{"string": "uAeAKXYqkrlOVnTTbaF4qw==$wztsL+JtOfbVsqqTKOVEQ0c="}`
- RESPONSE BODY: `{"decrypted_string": "string to encrypt"}`
