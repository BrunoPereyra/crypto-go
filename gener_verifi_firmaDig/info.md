# Criptografía asimétrica: El código utiliza el algoritmo de criptografía asimétrica RSA, que utiliza un par de claves matemáticamente relacionadas: una clave pública y una clave privada.

# Generación de claves: La función generateKeyPair genera un par de claves RSA de 2048 bits utilizando rsa.GenerateKey. La clave privada se mantiene en secreto, mientras que la clave pública se puede compartir con otros.

# Función de resumen criptográfico: Antes de firmar el mensaje, se calcula un resumen criptográfico del mensaje utilizando SHA-256 mediante sha256.Sum256. El resumen es una representación condensada del mensaje.

# Firma digital: La función signMessage toma el mensaje y la clave privada y firma el resumen del mensaje utilizando rsa.SignPKCS1v15. Esto genera la firma digital, que es el resultado cifrado del resumen utilizando la clave privada.

# Verificación de firma: La función verifySignature toma el mensaje, la firma y la clave pública. Calcula el resumen del mensaje y luego utiliza rsa.VerifyPKCS1v15 para verificar la firma utilizando la clave pública correspondiente.
