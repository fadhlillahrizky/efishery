const router = require("express").Router();

const auth = require("./auth");

router.get("/", function (req, res) {
    res.send("Hi, is there anything I can help you with?");
});


router.use("/auth", auth);

module.exports = router;
