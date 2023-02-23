const { httpResponse } = require("../../utilities/response");
const Service = require("./service");

module.exports = class controller {
    constructor(option) {
        this.service = option?.service || new Service();
    }

    async list(req, res) {
        httpResponse(await this.service.list(req.nodeCache), res);
    }

    async aggregate(req, res) {
        httpResponse(await this.service.aggregate(), res);
    }
}
