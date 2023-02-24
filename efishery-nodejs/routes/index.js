const router = require("express").Router();

const auth = require("./auth");
const resources = require("./resources");

router.get("/", function (req, res) {
    res.send("Hi, is there anything I can help you with?");
});


router.use("/auth", auth);
router.use("/resources", resources);

module.exports = router;
