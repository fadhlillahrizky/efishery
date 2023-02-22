const router = require("express").Router();

router.get("/", function (req, res) {
    res.send("Hi, is there anything I can help you with?");
});

module.exports = router;
