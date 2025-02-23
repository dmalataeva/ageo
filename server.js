const express = require('express');
const app = express();
const port = 8000;

app.use(express.static('.', {
  setHeaders: (res, path, stat) => {
    if (path.endsWith('.wasm')) {
      res.setHeader('content-type', 'application/wasm');
    }
  }
}));

app.listen(port, () => {
  console.log(`Server listening at http://localhost:${port}`);
});
