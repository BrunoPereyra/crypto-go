# primer y sugundo bloque

En resumen, rand.Read(key) se utiliza para generar una clave aleatoria,
mientras que el cifrado AES se utiliza para cifrar o descifrar datos
utilizando una clave previamente generada. La generación de la clave
aleatoria con rand.Read es un paso previo necesario antes de utilizar
el cifrado AES para proteger los datos.
osea que aes.NewCipher(key) hace que ese slice de bytes que generamos(nuestra clave)
sea caaz de cifrar y desifrar informacion de forma simetrica

# tercer bloques

rand.Read(iv): Esta línea de código genera y llena el vector de inicialización (IV)
con valores aleatorios. El IV se utiliza como el primer bloque de datos en el proceso
de cifrado en modo CBC (Cipher Block Chaining). El IV es necesario para proporcionar
aleatoriedad y evitar patrones repetitivos en el cifrado.

## ...

AES es el algoritmo de cifrado utilizado y el modo CBC es un modo de operación que se aplica sobre AES para proporcionar mayor seguridad y aleatoriedad en el cifrado.
