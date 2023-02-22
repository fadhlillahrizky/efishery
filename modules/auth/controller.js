const { httpResponse } = require("../../utilities/response");
const Service = require("./service");

module.exports = class controller {
    constructor(option) {
        this.service = option?.service || new Service();
    }

    async login(req, res) {
        httpResponse(await this.service.login(req.body), res);
    }

    async register(req, res) {
        httpResponse(await this.service.register(req.body), res);
    }

    async checkToken(req, res) {
        const token = req.header("authorization").split(' ')[1] ?? '';
        httpResponse(await this.service.checkToken(token), res);
    }
}
