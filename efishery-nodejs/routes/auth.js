const router = require("express").Router();
const Controller = require("../modules/auth/controller");


router.post("/register", async (req, res, next) => {
    try {
        await new Controller().register(req, res);
    } catch (e) {
        next(e);
    }
});

router.post("/login", async (req, res, next) => {
    try {
        await new Controller().login(req, res);
    } catch (e) {
        next(e);
    }
});

router.get("/check-token", async (req, res, next) => {
    try {
        await new Controller().checkToken(req, res);
    } catch (e) {
        next(e);
    }
});

module.exports = router;
