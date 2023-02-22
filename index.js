require("dotenv").config();

const express = require("express");
const path = require("path");
const cors = require("cors");
const helmet = require("helmet");
const routers = require("./routes");


const app = express();


app.use(cors()); // include before other routes
app.use(helmet());

app.use(
    express.json({
        limit: "3mb",
    })
);
app.use(express.urlencoded({ extended: false }));

app.use("/api", routers);


// client error handler
app.use((err, req, res, next) => {
    if (req.xhr) {
        res.status(500).send({
            success: false,
            message: "Something failed!",
        });
    } else {
        next(err);
    }
});


module.exports = app;
