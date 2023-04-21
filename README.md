
# Go Redis HTTPS Docker
go redis https docker

## openssl Installation
you can run OpenSSL on Windows. You can either install it directly or use the version that comes with Git for Windows. Here are the steps for both methods:

### Method 1: Using Git for Windows

If you haven't already, download and install Git for Windows.
Run the "Git Bash" application that comes with Git for Windows.
In Git Bash, you can use the openssl command just like you would on a Unix-based system. For example, to create a self-signed certificate and private key, you can use the same command mentioned before:

```bash
 openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365

```
###  Method 2: Installing OpenSSL directly on Windows

Download the precompiled OpenSSL binaries for Windows from this link. https://slproweb.com/products/Win32OpenSSL.html  Choose the appropriate version for your system (either 32-bit or 64-bit) and install it.

Add the OpenSSL installation folder to the Windows PATH environment variable:

Right-click on "Computer" or "This PC" on the desktop or in File Explorer and select "Properties."
Click on "Advanced system settings" on the left side of the System Properties window.
In the System Properties window, click on the "Environment Variables" button.
In the Environment Variables window, under "System variables," find the variable named Path and click on the "Edit" button.
Add the path to the OpenSSL installation folder (e.g., C:\Program Files\OpenSSL-Win64\bin) to the end of the variable value. Make sure to separate it from the previous path with a semicolon (;).
Click "OK" to close the windows and apply the changes.
Open a new Command Prompt or PowerShell window, and you should be able to use the openssl command directly. For example:


 ```bash
 openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365

```
Now you can use OpenSSL on Windows to create a self-signed certificate and private key for your application.

## Issues
First, locate the openssl.cnf file on your system. It is usually found in the bin or ssl subdirectory of the OpenSSL installation folder. For example, it might be located at C:\Program Files\OpenSSL-Win64\bin\openssl.cnf or C:\Program Files\OpenSSL-Win64\ssl\openssl.cnf.

Once you have located the openssl.cnf file, modify the command to include the -config option followed by the path to the file:
 ```bash
 openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -config "C:\Program Files\OpenSSL-Win64\bin\openssl.cnf"


```
Replace "C:\Program Files\OpenSSL-Win64\bin\openssl.cnf" with the correct path to the openssl.cnf file on your system. This should resolve the error, and OpenSSL should now create the certificate and private key without any issues.

## No PEM pass phrase
When you run the openssl command to generate a self-signed certificate and private key, it asks for a PEM pass phrase to encrypt the private key file. This adds an extra layer of security, as anyone who wants to use the private key will need to provide the pass phrase to decrypt it.

You can choose any pass phrase you'd like, but make sure it's something you can remember, as you'll need it later when using the private key. After entering the pass phrase, you'll be asked to verify it by entering it again.

If you want to generate a private key without a pass phrase (not recommended for production environments), you can add the -nodes option to the command:
 ```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -config "C:\Program Files\OpenSSL-Win64\bin\openssl.cnf" -nodes



```

This command will create a private key without encryption, and you won't be prompted for a pass phrase. However, this also means that anyone with access to the private key file can use it without restriction, so use this option with caution.
## License

[MIT](https://choosealicense.com/licenses/mit/)

