const express = require('express');
const app = express();
const port = 3000;

app.use(express.json());

//  In-memory store for assets
const assets = {};

//  POST /createAsset
app.post('/createAsset', (req, res) => {
  const asset = req.body;
  assets[asset.id] = asset;
  console.log('Asset stored:', asset);
  res.send(`Asset ${asset.id} created successfully.`);
});

//  GET /readAsset/:id
app.get('/readAsset/:id', (req, res) => {
  const assetId = req.params.id;
  const asset = assets[assetId];

  if (asset) {
    res.json(asset);
  } else {
    res.status(404).send('Asset not found.');
  }
});

// 🔹 Start server
app.listen(port, () => {
  console.log(` API server running at http://localhost:${port}`);
});
