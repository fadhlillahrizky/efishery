const router = require("express").Router();
const Verify = require("../utilities/verify-jwt");
const Controller = require("../modules/resources/controller");


router.get("", Verify, async (req, res, next) => {
    try {
        await new Controller().list(req, res);
    } catch (e) {
        next(e);
    }
});

router.get("/aggregate", Verify, async (req, res, next) => {
    try {
        await new Controller().aggregate(req, res);
    } catch (e) {
        next(e);
    }
});

module.exports = router;
