"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const express = require("express");
const app = express();
const port = 8000;
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.get('/', (req, res) => {
    res.status(200).send("Welcome to golang niggus!");
});
app.post('/post', (req, res) => {
    let jsonData = req.body;
    res.status(200).send(JSON.stringify({ success: true, message: "Data posted", data: jsonData }));
});
app.post('/postform', (req, res) => {
    let jsonData = req.body;
    res.status(200).send(JSON.stringify({ success: true, message: "Data posted", data: JSON.stringify(jsonData) }));
});
app.listen(port, () => {
    console.log(`Dera serving on port ${port}`);
});
