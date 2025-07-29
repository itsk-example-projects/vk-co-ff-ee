package web

import (
	"html/template"
)

type pageData struct {
	PlainInput    string
	CipherInput   string
	EncryptResult string
	DecryptResult string
}

var pageTemplate = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8"/>
    <title>VK CO FF EE</title>
    <style>
        body {
            font-family: Consolas, monospace;
            max-width: 800px;
            margin: 40px auto;
            background-color: #121212;
            color: #e0e0e0;
        }
        textarea, input[type=text] {
			font-family: Consolas, monospace;
            width: 100%;
            box-sizing: border-box;
            background-color: #1e1e1e;
            color: #ffffff;
            border: 1px solid #333;
            padding: 8px;
            border-radius: 4px;
        }
        form {
            margin-bottom: 30px;
        }
        label {
			font-family: Consolas, monospace;
            display: block;
            margin-bottom: 5px;
            font-weight: bold;
        }
        button {
			font-family: Consolas, monospace;
            padding: 8px 16px;
            font-size: 1em;
            background-color: #333;
            color: #e0e0e0;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            transition: background-color 0.2s;
        }
        button:hover {
            background-color: #555;
            color: #ffffff;
        }
        h1, h2, h3 {
            color: #fff;
        }
        textarea[readonly] {
			font-family: Consolas, monospace;
            opacity: 0.9;
        }
    </style>
</head>
<body>
    <h1>VK CO FF EE</h1>

    <h2>Encrypt</h2>
    <form method="post" action="/">
        <input type="hidden" name="action" value="encrypt"/>
        <label for="plain">Plain text</label>
        <textarea id="plain" name="input" rows="5" placeholder="Enter text to encrypt">{{.PlainInput}}</textarea>
        <label for="custom_enc">Custom key (optional)</label>
        <input id="custom_enc" type="text" name="customKey" placeholder="Enter custom key"/>
        <p><button type="submit">Encrypt</button></p>
    </form>
    {{if .EncryptResult}}
    <h3>Encrypted result</h3>
    <textarea rows="5" readonly>{{.EncryptResult}}</textarea>
    {{end}}

    <h2>Decrypt</h2>
    <form method="post" action="/">
        <input type="hidden" name="action" value="decrypt"/>
        <label for="cipher">Cipher text</label>
        <textarea id="cipher" name="input" rows="5" placeholder="Enter cipher text to decrypt">{{.CipherInput}}</textarea>
        <label for="custom_dec">Custom key (optional)</label>
        <input id="custom_dec" type="text" name="customKey" placeholder="Enter custom key"/>
        <p><button type="submit">Decrypt</button></p>
    </form>
    {{if .DecryptResult}}
    <h3>Decrypted result</h3>
    <textarea rows="5" readonly>{{.DecryptResult}}</textarea>
    {{end}}
</body>
</html>
`))
