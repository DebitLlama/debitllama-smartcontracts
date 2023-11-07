const { encrypt, decrypt, getEncryptionPublicKey } = require("@metamask/eth-sig-util");
const Buffer = require('buffer/').Buffer;

function encryptData(publicKey, data) {

    return encrypt({ publicKey, data, version: 'x25519-xsalsa20-poly1305' });
}

function decryptData(privateKey, encryptedData) {
    return decrypt({ encryptedData, privateKey: privateKey.substring(2) });
}

function packEncryptedMessage(encryptedMessage) {
    const nonceBuf = Buffer.from(encryptedMessage.nonce, 'base64')
    const ephemPublicKeyBuf = Buffer.from(encryptedMessage.ephemPublicKey, 'base64')
    const ciphertextBuf = Buffer.from(encryptedMessage.ciphertext, 'base64')
    const messageBuff = Buffer.concat([
        Buffer.alloc(24 - nonceBuf.length),
        nonceBuf,
        Buffer.alloc(32 - ephemPublicKeyBuf.length),
        ephemPublicKeyBuf,
        ciphertextBuf
    ])
    return '0x' + messageBuff.toString('hex')
}

function unpackEncryptedMessage(encryptedMessage) {
    if (encryptedMessage.slice(0, 2) === '0x') {
        encryptedMessage = encryptedMessage.slice(2)
    }
    const messageBuff = Buffer.from(encryptedMessage, 'hex')
    const nonceBuf = messageBuff.slice(0, 24)
    const ephemPublicKeyBuf = messageBuff.slice(24, 56)
    const ciphertextBuf = messageBuff.slice(56)
    return {
        version: 'x25519-xsalsa20-poly1305',
        nonce: nonceBuf.toString('base64'),
        ephemPublicKey: ephemPublicKeyBuf.toString('base64'),
        ciphertext: ciphertextBuf.toString('base64')
    }
}

function getPublicKeyFromPrivateKey(privkey) {
    // Get the public encryption key from a private key
    // Used for passkey encryption
    return getEncryptionPublicKey(privkey)
}

module.exports = { encryptData, decryptData, packEncryptedMessage, unpackEncryptedMessage, getPublicKeyFromPrivateKey }