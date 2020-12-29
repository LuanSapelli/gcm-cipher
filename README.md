# Golang GCM Cipher 🚀

<p>
Simple API to encrypt and decrypt strings using the AES/GCM method
</p>

### Step by step

##### Executing the API
1. Set the enviroment vars (`.env.example`), these vars are encryption keys, you can change them as you wish
2. Execute `go run main.go`
3. In postman or browser, use the `health-check` endpoint (`http://localhost/gcm-cipher/api/health-check`), if you get this response, smile! (`"status": "alive and kicking!"`)

##### Encrypt Endpoint
1. URL: `http://localhost/gcm-cipher/api/encrypt-data`
2. METHOD: `POST`
3. REQUEST BODY: `{"string": "string to encrypt"}`
4. RESPONSE BODY: `{"encrypted_string": "uAeAKXYqkrlOVnTTbaF4qw==$wztsL+JtOfbVsqqTKOVEQ0c="}`

##### Decrypt Endpoint
1. URL: `http://localhost/gcm-cipher/api/decrypt-data`
2. METHOD: `POST`
3. REQUEST BODY: `{"string": "uAeAKXYqkrlOVnTTbaF4qw==$wztsL+JtOfbVsqqTKOVEQ0c="}`
4. RESPONSE BODY: `{"decrypted_string": "string to encrypt"}`
