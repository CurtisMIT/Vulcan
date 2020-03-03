package tpl

// AddNodeTemplate creates a node template
func AddNodeTemplate() []byte {
	return []byte(`const http = require('http');

const hostname = '127.0.0.1';
const port = 3000;

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end('Hello World');
});

server.listen(port, hostname, () => {
  console.log(` + "`Server running at http://${hostname}:${port}/`" + `);
});
	`)
}

// AddFlaskTemplate creates a flask template
func AddFlaskTemplate() []byte {
	return []byte(`from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, World!'
	`)
}
