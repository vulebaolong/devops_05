const http = require('node:http');
const os = require('node:os');

http.createServer((req, res) => {
  console.log('Request xu ly boi: ' + os.hostname());
  res.setHeader('Content-Type', 'text/plain; charset=utf-8');
  res.setHeader('Cache-Control', 'no-store');
  res.end('Container: ' + os.hostname() + '\n');
}).listen(3000);
